module github.com/odigos-io/odigos/procdiscovery

go 1.22.0

require github.com/odigos-io/odigos/common v1.0.48

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
)

replace github.com/odigos-io/odigos/common => ../common
