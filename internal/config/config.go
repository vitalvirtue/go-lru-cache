package config

// AppConfig holds configuration settings
type AppConfig struct {
    CacheCapacity int
}

// NewConfig initializes and returns the application configuration
func NewConfig() *AppConfig {
    return &AppConfig{CacheCapacity: 5} // Varsayılan kapasite
}
