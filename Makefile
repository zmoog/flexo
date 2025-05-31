collector-source:
	cd collector && go run go.opentelemetry.io/collector/cmd/builder@v0.127.0 --config ./builder-config.yaml
