{{/*
#######################
KDL SERVER SECTION
#######################
*/}}

{{/*
Validate that if 'env.MINIO_ENDPOINT' exists
*/}}
{{- define "validate.server_environment" -}}
{{- if or (not .Values.env.KDL_ADMIN_EMAIL) (not .Values.env.KDL_ADMIN_USERNAME) (not .Values.env.MINIO_ENDPOINT) -}}
  {{- fail "Error: environment env.MINIO_ENDPOINT doesn't exist" -}}
{{- end -}}
{{- end -}}

{{/*
#######################
CLEANER SECTION
#######################
*/}}

{{/*
Validate that if 'cleaner.enabled' is true and 'sharedVolume.enabled' is false,
then 'volumes' and 'volumeMounts' must not be empty.
*/}}
{{- define "validate.cleaner" -}}
{{- if and .Values.cleaner.enabled (not .Values.sharedVolume.enabled) (or (empty .Values.volumes) (empty .Values.volumeMounts)) -}}
  {{- fail "Error: When 'cleaner.enabled' is true and 'sharedVolume.enabled' is false, 'volumes' and 'volumeMounts' must not be empty. Please enable 'sharedVolume' or provide appropriate 'volumes' and 'volumeMounts' configurations." -}}
{{- end -}}
{{- end -}}
