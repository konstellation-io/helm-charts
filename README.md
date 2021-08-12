# KDL Server

## App

|  Component  | Coverage  |  Bugs  |  Maintainability Rating  |
| :---------: | :-----:   |  :---: |  :--------------------:  |
|  App API  | [![coverage][app-api-coverage]][app-api-coverage-link] | [![bugs][app-api-bugs]][app-api-bugs-link] | [![mr][app-api-mr]][app-api-mr-link] |
|  App UI  | [![coverage][app-ui-coverage]][app-ui-coverage-link] | [![bugs][app-ui-bugs]][app-ui-bugs-link] | [![mr][app-ui-mr]][app-ui-mr-link] |

## Knowledge Graph

|  Component  | Coverage  |  Bugs  |  Maintainability Rating  |
| :---------: | :-----:   |  :---: |  :--------------------:  |
|  Knowledge Graph  | [![coverage][kg-coverage]][kg-coverage-link] | [![bugs][kg-bugs]][kg-bugs-link] | [![mr][kg-mr]][kg-mr-link] |


[app-api-coverage]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_api&metric=coverage

[app-api-coverage-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_api&metric=Coverage

[app-api-bugs]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_api&metric=bugs

[app-api-bugs-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_api&metric=Reliability

[app-api-loc]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_api&metric=ncloc

[app-api-loc-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_api&metric=Coverage

[app-api-mr]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_api&metric=sqale_rating

[app-api-mr-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_api&metric=Maintainability

[app-ui-coverage]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_ui&metric=coverage

[app-ui-coverage-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_ui&metric=Coverage

[app-ui-bugs]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_ui&metric=bugs

[app-ui-bugs-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_ui&metric=Reliability

[app-ui-loc]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_ui&metric=ncloc

[app-ui-loc-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_ui&metric=Coverage

[app-ui-mr]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_server_app_ui&metric=sqale_rating

[app-ui-mr-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_server_app_ui&metric=Maintainability

[kg-coverage]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_konwledge_graph&metric=coverage

[kg-coverage-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_konwledge_graph&metric=Coverage

[kg-bugs]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_konwledge_graph&metric=bugs

[kg-bugs-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_konwledge_graph&metric=Reliability

[kg-loc]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_konwledge_graph&metric=ncloc

[kg-loc-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_konwledge_graph&metric=Coverage

[kg-mr]: https://sonarcloud.io/api/project_badges/measure?project=konstellation_kdl_konwledge_graph&metric=sqale_rating

[kg-mr-link]: https://sonarcloud.io/component_measures?id=konstellation_kdl_konwledge_graph&metric=Maintainability

## Development

### Dependencies

#### Microk8s

The local version of Kubernetes to deploy KDL. The version required is **1.19**.

Linux installation:
```
snap install --channel=1.19/stable microk8s --classic
```

In Mac, if the multipass vm doesn't exist the kdlctl.sh will create it automatically.
You also can do it manually using:

```
source .kdlctl.conf
microk8s install --cpu ${MICROK8S_CPUS} --mem ${MICROK8S_MEMORY} --disk ${MICROK8S_DISK} --channel ${MICROK8S_CHANNEL}
```

#### Docker

Needed to build the KDL images. Installation:

https://docs.docker.com/get-docker/

#### Helm

K8s package manager. Make sure you have v3+. Installation:

https://helm.sh/docs/intro/install/

#### gettext

OS package to fill templates during deployment. Usually it is installed in Mac and Linux.

Ubuntu:
```
sudo apt-get install gettext
```

Mac:
```
brew install gettext
```

#### kubectl

The Kubernetes command-line tool is useful to run commands against Kubernetes clusters.

https://kubernetes.io/docs/tasks/tools/

## Local Environment

This repo contains a tool called `./kdlctl.sh` to handle common actions you need during development.

All the configuration needed to run KDL locally can be found in `.kdlctl.conf` file. Usually you'd be ok with the
default values. Check Microk8s parameters if you need to tweak the resources assigned to the VM (only in Mac).

Run help to get info for each command:

```
./kdlctl.sh --help

  kdlctl.sh -- a tool to manage KDL environment during development.

  syntax: kdlctl.sh <command> [options]

    commands:
      dev     creates a complete local environment.
      start   starts microk8s.
      stop    stops microk8s.
      build   calls docker to build all images and push them to microk8s registry.
      deploy  calls helm to create install/upgrade a kdl release on microk8s.
      restart restarts kdl pods or microk8s useful after build command.

    global options:
      h     prints this help.
      q     silent mode.
```

### Install local environment

To install KDL in your local environment:

```
$ ./kdlctl.sh dev
```

It will install everything in the namespace specified in your development `.kdlconf` file.

### Login to local environment

In order to access the admin app, the login process can be done automatically using this script:

```
$ ./kdlctl.sh login
```

You will see an output like this:

```
OS: Darwin

Login link: https://kdlapp.kdl.192.168.64.2.nip.io
```

You can find the admin credentials `GITEA_ADMIN_USER` and `GITEA_ADMIN_PASSWORD` in the `.kdlctl.conf` file.

### Uninstall local environment

If you want to delete all resources generated into your microk8s run the following command:

```
$ ./kdlctl.sh uninstall
```

## Versioning lifecycle

In the development lifecycle of KLI there are three main stages depend if we are going to add a new feature, release a new version with some features or apply a fix to a current release.

### Alphas

In order to add new features just create a feature branch from master, and after the merger the Pull Request a workflow will run the tests and if everything passes a new alpha tag will be created (like *v0.0-alpha.0*), and a new release will be generated with this tag.

### Releases

After some alpha versions we can create what we call a release, and to do that we have to run manual the Release Action. This workflow will create a new release branch and a new tag like *v0.0.0*. With this tag, a new release will be generated.

### Fixes

If we find out a bug in a release, we can apply a bugfix just by creating a fixed branch from the specific release branch, and creating a Pull Request to the same release branch. When the Pull Request is merged, after passing the tests, a new fix tag will be created just by increasing the patch number of the version, and a new release will be build and released.
