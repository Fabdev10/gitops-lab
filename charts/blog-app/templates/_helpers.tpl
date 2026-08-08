{{- define "blog-app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "blog-app.labels" -}}
helm.sh/chart: {{ include "blog-app.name" . }}-{{ .Chart.Version | replace "+" "_" }}
app.kubernetes.io/name: {{ include "blog-app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "blog-app.selectorLabels" -}}
app.kubernetes.io/name: {{ include "blog-app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}
