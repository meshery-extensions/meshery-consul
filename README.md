<p style="text-align:center;" align="center"><a href="https://meshery.io"><picture align="center">
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-dark-text-side.svg"  width="70%" align="center" style="margin-bottom:20px;">
  <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
  <img alt="Meshery logo" src="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-tag-light-text-side.png" width="70%" align="center" style="margin-bottom:20px;">
</picture></a><br /><br /></p>

# Meshery Adapter for Consul

[![Docker Pulls](https://img.shields.io/docker/pulls/meshery/meshery-consul.svg)](https://hub.docker.com/r/meshery/meshery-consul)
[![Go Report Card](https://goreportcard.com/badge/github.com/meshery/meshery-consul)](https://goreportcard.com/report/github.com/meshery/meshery-consul)
[![Build Status](https://github.com/meshery/meshery-consul/workflows/Meshery%20Consul/badge.svg)](https://github.com/meshery/meshery-consul/actions)
[![GitHub](https://img.shields.io/github/license/meshery/meshery-consul.svg)](https://opensource.org/licenses/Apache-2.0)
[![GitHub issues by-label](https://img.shields.io/github/issues/meshery/meshery-consul/help%20wanted.svg)](https://github.com/issues?q=is%3Aopen%20is%3Aissue%20archived%3Afalse%20(org%3Ameshery%20OR%20org%3Aservice-mesh-performance%20OR%20org%3Aservice-mesh-patterns%20OR%20org%3Ameshery-extensions)%20label%3A%22help%20wanted%22)
[![Website](https://img.shields.io/website/https/meshery.io/meshery.svg)](https://meshery.io/)
[![Twitter Follow](https://img.shields.io/twitter/follow/mesheryio.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=mesheryio)
[![Discuss Users](https://img.shields.io/discourse/users?server=https%3A%2F%2Fdiscuss.layer5.io)](https://meshery.io/community#discussion-forums)
[![Slack](https://img.shields.io/badge/Slack-@meshery.svg?logo=slack)](https://slack.meshery.io/)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3564/badge)](https://bestpractices.coreinfrastructure.org/projects/3564)

<p style="clear:both;">
<h2><a href="https://meshery.io">Meshery</a></h2>
<a href="https://meshery.io"><img src="/.github/img/readme/meshery-logo-light-text.svg"
style="margin:10px;" width="125px" 
alt="Meshery - the Cloud Native Management Plane" align="left" /></a>
As a self-service engineering platform, <a href="https://meshery.io">Meshery</a> enables collaborative design and operation of cloud native infrastructure. Through it's extension points, Meshery offers the ability to optionally plugin adapters in order to more deeply integrate with specific systems like Consul. 
<br /><br /><p align="center"><i>If you’re using Meshery or if you like the project, please <a href="https://github.com/meshery/meshery/stargazers">★</a> star this repository to show your support! 🤩</i></p>
</p>

<p style="clear:both;">
<h2><a name="contributing"></a><a name="community"></a> <a href="https://slack.meshery.io">Community</a> and <a href="https://docs.meshery.io/project/contributing">Contributing</a></h2>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://layer5.io/community/newcomers">Contributor Journey Map</a> for a tour of resources available to you and jump into our <a href="https://slack.meshery.io">Slack</a>! Contributors are expected to adhere to the <a href="https://github.com/cncf/foundation/blob/master/code-of-conduct.md">CNCF Code of Conduct</a>.

<a href="https://slack.meshery.io"><img alt="Meshery Slack" src="/.github/img/readme/slack-128.png" style="margin-left:10px;padding-top:5px;" width="110px" align="right" /></a>

<a href="https://meshery.io/community"><img alt="Meshery Community" src="/.github/img/readme/community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>

<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on the <a href="https://meshery.io/calendar">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/@mesheryio?sub_confirmation=1">meeting recordings</a>.<br />
✔️ <em><strong>To access the Community Drive</strong></em>, fill <a href="https://layer5.io/newcomer">Community Member Form</a>.<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://meshery.io/community#discussion-forums">Community Forum</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen%20is%3Aissue%20archived%3Afalse%20(org%3Ameshery%20OR%20org%3Aservice-mesh-performance%20OR%20org%3Aservice-mesh-patterns%20OR%20org%3Ameshery-extensions)%20label%3A%22help%20wanted%22">help-wanted label</a>.
</p>

**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
