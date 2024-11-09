# Add new Helm Chart

This document provides guidelines for adding a new Helm chart to the repository. By following these steps, you can ensure that the new chart is well-integrated with existing workflows.

## Prerequisites

Installed tools:

* [`Helm`](https://helm.sh/docs/intro/install/)
* [`kubectl`](https://kubernetes.io/docs/tasks/tools/)
* Kubernetes platform:
  * [`KinD`](https://kind.sigs.k8s.io)
  * [`Minikube`](https://minikube.sigs.k8s.io/docs/start/)
  * [`k3s`](https://k3s.io)

## Steps

### 1. Create basic chart structure

* Go to the `charts/` directory
* Use the `helm` to create the new chart:

  ```bash
  helm create my-chart
  ```

Output will a structure for your chart in `charts/my-chart`.

### 2. Add dependencies

* If the chart depends on other charts, add these dependencies in `Chart.yaml` under the `dependencies` section:

```yaml
dependencies:
- name: postgresql
  version: 15.5.38
  repository: "https://charts.bitnami.com/bitnami"
- name: redis
  version: 6.3.5
  repository: "https://charts.bitnami.com/bitnami"
```

* Run helm dependency update `charts/my-chart` to download the dependencies into the `charts/` directory.

### 3. Configure templates

* Add manifests (like `configmap.yaml` or `secret.yaml`) in a `templates/` subdirectory with properly parametrization
* Use `_helpers.tpl` for reusable template snippets like `labels` or `annotations`
* Include custom resource definitions (`CRDs`) in a `crds/` directory if required

### 4. Test same tools like GitHub Actions

* Chart will automatically be included in the [`[Helm Charts] Lint and Test PR workflow`](./../.github/workflows/helm-lint-test.yml) workflow which validates and tests all charts in the repository
* Use the [`ct tool`](https://github.com/helm/chart-testing) locally to verify the chart:

```bash
ct lint --config .github/ct.yml --lint-conf .github/helmlintconf.yml
```

* Releases wil be generate with [`[Helm Charts] Releases`](./../.github/workflows/helm-release.yml) your chart when changes are pushed to the `main` branch based on `Chart.yaml` version field

### 5. Add documentation

* Create `README.md.gotmpl` with the sections
* Use the [`helm-docs`](https://github.com/norwoodj/helm-docs) to auto-generate the final `README.md`:

  ```bash
  helm-docs .
  ```

* To add examples use `examples/` directory with example values

### 6. Testing

* Add integration tests in the `templates/tests/` directory
* Example:

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: test-connection
  spec:
    containers:
    - name: curl
      image: curlimages/curl
      args:
      - -f
      - http://my-service
  ```

* Use the `KinD` cluster setup and deploy your chart to test if is functional in a local Kubernetes cluster

## Additional tips

* Follow semantic versioning in `Chart.yaml` to distinguish between breaking changes, features and fix patches
* Reuse templates or values from existing charts (`kdl-server` or `konstellation-base`) to maintain consistency
* Testing locally with a local `KinD` (u other platform) cluster to catch issues early

  ```bash
  helm lint charts/my-chart
  kind create cluster
  helm install my-chart charts/my-chart
  ```

### Linting

* Use `helm lint` to detect common issues in your chart:

  ```bash
  helm lint charts/my-chart
  ```

* Run [`ct tool`](https://github.com/helm/chart-testing) locally to validate the chart:

  ```bash
  ct lint --config .github/ct.yml
  ```

* Include a `values.yaml` file with realistic default values to avoid empty or invalid configurations during linting

### Dependency Validation

* Verify that all dependencies:

  ```bash
  helm dependency update charts/my-chart
  ```

* Test with different dependency versions to ensure compatibility

### Template Rendering

* Use `helm template` command to render templates locally and inspect the output:

  ```bash
  helm template my-chart charts/my-chart
  ```

* Check for:
  * Syntax issues
  * Correct API versions for Kubernetes resources
  * Proper indentation

### Unit Testing

* Include unit tests for templates using tools like [helm-unittest](https://github.com/helm-unittest/helm-unittest):

  ```bash
  helm unittest charts/my-chart
  ```

* Create realistic test cases for different values in `values.yaml`:
  * Default configurations
  * Missing required parameters

### Integration testing

* Use a local Kubernetes cluster (for example: `KinD`) to deploy the chart and validate:

  ```bash
  kind create cluster
  helm install my-chart charts/my-chart
  kubectl get pods
  ```

* Validate that:
  * All resources are created successfully
  * `Pods` are running and healthy
  * `Services`, `ConfigMaps` and `Secrets` are properly configured

### Resource cleanup

* Verify that all resources created by the chart are properly cleaned up when uninstalled:

  ```bash
  helm uninstall my-chart
  kubectl get all
  ```

* Ensure no dangling resources (for example: `PersistentVolumeClaims`) remain

### Chart Versioning

* MAJOR for breaking changes
* MINOR for features
* PATCH for bug fixes
