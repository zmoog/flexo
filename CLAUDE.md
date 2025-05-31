# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Flexo is a custom OpenTelemetry (OTel) Collector distribution with Kubernetes deployment support. The project uses the OTel Collector Builder to generate a custom binary with specific components for telemetry data collection and processing.

## Development Commands

### Building the Collector
```bash
make collector-source
```
This builds the custom OTel collector binary using the OpenTelemetry Collector Builder with the configuration in `collector/builder-config.yaml`.

### Local Development with Tilt
```bash
tilt up
```
Starts the development environment which:
- Builds the collector using ko (for container images)
- Deploys to Kubernetes cluster
- Exposes ports 4317 (gRPC) and 4318 (HTTP) for OTLP ingestion

### Working with the Collector Binary
The built collector binary is located at `collector/otelcol/otelcol` after running the build command.

## Architecture

### Collector Components
The custom collector includes these OpenTelemetry components (defined in `collector/builder-config.yaml`):

**Receivers:**
- `otlpreceiver` - Receives telemetry data via OTLP protocol

**Processors:**
- `batchprocessor` - Batches telemetry data for efficient processing

**Exporters:**
- `debugexporter` - Outputs telemetry data to console for debugging
- `otlpexporter` - Forwards telemetry data via OTLP protocol

**Configuration Providers:**
- File, environment, HTTP/HTTPS, and YAML providers for flexible configuration

### Kubernetes Deployment
- Deployment uses the custom `flexo-otelcol` image built with ko
- Configuration is provided via ConfigMap (`flexo-collector-config`)
- Default collector config processes traces, metrics, and logs through the same pipeline: `otlp receiver → batch processor → debug exporter`

### Development Environment
- Uses Tilt for local Kubernetes development workflow
- Ko builds Go binaries directly into container images
- Port forwarding enables local testing on ports 4317/4318

## File Structure Context

- `collector/builder-config.yaml` - Defines which OTel components to include in the custom distribution
- `collector/otelcol/` - Generated Go module and binary output directory
- `k8s/k8s.yaml` - Kubernetes manifests for deployment and configuration
- `Tiltfile` - Development environment configuration for local k8s deployment
- `tools/tilt/ko.star` - Ko build configuration for Tilt