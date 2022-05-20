package vimeo

// Config provides a way to configure the Client depending on your needs.
type Config struct {
	// Uploader
	Uploader            Uploader
	PersonalAccessToken string
}

// DefaultConfig return the default Client configuration.
func DefaultConfig() *Config {
	return &Config{
		Uploader: nil,
	}
}
