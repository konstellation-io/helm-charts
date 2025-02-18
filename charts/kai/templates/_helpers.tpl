{{/*
Expand the name of the chart.
*/}}
{{- define "kai.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kai.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- if .Values.nameOverride }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kai.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kai.labels" -}}
helm.sh/chart: {{ include "kai.chart" . }}
{{ include "kai.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kai.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kai.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
#######################
# ADMIN API SECTION
#######################
*/}}

{{/*
Create the name for the adminApi
*/}}
{{- define "kai.adminApi.name" -}}
{{- printf "%s-admin-api" (include "kai.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "kai.adminApiServiceAccountName" -}}
{{- if .Values.adminApi.serviceAccount.create -}}
{{- default (include "kai.fullname" .) .Values.adminApi.serviceAccount.name -}}
{{- else -}}
{{- default "default" .Values.adminApi.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Default adminApi component
*/}}
{{- define "kai.adminApiComponentLabel" -}}
kai/component: admin-api
{{- end -}}

{{/*
Generate labels for adminApi component
*/}}
{{- define "kai.adminApiLabels" -}}
{{- toYaml (merge ((include "kai.labels" .) | fromYaml) ((include "kai.adminApiComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Generate selectorLabels for adminApi component
*/}}
{{- define "kai.selectorKaiAdminApiLabels" -}}
{{- toYaml (merge ((include "kai.selectorLabels" .) | fromYaml) ((include "kai.adminApiComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchSelectorAdminApiLabels" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kai.selectorKaiAdminApiLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kai.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchTopologySpreadConstraintsAdminApi" -}}
{{- range $constraint := .Values.topologySpreadConstraints }}
{{- include "kai.patchSelectorAdminApiLabels" (merge (dict "_target" $constraint (include "kai.selectorKaiAdminApiLabels" $)) $) }}
{{- end }}
{{- end }}

{{/*
#########################
# K8S MANAGER SECTION
#########################
*/}}

{{/*
Create the name for the k8sManager
*/}}
{{- define "kai.k8sManager.name" -}}
{{- printf "%s-k8s-manager" (include "kai.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "kai.k8sManagerServiceAccountName" -}}
{{- if .Values.k8sManager.serviceAccount.create -}}
{{- default (printf "%s-k8s-manager" (include "kai.fullname" .)) .Values.k8sManager.serviceAccount.name | quote -}}
{{- else -}}
{{- default "default" .Values.k8sManager.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Default k8sManager component
*/}}
{{- define "kai.k8sManagerComponentLabel" -}}
kai/component: k8s-manager
{{- end -}}

{{/*
Generate labels for k8sManager component
*/}}
{{- define "kai.k8sManagerLabels" -}}
{{- toYaml (merge ((include "kai.labels" .) | fromYaml) ((include "kai.k8sManagerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Generate selectorLabels for k8sManager component
*/}}
{{- define "kai.selectorK8sManagerLabels" -}}
{{- toYaml (merge ((include "kai.selectorLabels" .) | fromYaml) ((include "kai.k8sManagerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchSelectorK8sManagerLabels" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kai.selectorK8sManagerLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kai.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchTopologySpreadConstraintsK8sManager" -}}
{{- range $constraint := .Values.k8sManager.topologySpreadConstraints }}
{{- include "kai.patchSelectorK8sManagerLabels" (merge (dict "_target" $constraint (include "kai.selectorK8sManagerLabels" $)) $) }}
{{- end }}
{{- end }}

{{/*
###########################
# NATS MANAGER SECTION
###########################
*/}}

{{/*
Create the name for the natsManager
*/}}
{{- define "kai.natsManager.name" -}}
{{- printf "%s-nats-manager" (include "kai.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "kai.natsManagerServiceAccountName" -}}
{{- if .Values.natsManager.serviceAccount.create -}}
{{- default (printf "%s-nats-manager" (include "kai.fullname" .)) .Values.natsManager.serviceAccount.name | quote -}}
{{- else -}}
{{- default "default" .Values.natsManager.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Default natsManager component
*/}}
{{- define "kai.natsManagerComponentLabel" -}}
kai/component: nats-manager
{{- end -}}

{{/*
Generate labels for natsManager component
*/}}
{{- define "kai.natsManagerLabels" -}}
{{- toYaml (merge ((include "kai.labels" .) | fromYaml) ((include "kai.natsManagerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Generate selectorLabels for natsManager component
*/}}
{{- define "kai.selectorNatsManagerLabels" -}}
{{- toYaml (merge ((include "kai.selectorLabels" .) | fromYaml) ((include "kai.natsManagerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchSelectorNatsManagerLabels" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kai.selectorNatsManagerLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kai.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kai.patchTopologySpreadConstraintsNatsManager" -}}
{{- range $constraint := .Values.natsManager.topologySpreadConstraints }}
{{- include "kai.patchSelectorNatsManagerLabels" (merge (dict "_target" $constraint (include "kai.selectorNatsManagerLabels" $)) $) }}
{{- end }}
{{- end }}
