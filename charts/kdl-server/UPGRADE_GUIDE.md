> [!IMPORTANT]
> Upgrading an existing Release to a new major version (`v0.15.X` -> `v1.0.0`) indicates that there is an incompatible **BREAKING CHANGES** needing manual actions.

### From `6.2.0` to `6.2.1`

* Remove `PersistentVolumeClaim` values from KDL server. Don't need.

### From `6.1.0` to `6.2.0`

> [!IMPORTANT]
> Execute the following actions to update the CRDs before applying the upgrade.
> ```bash
> kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v6.2.0/helm/kdl-server/crds/project-operator-crd.yaml
> kubectl apply --server-side -f https://raw.githubusercontent.com/konstellation-io/kdl-server/v6.2.0/helm/kdl-server/crds/user-tools-operator-crd.yaml
> ```

This release introduces several architectural improvements and updates to core components. The main changes include enhanced security configurations, streamlined HTTPS management, and updated component versions.

#### Breaking Changes

* **Global HTTPS Configuration**
  - The `global.ingress.tls.enabled` has been replaced with `global.enableHttps`
  - TLS secret configurations have been moved to individual component sections
  - All components now use the new global HTTPS configuration for consistent behavior

* **Component Version Updates**
  - KDL Server updated to version `1.40.0`
  - Project Operator updated to version `0.21.0`
  - User Tools Operator updated to version `0.32.0`
  - MLflow updated to version `0.15.0`
  - Filebrowser updated to version `1.0.0`

* **Filebrowser Changes**
  - Repository changed from `filebrowser/filebrowser` to `konstellation/kdl-filebrowser`
  - Added comprehensive S3 integration configurations
  - New security context settings for FUSE mount support

#### Deprecations

* **Cleaner Job**
  - The cleaner cronjob configuration has been removed
  - Users should migrate to alternative cleanup solutions

* **Legacy Helpers**
  - Removed deprecated TLS helper functions
  - Removed legacy global TLS configurations

#### Other Changes

* **Enhanced S3 Integration**
  - Added detailed S3 configuration for Filebrowser
  - Improved S3 mount options for better performance
  - Added cache configuration for S3 operations

* **Security Improvements**
  - Added pod security context configurations for Filebrowser
  - Enhanced FUSE device mounting capabilities
  - Improved S3 credential management

#### Default Changes

The following components are now disabled by default to allow better integration with existing infrastructure:
- Keycloak
- MinIO
- PostgreSQL
- MongoDB

Users should enable these components explicitly if needed or configure external services.

#### Migration Steps

* [KDL - From 6.1.0 to 6.2.0](https://intelygenz.atlassian.net/wiki/spaces/K/pages/420446216/KDL+-+From+6.1.0+to+6.2.0)

CHANGELOG: [6.2.0](https://github.com/konstellation-io/helm-charts/releases/tag/kdl-server-6.2.0)

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

#### Migration Steps

* [KDL - From 6.0.2 to 6.1.0](https://intelygenz.atlassian.net/wiki/spaces/K/pages/420446216/KDL+-+From+6.0.2+to+6.1.0)

CHANGELOG: [6.1.0](https://github.com/konstellation-io/helm-charts/releases/tag/kdl-server-6.1.0)

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
