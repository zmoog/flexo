.PHONY: generate
generate:
	# look inside the receiver directory
	# and run mdatagen against the metadata.yaml
	# found there
	find receiver -name go.mod -execdir go tool mdatagen metadata.yaml \;

collector-source:
	cd collector && go run go.opentelemetry.io/collector/cmd/builder@v0.127.0 --config ./builder-config.yaml
