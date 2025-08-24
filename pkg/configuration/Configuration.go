package configuration

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func New() *Configuration {
	pflag.String("provider-port", "7431", "Provider port")
	pflag.String("endpoint", "node.private:2371", "Etcd endpoint")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	viper.AutomaticEnv()
	viper.BindEnv("provider-port", "PROVIDER_PORT")
	viper.BindEnv("endpoint", "ENDPOINT")

	viper.SetDefault("provider-port", "7431")
	viper.SetDefault("endpoint", "node.private:2371")

	config := &Configuration{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Errorf("error unmarshaling config: %w", err))
	}

	if config.ProviderPort == "" {
		panic(fmt.Errorf("provider port must be non-empty"))
	}

	if config.Endpoint == "" {
		panic(fmt.Errorf("etcd endpoint must be non-empty"))
	}

	return config
}
