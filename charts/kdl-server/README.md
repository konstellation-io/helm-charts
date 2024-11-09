# kdl-server

A Helm chart to deploy KDL server

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| ialejandro | <ivan.alejandro@intelygenz.com> |  |

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
| backup | object | `{"activeDeadlineSeconds":3600,"backoffLimit":3,"concurrencyPolicy":"Forbid","enabled":false,"extraVolumeMounts":[],"extraVolumes":[],"failedJobsHistoryLimit":1,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-backup","tag":"0.23.0"},"name":"backup-gitea","resources":{"limits":{"cpu":"100m","memory":"256Mi"},"requests":{"cpu":"100m","memory":"100Mi"}},"s3":{"awsAccessKeyID":"aws-access-key-id","awsSecretAccessKey":"aws-secret-access-key","bucketName":"s3-bucket-name"},"schedule":"0 1 * * 0","startingDeadlineSeconds":60,"successfulJobsHistoryLimit":0,"ttlSecondsAfterFinished":""}` | Backup job configuration |
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
| backup.resources | object | `{"limits":{"cpu":"100m","memory":"256Mi"},"requests":{"cpu":"100m","memory":"100Mi"}}` | Resource requests and limits for backup container. Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
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
| drone | object | `{"adminToken":"7GSipOV0wJZQioZNBxaw3AotHV1tA4K4","affinity":{},"enabled":true,"image":{"pullPolicy":"IfNotPresent","repository":"drone/drone","tag":"1.10.1"},"ingress":{"annotations":{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n","nginx.ingress.kubernetes.io/proxy-body-size":"100m"},"className":"nginx","tls":{"secretName":null}},"nodeSelector":{},"rpcSecret":"runner-shared-secret","runnerCapacity":5,"storage":{"size":"10Gi","storageClassName":"standard"},"tolerations":[]}` | Drone deployment (DEPRECATION) Remove in future versions |
| drone.adminToken | string | `"7GSipOV0wJZQioZNBxaw3AotHV1tA4K4"` | Drone Server admin token |
| drone.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| drone.enabled | bool | `true` | Whether to enable Drone |
| drone.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n","nginx.ingress.kubernetes.io/proxy-body-size":"100m"}` | Ingress annotations |
| drone.ingress.className | string | `"nginx"` | The ingress class name |
| drone.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| drone.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| drone.rpcSecret | string | `"runner-shared-secret"` | Drone RPC secret for allowing Drone runners to authentiticate the RPC connection to the server |
| drone.runnerCapacity | int | `5` | The max number of concurrent jobs that a Drone runner can run |
| drone.storage.size | string | `"10Gi"` | Storage size |
| drone.storage.storageClassName | string | `"standard"` | The Storage ClassName |
| drone.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| droneAuthorizer | object | `{"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-drone-authorizer","tag":"0.16.0"}}` | Drone Authorizer deployment (DEPRECATION) Remove in future versions |
| droneRunner | object | `{"affinity":{},"debug":"true","droneRunnerEnviron":"","image":{"pullPolicy":"IfNotPresent","repository":"drone/drone-runner-kube","tag":"1.0.0-beta.6"},"nodeSelector":{},"pluginSecret":"my-secret","serviceAccountJob":{"annotations":{},"create":false,"enabled":false,"name":"drone-runner-job"},"tolerations":[],"trace":"true"}` | Drone Runner deployment (DEPRECATION) Remove in future versions |
| droneRunner.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| droneRunner.debug | string | `"true"` | Sets DRONE_DEBUG environment variable |
| droneRunner.droneRunnerEnviron | string | `""` | Configures the DRONE_RUNNER_ENVIRON environment variable. Ref: https://docs.drone.io/runner/kubernetes/configuration/reference/drone-runner-environ/ |
| droneRunner.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| droneRunner.pluginSecret | string | `"my-secret"` | Provides the secret token used to authenticate http requests to the Kubernetes Secrets Extension endpoint |
| droneRunner.serviceAccountJob.annotations | object | `{}` | If `.Values.droneRunner.serviceAccountJob.create` is set to `true`, sets annotations to the service account |
| droneRunner.serviceAccountJob.create | bool | `false` | If `.Values.droneRunner.serviceAccountJob.enabled` is set to `true`, creates the service account |
| droneRunner.serviceAccountJob.enabled | bool | `false` | Whether to enable the service account for Drone job pods |
| droneRunner.serviceAccountJob.name | string | `"drone-runner-job"` | The name of the Drone job service account |
| droneRunner.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| droneRunner.trace | string | `"true"` | Sets DRONE_TRACE environment variable |
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
| kdlServer | object | `{"affinity":{},"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-server","tag":"1.38.0"},"ingress":{"annotations":{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m","nginx.ingress.kubernetes.io/proxy-connect-timeout":"3600","nginx.ingress.kubernetes.io/proxy-read-timeout":"3600","nginx.ingress.kubernetes.io/proxy-send-timeout":"3600"},"className":"","enabled":false,"tls":{"secretName":null}},"nodeSelector":{},"tolerations":[]}` | KDL Server deployment |
| kdlServer.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| kdlServer.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kdl-server","tag":"1.38.0"}` | Image registry The image configuration for the base service |
| kdlServer.ingress | object | `{"annotations":{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m","nginx.ingress.kubernetes.io/proxy-connect-timeout":"3600","nginx.ingress.kubernetes.io/proxy-read-timeout":"3600","nginx.ingress.kubernetes.io/proxy-send-timeout":"3600"},"className":"","enabled":false,"tls":{"secretName":null}}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| kdlServer.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m","nginx.ingress.kubernetes.io/proxy-connect-timeout":"3600","nginx.ingress.kubernetes.io/proxy-read-timeout":"3600","nginx.ingress.kubernetes.io/proxy-send-timeout":"3600"}` | Ingress annotations |
| kdlServer.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| kdlServer.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| kdlServer.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| keycloak | object | `{"command":[],"enabled":true,"env":{},"image":{"repository":"keycloak/keycloak","tag":"26.0"},"ingress":{"annotations":{},"className":"","enabled":true,"hosts":[{"host":"keycloak.mydomain.com","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]},"livenessProbe":{"enabled":true},"readinessProbe":{"enabled":true,"httpGet":{"path":"/realms/master"}},"service":{"healthPath":"/realms/master","targetPort":8080},"serviceAccount":{"create":true}}` | Keycloak subchart deployment </br> Ref: https://github.com/konstellation-io/helm-charts/blob/konstellation-base-1.0.2/charts/konstellation-base/values.yaml |
| keycloak.enabled | bool | `true` | Enable or disable Keycloak subchart |
| knowledgeGalaxy | object | `{"affinity":{},"config":{"descriptionMinWords":50,"logLevel":"INFO","numberOfOutputs":1000,"workers":1},"enabled":false,"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"},"nodeSelector":{},"serviceaccount":{"annotations":{},"create":true,"imagePullSecrets":[],"name":""},"tolerations":[]}` | knowledge-galaxy deployment </br> Ref: https://github.com/konstellation-io/knowledge-galaxy |
| knowledgeGalaxy.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| knowledgeGalaxy.config | object | `{"descriptionMinWords":50,"logLevel":"INFO","numberOfOutputs":1000,"workers":1}` | Configuration |
| knowledgeGalaxy.config.descriptionMinWords | int | `50` | Minimum number of words to use for project description |
| knowledgeGalaxy.config.logLevel | string | `"INFO"` | Log level |
| knowledgeGalaxy.config.numberOfOutputs | int | `1000` | Number of outputs that the recommender returns |
| knowledgeGalaxy.config.workers | int | `1` | Number of threads for the server |
| knowledgeGalaxy.enabled | bool | `false` | Whether to enable Knowledge Galaxy |
| knowledgeGalaxy.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/knowledge-galaxy","tag":"v1.2.1"}` | Image registry The image configuration for the base service |
| knowledgeGalaxy.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| knowledgeGalaxy.serviceaccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[],"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| knowledgeGalaxy.serviceaccount.imagePullSecrets | list | `[]` | Reference to one or more secrets to be used when pulling images. Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/ |
| knowledgeGalaxy.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| minio | object | `{"auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":true,"mode":"standalone","persistence":{"enabled":false}}` | MinIO subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/minio/values.yaml |
| minio.enabled | bool | `true` | Enable or disable MinIO subchart |
| mongodb | object | `{"architecture":"standalone","auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":true,"persistence":{"enabled":false}}` | MongoDB subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml |
| mongodb.enabled | bool | `true` | Enable or disable MongoDB subchart |
| oauth2Proxy.config.cookieSecret | string | `"mycookiesecret16"` | The seed string for secure cookies. Ref: https://oauth2-proxy.github.io/oauth2-proxy/docs/configuration/overview |
| oauth2Proxy.customConfig | string | `nil` | The OAuth2-Proxy custom configuration file |
| oauth2Proxy.image.pullPolicy | string | `"IfNotPresent"` |  |
| oauth2Proxy.image.repository | string | `"quay.io/oauth2-proxy/oauth2-proxy"` |  |
| oauth2Proxy.image.tag | string | `"v7.0.1-amd64"` |  |
| oauth2proxy | object | `{"clientID":"XXXXXXX","clientSecret":"XXXXXXXX","cookieName":"","cookieSecret":"XXXXXXXX","enabled":false}` | OAuth2-Proxy subchart deployment </br> Ref: https://github.com/oauth2-proxy/manifests/blob/main/helm/oauth2-proxy/values.yaml |
| oauth2proxy.enabled | bool | `false` | Enable or disable OAuth2-Proxy subchart |
| postgres.dbName | string | `"gitea"` |  |
| postgres.dbPassword | string | `"gitea"` |  |
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
| sharedVolume | object | `{"name":"received-data","size":"10Gi","storageClassName":"standard"}` | Mount volume to share data between components |
| sharedVolume.name | string | `"received-data"` | The name of the shared volume |
| sharedVolume.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| sharedVolume.storageClassName | string | `"standard"` | Storage class to use for persistence |
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
