package discordreceiver

import "errors"

type Config struct {
	Token string `mapstructure:"token"`
}

func (c *Config) Validate() error {
	if c.Token == "" {
		return errors.New("token is required")
	}
	return nil
}
