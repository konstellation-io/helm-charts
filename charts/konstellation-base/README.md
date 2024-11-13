# konstellation-base

A Helm chart to deploy konstellation-base for Kubernetes

## Description

`konstellation-base` Helm chart is designed to facilitate the deployment of applications that lack their own Helm charts or have deprecated ones. It also serves as a foundation for deploying custom-developed applications.

* **Replica management**: configure the number of replicas for your service to ensure scalability and reliability.
* **Image configuration**: specify repository, tag and pull policy to control the application's container image.
* **Service account management**: option to create and configure a Kubernetes `serviceaccount`, including setting `annotations` and specifying whether to automatically mount api credentials.
* **Environment variables**: define environment variables directly or source them from existing Kubernetes `Secrets` or `ConfigMaps`.
* **Probes and lifecycle hooks**: set up `liveness`, `readiness` and `startup probes` to monitor the application's health and define lifecycle hooks for graceful startup and shutdown processes.
* **Resource management**: allocate `cpu` and `memory` resources, set up `autoscaling` based on utilization metrics and define `pod disruption budgets` to maintain application availability during maintenance events.
* **Storage configuration**: manage persistent storage needs by configuring `PVC`, specifying access modes, storage classes and other relevant parameters.
* **Networking and exposure**: configure Kubernetes services to expose your application, with options for service types like `ClusterIP`, `NodePort` or `LoadBalancer`. Additionally, set up `Ingress` resources to manage external access to the application.
* **Security contexts**: define security contexts at both the pod and container levels to enforce security policies, such as running containers as non-root users or setting read-only root filesystems.
* **Node scheduling**: control pod placement using `node selectors`, `tolerations`, `affinities` and `topology spread constraints` to optimize resource utilization and maintain high availability.

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| ialejandro | <ivan.alejandro@intelygenz.com> |  |

## Prerequisites

* Helm 3+
* Kubernetes 1.24+

## Add repository

```console
helm repo add konstellation-io https://charts.konstellation.io
helm repo update
```

## Install Helm chart (repository mode)

```console
helm install [RELEASE_NAME] konstellation-io/konstellation-base
```

This install all the Kubernetes components associated with the chart and creates the release.

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

## Uninstall Helm chart

```console
helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Configuration

See [Customizing the chart before installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with comments:

```console
helm show values konstellation-io/konstellation-base
```

## Examples

Go to [examples](./examples) directory to see some examples of how to use this chart.

```console
# local chart
helm template test . -f examples/XX-example.yaml

# remote chart
helm template test konstellation-io/konstellation-base -f examples/XX-example.yaml
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | Affinity for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity |
| args | list | `[]` | Configure args </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling with CPU or memory utilization percentage </br> Ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/ |
| command | list | `[]` | Configure command </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/ |
| configMaps | object | `{}` | ConfigMap values to create configuration files Generate ConfigMap with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/configmap/ |
| env | object | `{}` | Environment variables to configure application |
| envFromConfigMap | object | `{}` | Variables from configMap |
| envFromSecrets | object | `{}` | Variables from secrets |
| fullnameOverride | string | `""` | String to fully override konstellation-base.fullname template |
| image | object | `{"pullPolicy":"IfNotPresent","repository":"nginx","tag":""}` | Image registry The image configuration for the base service |
| imagePullSecrets | list | `[]` | Docker registry secret names as an array |
| ingress | object | `{"annotations":{},"className":"","enabled":false,"hosts":[{"host":"chart-example.local","paths":[{"path":"/","pathType":"ImplementationSpecific"}]}],"tls":[]}` | Ingress configuration to expose app </br> Ref: https://kubernetes.io/docs/concepts/services-networking/ingress/ |
| initContainers | list | `[]` | Configure additional containers </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ |
| lifecycle | object | `{}` | Configure lifecycle hooks </br> Ref: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/ </br> Ref: https://learnk8s.io/graceful-shutdown |
| livenessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure liveness checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| livenessProbeCustom | object | `{}` | Custom livenessProbe |
| nameOverride | string | `""` | String to partially override konstellation-base.fullname template (will maintain the release name) |
| networkPolicy | object | `{"egress":[],"enabled":false,"ingress":[],"policyTypes":[]}` | NetworkPolicy configuration </br> Ref: https://kubernetes.io/docs/concepts/services-networking/network-policies/ |
| networkPolicy.enabled | bool | `false` | Enable or disable NetworkPolicy |
| networkPolicy.policyTypes | list | `[]` | Policy types |
| nodeSelector | object | `{}` | Node labels for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodeselector |
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
| readinessProbe | object | `{"enabled":false,"failureThreshold":3,"initialDelaySeconds":10,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":1}` | Configure readinessProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| readinessProbeCustom | object | `{}` | Custom readinessProbe |
| replicaCount | int | `1` | Number of replicas Specifies the number of replicas for the service |
| resources | object | `{}` | Resources limits and requested </br> Ref: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ |
| secrets | object | `{}` | Secrets values to create credentials and reference by envFromSecrets Generate Secret with following name: <release-name>-<name> </br> Ref: https://kubernetes.io/docs/concepts/configuration/secret/ |
| securityContext | object | `{}` | Defines privilege and access control settings for a Container </br> Ref: https://kubernetes.io/docs/concepts/security/pod-security-standards/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ |
| service | object | `{"port":80,"type":"ClusterIP"}` | Kubernetes service to expose Pod </br> Ref: https://kubernetes.io/docs/concepts/services-networking/service/ |
| service.port | int | `80` | Kubernetes Service port |
| service.type | string | `"ClusterIP"` | Kubernetes Service type. Allowed values: NodePort, LoadBalancer or ClusterIP |
| serviceAccount | object | `{"annotations":{},"automount":true,"create":true,"name":""}` | Enable creation of ServiceAccount </br> Ref: https://kubernetes.io/docs/concepts/security/service-accounts/ |
| serviceMonitor | object | `{"enabled":false,"interval":"30s","metricRelabelings":[],"relabelings":[],"scrapeTimeout":"10s"}` | Enable ServiceMonitor to get metrics </br> Ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor |
| serviceMonitor.enabled | bool | `false` | Enable or disable |
| startupProbe | object | `{"enabled":false,"failureThreshold":30,"initialDelaySeconds":180,"periodSeconds":10,"successThreshold":1,"timeoutSeconds":5}` | Configure startupProbe checker </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes |
| startupProbeCustom | object | `{}` | Custom startupProbe |
| terminationGracePeriodSeconds | int | `30` | Configure Pod termination grace period </br> Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination |
| testConnection | object | `{"enabled":false,"repository":"busybox","tag":""}` | Enable or disable test connection |
| tolerations | list | `[]` | Tolerations for pod assignment </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/ |
| topologySpreadConstraints | list | `[]` | Control how Pods are spread across your cluster </br> Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/#example-multiple-topologyspreadconstraints |
| volumeMounts | list | `[]` | Additional volumeMounts on the output Deployment definition |
| volumes | list | `[]` | Additional volumes on the output Deployment definition </br> Ref: https://kubernetes.io/docs/concepts/storage/volumes/ </br> Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ </br> Ref: https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/#create-a-pod-that-has-access-to-the-secret-data-through-a-volume |
