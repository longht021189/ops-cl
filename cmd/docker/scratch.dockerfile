FROM scratch

{{$builder := .Builder}}
{{range $file := .Files}}
COPY --from={{$builder}} {{$file}} {{$file}}
{{end}}