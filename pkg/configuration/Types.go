package configuration

type Configuration struct {
	Endpoint     string `mapstructure:"endpoint"`
	ProviderPort string `mapstructure:"provider-port"`
}
