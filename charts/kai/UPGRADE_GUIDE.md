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
