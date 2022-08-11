// Copyright 2020 Layer5, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consul

import (
	"context"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-consul/consul/oam"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models"
	"github.com/layer5io/meshkit/models/oam/core/v1alpha1"
	"gopkg.in/yaml.v2"
)

type Consul struct {
	adapter.Adapter
}

func New(config config.Handler, log logger.Handler, kubeConfig config.Handler) adapter.Handler {
	return &Consul{
		adapter.Adapter{Config: config, Log: log, KubeconfigHandler: kubeConfig},
	}
}

//CreateKubeconfigs creates and writes passed kubeconfig onto the filesystem
func (h *Consul) CreateKubeconfigs(kubeconfigs []string) error {
	var errs = make([]error, 0)
	for _, kubeconfig := range kubeconfigs {
		kconfig := models.Kubeconfig{}
		err := yaml.Unmarshal([]byte(kubeconfig), &kconfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// To have control over what exactly to take in on kubeconfig
		h.KubeconfigHandler.SetKey("kind", kconfig.Kind)
		h.KubeconfigHandler.SetKey("apiVersion", kconfig.APIVersion)
		h.KubeconfigHandler.SetKey("current-context", kconfig.CurrentContext)
		err = h.KubeconfigHandler.SetObject("preferences", kconfig.Preferences)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("clusters", kconfig.Clusters)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("users", kconfig.Users)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = h.KubeconfigHandler.SetObject("contexts", kconfig.Contexts)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return mergeErrors(errs)
}

// ProcessOAM will handles the grpc invocation for handling OAM objects
func (h *Consul) ProcessOAM(ctx context.Context, oamReq adapter.OAMRequest, hchan *chan interface{}) (string, error) {
	err := h.CreateKubeconfigs(oamReq.K8sConfigs)
	if err != nil {
		return "", err
	}
	h.SetChannel(hchan)
	kubeconfigs := oamReq.K8sConfigs
	var comps []v1alpha1.Component
	for _, acomp := range oamReq.OamComps {
		comp, err := oam.ParseApplicationComponent(acomp)
		if err != nil {
			h.Log.Error(ErrParseOAMComponent)
			continue
		}

		comps = append(comps, comp)
	}

	config, err := oam.ParseApplicationConfiguration(oamReq.OamConfig)
	if err != nil {
		h.Log.Error(ErrParseOAMConfig)
	}
	// If operation is delete then first HandleConfiguration and then handle the deployment
	if oamReq.DeleteOp {
		// Process configuration
		msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg2, ErrProcessOAM(err)
		}

		// Process components
		msg1, err := h.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg1 + "\n" + msg2, ErrProcessOAM(err)
		}

		return msg1 + "\n" + msg2, nil
	}
	// Process components
	msg1, err := h.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1, ErrProcessOAM(err)
	}

	// Process configuration
	msg2, err := h.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1 + "\n" + msg2, ErrProcessOAM(err)
	}

	return msg1 + "\n" + msg2, nil
}
