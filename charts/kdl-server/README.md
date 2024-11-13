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
* Kubernetes 1.24+

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| oci://ghcr.io/konstellation-io/helm-charts | keycloak(konstellation-base) | 1.0.2 |
| oci://ghcr.io/oauth2-proxy/charts | oauth2proxy(oauth2-proxy) | 7.7.28 |
| oci://registry-1.docker.io/bitnamicharts | minio | 14.8.1 |
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
> Upgrading an existing Release to a new major version (`v0.15.X` -> ``v1.0.0`) indicates that there is an incompatible **BREAKING CHANGES** needing manual actions.

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
| backup | object | `{"activeDeadlineSeconds":3600,"backoffLimit":3,"concurrencyPolicy":"Forbid","enabled":false,"extraVolumeMounts":[],"extraVolumes":[],"failedJobsHistoryLimit":1,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-backup","tag":"0.23.0"},"name":"backup-gitea","resources":{},"s3":{"awsAccessKeyID":"aws-access-key-id","awsSecretAccessKey":"aws-secret-access-key","bucketName":"s3-bucket-name"},"schedule":"0 1 * * 0","startingDeadlineSeconds":60,"successfulJobsHistoryLimit":0,"ttlSecondsAfterFinished":""}` | Backup job configuration |
| backup.activeDeadlineSeconds | int | `3600` | Sets the activeDeadlineSeconds param for the backup cronjob. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/job/#job-termination-and-cleanup |
| backup.backoffLimit | int | `3` | Sets the backoffLimit param for the backup cronjob. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/job/#pod-backoff-failure-policy |
| backup.concurrencyPolicy | string | `"Forbid"` | Specifies how to treat concurrent executions of a Job. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#concurrency-policy |
| backup.enabled | bool | `false` | Whether to enable backup |
| backup.extraVolumeMounts | list | `[]` | Extra volume mounts for backup pods |
| backup.extraVolumes | list | `[]` | Extra volumes for backup pods |
| backup.failedJobsHistoryLimit | int | `1` | The number of failed finished jobs to retain. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#jobs-history-limits |
| backup.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| backup.image.repository | string | `"konstellation/kdl-backup"` | Image repository |
| backup.image.tag | string | `"0.23.0"` | Image tag |
| backup.name | string | `"backup-gitea"` | Name of the backup cronjob |
| backup.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| backup.s3 | object | `{"awsAccessKeyID":"aws-access-key-id","awsSecretAccessKey":"aws-secret-access-key","bucketName":"s3-bucket-name"}` | AWS S3 Bucket configuration |
| backup.s3.awsAccessKeyID | string | `"aws-access-key-id"` | AWS Access Key ID for acceding backup bucket |
| backup.s3.awsSecretAccessKey | string | `"aws-secret-access-key"` | AWS Secret Access Key for acceding backup bucket |
| backup.s3.bucketName | string | `"s3-bucket-name"` | The S3 bucket that will store all backups |
| backup.schedule | string | `"0 1 * * 0"` | Backup cronjob schedule |
| backup.startingDeadlineSeconds | int | `60` | Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#job-creation |
| backup.successfulJobsHistoryLimit | int | `0` | The number of successful finished jobs to retain. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#jobs-history-limits |
| backup.ttlSecondsAfterFinished | string | `""` | Limits the lifetime of a Job that has finished execution (either Complete or Failed). |
| cleaner | object | `{"enabled":false,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-cleaner","tag":"0.16.0"},"schedule":"0 1 * * 0","threshold":5,"trashPath":"/shared-storage/.trash"}` | Cleaner job configuration |
| cleaner.enabled | bool | `false` | Whether to enable cleaner cronjob |
| cleaner.schedule | string | `"0 1 * * 0"` | Celaner cronjob schedule |
| cleaner.threshold | int | `5` | The minimun age of files to be removed |
| cleaner.trashPath | string | `"/shared-storage/.trash"` | The name of the trash path |
| command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| env | object | `{}` | Environment variables to configure application |
| envFromConfigMap | object | `{}` | Variables from configMap |
| envFromFiles | object | `{}` | Variables from files managed by you </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| envFromSecrets | object | `{}` | Variables from secrets |
| fullnameOverride | string | `""` | String to fully override konstellation-base.fullname template |
| gitea | object | `{"admin":{"email":"test@test.com","password":"123456","username":"kdladmin"},"affinity":{},"enabled":true,"image":{"pullPolicy":"IfNotPresent","repository":"gitea/gitea","tag":"1.14.4"},"ingress":{"annotations":{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n"},"className":"nginx","tls":{"secretName":null}},"nodeSelector":{},"storage":{"size":"10Gi","storageClassName":"standard"},"tolerations":[]}` | Gitea deployment (DEPRECATION) Remove in future versions |
| gitea.admin.email | string | `"test@test.com"` | Admin user email |
| gitea.admin.password | string | `"123456"` | Admin password |
| gitea.admin.username | string | `"kdladmin"` | Admin username |
| gitea.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| gitea.enabled | bool | `true` | Whether to enable Gitea |
| gitea.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n"}` | Ingress annotations |
| gitea.ingress.className | string | `"nginx"` | The ingress class name |
| gitea.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| gitea.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| gitea.storage.size | string | `"10Gi"` | Storage size |
| gitea.storage.storageClassName | string | `"standard"` | Storage class name |
| gitea.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| giteaOauth2Setup.image.pullPolicy | string | `"IfNotPresent"` |  |
| giteaOauth2Setup.image.repository | string | `"konstellation/kdl-gitea-oauth2-setup"` |  |
| giteaOauth2Setup.image.tag | string | `"0.16.0"` |  |
| global.domain | string | `"kdl.local"` | The DNS domain name that will serve the application |
| global.ingress | object | `{"tls":{"caSecret":{},"enabled":true,"secretName":null}}` | Ingress configuration |
| global.ingress.tls.caSecret | object | `{}` | A secret containing the the CA cert is needed in order to use a self-signed certificate. Check [values.yaml](./values.yaml) for usage details. |
| global.ingress.tls.enabled | bool | `true` | Whether to enable TLS |
| global.ingress.tls.secretName | string | If not defined, for each chart component that uses an ingress, an autogenerated secret name based on the `.Values.global.domain` and the component name will be used. Example: for gitea `kdl.local-gitea-tls` will be used | The name of the TLS secret to use for all ingresses. Specific component ingress secret names take precedence over this. |
| global.mongodb | DEPRECATION | `{"connectionString":{"secretKey":"","secretName":""}}` | Configure MongoDB string URI |
| global.mongodb.connectionString.secretKey | string | `""` | The name of the secret key that contains the MongoDB connection string. |
| global.mongodb.connectionString.secretName | string | `""` | The name of the secret that contains a key with the MongoDB connection string. |
| global.serverName | string | `"local-server"` | KDL Server instance name |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-server","tag":"1.38.0"}` | Image registry The image configuration for the base service |
| imagePullSecrets | list | `[]` | Docker registry secret names as an array |
| ingress | object | `{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| keycloak | object | `{"command":[],"enabled":true,"env":{},"fullnameOverride":"keycloak","image":{"repository":"keycloak/keycloak","tag":"26.0"},"ingress":{"annotations":{},"className":"","enabled":true,"hosts":[{"host":"keycloak.mydomain.com","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]},"livenessProbe":{"enabled":true},"readinessProbe":{"enabled":true,"httpGet":{"path":"/realms/master"}},"service":{"healthPath":"/realms/master","targetPort":8080},"serviceAccount":{"create":true}}` | Keycloak subchart deployment </br> Ref: https://github.com/konstellation-io/helm-charts/blob/konstellation-base-1.0.2/charts/konstellation-base/values.yaml |
| keycloak.enabled | bool | `true` | Enable or disable Keycloak subchart |
| knowledgeGalaxy | object | `{"affinity":{},"args":[],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"config":{"descriptionMinWords":50,"logLevel":"INFO","numberOfOutputs":1000,"workers":1},"enabled":false,"env":{},"envFromConfigMap":{},"envFromFiles":{},"envFromSecrets":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"},"initContainers":[],"lifecycle":{},"livenessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"livenessProbeCustom":{},"networkPolicy":{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"readinessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"readinessProbeCustom":{},"resources":{},"secrets":{},"securityContext":{},"service":{"port":80,"targetPort":8080,"type":"ClusterIP"},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"startupProbe":{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"startupProbeCustom":{},"terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[],"volumeMounts":[],"volumes":[]}` | knowledge-galaxy deployment </br> Ref: https://github.com/konstellation-io/knowledge-galaxy |
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
| knowledgeGalaxy.envFromFiles | object | `{}` | Variables from files managed by you </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| knowledgeGalaxy.envFromSecrets | object | `{}` | Variables from secrets |
| knowledgeGalaxy.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"}` | Image registry The image configuration for the base service |
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
| knowledgeGalaxy.secrets | object | `{}` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
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
| minio | object | `{"auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":true,"mode":"standalone","persistence":{"enabled":false}}` | MinIO subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/minio/values.yaml TODO: pending to remove legacy minio |
| minio.enabled | bool | `true` | Enable or disable MinIO subchart |
| mongodb | object | `{"architecture":"standalone","auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":false,"persistence":{"enabled":false}}` | MongoDB subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml |
| mongodb.enabled | bool | `false` | Enable or disable MongoDB subchart |
| nameOverride | string | `""` | String to partially override konstellation-base.fullname template (will maintain the release name) |
| networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| networkPolicy.policyTypes | list | `[]` | Policy types |
| nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| oauth2Proxy.config.cookieSecret | string | `"mycookiesecret16"` | The seed string for secure cookies. Ref: https://oauth2-proxy.github.io/oauth2-proxy/docs/configuration/overview |
| oauth2Proxy.customConfig | string | `nil` | The OAuth2-Proxy custom configuration file |
| oauth2Proxy.image.pullPolicy | string | `"IfNotPresent"` |  |
| oauth2Proxy.image.repository | string | `"quay.io/oauth2-proxy/oauth2-proxy"` |  |
| oauth2Proxy.image.tag | string | `"v7.0.1-amd64"` |  |
| oauth2proxy | object | `{"clientID":"XXXXXXX","clientSecret":"XXXXXXXX","cookieName":"","cookieSecret":"XXXXXXXX","enabled":false,"extraContainers":[],"extraObjects":[],"extraVolumeMounts":[],"extraVolumes":[],"httpScheme":"http"}` | OAuth2-Proxy subchart deployment </br> Ref: https://github.com/oauth2-proxy/manifests/blob/main/helm/oauth2-proxy/values.yaml |
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
| postgres.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| postgres.dbName | string | `"gitea"` | The name of the Postgres database for Gitea |
| postgres.dbPassword | string | `"test"` | The password for the Gitea's database |
| postgres.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| postgres.image.repository | string | `"postgres"` | The image repository |
| postgres.image.tag | float | `12.1` | The image tag |
| postgres.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| postgres.storage.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| postgres.storage.storageClassName | string | `""` | Storage class to use for persistence |
| postgres.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| postgresql | object | `{"auth":{"database":"kdl","password":"ChangeMe","username":"user"},"enabled":true,"primary":{"persistence":{"enabled":false}},"replicaCount":1}` | PostgreSQL subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/postgresql/values.yaml |
| postgresql.enabled | bool | `true` | Enable or disable PostgreSQL subchart |
| projectOperator | object | `{"affinity":{},"filebrowser":{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"filebrowser/filebrowser","tag":"v2"},"nodeSelector":{},"tolerations":[]},"kubeRbacProxy":{"image":{"pullPolicy":"IfNotPresent","repository":"gcr.io/kubebuilder/kube-rbac-proxy","tag":"v0.8.0"}},"manager":{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-project-operator","tag":"0.19.0"}},"mlflow":{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-mlflow","tag":"v0.13.5"},"ingress":{"annotations":{},"className":"nginx","tls":{"secretName":null}},"nodeSelector":{},"tolerations":[],"volume":{"size":"1Gi","storageClassName":"standard"}},"nodeSelector":{},"resources":{},"tolerations":[]}` | project-operator operator |
| projectOperator.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| projectOperator.filebrowser | object | `{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"filebrowser/filebrowser","tag":"v2"},"nodeSelector":{},"tolerations":[]}` | filebrowser configuration |
| projectOperator.filebrowser.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| projectOperator.filebrowser.image | object | `{"pullPolicy":"IfNotPresent","repository":"filebrowser/filebrowser","tag":"v2"}` | Image registry The image configuration for the base service |
| projectOperator.filebrowser.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| projectOperator.filebrowser.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| projectOperator.kubeRbacProxy | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"gcr.io/kubebuilder/kube-rbac-proxy","tag":"v0.8.0"}}` | kube-rbac-proxy container specs |
| projectOperator.kubeRbacProxy.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| projectOperator.kubeRbacProxy.image.repository | string | `"gcr.io/kubebuilder/kube-rbac-proxy"` | Image repository |
| projectOperator.kubeRbacProxy.image.tag | string | `"v0.8.0"` | Image tag |
| projectOperator.manager | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-project-operator","tag":"0.19.0"}}` | The image configuration for the manager container |
| projectOperator.manager.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-project-operator","tag":"0.19.0"}` | Image registry The image configuration for the base service |
| projectOperator.mlflow | object | `{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-mlflow","tag":"v0.13.5"},"ingress":{"annotations":{},"className":"nginx","tls":{"secretName":null}},"nodeSelector":{},"tolerations":[],"volume":{"size":"1Gi","storageClassName":"standard"}}` | mlflow configuration |
| projectOperator.mlflow.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| projectOperator.mlflow.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-mlflow","tag":"v0.13.5"}` | Image registry The image configuration for the base service |
| projectOperator.mlflow.ingress | object | `{"annotations":{},"className":"nginx","tls":{"secretName":null}}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| projectOperator.mlflow.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| projectOperator.mlflow.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| projectOperator.mlflow.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| projectOperator.mlflow.volume.size | string | `"1Gi"` | The storage size for the persistent volume claim |
| projectOperator.mlflow.volume.storageClassName | string | `"standard"` | Storage class to use for persistence |
| projectOperator.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| projectOperator.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| projectOperator.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| readinessProbeCustom | object | `{}` | Custom readinessProbe |
| resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| secrets | object | `{}` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| service | object | `{"port":80,"targetPort":3000,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| service.port | int | `80` | Kubernetes Service port |
| service.targetPort | int | `3000` | Pod expose port |
| service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| serviceMonitor.enabled | bool | `false` | Enable or disable |
| sharedVolume | object | `{"accessModes":["ReadWriteMany"],"annotations":{},"enabled":false,"labels":{},"selector":{},"size":"10Gi","storageClass":"","volumeBindingMode":"","volumeName":""}` | Persistent Volume configuration Mount volume to share data between components </br> Ref: https://kubernetes.io/docs/concepts/storage/persistent-volumes/ |
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
| testConnection | object | `{"enabled":false,"repository":"busybox","tag":""}` | Enable or disable test connection |
| tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| userToolsOperator | object | `{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-user-tools-operator","tag":"0.29.0"},"ingress":{"annotations":{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n","nginx.ingress.kubernetes.io/proxy-body-size":"1000000m"},"className":"nginx","enabled":false,"tls":{"secretName":null}},"kubeconfig":{"enabled":false,"externalServerUrl":""},"nodeSelector":{},"oauth2Proxy":{"image":{"pullPolicy":"IfNotPresent","repository":"quay.io/oauth2-proxy/oauth2-proxy","tag":"v7.0.1-amd64"}},"repoCloner":{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-repo-cloner","tag":"0.18.0"}},"storage":{"size":"10Gi","storageClassName":"standard"},"tolerations":[],"vscode":{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-vscode","tag":"v0.15.0"}},"vscodeRuntime":{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-py","tag":"3.9"}}}` | User Tools Operator deployment |
| userToolsOperator.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| userToolsOperator.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-user-tools-operator","tag":"0.29.0"}` | Image registry The image configuration for the base service |
| userToolsOperator.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| userToolsOperator.kubeconfig.enabled | bool | `false` | Whether to enable kubeconfig for using with VSCode remote development. |
| userToolsOperator.kubeconfig.externalServerUrl | string | `""` | The Kube API Server URL for using with VSCode remote development |
| userToolsOperator.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| userToolsOperator.oauth2Proxy | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"quay.io/oauth2-proxy/oauth2-proxy","tag":"v7.0.1-amd64"}}` | oauth2-proxy configuration |
| userToolsOperator.repoCloner | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-repo-cloner","tag":"0.18.0"}}` | repocloner configuration The following components are managed by the manager container when `usertool` custom resources are detected |
| userToolsOperator.repoCloner.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-repo-cloner","tag":"0.18.0"}` | Image registry The image configuration for the base service |
| userToolsOperator.storage | object | `{"size":"10Gi","storageClassName":"standard"}` | Storage configuration |
| userToolsOperator.storage.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| userToolsOperator.storage.storageClassName | string | `"standard"` | Storage class to use for persistence |
| userToolsOperator.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| userToolsOperator.vscode | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-vscode","tag":"v0.15.0"}}` | vscode configuration |
| userToolsOperator.vscode.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-vscode","tag":"v0.15.0"}` | Image registry The image configuration for the base service |
| userToolsOperator.vscodeRuntime | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-py","tag":"3.9"}}` | vscodeRuntime configuration |
| userToolsOperator.vscodeRuntime.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-py","tag":"3.9"}` | Image registry The image configuration for the base service |
| volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
