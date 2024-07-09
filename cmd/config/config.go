package config

type Config struct {
	Locations LocationsConfig
}

type LocationsConfig struct {
	Next string
	Prev string
}

func GetConfig() Config {
	return Config{
		Locations: LocationsConfig{
			Next: "",
			Prev: "",
		},
	}
}
