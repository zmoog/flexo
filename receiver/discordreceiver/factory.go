package discordreceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("discord")
)

func createDefaultConfig() component.Config {
	return Config{}
}

// createLogsReceiver creates a new discord receiver
func createLogsReceiver(ctx context.Context, settings receiver.Settings, baseCfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	logger := settings.Logger
	config := baseCfg.(Config)

	rcvr := NewDiscordReceiver(logger, consumer, &config)

	return rcvr, nil
}

// NewFactory creates a new receiver factory
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelAlpha),
	)
}
