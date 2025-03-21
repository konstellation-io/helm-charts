# kai

A Helm chart to deploy KAI

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| ialejandro | <ivan.alejandro@intelygenz.com> |  |
| alpiquero | <angelluis.piquero@intelygenz.com> |  |

## Prerequisites

* Helm 3+
* Kubernetes 1.26+
* Nginx ingress controller. See [Ingress Controller](#ingress-controller).

## Compatibility matrix

| Release ↓ / Kubernetes → | 1.24 | 1.25 | 1.26 | 1.27 | 1.28 | 1.29 | 1.30 |
|:------------------------:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
| 0.2.0                    | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   | ✅   |

| Symbol | Description |
|:------:|-------------|
| ✅     | Perfect match: all features are supported. Client and server versions have exactly the same features/APIs. |
| 🟠     | Forward compatibility: the client will work with the server, but not all new server features are supported. The server has features that the client library cannot use. |
| ❌     | Backward compatibility/Not applicable: the client has features that may not be present in the server. Common features will work, but some client APIs might not be available in the server. |
| -      | Not tested: this combination has not been verified or is not applicable. |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | redis | 18.2.1 |
| https://charts.min.io/ | minio | 5.0.14 |
| https://grafana.github.io/helm-charts | grafana | 7.0.3 |
| https://grafana.github.io/helm-charts | loki | 5.36.3 |
| https://nats-io.github.io/k8s/helm/charts | nats | 1.1.9 |
| https://prometheus-community.github.io/helm-charts | prometheus | 25.4.0 |

## Add repository

```console
helm repo add konstellation-io https://charts.konstellation.io
helm repo update
```

## Install Helm chart (repository mode)

```console
helm install [RELEASE_NAME] konstellation-io/kai
```

This install all the Kubernetes components associated with the chart and creates the release.

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Install Helm chart (OCI mode)

Charts are also available in OCI format. The list of available charts can be found [here](https://github.com/konstellation-io/helm-charts/pkgs/container/helm-charts%2Fkai).

```console
helm install [RELEASE_NAME] oci://ghcr.io/konstellation-io/helm-charts/kai --version=[version]
```

## Included dependencies

* minio
* grafana
* loki
* prometheus
* redis

## Uninstall Helm chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

### Upgrading an existing Release to a new major version

A major chart version change (like v0.15.3 -> v1.0.0) indicates that there is an incompatible breaking change needing
manual actions.

### Legacy changes

* The MongoDB database that was being deployed within the chart has been removed. An external database is needed now. If you come from previous versions of this chart, a MongoDB data migration is necessary. Ref: https://www.mongodb.com/docs/manual/tutorial/backup-and-restore-tools/

Changes in `values.yaml`:

* `mongodb` has been removed in favour of `config.mongodb`
* `mongoExpress` has been added

See [MongoDB](#mongodb) for related info.

* MongoDB Kubernetes resources have been renamed. That also renames the generated mongodb PVC that stores the MongoDB data. A database data migration will be necessary if you come from previous KAI releases.
* The Mongo Express credentials Kubernetes secret has been modified. This secret will only be created if you are deploying the chart for the first time because it uses Helm hooks to avoid secret recreation on chart's upgrades. If you come from a previous release of KAI, execute the following script before upgrading:

```shell
#!/bin/bash
RELEASE_NAME=<release_name>
NAMESPACE=<release_namespace>
ME_CONFIG_MONGODB_ADMINUSERNAME=$(kubectl -n $NAMESPACE get secret kai-mongo-express-secret -o jsonpath='{.data.ME_CONFIG_MONGODB_AUTH_USERNAME}'| base64 -d)
ME_CONFIG_MONGODB_ADMINPASSWORD=$(kubectl -n $NAMESPACE get secret kai-mongo-express-secret -o jsonpath='{.data.ME_CONFIG_MONGODB_AUTH_PASSWORD}'| base64 -d)
kubectl create secret -n $NAMESPACE generic --from-literal ME_CONFIG_MONGODB_ADMINUSERNAME=$ME_CONFIG_MONGODB_ADMINUSERNAME --from-literal ME_CONFIG_MONGODB_ADMINPASSWORD=$ME_CONFIG_MONGODB_ADMINPASSWORD $RELEASE_NAME-mongo-express -o yaml --dry-run=client | kubectl apply -f -
kubectl -n $NAMESPACE annotate secret $RELEASE_NAME-mongo-express helm.sh/hook='pre-install' helm.sh/hook-delete-policy='before-hook-creation'
```

* Minimal Kubernetes supported version is now **v1.19.x**

* Moved `.Values.entrypoints` block to `.Values.k8sManager.generatedEntrypoints` in `values.yaml`.

* k8s-manager Service Account settings have been moved to `k8sManager.serviceAccount` in `values.yaml`

* Removed `mongodb.mongodbUsername` and `mongodb.mongodbPassword` from **values.yaml** in favour of `mongodb.auth.adminUser` and `mongodb.auth.adminpassword`
* Removed `rbac.createServiceAccount` and `rbac.serviceAccount`
* Added `rbac.create` (defaults to true) and added Service Account related block under `k8sManager.serviceAccount`
* Removed other unused values from `values.yaml`.

Check commits [1fab33b](https://github.com/konstellation-io/kai/pull/593/commits/1fab33b8351cae317753017373ac2dab4817c36f) and [a280847](https://github.com/konstellation-io/kai/pull/598/commits/59e7365350d67d30984a2554a28d0241cf74f13e) for more details.

This major version comes with the following changes:

* **Resource label refactor**: Labels have changed for some resources, so the following resources must be manually deleted before updating.

    Affected deployment resources:
    * Admin API
    * Chronograf
    * k8s-manager
    * MongoDB
    * MongoExpress

    Affected statefulset resources:
    * MongoDB
    * NATS

    The commit that introduces the changes is [located here](https://github.com/konstellation-io/kai/pull/585).

* **Ingress annotations are taken from values.yaml**: Now default ingress annotations are specified from [values.yaml](values.yaml) file. If additional ingress annotations are required, those must be appended to the default ones via extra values files or by using the `--set` argument.

* **Openshift routes have been removed**: All Openshift route manifests have been removed from chart. Extend it if you are planning to install it on Openshift platforms.

* **Prometheus Operator have been removed**: Application functionallity has been decoupled from Prometheus so this component is no longer necessary. Use [kube-prometheus-stack](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack) for platform monitoring if needed.

## Configuration

See [Customizing the chart before installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with comments:

```console
helm show values konstellation-io/kai
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| adminApi.affinity | object | `{}` | Assign custom affinity rules to the Admin API pods # ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ # |
| adminApi.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| adminApi.host | string | `"api.kai.local"` | Hostname. This will be used to create the ingress rule and must be a subdomain of `.config.baseDomainName` |
| adminApi.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| adminApi.image.repository | string | `"konstellation/kai-admin-api"` | Image repository |
| adminApi.image.tag | string | `"0.3.0-develop.17"` | Image tag |
| adminApi.imagePullSecrets | list | `[]` | Image pull secrets |
| adminApi.ingress.annotations | object | See `adminApi.ingress.annotations` in [values.yaml](./values.yaml) | Ingress annotations |
| adminApi.ingress.className | string | `"kong"` | The name of the ingress class to use |
| adminApi.logLevel | string | `"INFO"` | Default application log level |
| adminApi.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. # ref: https://kubernetes.io/docs/user-guide/node-selection/ # |
| adminApi.resources | object | `{}` | Container resources |
| adminApi.serviceAccount.annotations | object | `{}` |  |
| adminApi.serviceAccount.create | bool | `true` |  |
| adminApi.serviceAccount.name | string | `""` |  |
| adminApi.tolerations | list | `[]` | Tolerations for use with node taints # ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ # |
| config.admin.corsEnabled | bool | `true` | Whether to enable CORS on Admin API |
| config.baseDomainName | string | `"kai.local"` | Base domain name for Admin API and K8S Manager apps |
| config.loki.datasource | object | `{"jsonData":"{}","uid":""}` | Only when `loki.enabled: true` and `grafana.enabled: true`. Grafana datasource json data config. |
| config.loki.host | string | `"{{ include \"loki.singleBinaryFullname\" .Subcharts.loki }}"` | Loki host. Change this to your own URL when `loki.enabled: false` |
| config.loki.isDefault | bool | `true` | Only when `loki.enabled: true` and `grafana.enabled: true`. Set loki as default datasource for Grafana. |
| config.loki.port | string | `"{{ .Values.loki.loki.server.http_listen_port }}"` | Loki port. Change this to your own URL when `loki.enabled: false` |
| config.minio.defaultRegion | string | us-east-1 | Default region (only affect to Minio buckets) |
| config.minio.tier.aws | object | `{"auth":{"accessKeyID":"","secretAccessKey":"","secretKeyNames":{"accessKey":"","secretKey":""},"secretName":""},"endpointURL":"","region":""}` | Transition Objects from MinIO to AWS S3 |
| config.minio.tier.aws.auth | object | `{"accessKeyID":"","secretAccessKey":"","secretKeyNames":{"accessKey":"","secretKey":""},"secretName":""}` | AWS authentication config @default: first look for the keys in pre-existing kubernetes secret object (secretName and secretKeyNames), if not set, look for the keys in values.yaml (accessKeyID and secretAccessKey) |
| config.minio.tier.aws.auth.accessKeyID | string | `""` | S3 Access Key ID if no secret is used |
| config.minio.tier.aws.auth.secretAccessKey | string | `""` | S3 Secret Access Key if no secret is used |
| config.minio.tier.aws.auth.secretKeyNames | object | `{"accessKey":"","secretKey":""}` | Secret reference for AWS access keys |
| config.minio.tier.aws.auth.secretKeyNames.accessKey | string | `""` | Name of the key in the secret that contains the access key ID |
| config.minio.tier.aws.auth.secretKeyNames.secretKey | string | `""` | Name of the key in the secret that contains the secret access key |
| config.minio.tier.aws.auth.secretName | string | `""` | Name of the secret that contains the credentials for S3 |
| config.minio.tier.aws.endpointURL | string | https://s3.amazonaws.com | S3 Service endpoint URL |
| config.minio.tier.aws.region | string | us-east-1 | The Region where the remote bucket was created. |
| config.minio.tier.enabled | bool | `false` | Whether to enable MinIO Tiering @default: If is disable MinIO will use only local storage |
| config.minio.tier.name | string | KAI-REMOTE-STORAGE | Tier name |
| config.minio.tier.remoteBucketName | string | `""` | Remote storage bucket name (must exist) |
| config.minio.tier.remotePrefix | string | DATA | Prefix or path in bucket where object transition will happen (will be created if not exist) |
| config.mongodb.connectionString.secretKey | string | `""` | The name of the secret key that contains the MongoDB connection string. |
| config.mongodb.connectionString.secretName | string | `""` | The name of the secret that contains a key with the MongoDB connection string. |
| config.prometheus.datasource | object | `{"jsonData":"{}"}` | Only when `prometheus.enabled: true` and `grafana.enabled: true`. Grafana datasource json data config. |
| config.prometheus.isDefault | bool | `false` | Only when `prometheus.enabled: true` and `grafana.enabled: true`. Set prometheus as default datasource for Grafana. |
| config.prometheus.kaiScrapeConfigs.configmapName | string | `"prometheus-additional-scrape-configs"` | configmap name for additional scrape configs |
| config.prometheus.kaiScrapeConfigs.enabled | bool | `true` | Enable creation of configmap that contains custom prometheus scrape configs for KAI metrics. Usefull to use with external prometheus instance. If `prometheus.enabled: true` this cannot be disabled |
| config.prometheus.url | string | `"http://{{ include \"prometheus.fullname\" .Subcharts.prometheus }}-{{ .Values.prometheus.server.name }}:{{ .Values.prometheus.server.service.servicePort }}{{ .Values.prometheus.server.prefixURL }}"` | Prometheus endpoint url. Change this to your own URL when `prometheus.enabled: false` |
| config.redis.architecture | string | `"standalone"` | architecture. Allowed values: `standalone` or `replication`. Only apply when use your own redis. This config allow send replicas urls to admin-api when replication is activated |
| config.redis.auth.enabled | bool | `false` | Whether to enable auth to redis. Only apply when use your own redis. This allow send credentials to admin-api |
| config.redis.auth.existingSecret | string | `""` | Name of the secret that contains the redis password |
| config.redis.auth.existingSecretPasswordKey | string | `""` | Name of the key in the secret that contains the redis password |
| config.redis.auth.password | string | `""` | Redis password if no existingSecret is used and `redis.enabled: false`. (create a secret with this password and send credentials to admin-api) |
| config.redis.master.url | string | `"{{ printf \"%s-master\" (include \"common.names.fullname\" .Subcharts.redis) }}:{{ .Values.redis.master.service.ports.redis }}"` | Redis Master endpoint url. Change this to your own URL when `redis.enabled: false` |
| config.redis.replicas.url | string | `"{{ printf \"%s-replicas\" (include \"common.names.fullname\" .Subcharts.redis) }}:{{ .Values.redis.replica.service.ports.redis }}"` | Redis Replicas endpoint url. Change this to your own URL when `redis.enabled: false` |
| config.tls.certSecretName | string | `""` | An existing secret containing a valid wildcard certificate for the value provissioned in `.config.baseDomainName`. Required if `config.tls.enabled = true` |
| config.tls.enabled | bool | `false` | Whether to enable TLS |
| developmentMode | bool | `false` | Whether to setup developement mode |
| grafana.admin | object | `{"existingSecret":"","passwordKey":"admin-password","userKey":"admin-user"}` | Use an existing secret for the admin user |
| grafana.admin.passwordKey | string | `"admin-password"` | Name of the key in the secret that contains the password |
| grafana.admin.userKey | string | `"admin-user"` | Name of the key in the secret that contains the admin user |
| grafana.adminPassword | string | Randomly generated value | Set admin password (ommited if existingSecret is set) |
| grafana.adminUser | string | `"admin"` | Admin user name |
| grafana.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| grafana.enabled | bool | `true` | Whether to enable Grafana |
| grafana.image.tag | string | `"10.2.0"` | Grafana version |
| grafana.ingress.annotations | object | `{}` | Ingress annotations |
| grafana.ingress.enabled | bool | `true` | Enable ingress for MinIO Web Console |
| grafana.ingress.hosts | list | `["monitoring.kai.local"]` | Ingress hostnames |
| grafana.ingress.ingressClassName | string | `"kong"` | The name of the ingress class to use |
| grafana.ingress.labels | object | `{}` | Ingress labels |
| grafana.ingress.tls | list | `[]` | Ingress TLS configuration |
| grafana.persistence.accessMode | string | `"ReadWriteOnce"` | Access mode for the volume |
| grafana.persistence.enabled | bool | `false` | Enables persistent storage using PVC |
| grafana.persistence.size | string | `"1Gi"` | Storage size |
| grafana.persistence.storageClass | string | `""` | Storage class name |
| grafana.plugins[0] | string | `"redis-datasource"` |  |
| grafana.service.port | int | `80` | Internal port number for Grafana service |
| grafana.service.type | string | `"ClusterIP"` | Service type |
| grafana.sidecar | object | `{"datasources":{"enabled":true,"label":"grafana_datasource","labelValue":"1","maxLines":1000}}` | sidecar config (required for datasource config section in loki and prometheus) |
| k8sManager.affinity | object | `{}` | Assign custom affinity rules to the K8S Manager pods # ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ # |
| k8sManager.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| k8sManager.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| k8sManager.image.repository | string | `"konstellation/kai-k8s-manager"` | Image repository |
| k8sManager.image.tag | string | `"0.3.0-develop.17"` | Image tag |
| k8sManager.imageBuilder.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for image builder's jobs |
| k8sManager.imageBuilder.image.repository | string | `"gcr.io/kaniko-project/executor"` | Image repository for image builder's jobs |
| k8sManager.imageBuilder.image.tag | string | `"v1.18.0"` | Image tag for image builder's jobs |
| k8sManager.imageBuilder.logLevel | string | `"info"` | Log level for image builder's jobs |
| k8sManager.imageBuilder.netrc.content | string | `""` | .netrc file content. Ref: https://everything.curl.dev/usingcurl/netrc |
| k8sManager.imageBuilder.netrc.enabled | bool | `false` | Whether to create .netrc file for authentication for private dependency repositories |
| k8sManager.imageBuilder.pullCredentials.enabled | bool | `false` | Whether to add pull credentials to the image builder's jobs (needed to pull base images from private registries or proxies) |
| k8sManager.imageBuilder.pullCredentials.password | string | `""` | The registry password |
| k8sManager.imageBuilder.pullCredentials.registry | string | `"https://index.docker.io/v1/"` | The registry server |
| k8sManager.imageBuilder.pullCredentials.username | string | `""` | The registry username |
| k8sManager.imagePullSecrets | list | `[]` | Image pull secrets |
| k8sManager.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. # ref: https://kubernetes.io/docs/user-guide/node-selection/ # |
| k8sManager.processes.sidecars.fluentbit.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for Fuent Bit sidecar |
| k8sManager.processes.sidecars.fluentbit.image.repository | string | `"fluent/fluent-bit"` | Image repository for Fuent Bit sidecar |
| k8sManager.processes.sidecars.fluentbit.image.tag | string | `"2.2.0"` | Image tag for Fuent Bit sidecar |
| k8sManager.processes.sidecars.telegraf.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy for Fuent Bit sidecar |
| k8sManager.processes.sidecars.telegraf.image.repository | string | `"telegraf"` | Image repository for Fuent Bit sidecar |
| k8sManager.processes.sidecars.telegraf.image.tag | string | `"1.28.5"` | Image tag for Fuent Bit sidecar |
| k8sManager.processes.triggers.ingress.annotations | object | `{}` | The annotations that all the generated ingresses for the entrypoints will have |
| k8sManager.processes.triggers.ingress.className | string | `"kong"` | The ingressClassName to use for the enypoints' generated ingresses |
| k8sManager.resources | object | `{}` | Container resources |
| k8sManager.serviceAccount.annotations | object | `{}` | The Service Account annotations |
| k8sManager.serviceAccount.create | bool | `true` | Whether to create the Service Account |
| k8sManager.serviceAccount.name | string | `""` | The name of the service account. @default: A pre-generated name based on the chart relase fullname sufixed by `-k8s-manager` |
| k8sManager.tolerations | list | `[]` | Tolerations for use with node taints # ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ # |
| keycloak.adminApi.oidcClient.clientId | string | `"admin-cli"` | The name of the OIDC client in Keycloak for the master realm admin |
| keycloak.affinity | object | `{}` | Assign custom affinity rules to the Keycloak pods |
| keycloak.argsOverride | object | `{}` | Args to pass to the Keycloak startup command. This takes precedence over options passed through env variables |
| keycloak.auth.adminPassword | string | `"123456"` | Keycloak admin password |
| keycloak.auth.adminUser | string | `"admin"` | Keycloak admin username |
| keycloak.auth.existingSecret.name | string | `""` | The name of the secret that contains a key with the Keycloak admin password. Existing secret takes precedence over `adminUser` and `adminPassword` |
| keycloak.auth.existingSecret.passwordKey | string | `""` | The name of the secret key that contains the Keycloak admin password. |
| keycloak.auth.existingSecret.userKey | string | `""` | The name of the secret key that contains the Keycloak admin username. |
| keycloak.config.healthEnabled | string | `"true"` | If the server should expose health check endpoints. If set to "false", container liveness and readiness probes should be disabled. |
| keycloak.config.hostnameStrict | string | `"false"` | Disables dynamically resolving the hostname from request headers. Should always be set to true in production, unless proxy verifies the Host header. |
| keycloak.config.httpEnabled | string | `"true"` | Whether to enable http |
| keycloak.config.metricsEnabled | string | `"false"` | Whether to enable metrics |
| keycloak.config.proxy | string | `"edge"` | The proxy address forwarding mode if the server is behind a reverse proxy. Valid values are `none`, `edge`, `reencrypt` and `passthrough` |
| keycloak.db.auth.database | string | `""` | The database name |
| keycloak.db.auth.host | string | `""` | The database hostname |
| keycloak.db.auth.password | string | `""` | The database password |
| keycloak.db.auth.port | string | `""` | The database port |
| keycloak.db.auth.secretDatabaseKey | string | `""` | The name of the secret key that contains the database name. Takes precedence over `database` |
| keycloak.db.auth.secretHostKey | string | `""` | The name of the secret key that contains the database host. |
| keycloak.db.auth.secretName | string | `""` | The name of the secret that contains the database connection config keys. |
| keycloak.db.auth.secretPasswordKey | string | `""` | The name of the secret key that contains the database password. |
| keycloak.db.auth.secretPortKey | string | `""` | The name of the secret key that contains the database port. Takes precedence over `host` |
| keycloak.db.auth.secretUserKey | string | `""` | The name of the secret key that contains the database username. Takes precedence over `port` |
| keycloak.db.auth.username | string | `""` | The database username |
| keycloak.db.type | string | `"postgres"` | Keycloak database type |
| keycloak.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| keycloak.extraEnv | object | `{}` | Keycloak extra env vars in the form of a list of key-value pairs |
| keycloak.extraVolumeMounts | list | `[]` | Extra volume mounts |
| keycloak.extraVolumes | list | `[]` | Extra volumes |
| keycloak.host | string | `"auth.kai.local"` | Hostname. This will be used to create the ingress rulem and to configure Keycloak and must be a subdomain of `.config.baseDomainName` |
| keycloak.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| keycloak.image.repository | string | `"quay.io/keycloak/keycloak"` | The image repository |
| keycloak.image.tag | string | `"24.0.1"` | The image tag |
| keycloak.imagePullSecrets | list | `[]` | Image pull secrets |
| keycloak.ingress.annotations | object | See `keycloak.ingress.annotations` in [values.yaml](./values.yaml) | Ingress annotations |
| keycloak.ingress.className | string | `"kong"` | The name of the ingress class to use |
| keycloak.kli.oidcClient.clientId | string | `"kai-kli-oidc"` | The name of the OIDC client in Keycloak for KLI |
| keycloak.kliCI.oidcClient.clientId | string | `"kai-kli-ci-oidc"` | The name of the OIDC client in Keycloak for KLI CI |
| keycloak.kliCI.oidcClient.secret | string | `""` | The secret for the OIDC client that will be created on Keycloak first startup |
| keycloak.kong.oidcClient.clientId | string | `"kong-oidc"` | The name of the OIDC client in Keycloak for Kong |
| keycloak.kong.oidcClient.secret | string | `""` | The secret for the OIDC client that will be created on Keycloak first startup |
| keycloak.kong.oidcPluginName | string | `"oidc"` | The name of the OIDC Kong plugin that should be installed on Kong ingress controller |
| keycloak.livinessProbe | object | `{"failureThreshold":3,"httpGet":{"path":"/health/live","port":"http"},"initialDelaySeconds":30,"periodSeconds":10,"timeoutSeconds":5}` | Container liveness probe |
| keycloak.minio.oidcClient | object | `{"clientId":"minio","secret":""}` | The name of the OIDC client in Keycloak for MinIO |
| keycloak.minio.oidcClient.clientId | string | `"minio"` | The name of the OIDC client in Keycloak for Kong |
| keycloak.minio.oidcClient.secret | string | `""` | The secret for the OIDC client that will be created on Keycloak first startup |
| keycloak.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. |
| keycloak.podAnnotations | object | `{}` | Pod annotations |
| keycloak.podSecurityContext | object | `{}` | Pod security context |
| keycloak.readinessProbe | object | `{"failureThreshold":3,"httpGet":{"path":"/health/ready","port":"http"},"initialDelaySeconds":30,"periodSeconds":10,"timeoutSeconds":5}` | Container readiness probe |
| keycloak.realmName | string | `"konstellation"` | The name of the realm that will be crated on Keycloak first startup |
| keycloak.resources | object | `{}` | Container resources |
| keycloak.securityContext | object | `{}` |  |
| keycloak.service.ports.http | int | `8080` | The http port the service will listen on. Only |
| keycloak.service.ports.https | int | `8443` | The https port the service will listen on |
| keycloak.service.type | string | `"ClusterIP"` | Service type |
| keycloak.serviceAccount.annotations | object | `{}` |  |
| keycloak.serviceAccount.create | bool | `true` |  |
| keycloak.serviceAccount.name | string | `""` |  |
| keycloak.ssoSessionIdleTimeout | int | `14400` | The time in seconds that a user has to be inactive to expire the session |
| keycloak.tolerations | list | `[]` | Assign custom tolerations to the Keycloak pods |
| loki.enabled | bool | `true` | Whether to enable Loki |
| loki.gateway.enabled | bool | `false` | Specifies whether the gateway should be enabled |
| loki.loki.auth_enabled | bool | `false` | Should authentication be enabled |
| loki.loki.commonConfig | object | `{"replication_factor":1}` | monolithic loki |
| loki.loki.image.tag | string | `"2.9.2"` | Loki version |
| loki.loki.persistence.accessMode | string | `"ReadWriteOnce"` | Access mode for the volume |
| loki.loki.persistence.enabled | bool | `false` | Enables persistent storage using PVC |
| loki.loki.persistence.size | string | `"2Gi"` | Storage size |
| loki.loki.persistence.storageClass | string | `""` | Storage class name |
| loki.loki.server.http_listen_port | int | `3100` |  |
| loki.loki.service.port | int | `3100` | Internal port number for Grafana service |
| loki.loki.service.type | string | `"ClusterIP"` | Service type |
| loki.loki.storage.type | string | `"filesystem"` |  |
| loki.monitoring.lokiCanary | object | `{"enabled":false}` | Whether to enable lokiCanary |
| loki.monitoring.selfMonitoring | object | `{"enabled":false,"grafanaAgent":{"installOperator":false}}` | scrape its own Loki logs |
| loki.singleBinary.replicas | int | `1` |  |
| loki.test.enabled | bool | `false` |  |
| minio.consoleIngress.annotations | object | `{}` | Ingress annotations |
| minio.consoleIngress.enabled | bool | `true` | Enable ingress for MinIO Web Console |
| minio.consoleIngress.hosts | list | `["storage-console.kai.local"]` | Ingress hostnames |
| minio.consoleIngress.ingressClassName | string | `"kong"` | The name of the ingress class to use |
| minio.consoleIngress.labels | object | `{}` | Ingress labels |
| minio.consoleIngress.tls | list | `[]` | Ingress TLS configuration |
| minio.existingSecret | string | `""` | Use an exising secret for root user and password |
| minio.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| minio.image.repository | string | `"quay.io/minio/minio"` | Image repository |
| minio.image.tag | string | `"RELEASE.2023-09-30T07-02-29Z"` | Image tag |
| minio.ingress.annotations | object | `{}` | Ingress annotations |
| minio.ingress.enabled | bool | `true` | Enable ingress for MinIO API |
| minio.ingress.hosts | list | `["storage.kai.local"]` | Ingress hostnames |
| minio.ingress.ingressClassName | string | `"kong"` | The name of the ingress class to use |
| minio.ingress.labels | object | `{}` | Ingress labels |
| minio.ingress.tls | list | `[]` | Ingress TLS configuration |
| minio.mcImage.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| minio.mcImage.repository | string | `"quay.io/minio/mc"` | Image repository |
| minio.mcImage.tag | string | `"RELEASE.2023-09-29T16-41-22Z"` | Image tag |
| minio.minioAPIPort | string | `"9000"` | Internal port number for MinIO S3 API container |
| minio.minioConsolePort | string | `"9001"` | Internal port number for MinIO Browser Console container |
| minio.mode | string | `"standalone"` | Sets minio mode |
| minio.persistence.accessMode | string | `"ReadWriteOnce"` | Access mode for the volume |
| minio.persistence.enabled | bool | `true` | Enables persistent storage using PVC |
| minio.persistence.size | string | `"10Gi"` | Storage size @note: If Tier is enabled, ensure enough space to MinIO have time to transfer objects to external storage and free space in local persistence storage |
| minio.persistence.storageClass | string | `""` | Storage class name |
| minio.resources | object | `{"requests":{"memory":"256Mi"}}` | Sets pods resources |
| minio.rootPassword | string | Randomly generated value | Sets Root password |
| minio.rootUser | string | Randomly generated value | Sets Root user |
| minio.service.port | string | `"9000"` | Internal port number for MinIO S3 API service |
| minio.service.type | string | `"ClusterIP"` | Service type |
| nameOverride | string | `""` | Provide a name in place of kai for `app.kubernetes.io/name` labels |
| nats.config | object | `{"cluster":{"enabled":false,"replicas":3},"jetstream":{"enabled":true,"fileStore":{"enabled":true,"pvc":{"enabled":true,"size":"10Gi","storageClassName":null}},"memoryStore":{"enabled":true,"maxSize":"2Gi"}},"merge":{"debug":false,"logtime":true,"trace":false},"nats":{"port":4222}}` | The NATS config as described at https://github.com/nats-io/k8s/tree/main/helm/charts/nats#nats-server |
| nats.config.cluster | object | `{"enabled":false,"replicas":3}` | The NATS server configuration |
| nats.config.jetstream | object | `{"enabled":true,"fileStore":{"enabled":true,"pvc":{"enabled":true,"size":"10Gi","storageClassName":null}},"memoryStore":{"enabled":true,"maxSize":"2Gi"}}` | The NATS JetStream configuration |
| nats.config.jetstream.fileStore | object | `{"enabled":true,"pvc":{"enabled":true,"size":"10Gi","storageClassName":null}}` | The NATS JetStream storage configuration |
| nats.config.jetstream.memoryStore | object | `{"enabled":true,"maxSize":"2Gi"}` | The NATS JetStream memory storage configuration |
| nats.config.merge | object | `{"debug":false,"logtime":true,"trace":false}` | Merge the NATS server configuration |
| nats.monitor.enabled | bool | `false` | Whether to enable monitoring |
| nats.monitor.port | int | `8222` | Monitoring service port |
| nats.natsBox.enabled | bool | `false` | Whether to enable the NATS Box |
| nats.podTemplate.merge | object | `{"spec":{"affinity":{},"nodeSelector":{},"tolerations":[]}}` | Merge the pod template: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#pod-v1-core |
| nats.podTemplate.merge.spec.affinity | object | `{}` | Assign custom affinity rules to the NATS pods |
| nats.podTemplate.merge.spec.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. |
| nats.podTemplate.merge.spec.tolerations | list | `[]` | Assign custom tolerations to the NATS pods |
| nats.promExporter.enabled | bool | `false` | Whether to enable the Prometheus Exporter |
| nats.promExporter.port | int | `7777` | Prometheus Exporter service port |
| nats.service.name | string | `nil` | nats service name |
| nats.serviceAccount.enabled | bool | `true` | Whether to enable the service account |
| natsManager.affinity | object | `{}` | Assign custom affinity rules to the NATS pods # ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ # |
| natsManager.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| natsManager.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| natsManager.image.repository | string | `"konstellation/kai-nats-manager"` | Image repository |
| natsManager.image.tag | string | `"0.3.0-develop.17"` | Image tag |
| natsManager.imagePullSecrets | list | `[]` | Image pull secrets |
| natsManager.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. # ref: https://kubernetes.io/docs/user-guide/node-selection/ # |
| natsManager.resources | object | `{}` | Container resources |
| natsManager.serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| natsManager.serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| natsManager.serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |
| natsManager.tolerations | list | `[]` | Tolerations for use with node taints # ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ # |
| prometheus.alertmanager.enabled | bool | `true` | Whether to enable alertmanager |
| prometheus.alertmanager.image.tag | string | `"v0.26.0"` | alertmanager server version |
| prometheus.alertmanager.persistence.accessModes | list | `["ReadWriteOnce"]` | Access mode for the volume |
| prometheus.alertmanager.persistence.enabled | bool | `true` |  |
| prometheus.alertmanager.persistence.size | string | `"2Gi"` | Storage size |
| prometheus.alertmanager.persistence.storageClass | string | `""` | Storage class name |
| prometheus.enabled | bool | `true` | Whether to enable Prometheus |
| prometheus.kube-state-metrics.enabled | bool | `false` |  |
| prometheus.prometheus-node-exporter.enabled | bool | `false` |  |
| prometheus.prometheus-pushgateway.enabled | bool | `false` |  |
| prometheus.scrapeConfigFiles | list | `["/etc/config/additional-scrape-configs/*.yaml"]` | Files to get additional scrapeConfigs. This allows scrape configs defined in `prometheus.server.extraConfigmapMounts` |
| prometheus.server.extraConfigmapMounts[0] | object | `{"configMap":"prometheus-additional-scrape-configs","mountPath":"/etc/config/additional-scrape-configs/","name":"additional-scrape-configs","readOnly":true,"subPath":""}` | name for volumes and volumeMount config |
| prometheus.server.extraConfigmapMounts[0].configMap | string | `"prometheus-additional-scrape-configs"` | name of the configmap. Must be the same name set in `config.prometheus.kaiScrapeConfigs.configmapName` |
| prometheus.server.image.tag | string | `"v2.47.2"` | prometheus server version |
| prometheus.server.ingress.annotations | object | `{}` | Ingress annotations |
| prometheus.server.ingress.enabled | bool | `true` | Enable ingress for MinIO Web Console |
| prometheus.server.ingress.extraLabels | object | `{}` | Ingress labels |
| prometheus.server.ingress.hosts | list | `["prometheus.kai.local"]` | Ingress hostnames |
| prometheus.server.ingress.ingressClassName | string | `"kong"` | The name of the ingress class to use |
| prometheus.server.ingress.tls | list | `[]` | Ingress TLS configuration |
| prometheus.server.name | string | `"server"` | name of the prometheus server |
| prometheus.server.persistentVolume.accessModes | list | `["ReadWriteOnce"]` | Access mode for the volume |
| prometheus.server.persistentVolume.enabled | bool | `true` | Enables persistent storage using PVC |
| prometheus.server.persistentVolume.size | string | `"2Gi"` | Storage size |
| prometheus.server.persistentVolume.storageClass | string | `""` | Storage class name |
| prometheus.server.service.servicePort | int | `80` | Internal port number for grafana service |
| prometheus.server.service.type | string | `"ClusterIP"` | Service type |
| rbac.create | bool | `true` | Whether to create the roles for the services that could use custom Service Accounts |
| redis.architecture | string | `"standalone"` | architecture. Allowed values: `standalone` or `replication` |
| redis.auth.enabled | bool | `true` | Whether to enable auth to redis |
| redis.auth.existingSecret | string | `""` | The name of an existing secret with redis credentials |
| redis.auth.existingSecretPasswordKey | string | `""` | Password key to be retrieved from existing secret |
| redis.auth.password | string | `""` |  |
| redis.enabled | bool | `true` | Whether to enable redis |
| redis.image.repository | string | `"redis/redis-stack-server"` |  |
| redis.image.tag | string | `"7.2.0-v6"` | Redis server version |
| redis.master | object | `{"args":["-c","/opt/bitnami/scripts/start-script/start-master.sh"],"containerPorts":{"redis":6379},"extraVolumeMounts":[{"mountPath":"/opt/bitnami/scripts/start-script","name":"redis-master-start-script"}],"extraVolumes":[{"configMap":{"defaultMode":493,"name":"redis-stack-master-config"},"name":"redis-master-start-script"}],"persistence":{"accessModes":["ReadWriteOnce"],"enabled":true,"size":"2Gi","storageClass":""},"service":{"ports":{"redis":6379},"type":"ClusterIP"}}` | number of replicas |
| redis.master.containerPorts | object | `{"redis":6379}` | redis Container port to open on master nodes |
| redis.master.extraVolumeMounts | list | `[{"mountPath":"/opt/bitnami/scripts/start-script","name":"redis-master-start-script"}]` | Extra volume mounts for additional config |
| redis.master.extraVolumes | list | `[{"configMap":{"defaultMode":493,"name":"redis-stack-master-config"},"name":"redis-master-start-script"}]` | Extra volumes for additional config |
| redis.master.persistence.accessModes | list | `["ReadWriteOnce"]` | Access mode for the volume |
| redis.master.persistence.size | string | `"2Gi"` | Storage size |
| redis.master.persistence.storageClass | string | `""` | Storage class name |
| redis.master.service.ports | object | `{"redis":6379}` | Internal port number for master redis service |
| redis.master.service.type | string | `"ClusterIP"` | Service type |
| redis.metrics.enabled | bool | `false` | Start a sidecar prometheus exporter to expose redis metrics |
| redis.replica.args[0] | string | `"-c"` |  |
| redis.replica.args[1] | string | `"/opt/bitnami/scripts/start-script/start-replicas.sh"` |  |
| redis.replica.autoscaling.enabled | bool | `false` |  |
| redis.replica.autoscaling.maxReplicas | int | `5` | Max replicas for the pod autoscaling |
| redis.replica.autoscaling.minReplicas | int | `1` | Min replicas for the pod autoscaling |
| redis.replica.extraVolumeMounts | list | `[{"mountPath":"/opt/bitnami/scripts/start-script","name":"redis-replicas-start-script"}]` | Extra volume mounts for additional config |
| redis.replica.extraVolumes | list | `[{"configMap":{"defaultMode":493,"name":"redis-stack-replicas-config"},"name":"redis-replicas-start-script"}]` | Extra volumes for additional config |
| redis.replica.persistence.accessModes | list | `["ReadWriteOnce"]` | Access mode for the volume |
| redis.replica.persistence.enabled | bool | `true` |  |
| redis.replica.persistence.size | string | `"2Gi"` | Storage size |
| redis.replica.persistence.storageClass | string | `""` | Storage class name |
| redis.replica.replicaCount | int | `1` | Number of replicas |
| redis.replica.service.ports | object | `{"redis":6379}` | Internal port number for master redis service |
| redis.replica.service.type | string | `"ClusterIP"` | Service type |
| redis.tls.authClients | bool | `false` | Require clients to authenticate |
| redis.tls.autoGenerated | bool | `false` | Enable autogenerated certificates |
| redis.tls.enabled | bool | `false` | Enabled Enable TLS traffic |
| redis.tls.existingSecret | string | `""` | The name of the existing secret that contains the TLS certificates |
| registry.affinity | object | `{}` | Assign custom affinity rules to the pods # ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ # |
| registry.auth.password | string | password | Registry password |
| registry.auth.user | string | user | Registry username |
| registry.config | string | `""` | A string contaning the config for Docker Registry. Ref: https://docs.docker.com/registry/configuration/. |
| registry.configSecret.key | string | `""` | The name of the secret key that contains the registry config file |
| registry.configSecret.name | string | `""` | Takes precedence over 'registry.config'. The name of the secret that contains the registry config file. |
| registry.containerPort | int | `5000` | The container port |
| registry.deploymentStrategy | object | `{"type":"Recreate"}` | Deployment Strategy |
| registry.extraVolumeMounts | list | `[]` | Extra volume mounts for the registry deployment |
| registry.extraVolumes | list | `[]` | Extra volumes for the registry deployment |
| registry.host | string | `"registry.kai.local"` | Hostname. This will be used to create the ingress rule and must be a subdomain of `.config.baseDomainName` |
| registry.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| registry.image.repository | string | `"distribution/distribution"` | Image repository |
| registry.image.tag | string | `"3.0.0-alpha.1"` | Image tag |
| registry.imagePullSecrets | list | `[]` | Image pull secrets |
| registry.ingress.annotations | object | See `adminApi.ingress.annotations` in [values.yaml](./values.yaml) | Ingress annotations |
| registry.ingress.className | string | `"kong"` | The name of the ingress class to use |
| registry.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. # ref: https://kubernetes.io/docs/user-guide/node-selection/ # |
| registry.podAnnotations | object | `{}` | Pod annotations |
| registry.podSecurityContext | object | `{}` | Pod security context |
| registry.resources | object | `{}` | Container resources |
| registry.securityContext | object | `{}` |  |
| registry.service.ports.http | int | `5000` | The http port the service will listen on. Only |
| registry.service.type | string | `"ClusterIP"` | Service type |
| registry.serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| registry.serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| registry.serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |
| registry.storage.accessMode | string | `"ReadWriteOnce"` | Access mode for the volume |
| registry.storage.enabled | bool | `true` | Whether to enable persistence. This only is used when storageDriver is set to "filesystem" |
| registry.storage.path | string | `"/var/lib/registry"` | Persistent volume mount point. This will define Registry app workdir too. |
| registry.storage.size | string | `"10Gi"` | Storage size |
| registry.storage.storageClass | string | `""` | Storage class name |
| registry.storageDriver.azure.config | object | `{}` | Azure Storage driver config block as defined at https://distribution.github.io/distribution/storage-drivers/azure/ |
| registry.storageDriver.azure.enabled | bool | `false` | Whether to enable the Azure storage driver |
| registry.storageDriver.filesystem.config.rootDirectory | string | `"/var/lib/registry"` |  |
| registry.storageDriver.filesystem.enabled | bool | `true` | Whether to enable the filesystem storage driver |
| registry.storageDriver.gcs.config | object | `{}` | GCS Storage driver config block as defined at https://distribution.github.io/distribution/storage-drivers/gcs/ |
| registry.storageDriver.gcs.enabled | bool | `false` | Whether to enable the GCS storage driver |
| registry.storageDriver.inmemory.enabled | bool | `false` | Whether to enable the in-memory storage driver. Development only |
| registry.storageDriver.s3.config | object | `{}` | S3 Storage driver config block as defined at https://distribution.github.io/distribution/storage-drivers/s3/ |
| registry.storageDriver.s3.enabled | bool | `false` | Whether to enable the S3 storage driver |
| registry.tolerations | list | `[]` | Tolerations for use with node taints # ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ # |

## Ingress controller

This Chart has been developed using **Nginx Ingress Controller**. So using the default ingress annotations ensures its correct operation. .

*See [values.yaml](values.yaml) file and [Nginx Ingress controller](https://kubernetes.github.io/ingress-nginx/) for additional documentation**.

However, users could use any other ingress controller (for example, [Traefik](https://doc.traefik.io/traefik/providers/kubernetes-ingress/)). In that case, ingress configurations equivalent to the default ones must be povided.

Notice that even using equivalent ingress configurations the correct operation of the appliance is not guaranteed.

## MongoDB

This chart needs an external MongoDB compatible database to work. Following user and permissions are recomended for a correct and secure application opration:

* User **kai**
  * Purposse: KAI main database user
  * Database: **admin**
  * Attached Roles:
    * *userAdminAnyDatabase* (admin)
    * *readWriteAnyDatabase* (admin)
    * *dbAdminAnyDatabase* (admin)
