# kdl-server

A Helm chart to deploy KDL server

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| ialejandro | <ivan.alejandro@intelygenz.com> |  |
| alpiquero | <angelluis.piquero@intelygenz.com> |  |
| danielchg | <daniel.chavero@intelygenz.com> |  |

## Prerequisites

* Helm 3+
* Kubernetes 1.26+

## Compatibility matrix

| Release ↓ / Kubernetes → | 1.24 | 1.25 | 1.26 | 1.27 | 1.28 | 1.29 | 1.30 | 1.31 |
|:------------------------:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
| 6.0.2                    | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |
| 6.1.0                    | ❌   | ❌   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |
| 6.2.0                    | ❌   | ❌   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |

| Release ↓ / kdl-app → | 1.38.0 | 1.38.1 | 1.39.0 | 1.40.0 |
|:---------------------:|:------:|:------:|:------:|:------:|
| 6.0.2                 | ✅     | ✅     | ❌     | ❌     |
| 6.1.0                 | ❌     | ❌     | ✅     | ❌     |
| 6.2.0                 | ❌     | ❌     | ❌     | ✅     |

| Release ↓ / project-operator → | 0.19.0 | 0.20.0 | 0.21.0 |
|:------------------------------:|:------:|:------:|:------:|
| 6.0.2                          | ✅     | ❌     | ❌     |
| 6.1.0                          | ❌     | ✅     | ❌     |
| 6.2.0                          | ❌     | ❌     | ✅     |

| Release ↓ / user-tools-operator → | 0.30.0 | 0.31.0 | 0.32.0 |
|:---------------------------------:|:------:|:------:|:------:|
| 6.0.2                             | ✅     | ❌     | ❌     |
| 6.1.0                             | ❌     | ✅     | ❌     |
| 6.2.0                             | ❌     | ❌     | ✅     |

| Symbol | Description |
|:------:|-------------|
| ✅     | Perfect match: all features are supported. Client and server versions have exactly the same features/APIs. |
| 🟠     | Forward compatibility: the client will work with the server, but not all new server features are supported. The server has features that the client library cannot use. |
| ❌     | Backward compatibility/Not applicable: the client has features that may not be present in the server. Common features will work, but some client APIs might not be available in the server. |
| -      | Not tested: this combination has not been verified or is not applicable. |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.min.io | minio | 3.2.0 |
| oci://ghcr.io/konstellation-io/helm-charts | keycloak(konstellation-base) | 1.1.2 |
| oci://ghcr.io/oauth2-proxy/charts | oauth2proxy(oauth2-proxy) | 7.7.28 |
| oci://registry-1.docker.io/bitnamicharts | mongodb | 16.2.1 |
| oci://registry-1.docker.io/bitnamicharts | postgresql | 15.5.38 |

## Add repository

```console
helm repo add konstellation-io https://charts.konstellation.io
helm repo update
```

## Install Helm chart (repository mode)

```console
helm install [RELEASE_NAME] konstellation-io/kdl-server
```

This install all the Kubernetes components associated with the chart and creates the release.

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Install Helm chart (OCI mode)

Charts are also available in OCI format. The list of available charts can be found [here](https://github.com/konstellation-io/helm-charts/pkgs/container/helm-charts%2Fkdl-server).

```console
helm install [RELEASE_NAME] oci://ghcr.io/konstellation-io/helm-charts/kdl-server --version=[version]
```

## Uninstall Helm chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

> [!IMPORTANT]
> Upgrading an existing Release to a new major version (`v0.15.X` -> `v1.0.0`) indicates that there is an incompatible **BREAKING CHANGES** needing manual actions.

### From `6.0.2` to `6.1.0`

**global**

The new `global` section consolidates commonly shared configurations across all components. Key additions include:

* `imageRegistry`: default container image registry for all components
* `imagePullSecrets`: enables authentication for pulling images from private registries
* `env`: Global environment variables applicable to all components
* `envFromSecrets` and `envFromConfigMap`: allow defining environment variables from Kubernetes Secrets and ConfigMaps
* `envFromFiles`: adds support for loading environment variables from external files, enhancing flexibility in environment management

**kdl-server**

* **Reorganization**
  * `resources`, `image`, `env`, `ingress`, `service`, and `persistentVolume` configurations have been moved to root values

* **Additions**
  * `nameOverride` and `fullnameOverride`: allow overriding naming conventions for components
  * `autoscaling`: support for horizontal pod autoscaler with configurable CPU and memory thresholds, minimum and maximum replicas
  * `pdb` (Pod Disruption Budget): configurable to ensure high availability during voluntary disruptions
  * `volumeMounts`: support for attaching custom volume mounts to containers
  * `podSecurityContext`: defines pod-level security settings, such as `fsGroup`
  * `securityContext`: configurable container-level security options, such as dropping capabilities or running as a non-root user
  * `livenessProbe`, `readinessProbe`, and `startupProbe`: added for container lifecycle management
  * `extraContainers` and `initContainers`: allow additional functionality and custom initialization processes
  * `serviceAccount`: support for annotations, custom names, and API credential management
  * `networkPolicy`: configurable global ingress and egress policies with support for IP blocks, namespaces, and pod selectors
  * `terminationGracePeriodSeconds`: configurable termination grace period for pods across all components

* **Deprecations**
  * Legacy MongoDB connection string configuration has been deprecated
  * Simplified ingress annotations under global configurations

**cleaner**

* **Additions**
  * `schedule`: configurable cronjob schedule for cleaning up old files
  * `trashPath`: allows specifying the path to be cleaned
  * `threshold`: defines the minimum file age before cleanup
  * Resource limits and requests: support for defining CPU and memory usage
  * Have been removed on future releases

**knowledgeGalaxy**

* **Additions**
  * `imagePullSecrets`: support for pulling images from private registries
  * Enhanced environment variable management with `envFromFiles` and `envFromSecrets`
  * `livenessProbe` and `readinessProbe`: added for improved health monitoring
  * `networkPolicy`: support for ingress and egress controls
  * Customization of `serviceAccount` with annotations and name overrides

**userToolsOperator**

* **Additions**
  * Resource configuration for CPU and memory limits
  * `env` and `envFromFiles`: enhanced environment variable management
  * `livenessProbe` and `readinessProbe`: support for container health checks
  * `networkPolicy`: added for more secure communication
  * `serviceAccount`: customization options for API credential management

**projectOperator**

* **Additions**
  * Resource limits and requests for components
  * `serviceMonitor` integration for Prometheus monitoring
  * `affinity`, `tolerations`, and `nodeSelector` support for pod scheduling
  * Lifecycle hooks for managing pod startup and termination processes

**gitea**

* **Deprecations**
  * Legacy ingress and secret configurations have been removed
* **Additions**
  * Improved resource management for Gitea pods
  * Enhanced `networkPolicy` for better control of ingress and egress

**keycloak**

* **Additions**
  * Based on `konstellation-io/konstellation-base` chart
  * `fullnameOverride`: support for custom naming conventions
  * `imagePullSecrets`: added for private image registries
  * Enhanced handling of environment variables through `envFromFiles` and `envFromSecrets`
  * Persistent volume support with flexible options for storage classes and access modes

**minio**

* **Deprecations**
  * Legacy ingress configurations have been deprecated
* **Additions**
  * Change dependecy to `bitnami/minio` chart
  * Enhanced volume configurations for improved persistence
  * Added support for `networkPolicy` to control access

**mongodb**

* **Deprecations**
  * Legacy connection string configurations are now deprecated
* **Additions**
  * Change dependecy to `bitnami/mongodb` chart
  * Improved secret-based MongoDB connection string management
  * Enhanced integration with shared volumes for persistence

**oauth2-proxy**

* **New Features**
  * Change dependecy to `oauth2-proxy/oauth2-proxy` chart
  * Introduced a new, centralized `oauth2` configuration section to replace legacy configurations
  * `clientID` and `clientSecret` settings added for more secure integration with OAuth2 providers
  * Support for multiple OAuth2 providers with distinct configurations
  * `redirectURIs` and `scopes` now configurable at a granular level
  * Enhanced token validation and refresh capabilities with support for advanced OAuth2 flows
  * Added support for `openid` integration, improving compatibility with identity providers
* **Legacy oauth2Proxy**
  * Legacy OAuth2 configurations have been deprecated
  * Removed hardcoded `clientID` and `clientSecret` options in favor of more flexible secret-based configurations
  * Updated callback and redirect URIs to adhere to modern OAuth2 specifications

**postgresql**

* **Additions**
  * Change dependecy to `bitnami/postgresql` chart
  * Introduced configuration for PostgreSQL integration to Keycloak
  * Support for environment variable customization specific to PostgreSQL
  * Enhanced persistent volume support for PostgreSQL data storage
  * Added compatibility with `serviceAccount` for PostgreSQL pods
  * `securityContext` and `podSecurityContext` configurations added for PostgreSQL security

**sharedVolume**

* **Additions**
  * Support for shared persistent volumes with options for access modes and storage classes
  * Label-based volume bindings for using pre-provisioned volumes

**deprecated features**

* **Drone**
  * Removed `drone`, `droneAuthorizer`, and `droneRunner` configurations
* **Legacy MongoDB**
  * Connection strings have been deprecated in favor of secret-based management
* **Legacy oauth2-Proxy**
  * Replaced with new `oauth2-proxy` configurations
* **MinIO legacy Configurations**
  * Deprecated older ingress and volume configurations

Execute the following actions to update the CRDs before applying the upgrade.

* Run the following script to update CRDs:

  ```bash
  kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v6.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
  ```

CHANGELOG: [6.0.2](https://github.com/konstellation-io/helm-charts/releases/tag/kdl-server-6.1.0)

### From `6.0.1` to `6.0.2`

Changes in values:

* `konstellation/mlflow` -> `konstellation/kdl-mlflow`
* `konstellation/repo-cloner` -> `konstellation/kdl-repo-cloner`
* `konstellation/vscode` -> `konstellation/kdl-vscode`
* `konstellation/project-operator` -> `konstellation/kdl-project-operator`
* `konstellation/gitea-oauth2-setup` -> `konstellation/kdl-gitea-oauth2-setup`
* `konstellation/drone-authorizer` -> `konstellation/kdl-drone-authorizer`
* `konstellation/cleaner` -> `konstellation/kdl-cleaner`

Bump versions:

* `konstellation/kdl-server`: from `1.35.0` -> `1.38.0`
* `konstellation/kdl-repo-cloner`: from `0.15.0` -> `0.18.0`

CHANGELOG: [6.0.2](https://github.com/konstellation-io/helm-charts/releases/tag/kdl-server-6.0.2)

### From `5.X` to `6.X`

New requirements:

* An existing running MongoDB database must be accessible as internal MongoDB database has been removed. Check [MongoDB](#mongodb)

Changes in values:

* `mongodb` has been removed
* `global.mongodb.connectionString.uri`, `global.mongodb.connectionString.secretName` and `global.mongodb.connectionString.secretKey`

Execute the following actions to update the CRDs before applying the upgrade.

* Remove all `UserTools` resources from your cluster.
* Run the following script to update CRDs:

  ```bash
  kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v6.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
  kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v6.0.0/helm/kdl-server/crds/project-operator-crd.yaml
  ```

Existing `KDLProject` resources won't be updated with the new fields after upgrading the chart.

Run the following script to patch all existing `KDLProject` resources:

```bash
#!/bin/bash
NAMESPACE=<release_namespace>
RELEASE_NAME=<release_name>

cat << EOF > patch-file.yaml
spec:
  minio:
    endpointURL: http://${RELEASE_NAME}-minio:9000
EOF

for project in $(kubectl -n ${NAMESPACE} get kdlprojects.project.konstellation.io -o name); do kubectl -n ${NAMESPACE} patch ${project} --type merge --patch-file patch-file.yaml; done
rm -f patch-file.yaml
```

### From `4.X` to `5.X`

Changes in values:

* `domain` moved to `global.domain`
* `serverName` moved to `global.serverName`
* `tls` moved to `global.ingress.tls`
* `cert-manager` has been removed. TLS certificate management has been lend to users own management. Use `global.tls.secretName` or `<component>.tls.secretName` to set the secret name containing the TLS certificate.
* Added `projectOperator.mlflow.ingress.annotations`, `projectOperator.mlflow.ingress.className`
* Added `drone.ingress.tls`, `gitea.ingress.tls`, `kdlServer.ingress.tls`, `minio.ingress.tls`, `minio.consoleIngress.tls`, `projectOperator.mlflow.ingress.tls` and `userToolsOperator.ingress.tls` for individual tls config. These values take precedence over `global.ingress.tls`.
* `drone.ingress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `drone.ingress.className: "nginx"`
* `gitea.ingress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `gitea.ingress.className: "nginx"`
* `kdlServer.ingress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `kdlServer.ingress.className: "nginx"`
* `minio.ingress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `minio.ingress.className: "nginx"`
* `minio.consoleIngress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `minio.consoleIngress.className: "nginx"`
* `userToolsOperator.ingress.annotations.kubernetes.io/ingress.class: "nginx"` removed in favour of `userToolsOperator.ingress.className: "nginx"`

Run these commands to update the CRDs before applying the upgrade.

```bash
kubectl apply --server-side --force-conflicts -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v5.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
kubectl apply --server-side --force-conflicts -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v5.0.0/helm/kdl-server/crds/project-operator-crd.yaml
```

Existing `KDLProject` resources won't be updated with the new fields after upgrading the chart.

Run the following script to patch all existing `KDLProject` resources:

```bash
#!/bin/bash
NAMESPACE=<release_namespace>
INGRESS_CLASS=<ingress_class_name>
TLS_ENABLED=<true|false>
TLS_SECRET_NAME=<secret_name>

cat << EOF > patch-file.yaml
spec:
  mlflow:
    ingress:
      className: "${INGRESS_CLASS}"
      tls:
        enabled: ${TLS_ENABLED}
        secretName: "${TLS_SECRET_NAME}"
      annotations:
        # place your custom annotations here
EOF

for project in $(kubectl -n ${NAMESPACE} get kdlprojects.project.konstellation.io -o name); do kubectl -n ${NAMESPACE} patch ${project} --type merge --patch-file patch-file.yaml; done
rm -f patch-file.yaml
```

### From `3.X` to `4.X`

This major version comes with the following breaking changes:

* Fixed an issue with **usertools.kdl.konstellation.io** CRD that produced errors in **user-tools-operator** with *UserTools* resources during the reconciling process.
* Added `minio.consoleIngress.annotations` to *values.yaml*

Run these commands to update the CRDs before applying the upgrade.

```bash
kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v4.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
```

### From `2.X` to `3.X`

This major version comes with the following breaking changes:

* Ingress configuration changed from *values.yaml*:
  * removed `ingress.type`
  * added `drone.ingress.annotations`, `kdlApp.ingress.annotations`, `gitea.ingress.annotations`, `minio.ingress.annotations`, `userToolsOperator.ingress.annotations`

* Upgrade user-tools-operator to `v0.20.0`.
  * TLS secret name and Ingress annotations are now received from the operator values

* Upgrade app to `1.17.0`:
  * pass the name of the TLS secret and Ingress annotations through `userTools` resources.
  * pass Ingress annotations through `userTools`.

Run these commands to update the CRDs before applying the upgrade.

```bash
kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v3.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
```

### From `1.X` to `2.X`

This major version comes with the following breaking changes:

* This upgrades user-tools-operator to `v0.17.0`.
  * users service accounts are now managed by `kdlServer` instead the `user-tools-operator`

Run these commands to update the CRDs before applying the upgrade.

```bash
kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v2.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
```

### From `0.X` to `1.X`

This major version comes with the following breaking changes:

* UserTools CRD metadata changes:
  * `metadata.name` changed to `usertools.kdl.konstellation.io`
  * `spec.groups` changed to `kdl.konstellation.io`
* KDL Runtimes support

Run these commands to update the CRDs before applying the upgrade:

```bash
kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v1.0.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
```

## Configuration

See [Customizing the chart before installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with comments:

```console
helm show values konstellation-io/kdl-server
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| env | object | `{}` | Environment variables to configure application </br> Ref: https://github.com/konstellation-io/kdl-server/tree/main/app/api |
| envFromConfigMap | object | `{}` | Variables from configMap |
| envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| envFromSecrets | object | `{}` | Variables from secrets |
| extraContainers | list | `[]` | Configure extra containers |
| fullnameOverride | string | `""` | String to fully override kdl-server.fullname template |
| global.domain | string | `"kdl.local"` | The DNS domain name that will serve the application |
| global.enableHttps | bool | `true` | Enable HTTPs Use to enable or disable HTTPS on the endpoints |
| global.env | object | `{}` | Environment variables to configure application |
| global.envFromConfigMap | object | `{}` | Variables from configMap |
| global.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| global.envFromSecrets | object | `{}` | Variables from secrets |
| global.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| global.imageRegistry | string | `""` | Specifies the registry to pull images from. Leave empty for the default registry |
| global.serverName | string | `"local-server"` | KDL Server instance name |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-server","tag":"1.40.0"}` | Image registry The image configuration for the base service |
| imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| ingress | object | `{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| keycloak | object | `{"command":[],"enabled":true,"env":{},"fullnameOverride":"keycloak","image":{"repository":"keycloak/keycloak","tag":"26.0"},"ingress":{"annotations":{},"className":"","enabled":true,"hosts":[{"host":"keycloak.mydomain.com","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]},"livenessProbe":{"enabled":true},"readinessProbe":{"enabled":true,"httpGet":{"path":"/realms/master"}},"service":{"healthPath":"/realms/master","targetPort":8080},"serviceAccount":{"create":true}}` | Keycloak subchart deployment </br> Ref: https://github.com/konstellation-io/helm-charts/blob/kdl-server-1.0.2/charts/kdl-server/values.yaml |
| keycloak.enabled | bool | `true` | Enable or disable Keycloak subchart |
| knowledgeGalaxy | object | `{"affinity":{},"args":[],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"config":{"descriptionMinWords":50,"logLevel":"INFO","numberOfOutputs":1000,"workers":1},"enabled":false,"env":{},"envFromConfigMap":{},"envFromFiles":[],"envFromSecrets":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"},"imagePullSecrets":[],"initContainers":[],"lifecycle":{},"livenessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"livenessProbeCustom":{},"networkPolicy":{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"readinessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"readinessProbeCustom":{},"resources":{},"secrets":[],"securityContext":{},"service":{"port":80,"targetPort":8080,"type":"ClusterIP"},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"startupProbe":{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"startupProbeCustom":{},"terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[],"volumeMounts":[],"volumes":[]}` | knowledge-galaxy deployment </br> Ref: https://github.com/konstellation-io/knowledge-galaxy |
| knowledgeGalaxy.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| knowledgeGalaxy.args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| knowledgeGalaxy.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| knowledgeGalaxy.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| knowledgeGalaxy.config | object | `{"descriptionMinWords":50,"logLevel":"INFO","numberOfOutputs":1000,"workers":1}` | Configuration TODO: legacy backard compatibility, remove in future versions |
| knowledgeGalaxy.config.descriptionMinWords | int | `50` | Minimum number of words to use for project description |
| knowledgeGalaxy.config.logLevel | string | `"INFO"` | Log level |
| knowledgeGalaxy.config.numberOfOutputs | int | `1000` | Number of outputs that the recommender returns |
| knowledgeGalaxy.config.workers | int | `1` | Number of threads for the server |
| knowledgeGalaxy.enabled | bool | `false` | Whether to enable Knowledge Galaxy |
| knowledgeGalaxy.env | object | `{}` | Environment variables to configure application Ref: https://github.com/konstellation-io/knowledge-galaxy?tab=readme-ov-file#environment-variables |
| knowledgeGalaxy.envFromConfigMap | object | `{}` | Variables from configMap |
| knowledgeGalaxy.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| knowledgeGalaxy.envFromSecrets | object | `{}` | Variables from secrets |
| knowledgeGalaxy.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"}` | Image registry The image configuration for the base service |
| knowledgeGalaxy.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| knowledgeGalaxy.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| knowledgeGalaxy.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| knowledgeGalaxy.livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| knowledgeGalaxy.livenessProbeCustom | object | `{}` | Custom livenessProbe |
| knowledgeGalaxy.networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| knowledgeGalaxy.networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| knowledgeGalaxy.networkPolicy.policyTypes | list | `[]` | Policy types |
| knowledgeGalaxy.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| knowledgeGalaxy.podAnnotations | object | `{}` | Configure annotations on Pods |
| knowledgeGalaxy.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| knowledgeGalaxy.podLabels | object | `{}` | Configure labels on Pods |
| knowledgeGalaxy.podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| knowledgeGalaxy.readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| knowledgeGalaxy.readinessProbeCustom | object | `{}` | Custom readinessProbe |
| knowledgeGalaxy.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| knowledgeGalaxy.secrets | list | `[]` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| knowledgeGalaxy.securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| knowledgeGalaxy.service | object | `{"port":80,"targetPort":8080,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| knowledgeGalaxy.service.port | int | `80` | Kubernetes Service port |
| knowledgeGalaxy.service.targetPort | int | `8080` | Pod expose port |
| knowledgeGalaxy.service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| knowledgeGalaxy.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| knowledgeGalaxy.startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| knowledgeGalaxy.startupProbeCustom | object | `{}` | Custom startupProbe |
| knowledgeGalaxy.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| knowledgeGalaxy.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| knowledgeGalaxy.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| knowledgeGalaxy.volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| knowledgeGalaxy.volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
| lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| livenessProbeCustom | object | `{}` | Custom livenessProbe |
| minio | object | `{"enabled":true,"mode":"standalone","persistence":{"enabled":false},"rootPassword":"ChangeMe","rootUser":"ChangeMe"}` | MinIO subchart deployment </br> Ref: https://github.com/minio/minio/blob/RELEASE.2021-10-13T00-23-17Z/helm/minio/values.yaml TODO: pending to remove legacy minio |
| minio.enabled | bool | `true` | Enable or disable MinIO subchart |
| mongodb | object | `{"architecture":"standalone","auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":false,"persistence":{"enabled":false}}` | MongoDB subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml |
| mongodb.enabled | bool | `false` | Enable or disable MongoDB subchart |
| nameOverride | string | `""` | String to partially override kdl-server.fullname template (will maintain the release name) |
| networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| networkPolicy.policyTypes | list | `[]` | Policy types |
| nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| oauth2proxy | object | `{"clientID":"XXXXXXX","clientSecret":"XXXXXXXX","cookieName":"","cookieSecret":"XXXXXXXXXXXXXXXX","enabled":false,"extraContainers":[],"extraObjects":[],"extraVolumeMounts":[],"extraVolumes":[],"httpScheme":"http"}` | OAuth2-Proxy subchart deployment </br> Ref: https://github.com/oauth2-proxy/manifests/blob/main/helm/oauth2-proxy/values.yaml |
| oauth2proxy.enabled | bool | `false` | Enable or disable OAuth2-Proxy subchart |
| persistentVolume | object | `{"accessModes":["ReadWriteOnce"],"annotations":{},"enabled":false,"labels":{},"selector":{},"size":"8Gi","storageClass":"","volumeBindingMode":"","volumeName":""}` | Persistent Volume configuration </br> Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/ |
| persistentVolume.accessModes | list | `["ReadWriteOnce"]` | Persistent Volume access modes Must match those of existing PV or dynamic provisioner </br> Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/ |
| persistentVolume.annotations | object | `{}` | Persistent Volume annotations |
| persistentVolume.enabled | bool | `false` | Enable or disable persistence |
| persistentVolume.labels | object | `{}` | Persistent Volume labels |
| persistentVolume.selector | object | `{}` | Persistent Volume Claim Selector Useful if Persistent Volumes have been provisioned in advance </br> Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#selector |
| persistentVolume.size | string | `"8Gi"` | Persistent Volume size |
| persistentVolume.storageClass | string | `""` | Persistent Volume Storage Class If defined, storageClassName: <storageClass> If set to "-", storageClassName: "", which disables dynamic provisioning If undefined (the default) or set to null, no storageClassName spec is   set, choosing the default provisioner.  (gp2 on AWS, standard on   GKE, AWS & OpenStack) |
| persistentVolume.volumeBindingMode | string | `""` | Persistent Volume Binding Mode If defined, volumeBindingMode: <volumeBindingMode> If undefined (the default) or set to null, no volumeBindingMode spec is set, choosing the default mode. |
| persistentVolume.volumeName | string | `""` | Persistent Volume Name Useful if Persistent Volumes have been provisioned in advance and you want to use a specific one |
| podAnnotations | object | `{}` | Configure annotations on Pods |
| podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| podLabels | object | `{}` | Configure labels on Pods |
| podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| postgresql | object | `{"auth":{"database":"kdl","password":"ChangeMe","username":"user"},"enabled":true,"primary":{"persistence":{"enabled":false}},"replicaCount":1}` | PostgreSQL subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/postgresql/values.yaml |
| postgresql.enabled | bool | `true` | Enable or disable PostgreSQL subchart |
| projectOperator | object | `{"affinity":{},"args":["--enable-http2","--health-probe-bind-address=:8081","--leader-elect","--leader-election-id=project-operator","--metrics-bind-address=:8080","--zap-log-level=error","--zap-stacktrace-level=error"],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"enabled":true,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-project-operator","tag":"0.21.0"},"imagePullSecrets":[],"initContainers":[],"lifecycle":{},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"resources":{},"securityContext":{"allowPrivilegeEscalation":false,"runAsNonRoot":true},"service":{"extraPorts":[{"name":"metrics","port":8080,"targetPort":8080}],"port":8081,"targetPort":8081,"type":"ClusterIP"},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"serviceMonitor":{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"},"templateCustomResource":"apiVersion: kdl.konstellation.io/v1\nkind: KDLProject\nmetadata:\n  name: replaced-by-kdl-api\n  namespace: replaced-by-kdl-api\nspec:\n  projectId: replaced-by-kdl-api\n  domain: kdl.local\n\n  mlflow:\n    image:\n      repository: konstellation/kdl-mlflow\n      pullPolicy: IfNotPresent\n      tag: \"0.15.0\"\n\n    ingress:\n      enabled: true\n      tls:\n        enabled: true\n      className: \"replace-class-name\"\n      annotations:\n        cert-manager.io/cluster-issuer: selfsigned-clusterissuer\n        nginx.ingress.kubernetes.io/proxy-body-size: \"4096m\"\n\n    persistentVolume:\n      enabled: true\n      size: \"1Gi\"\n      accessModes:\n        - ReadWriteOnce\n      storageClass: \"replace-storage-class-name\"\n\n    env:\n      AWS_ACCESS_KEY_ID: \"replace-minio-access-key\"\n      AWS_SECRET_ACCESS_KEY: \"replace-minio-secret-access-key\"\n      MLFLOW_S3_ENDPOINT_URL: \"http://minio.kdl:9000\"\n      MLFLOW_TRACKING_URI: \"sqlite:///database.db\"\n\n  filebrowser:\n    image:\n      repository: konstellation/kdl-filebrowser\n      pullPolicy: IfNotPresent\n      tag: \"1.0.0\"\n\n    podSecurityContext:\n      fsGroup: 1000\n\n    securityContext:\n      privileged: true\n      capabilities:\n        add:\n          - SYS_ADMIN\n\n    env:\n      AWS_S3_ACCESS_KEY_ID: \"replace-minio-access-key\"\n      AWS_S3_MOUNT: \"/srv\"\n      AWS_S3_SECRET_ACCESS_KEY: \"replace-minio-secret-access-key\"\n      AWS_S3_URL: \"http://minio:9000\"\n      FB_ADDRESS: \"0.0.0.0\"\n      FB_DATABASE: \"/home/filebrowser/database.db\"\n      FB_LOG: stdout\n      FB_ROOT: /srv\n      S3FS_ARGS: >-\n        -o use_path_request_style\n        -o use_cache=/cache\n        -o ensure_diskfree=2048\n        -o max_stat_cache_size=100000\n        -o stat_cache_expire=300\n        -o enable_noobj_cache\n        -o dbglevel=warn\n        -o multipart_size=52\n        -o parallel_count=32\n        -o max_dirty_data=512\n        -o multireq_max=30\n        -o complement_stat\n        -o notsup_compat_dir\n        -o enable_content_md5\n        -o ro\n\n    volumes:\n      - name: cache-volume\n        emptyDir: {}\n      - name: fuse-device\n        hostPath:\n          path: /dev/fuse\n          type: CharDevice\n\n    volumeMounts:\n      - name: cache-volume\n        mountPath: /cache\n      - name: fuse-device\n        mountPath: /dev/fuse\n","terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[]}` | project-operator operator |
| projectOperator.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| projectOperator.args | list | `["--enable-http2","--health-probe-bind-address=:8081","--leader-elect","--leader-election-id=project-operator","--metrics-bind-address=:8080","--zap-log-level=error","--zap-stacktrace-level=error"]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| projectOperator.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| projectOperator.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| projectOperator.enabled | bool | `true` | Enable or disable project-operator |
| projectOperator.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-project-operator","tag":"0.21.0"}` | Image registry The image configuration for the base service |
| projectOperator.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| projectOperator.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| projectOperator.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| projectOperator.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| projectOperator.podAnnotations | object | `{}` | Configure annotations on Pods |
| projectOperator.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| projectOperator.podLabels | object | `{}` | Configure labels on Pods |
| projectOperator.podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| projectOperator.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| projectOperator.securityContext | object | `{"allowPrivilegeEscalation":false,"runAsNonRoot":true}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| projectOperator.service | object | `{"extraPorts":[{"name":"metrics","port":8080,"targetPort":8080}],"port":8081,"targetPort":8081,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| projectOperator.service.extraPorts | list | `[{"name":"metrics","port":8080,"targetPort":8080}]` | Pod extra ports |
| projectOperator.service.port | int | `8081` | Kubernetes Service port |
| projectOperator.service.targetPort | int | `8081` | Pod expose port |
| projectOperator.service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| projectOperator.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| projectOperator.serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| projectOperator.serviceMonitor.enabled | bool | `false` | Enable or disable |
| projectOperator.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| projectOperator.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| projectOperator.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| rbac | object | `{"create":true}` | Creation of resources RBAC </br> Ref: https://kubernetes.io/docs/reference/access-authn-authz/rbac/ |
| readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| readinessProbeCustom | object | `{}` | Custom readinessProbe |
| readyChecker | object | `{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5}` | Check if dependencies are ready |
| readyChecker.enabled | bool | `false` | Enable or disable ready-checker |
| readyChecker.pullPolicy | string | `"IfNotPresent"` | Pull policy for the image |
| readyChecker.repository | string | `"busybox"` | Repository of the image |
| readyChecker.retries | int | `30` | Number of retries before giving up |
| readyChecker.services | list | `[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}]` | List services |
| readyChecker.tag | string | `"latest"` | Overrides the image tag |
| readyChecker.timeout | int | `5` | Timeout for each check |
| resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| secrets | list | `[]` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| service | object | `{"port":80,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| service.port | int | `80` | Kubernetes Service port |
| service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| serviceMonitor.enabled | bool | `false` | Enable or disable |
| sharedVolume | object | `{"accessModes":["ReadWriteMany"],"annotations":{},"enabled":false,"labels":{},"selector":{},"size":"10Gi","storageClass":"","volumeBindingMode":"","volumeName":""}` | Shared Volume configuration Mount volume to share data between components </br> Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/ |
| sharedVolume.accessModes | list | `["ReadWriteMany"]` | Persistent Volume access modes Must match those of existing PV or dynamic provisioner </br> Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/ |
| sharedVolume.annotations | object | `{}` | Persistent Volume annotations |
| sharedVolume.enabled | bool | `false` | Enable or disable persistence |
| sharedVolume.labels | object | `{}` | Persistent Volume labels |
| sharedVolume.selector | object | `{}` | Persistent Volume Claim Selector Useful if Persistent Volumes have been provisioned in advance </br> Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#selector |
| sharedVolume.size | string | `"10Gi"` | Persistent Volume size |
| sharedVolume.storageClass | string | `""` | Persistent Volume Storage Class If defined, storageClassName: <storageClass> If set to "-", storageClassName: "", which disables dynamic provisioning If undefined (the default) or set to null, no storageClassName spec is   set, choosing the default provisioner.  (gp2 on AWS, standard on   GKE, AWS & OpenStack) |
| sharedVolume.volumeBindingMode | string | `""` | Persistent Volume Binding Mode If defined, volumeBindingMode: <volumeBindingMode> If undefined (the default) or set to null, no volumeBindingMode spec is set, choosing the default mode. |
| sharedVolume.volumeName | string | `""` | Persistent Volume Name Useful if Persistent Volumes have been provisioned in advance and you want to use a specific one |
| startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| startupProbeCustom | object | `{}` | Custom startupProbe |
| terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| testConnection | object | `{"enabled":false,"repository":"busybox","tag":"latest"}` | Enable or disable test connection |
| tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| userToolsOperator | object | `{"affinity":{},"args":["--enable-http2","--health-probe-bind-address=:8081","--leader-elect","--leader-election-id=user-tools-operator","--metrics-bind-address=:8080","--zap-log-level=error","--zap-stacktrace-level=error"],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"enabled":true,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-user-tools-operator","tag":"0.32.0"},"imagePullSecrets":[],"initContainers":[],"lifecycle":{},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"resources":{},"securityContext":{},"service":{"extraPorts":[{"name":"metrics","port":8080,"targetPort":8080}],"port":8081,"targetPort":8081,"type":"ClusterIP"},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"serviceMonitor":{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"},"templateCustomResource":"apiVersion: kdl.konstellation.io/v1\nkind: KDLUserTools\nmetadata:\n  name: replaced-by-kdl-api\n  namespace: replaced-by-kdl-api\nspec:\n  username: replaced-by-kdl-api\n  usernameSlug: replaced-by-kdl-api\n\n  sharedVolume:\n    enabled: true\n    name: \"\"\n\n  # must be declared\n  podLabels: {}\n\n  persistentVolume:\n    enabled: true\n    size: \"1Gi\"\n    accessModes:\n      - ReadWriteOnce\n    storageClass: \"\"\n\n  kubeconfig:\n    enabled: false\n    # externalServerUrl: https://192.168.0.21:16443\n\n  repoCloner:\n    image:\n      repository: konstellation/kdl-repo-cloner\n      pullPolicy: IfNotPresent\n      tag: \"0.19.0\"\n\n    # ref: https://github.com/konstellation-io/kdl-server/tree/main/repo-cloner\n    env: {}\n      # KDL_SERVER_MONGODB_URI: \"mongodb://kdl:123456@mongodb-svc.mongodb:27017/kdl?authSource=kdl&authMechanism=SCRAM-SHA-256\"\n\n  # must be declared\n  vscodeRuntime:\n    image:\n      repository: konstellation/kdl-py\n      pullPolicy: IfNotPresent\n      tag: \"3.9\"\n\n    env: {}\n      # MINIO_ACCESS_KEY: \"\"\n      # MINIO_SECRET_KEY: \"\"\n      # MINIO_ENDPOINT: \"http://minio:9000\"\n\n  # must be declared\n  nodeSelector: {}\n  tolerations: []\n  affinity: {}\n","terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[]}` | User Tools Operator deployment ref: https://github.com/konstellation-io/kdl-server/tree/main/user-tools-operator |
| userToolsOperator.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| userToolsOperator.args | list | `["--enable-http2","--health-probe-bind-address=:8081","--leader-elect","--leader-election-id=user-tools-operator","--metrics-bind-address=:8080","--zap-log-level=error","--zap-stacktrace-level=error"]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| userToolsOperator.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| userToolsOperator.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| userToolsOperator.enabled | bool | `true` | Enable or disable User Tools Operator deployment |
| userToolsOperator.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-user-tools-operator","tag":"0.32.0"}` | Image registry The image configuration for the base service |
| userToolsOperator.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| userToolsOperator.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| userToolsOperator.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| userToolsOperator.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| userToolsOperator.podAnnotations | object | `{}` | Configure annotations on Pods |
| userToolsOperator.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| userToolsOperator.podLabels | object | `{}` | Configure labels on Pods |
| userToolsOperator.podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| userToolsOperator.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| userToolsOperator.securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| userToolsOperator.service | object | `{"extraPorts":[{"name":"metrics","port":8080,"targetPort":8080}],"port":8081,"targetPort":8081,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| userToolsOperator.service.extraPorts | list | `[{"name":"metrics","port":8080,"targetPort":8080}]` | Pod extra ports |
| userToolsOperator.service.port | int | `8081` | Kubernetes Service port |
| userToolsOperator.service.targetPort | int | `8081` | Pod expose port |
| userToolsOperator.service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| userToolsOperator.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| userToolsOperator.serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| userToolsOperator.serviceMonitor.enabled | bool | `false` | Enable or disable |
| userToolsOperator.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| userToolsOperator.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| userToolsOperator.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
