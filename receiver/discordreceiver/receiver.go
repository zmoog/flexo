package discordreceiver

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

// discordReceiver is a receiver for discord messages
type discordReceiver struct {
	logger   *zap.Logger
	consumer consumer.Logs
	config   *Config
	session  *discordgo.Session
}

// NewDiscordReceiver creates a new discord receiver
func newDiscordReceiver(logger *zap.Logger, consumer consumer.Logs, config *Config) *discordReceiver {
	return &discordReceiver{
		logger:   logger,
		consumer: consumer,
		config:   config,
	}
}

// Start starts the receiver
func (r *discordReceiver) Start(ctx context.Context, host component.Host) error {
	r.logger.Info("Starting discord receiver")
	session, err := discordgo.New("Bot " + r.config.Token)
	if err != nil {
		r.logger.Error("Failed to create discord session", zap.Error(err))
		return err
	}

	// Store the session and add the message handler
	r.session = session
	r.session.AddHandler(r.onMessage)

	// Start the session
	err = r.session.Open()
	if err != nil {
		r.logger.Error("Failed to open discord session", zap.Error(err))
		return err
	}

	r.logger.Info("Discord receiver started")
	return nil
}

// Stop stops the receiver
func (r *discordReceiver) Shutdown(ctx context.Context) error {
	r.logger.Info("Shutting down discord receiver")
	if r.session != nil {
		r.logger.Info("Closing discord session")
		r.session.Close()
	}
	r.logger.Info("Discord receiver stopped")
	return nil
}

// onMessage handles a new message
func (r *discordReceiver) onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	r.logger.Info("Received message", zap.String("message", message.Content))
}
