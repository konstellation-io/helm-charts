# kdl-server

A Helm chart to deploy KDL server

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| ialejandro | <ivan.alejandro@intelygenz.com> |  |

## Prerequisites

* Helm 3+

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.min.io/ | minio | 3.2.0 |

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

## Uninstall Helm chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

> [!IMPORTANT]
> Upgrading an existing Release to a new major version (`v0.15.X` -> `v1.0.0`) indicates that there is an incompatible **BREAKING CHANGES** needing manual actions.

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
| backup.s3.awsAccessKeyID | string | `"aws-access-key-id"` | AWS Access Key ID for acceding backup bucket |
| backup.s3.awsSecretAccessKey | string | `"aws-secret-access-key"` | AWS Secret Access Key for acceding backup bucket |
| backup.s3.bucketName | string | `"s3-bucket-name"` | The S3 bucket that will store all backups |
| backup.schedule | string | `"0 1 * * 0"` | Backup cronjob schedule |
| backup.startingDeadlineSeconds | int | `60` | Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#job-creation |
| backup.successfulJobsHistoryLimit | int | `0` | The number of successful finished jobs to retain. Ref: https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#jobs-history-limits |
| backup.ttlSecondsAfterFinished | string | `""` | Limits the lifetime of a Job that has finished execution (either Complete or Failed). |
| cleaner.enabled | bool | `false` | Whether to enable cleaner cronjob |
| cleaner.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| cleaner.image.repository | string | `"konstellation/cleaner"` | The image repository |
| cleaner.image.tag | string | `"0.15.0"` | The image tag |
| cleaner.schedule | string | `"0 1 * * 0"` | Celaner cronjob schedule |
| cleaner.threshold | int | `5` | The minimun age of files to be removed |
| cleaner.trashPath | string | `"/shared-storage/.trash"` | The name of the trash path |
| drone.adminToken | string | `"7GSipOV0wJZQioZNBxaw3AotHV1tA4K4"` | Drone Server admin token |
| drone.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| drone.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| drone.image.repository | string | `"drone/drone"` | The image repository |
| drone.image.tag | string | `"1.10.1"` | The image tag |
| drone.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n","nginx.ingress.kubernetes.io/proxy-body-size":"100m"}` | Ingress annotations |
| drone.ingress.className | string | `"nginx"` | The ingress class name |
| drone.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| drone.nodeSelector | object | `{}` |  |
| drone.rpcSecret | string | `"runner-shared-secret"` | Drone RPC secret for allowing Drone runners to authentiticate the RPC connection to the server |
| drone.runnerCapacity | int | `5` | The max number of concurrent jobs that a Drone runner can run |
| drone.storage.size | string | `"10Gi"` | Storage size |
| drone.storage.storageClassName | string | `"standard"` | The Storage ClassName |
| drone.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| droneAuthorizer.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| droneAuthorizer.image.repository | string | `"konstellation/drone-authorizer"` | The image repository |
| droneAuthorizer.image.tag | string | `"0.16.0"` | The image tag |
| droneRunner.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| droneRunner.debug | string | `"true"` | Sets DRONE_DEBUG environment variable |
| droneRunner.droneRunnerEnviron | string | `""` | Configures the DRONE_RUNNER_ENVIRON environment variable. Ref: https://docs.drone.io/runner/kubernetes/configuration/reference/drone-runner-environ/ |
| droneRunner.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| droneRunner.image.repository | string | `"drone/drone-runner-kube"` | The image repository |
| droneRunner.image.tag | string | `"1.0.0-beta.6"` | The image tag |
| droneRunner.nodeSelector | object | `{}` |  |
| droneRunner.pluginSecret | string | `"my-secret"` | Provides the secret token used to authenticate http requests to the Kubernetes Secrets Extension endpoint |
| droneRunner.serviceAccountJob.annotations | object | `{}` | If `.Values.droneRunner.serviceAccountJob.create` is set to `true`, sets annotations to the service account |
| droneRunner.serviceAccountJob.create | bool | `false` | If `.Values.droneRunner.serviceAccountJob.enabled` is set to `true`, creates the service account |
| droneRunner.serviceAccountJob.enabled | bool | `false` | Whether to enable the service account for Drone job pods |
| droneRunner.serviceAccountJob.name | string | `"drone-runner-job"` | The name of the Drone job service account |
| droneRunner.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| droneRunner.trace | string | `"true"` | Sets DRONE_TRACE environment variable |
| gitea.admin.email | string | `"test@test.com"` | Admin user email |
| gitea.admin.password | string | `"123456"` | Admin password |
| gitea.admin.username | string | `"kdladmin"` | Admin username |
| gitea.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| gitea.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| gitea.image.repository | string | `"gitea/gitea"` | The image repository |
| gitea.image.tag | string | `"1.14.4"` | The image tag |
| gitea.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n"}` | Ingress annotations |
| gitea.ingress.className | string | `"nginx"` | The ingress class name |
| gitea.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| gitea.nodeSelector | object | `{}` |  |
| gitea.storage.size | string | `"10Gi"` | Storage size |
| gitea.storage.storageClassName | string | `"standard"` | Storage class name |
| gitea.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| giteaOauth2Setup.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| giteaOauth2Setup.image.repository | string | `"konstellation/gitea-oauth2-setup"` | The image repository |
| giteaOauth2Setup.image.tag | string | `"0.17.0"` | The image tag |
| global | object | `{"domain":"kdl.local","ingress":{"tls":{"caSecret":{},"enabled":true,"secretName":null}},"mongodb":{"connectionString":{"secretKey":"","secretName":""}},"serverName":"local-server"}` | Global section contains configuration options that are applied to all services |
| global.domain | string | `"kdl.local"` | The DNS domain name that will serve the application |
| global.ingress.tls.caSecret | object | `{}` | A secret containing the the CA cert is needed in order to use a self-signed certificate. Check [values.yaml](./values.yaml) for usage details. |
| global.ingress.tls.enabled | bool | `true` | Whether to enable TLS |
| global.ingress.tls.secretName | string | If not defined, for each chart component that uses an ingress, an autogenerated secret name based on the `.Values.global.domain` and the component name will be used. Example: for gitea `kdl.local-gitea-tls` will be used | The name of the TLS secret to use for all ingresses. Specific component ingress secret names take precedence over this. |
| global.mongodb.connectionString.secretKey | string | `""` | The name of the secret key that contains the MongoDB connection string. |
| global.mongodb.connectionString.secretName | string | `""` | The name of the secret that contains a key with the MongoDB connection string. |
| global.serverName | string | `"local-server"` | KDL Server instance name |
| kdlServer.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| kdlServer.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| kdlServer.image.repository | string | `"konstellation/kdl-server"` | The image repository |
| kdlServer.image.tag | string | `"1.38.0"` | The image tag |
| kdlServer.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m","nginx.ingress.kubernetes.io/proxy-connect-timeout":"3600","nginx.ingress.kubernetes.io/proxy-read-timeout":"3600","nginx.ingress.kubernetes.io/proxy-send-timeout":"3600"}` | Ingress annotations |
| kdlServer.ingress.className | string | `"nginx"` | The ingress class name |
| kdlServer.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| kdlServer.nodeSelector | object | `{}` |  |
| kdlServer.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| knowledgeGalaxy.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| knowledgeGalaxy.config.descriptionMinWords | int | `50` | Minimum number of words to use for project description |
| knowledgeGalaxy.config.logLevel | string | `"INFO"` | Log level |
| knowledgeGalaxy.config.numberOfOutputs | int | `1000` | Number of outputs that the recommender returns |
| knowledgeGalaxy.config.workers | int | `1` | Number of threads for the server |
| knowledgeGalaxy.enabled | bool | `false` | Whether to enable Knowledge Galaxy |
| knowledgeGalaxy.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| knowledgeGalaxy.image.repository | string | `"konstellation/knowledge-galaxy"` | The image repository |
| knowledgeGalaxy.image.tag | string | `"v1.2.1"` | The image tag |
| knowledgeGalaxy.nodeSelector | object | `{}` |  |
| knowledgeGalaxy.serviceaccount.annotations | object | `{}` | The service account annotations |
| knowledgeGalaxy.serviceaccount.enabled | bool | `true` | Whether to create a service account |
| knowledgeGalaxy.serviceaccount.imagePullSecrets | list | `[]` | Reference to one or more secrets to be used when pulling images. Ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/ |
| knowledgeGalaxy.serviceaccount.name | string | knowledge-galaxy | The name of the service account to use |
| knowledgeGalaxy.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| minio | object | Check [values.yaml](./values.yaml) | MinIO chart's values. Check MinIO chart's [documentation](https://github.com/minio/minio/tree/master/helm/minio) for more info about values |
| minio.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| minio.consoleIngress.annotations | object | `{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m"}` | Ingress annotations |
| minio.consoleIngress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| minio.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/proxy-body-size":"1000000m"}` | Ingress annotations |
| minio.ingress.className | string | `"nginx"` | The ingress class name |
| minio.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| minio.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| minio.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| oauth2Proxy.config.cookieSecret | string | `"mycookiesecret16"` | The seed string for secure cookies. Ref: https://oauth2-proxy.github.io/oauth2-proxy/docs/configuration/overview |
| oauth2Proxy.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| oauth2Proxy.image.repository | string | `"quay.io/oauth2-proxy/oauth2-proxy"` | The image repository |
| oauth2Proxy.image.tag | string | `"v7.0.1-amd64"` | The image tag |
| postgres.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| postgres.dbName | string | `"gitea"` | The name of the Postgres database for Gitea |
| postgres.dbPassword | string | `"test"` | The password for the Gitea's database |
| postgres.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| postgres.image.repository | string | `"postgres"` | The image repository |
| postgres.image.tag | float | `12.1` | The image tag |
| postgres.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| postgres.storage.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| postgres.storage.storageClassName | string | `"standard"` | Storage class to use for persistence |
| postgres.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| projectOperator.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| projectOperator.filebrowser.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| projectOperator.filebrowser.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| projectOperator.filebrowser.image.repository | string | `"filebrowser/filebrowser"` | The image repository |
| projectOperator.filebrowser.image.tag | string | `"v2"` | The image tag |
| projectOperator.filebrowser.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| projectOperator.filebrowser.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| projectOperator.kubeRbacProxy.image.pullPolicy | string | `"IfNotPresent"` | Image pull policy |
| projectOperator.kubeRbacProxy.image.repository | string | `"gcr.io/kubebuilder/kube-rbac-proxy"` | Image repository |
| projectOperator.kubeRbacProxy.image.tag | string | `"v0.8.0"` | Image tag |
| projectOperator.manager.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| projectOperator.manager.image.repository | string | `"konstellation/project-operator"` | The image repository |
| projectOperator.manager.image.tag | string | `"0.19.0"` | The image tag |
| projectOperator.manager.resources | object | `{}` | Resource requests and limits for primary projectOperator container. Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| projectOperator.mlflow.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| projectOperator.mlflow.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| projectOperator.mlflow.image.repository | string | `"konstellation/mlflow"` | The image repository |
| projectOperator.mlflow.image.tag | string | `"v0.13.5"` | The image tag |
| projectOperator.mlflow.ingress.annotations | object | `{}` | Ingress annotations |
| projectOperator.mlflow.ingress.className | string | `"nginx"` | The ingress class name |
| projectOperator.mlflow.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| projectOperator.mlflow.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| projectOperator.mlflow.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| projectOperator.mlflow.volume.size | string | `"1Gi"` | The storage size for the persistent volume claim |
| projectOperator.mlflow.volume.storageClassName | string | `"standard"` | Storage class to use for persistence |
| projectOperator.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| projectOperator.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| sharedVolume.name | string | `"received-data"` | The name of the shared volume |
| sharedVolume.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| sharedVolume.storageClassName | string | `"standard"` | Storage class to use for persistence |
| userToolsOperator.affinity | object | `{}` | Assign custom affinity rules. Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ |
| userToolsOperator.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| userToolsOperator.image.repository | string | `"konstellation/user-tools-operator"` | The image repository |
| userToolsOperator.image.tag | string | `"0.29.0"` | The image tag |
| userToolsOperator.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/configuration-snippet":"more_set_headers \"Content-Security-Policy: frame-ancestors 'self' *\";\n","nginx.ingress.kubernetes.io/proxy-body-size":"1000000m"}` | Ingress annotations |
| userToolsOperator.ingress.className | string | `"nginx"` | The ingress class name |
| userToolsOperator.ingress.tls.secretName | string | `nil` | The TLS secret name that will be used. It takes precedence over `.Values.global.ingress.tls.secretName`. |
| userToolsOperator.kubeconfig.enabled | bool | `false` | Whether to enable kubeconfig for using with VSCode remote development. Ref: https://code.visualstudio.com/docs/remote/remote-overview |
| userToolsOperator.kubeconfig.externalServerUrl | string | `""` | The Kube API Server URL for using with VSCode remote development |
| userToolsOperator.nodeSelector | object | `{}` | Define which Nodes the Pods are scheduled on. Ref: https://kubernetes.io/docs/user-guide/node-selection/ |
| userToolsOperator.oauth2Proxy.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| userToolsOperator.oauth2Proxy.image.repository | string | `"quay.io/oauth2-proxy/oauth2-proxy"` | The image repository |
| userToolsOperator.oauth2Proxy.image.tag | string | `"v7.0.1-amd64"` | The image tag |
| userToolsOperator.repoCloner.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| userToolsOperator.repoCloner.image.repository | string | `"konstellation/repo-cloner"` | The image repository |
| userToolsOperator.repoCloner.image.tag | string | `"0.17.0"` | The image tag |
| userToolsOperator.storage.size | string | `"10Gi"` | The storage size for the persistent volume claim |
| userToolsOperator.storage.storageClassName | string | `"standard"` | Storage class to use for persistence |
| userToolsOperator.tolerations | list | `[]` | If specified, the pod's tolerations. Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/ |
| userToolsOperator.vscode.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| userToolsOperator.vscode.image.repository | string | `"konstellation/vscode"` | The image repository |
| userToolsOperator.vscode.image.tag | string | `"v0.15.0"` | The image tag |
| userToolsOperator.vscodeRuntime.image.pullPolicy | string | `"IfNotPresent"` | The image pull policy |
| userToolsOperator.vscodeRuntime.image.repository | string | `"konstellation/kdl-py"` | The image repository |
| userToolsOperator.vscodeRuntime.image.tag | string | `"3.9"` | The image tag |
