# helm-charts

This repository contains Helm charts for deploying Konstellation services.

| Chart | Description | Version | App Version |
|-------|-------------|---------|-------------|
| [kdl-server](charts/kdl-server) | Konstellation Development Lifecycle Server | 6.0.2 | 1.38.0 |
| [konstellation-base](charts/konstellation-base) | Konstellation template Helm chart | 1.1.2 | 1.0.0 |

## Usage

Charts are available in:

* [Chart Repository](https://helm.sh/docs/topics/chart_repository/)
* [OCI Artifacts](https://helm.sh/docs/topics/registries/)

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

### OCI Registry

Charts are also available in OCI format. The list of available charts can be found [here](https://github.com/orgs/konstellation-io/packages?repo_name=helm-charts).

#### Install Helm chart

```console
helm install [RELEASE_NAME] oci://ghcr.io/konstellation-io/helm-charts/<helm-chart-name> --version=[version]
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
