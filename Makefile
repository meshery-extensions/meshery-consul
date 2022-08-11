e2e-tests-localhost: export MESHERY_ADAPTER_ADDR=localhost

.PHONY: check
check: error
	golangci-lint run

.PHONY: check-clean-cache
check-clean-cache:
	golangci-lint cache clean

.PHONY: e2e-tests-localhost
e2e-tests-localhost:
	$(info Running end-to-end tests on '$(MESHERY_ADAPTER_ADDR)')
	cd tests && bats e2e/

.PHONY: e2e-tests
e2e-tests:
	$(info Running end-to-end tests on '$(MESHERY_ADAPTER_ADDR)')
	cd tests && bats e2e/

.PHONY: protoc-setup
protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/layer5io/meshery/master/meshes/meshops.proto

.PHONY: proto
proto:	
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

.PHONY: docker
docker:
	docker build -t layer5/meshery-consul .

.PHONY: docker-run
docker-run:
	(docker rm -f meshery-consul) || true
	docker run --name meshery-consul -d \
	-p 10002:10002 \
	-e DEBUG=true \
	layer5/meshery-consul

.PHONY: run
run:
	go mod tidy;
	DEBUG=true go run main.go

run-force-dynamic-reg:
	FORCE_DYNAMIC_REG=true DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go

.PHONY: error
error:
	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers
