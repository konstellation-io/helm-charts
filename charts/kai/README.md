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
| 1.0.0                    | ❌   | ❌   | ✅   | ✅   | ✅   | ✅   | ✅   |

| Symbol | Description |
|:------:|-------------|
| ✅     | Perfect match: all features are supported. Client and server versions have exactly the same features/APIs. |
| 🟠     | Forward compatibility: the client will work with the server, but not all new server features are supported. The server has features that the client library cannot use. |
| ❌     | Backward compatibility/Not applicable: the client has features that may not be present in the server. Common features will work, but some client APIs might not be available in the server. |
| -      | Not tested: this combination has not been verified or is not applicable. |

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.min.io/ | minio | 5.4.0 |
| https://helm.releases.hashicorp.com | vault | 0.29.1 |
| https://nats-io.github.io/k8s/helm/charts | nats | 1.2.10 |
| oci://ghcr.io/konstellation-io/helm-charts | keycloak(konstellation-base) | 1.1.2 |
| oci://registry-1.docker.io/bitnamicharts | mongodb | 16.2.1 |
| oci://registry-1.docker.io/bitnamicharts | postgresql | 15.5.38 |
| oci://registry-1.docker.io/bitnamicharts | redis | 20.7.0 |

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

## Uninstall Helm chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

> [!IMPORTANT]
> Upgrading an existing Release to a new major version (`v0.15.X` -> `v1.0.0`) indicates that there is an incompatible **BREAKING CHANGES** needing manual actions.

## Configuration

See [Customizing the chart before installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with comments:

```console
helm show values konstellation-io/kai
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| adminApi | object | `{"affinity":{},"args":[],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"deploymentStrategy":"Recreate","env":{},"envFromConfigMap":{},"envFromFiles":[],"envFromSecrets":{},"extraContainers":[],"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kai","tag":"0.3.0-develop.17"},"imagePullSecrets":[],"ingress":{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.kai.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]},"initContainers":[],"lifecycle":{},"livenessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"livenessProbeCustom":{},"networkPolicy":{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{"fsGroup":10001},"readinessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"readinessProbeCustom":{},"readyChecker":{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5},"replicaCount":1,"resources":{},"secrets":[],"securityContext":{},"service":{"port":8080,"type":"ClusterIP"},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"serviceMonitor":{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"},"startupProbe":{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"startupProbeCustom":{},"terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[],"volumeMounts":[],"volumes":[]}` | adminAPI |
| adminApi.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| adminApi.args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| adminApi.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| adminApi.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| adminApi.deploymentStrategy | string | `"Recreate"` | Deployment strategy Specifies the strategy used to replace old Pods by new ones |
| adminApi.env | object | `{}` | Environment variables to configure application </br> Ref: https://github.com/konstellation-io/kdl-server/tree/main/app/api |
| adminApi.envFromConfigMap | object | `{}` | Variables from configMap |
| adminApi.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| adminApi.envFromSecrets | object | `{}` | Variables from secrets |
| adminApi.extraContainers | list | `[]` | Configure extra containers |
| adminApi.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kai","tag":"0.3.0-develop.17"}` | Image registry The image configuration for the base service |
| adminApi.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| adminApi.ingress | object | `{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.kai.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| adminApi.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| adminApi.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| adminApi.livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| adminApi.livenessProbeCustom | object | `{}` | Custom livenessProbe |
| adminApi.networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| adminApi.networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| adminApi.networkPolicy.policyTypes | list | `[]` | Policy types |
| adminApi.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| adminApi.podAnnotations | object | `{}` | Configure annotations on Pods |
| adminApi.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| adminApi.podLabels | object | `{}` | Configure labels on Pods |
| adminApi.podSecurityContext | object | `{"fsGroup":10001}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| adminApi.readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| adminApi.readinessProbeCustom | object | `{}` | Custom readinessProbe |
| adminApi.readyChecker | object | `{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5}` | Check if dependencies are ready |
| adminApi.readyChecker.enabled | bool | `false` | Enable or disable ready-checker |
| adminApi.readyChecker.pullPolicy | string | `"IfNotPresent"` | Pull policy for the image |
| adminApi.readyChecker.repository | string | `"busybox"` | Repository of the image |
| adminApi.readyChecker.retries | int | `30` | Number of retries before giving up |
| adminApi.readyChecker.services | list | `[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}]` | List services |
| adminApi.readyChecker.tag | string | `"latest"` | Overrides the image tag |
| adminApi.readyChecker.timeout | int | `5` | Timeout for each check |
| adminApi.replicaCount | int | `1` | Number of replicas Specifies the number of replicas for the service |
| adminApi.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| adminApi.secrets | list | `[]` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| adminApi.securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| adminApi.service | object | `{"port":8080,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| adminApi.service.port | int | `8080` | Kubernetes Service port |
| adminApi.service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| adminApi.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| adminApi.serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| adminApi.serviceMonitor.enabled | bool | `false` | Enable or disable |
| adminApi.startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| adminApi.startupProbeCustom | object | `{}` | Custom startupProbe |
| adminApi.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| adminApi.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| adminApi.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| adminApi.volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| adminApi.volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
| fullnameOverride | string | `""` | String to fully override kai.fullname template |
| global.env | object | `{}` | Environment variables to configure application |
| global.envFromConfigMap | object | `{}` | Variables from configMap |
| global.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| global.envFromSecrets | object | `{}` | Variables from secrets |
| global.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| global.imageRegistry | string | `""` | Specifies the registry to pull images from. Leave empty for the default registry |
| k8sManager | object | `{"affinity":{},"args":[],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"deploymentStrategy":"Recreate","env":{},"envFromConfigMap":{},"envFromFiles":[],"envFromSecrets":{},"extraContainers":[],"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kai-k8s-manager","tag":"0.3.0-develop.17"},"imagePullSecrets":[],"initContainers":[],"lifecycle":{},"livenessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"livenessProbeCustom":{},"networkPolicy":{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"rbac":{"create":true},"readinessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"readinessProbeCustom":{},"readyChecker":{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5},"resources":{},"secrets":[],"securityContext":{},"service":{"port":50051},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"serviceMonitor":{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"},"startupProbe":{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"startupProbeCustom":{},"terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[],"volumeMounts":[],"volumes":[]}` | K8S manager |
| k8sManager.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| k8sManager.args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| k8sManager.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| k8sManager.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| k8sManager.deploymentStrategy | string | `"Recreate"` | Deployment strategy Specifies the strategy used to replace old Pods by new ones |
| k8sManager.env | object | `{}` | Environment variables to configure application </br> Ref: https://github.com/konstellation-io/kdl-server/tree/main/app/api |
| k8sManager.envFromConfigMap | object | `{}` | Variables from configMap |
| k8sManager.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| k8sManager.envFromSecrets | object | `{}` | Variables from secrets |
| k8sManager.extraContainers | list | `[]` | Configure extra containers |
| k8sManager.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kai-k8s-manager","tag":"0.3.0-develop.17"}` | Image registry The image configuration for the base service |
| k8sManager.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| k8sManager.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| k8sManager.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| k8sManager.livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| k8sManager.livenessProbeCustom | object | `{}` | Custom livenessProbe |
| k8sManager.networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| k8sManager.networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| k8sManager.networkPolicy.policyTypes | list | `[]` | Policy types |
| k8sManager.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| k8sManager.podAnnotations | object | `{}` | Configure annotations on Pods |
| k8sManager.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| k8sManager.podLabels | object | `{}` | Configure labels on Pods |
| k8sManager.podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| k8sManager.rbac | object | `{"create":true}` | Creation of resources RBAC </br> Ref: https://kubernetes.io/docs/reference/access-authn-authz/rbac/ |
| k8sManager.readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| k8sManager.readinessProbeCustom | object | `{}` | Custom readinessProbe |
| k8sManager.readyChecker | object | `{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5}` | Check if dependencies are ready |
| k8sManager.readyChecker.enabled | bool | `false` | Enable or disable ready-checker |
| k8sManager.readyChecker.pullPolicy | string | `"IfNotPresent"` | Pull policy for the image |
| k8sManager.readyChecker.repository | string | `"busybox"` | Repository of the image |
| k8sManager.readyChecker.retries | int | `30` | Number of retries before giving up |
| k8sManager.readyChecker.services | list | `[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}]` | List services |
| k8sManager.readyChecker.tag | string | `"latest"` | Overrides the image tag |
| k8sManager.readyChecker.timeout | int | `5` | Timeout for each check |
| k8sManager.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| k8sManager.secrets | list | `[]` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| k8sManager.securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| k8sManager.service | object | `{"port":50051}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| k8sManager.service.port | int | `50051` | Kubernetes Service port |
| k8sManager.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| k8sManager.serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| k8sManager.serviceMonitor.enabled | bool | `false` | Enable or disable |
| k8sManager.startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| k8sManager.startupProbeCustom | object | `{}` | Custom startupProbe |
| k8sManager.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| k8sManager.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| k8sManager.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| k8sManager.volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| k8sManager.volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
| keycloak | object | `{"command":[],"enabled":false,"env":{},"image":{"repository":"keycloak/keycloak","tag":"26.0"},"ingress":{"annotations":{},"className":"","enabled":true,"hosts":[{"host":"keycloak.kai.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}]},"livenessProbe":{"enabled":true},"readinessProbe":{"enabled":true,"httpGet":{"path":"/realms/master"}},"service":{"healthPath":"/realms/master","targetPort":8080},"serviceAccount":{"create":true}}` | Keycloak subchart deployment </br> Ref: https://github.com/konstellation-io/helm-charts/blob/konstellation-base-1.0.2/charts/konstellation-base/values.yaml |
| keycloak.enabled | bool | `false` | Enable or disable Keycloak subchart |
| minio | object | `{"certsPath":"/certs/","consoleIngress":{"annotations":{},"enabled":false,"hosts":["minio-console.kai.local"],"ingressClassName":null,"path":"/","tls":[]},"enabled":false,"extraVolumeMounts":[],"extraVolumes":[],"ingress":{"enabled":false},"minioAPIPort":"9000","minioConsolePort":"9001","mode":"standalone","oidc":{"enabled":false},"persistence":{"enabled":false},"resources":{"requests":{"memory":"1Gi"}},"rootPassword":"","rootUser":"","users":[]}` | MinIO subchart deployment </br> Ref: https://github.com/minio/minio/blob/RELEASE.2025-02-07T23-21-09Z/helm/minio/values.yaml |
| minio.enabled | bool | `false` | Enable or disable MinIO subchart |
| mongodb | object | `{"architecture":"standalone","auth":{"rootPassword":"ChangeMe","rootUser":"ChangeMe"},"enabled":false,"persistence":{"enabled":false}}` | MongoDB subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/mongodb/values.yaml |
| mongodb.enabled | bool | `false` | Enable or disable MongoDB subchart |
| nameOverride | string | `""` | String to partially override kai.fullname template (will maintain the release name) |
| nats | object | `{"config":{"jetstream":{"enabled":true,"fileStore":{"enabled":true,"pvc":{"enabled":false}},"memoryStore":{"enabled":true,"maxSize":"2Gi"}},"merge":{"debug":false,"logtime":true,"trace":false}},"enabled":false,"monitor":{"enabled":false},"natsBox":{"enabled":false},"serviceAccount":{"enabled":true}}` | NATS subchart deployment </br> Ref: https://github.com/nats-io/k8s/blob/nats-1.2.10/helm/charts/nats/values.yaml |
| nats.enabled | bool | `false` | Enable or disable NATS subchart |
| nats.natsBox.enabled | bool | `false` | Whether to enable the NATS Box |
| natsManager | object | `{"affinity":{},"args":[],"autoscaling":{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80},"command":[],"deploymentStrategy":"Recreate","env":{},"envFromConfigMap":{},"envFromFiles":[],"envFromSecrets":{},"extraContainers":[],"image":{"pullPolicy":"IfNotPresent","repository":"konstellation/kai-nats-manager","tag":"0.3.0-develop.17"},"imagePullSecrets":[],"initContainers":[],"lifecycle":{},"livenessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"livenessProbeCustom":{},"networkPolicy":{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]},"nodeSelector":{},"podAnnotations":{},"podDisruptionBudget":{"enabled":false,"maxUnavailable":1,"minAvailable":null},"podLabels":{},"podSecurityContext":{},"readinessProbe":{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1},"readinessProbeCustom":{},"readyChecker":{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5},"resources":{},"secrets":[],"securityContext":{},"service":{"port":50051},"serviceAccount":{"annotations":{},"automount":true,"create":true,"name":""},"serviceMonitor":{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"},"startupProbe":{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5},"startupProbeCustom":{},"terminationGracePeriodSeconds":30,"tolerations":[],"topologySpreadConstraints":[],"volumeMounts":[],"volumes":[]}` | NATS manager |
| natsManager.affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| natsManager.args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| natsManager.autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| natsManager.command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| natsManager.deploymentStrategy | string | `"Recreate"` | Deployment strategy Specifies the strategy used to replace old Pods by new ones |
| natsManager.env | object | `{}` | Environment variables to configure application </br> Ref: https://github.com/konstellation-io/kdl-server/tree/main/app/api |
| natsManager.envFromConfigMap | object | `{}` | Variables from configMap |
| natsManager.envFromFiles | list | `[]` | Load all variables from files </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/#configure-all-key-value-pairs-in-a-configmap-as-container-environment-variables |
| natsManager.envFromSecrets | object | `{}` | Variables from secrets |
| natsManager.extraContainers | list | `[]` | Configure extra containers |
| natsManager.image | object | `{"pullPolicy":"IfNotPresent","repository":"konstellation/kai-nats-manager","tag":"0.3.0-develop.17"}` | Image registry The image configuration for the base service |
| natsManager.imagePullSecrets | list | `[]` | Specifies the secrets to use for pulling images from private registries Leave empty if no secrets are required E.g. imagePullSecrets:   - name: myRegistryKeySecretName |
| natsManager.initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| natsManager.lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| natsManager.livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| natsManager.livenessProbeCustom | object | `{}` | Custom livenessProbe |
| natsManager.networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| natsManager.networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| natsManager.networkPolicy.policyTypes | list | `[]` | Policy types |
| natsManager.nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
| natsManager.podAnnotations | object | `{}` | Configure annotations on Pods |
| natsManager.podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":1,"minAvailable":null}` | Pod Disruption Budget </br> Ref: https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/ |
| natsManager.podLabels | object | `{}` | Configure labels on Pods |
| natsManager.podSecurityContext | object | `{}` | Defines privilege and access control settings for a Pod </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| natsManager.readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| natsManager.readinessProbeCustom | object | `{}` | Custom readinessProbe |
| natsManager.readyChecker | object | `{"enabled":false,"pullPolicy":"IfNotPresent","repository":"busybox","retries":30,"services":[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}],"tag":"latest","timeout":5}` | Check if dependencies are ready |
| natsManager.readyChecker.enabled | bool | `false` | Enable or disable ready-checker |
| natsManager.readyChecker.pullPolicy | string | `"IfNotPresent"` | Pull policy for the image |
| natsManager.readyChecker.repository | string | `"busybox"` | Repository of the image |
| natsManager.readyChecker.retries | int | `30` | Number of retries before giving up |
| natsManager.readyChecker.services | list | `[{"name":"mongodb","port":27017},{"name":"keycloak","port":80}]` | List services |
| natsManager.readyChecker.tag | string | `"latest"` | Overrides the image tag |
| natsManager.readyChecker.timeout | int | `5` | Timeout for each check |
| natsManager.resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| natsManager.secrets | list | `[]` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| natsManager.securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| natsManager.service | object | `{"port":50051}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| natsManager.service.port | int | `50051` | Kubernetes Service port |
| natsManager.serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| natsManager.serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| natsManager.serviceMonitor.enabled | bool | `false` | Enable or disable |
| natsManager.startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| natsManager.startupProbeCustom | object | `{}` | Custom startupProbe |
| natsManager.terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| natsManager.tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| natsManager.topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| natsManager.volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| natsManager.volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
| postgresql | object | `{"auth":{"database":"kai","password":"ChangeMe","username":"user"},"enabled":false,"primary":{"persistence":{"enabled":false}},"replicaCount":1}` | PostgreSQL subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/main/bitnami/postgresql/values.yaml |
| postgresql.enabled | bool | `false` | Enable or disable PostgreSQL subchart |
| redis | object | `{"architecture":"standalone","auth":{"enabled":false},"enabled":false,"master":{"count":1,"persistence":{"enabled":false}},"replica":{"persistence":{"enabled":false},"replicaCount":1}}` | Redis subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/redis/20.7.0/bitnami/redis/values.yaml |
| redis.enabled | bool | `false` | Enable or disable Redis subchart |
| testConnection | object | `{"enabled":false,"repository":"busybox","tag":"latest"}` | Enable or disable test connection |
| vault | object | `{"enabled":false,"server":{"affinity":"","dataStorage":{"enabled":false},"ingress":{"enabled":false,"hosts":[{"host":"vault.kai.local"}],"ingressClassName":"","pathType":"ImplementationSpecific"}},"ui":{"enabled":true}}` | Vault subchart deployment </br> Ref: https://github.com/bitnami/charts/blob/redis/20.7.0/bitnami/redis/values.yaml |
| vault.enabled | bool | `false` | Enable or disable Vault subchart |
