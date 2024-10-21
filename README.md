# helm-charts

This repository contains Helm charts for deploying Konstellation services.

| Chart | Description | Version | App Version | Repository |
|-------|-------------|---------|-------------|------------|
| kdl-server | Konstellation Development Lifecycle Server | 6.0.1 | 1.35.0 | [kdl-server](charts/kdl-server) |

## Usage

Charts are available in:

* [Chart Repository](https://helm.sh/docs/topics/chart_repository/)

### Chart Repository

#### Add repository

```console
helm repo add konstellation-io https://charts.konstellation.io
helm repo update
```

#### Search for available charts

```console
helm search repo konstellation-io
```

#### Install Helm chart

```console
helm install [RELEASE_NAME] konstellation-io/<helm-chart-name>
```

This install all the Kubernetes components associated with the chart and creates the release.

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._
