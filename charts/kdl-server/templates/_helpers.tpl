{{/*
Expand the name of the chart.
*/}}
{{- define "kdl-server.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kdl-server.fullname" -}}
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
{{- define "kdl-server.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kdl-server.labels" -}}
helm.sh/chart: {{ include "kdl-server.chart" . }}
{{ include "kdl-server.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kdl-server.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kdl-server.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kdl-server.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kdl-server.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchLabelSelector" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kdl-server.selectorLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kdl-server.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchTopologySpreadConstraints" -}}
{{- range $constraint := .Values.topologySpreadConstraints }}
{{- include "kdl-server.patchLabelSelector" (merge (dict "_target" $constraint) $) }}
{{- end }}
{{- end }}

{{/*
#######################
SERVER SECTION
#######################
*/}}

{{/*
Default server component
*/}}
{{- define "kdl-server.kdlServerComponentLabel" -}}
kdl-server.component: server
{{- end -}}

{{/*
Generate labels for server component
*/}}
{{- define "kdl-server.kdlServerLabels" -}}
{{- toYaml (merge ((include "kdl-server.labels" .) | fromYaml) ((include "kdl-server.kdlServerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Generate selectorLabels for server component
*/}}
{{- define "kdl-server.selectorKdlServerLabels" -}}
{{- toYaml (merge ((include "kdl-server.selectorLabels" .) | fromYaml) ((include "kdl-server.kdlServerComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchSelectorServerLabels" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kdl-server.selectorKdlServerLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kdl-server.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchTopologySpreadConstraintsServer" -}}
{{- range $constraint := .Values.topologySpreadConstraints }}
{{- include "kdl-server.patchSelectorServerLabels" (merge (dict "_target" $constraint (include "kdl-server.selectorKdlServerLabels" $)) $) }}
{{- end }}
{{- end }}

{{/*
#######################
KNOWLEDGEGALAXY SECTION
#######################
*/}}

{{/*
Default knowledgeGalaxy component
*/}}
{{- define "kdl-server.knowledgeGalaxyComponentLabel" -}}
kdl-server.component: knowledgeGalaxy
{{- end -}}

{{/*
Generate labels for knowledgeGalaxy component
*/}}
{{- define "kdl-server.knowledgeGalaxyLabels" -}}
{{- toYaml (merge ((include "kdl-server.labels" .) | fromYaml) ((include "kdl-server.knowledgeGalaxyComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Generate selectorLabels for knowledgeGalaxy component
*/}}
{{- define "kdl-server.selectorKnowledgeGalaxyLabels" -}}
{{- toYaml (merge ((include "kdl-server.selectorLabels" .) | fromYaml) ((include "kdl-server.knowledgeGalaxyComponentLabel" .) | fromYaml)) }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch the label selector on an object
This template will add a labelSelector using matchLabels to the object referenced at _target if there is no labelSelector specified.
The matchLabels are created with the selectorLabels template.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchSelectorKnowledgeGalaxyLabels" -}}
{{- if not (hasKey ._target "labelSelector") }}
{{- $selectorLabels := (include "kdl-server.selectorKnowledgeGalaxyLabels" .) | fromYaml }}
{{- $_ := set ._target "labelSelector" (dict "matchLabels" $selectorLabels) }}
{{- end }}
{{- end }}

{{/*
Ref: https://github.com/aws/karpenter-provider-aws/blob/main/charts/karpenter/templates/_helpers.tpl
Patch topology spread constraints
This template uses the kdl-server.selectorLabels template to add a labelSelector to topologySpreadConstraints if one isn't specified.
This works because Helm treats dictionaries as mutable objects and allows passing them by reference.
*/}}
{{- define "kdl-server.patchTopologySpreadConstraintsKnowledgeGalaxy" -}}
{{- range $constraint := .Values.knowledgeGalaxy.topologySpreadConstraints }}
{{- include "kdl-server.patchSelectorKnowledgeGalaxyLabels" (merge (dict "_target" $constraint (include "kdl-server.selectorKnowledgeGalaxyLabels" $)) $) }}
{{- end }}
{{- end }}

{{/*
##############
# WIP LEGACY #
##############
*/}}

{{/*
Add the protocol part to the uri
*/}}
{{- define "http.protocol" -}}
  {{ ternary "https" "http" .Values.global.ingress.tls.enabled }}
{{- end -}}

{{/* Create the name of knowledge-galaxy service account to use */}}
{{- define "knowledgeGalaxy.serviceAccountName" -}}
{{- if .Values.knowledgeGalaxy.serviceaccount.enabled -}}
    {{ default "knowledge-galaxy" .Values.knowledgeGalaxy.serviceaccount.name }}
{{- end -}}
{{- end -}}

{{/*
Global tls secret name
*/}}
{{- define "global.tlsSecretName" -}}
{{-  if kindIs "invalid" .Values.global.ingress.tls.secretName -}}
  {{- printf "%s-%s-tls" $.Values.global.domain $.appName -}}
{{- else -}}
  {{- .Values.global.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create minio tls secret name
*/}}
{{- define "minio.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.minio.ingress.tls.secretName -}}
  {{- $_ := set $ "appName" "minio" }}
  {{- include "global.tlsSecretName" . -}}
{{- else -}}
  {{- .Values.minio.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create minio tls secret name
*/}}
{{- define "minioConsole.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.minio.consoleIngress.tls.secretName -}}
  {{- $_ := set $ "appName" "minio-console" }}
  {{- include "global.tlsSecretName" . -}}
{{- else -}}
  {{- .Values.minio.consoleIngress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create kdlServer tls secret name
*/}}
{{- define "kdlServer.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.kdlServer.ingress.tls.secretName -}}
  {{- $_ := set $ "appName" "kdlapp" }}
  {{- include "global.tlsSecretName" . -}}
{{- else -}}
  {{- .Values.kdlServer.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create gitea tls secret name
*/}}
{{- define "gitea.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.gitea.ingress.tls.secretName -}}
  {{- $_ := set $ "appName" "gitea" }}
  {{- include "global.tlsSecretName" . -}}
{{- else -}}
  {{- .Values.gitea.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create user-tools tls secret name
*/}}
{{- define "userTools.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.userToolsOperator.ingress.tls.secretName -}}
  {{- .Values.global.ingress.tls.secretName -}}
{{- else -}}
  {{- .Values.userToolsOperator.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}

{{/*
Create projectOperator mlflow tls secret name
*/}}
{{- define "projectOperator.mlflow.tlsSecretName" -}}
{{- if kindIs "invalid" .Values.projectOperator.mlflow.ingress.tls.secretName -}}
  {{- .Values.global.ingress.tls.secretName -}}
{{- else -}}
  {{- .Values.projectOperator.mlflow.ingress.tls.secretName -}}
{{- end -}}
{{- end -}}
