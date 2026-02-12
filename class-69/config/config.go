package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config
type DBConfig struct {
	Host           string
	Port           int
	Name           string
	User           string
	Password       string
	EnableSSLMODE  bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JwtSecretKey is required")
		os.Exit(1)
	}

	// 1. Load Database Host
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host is required")
		os.Exit(1)
	}

	// 2. Load and Parse Database Port
	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		fmt.Println("Db port is required")
		os.Exit(1)
	}
	dbPrt, err := strconv.ParseInt(dbPortStr, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	// 3. Load Database Name and User
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB Use is required")
		os.Exit(1)
	}

	// 4. Load Password and SSL Mode
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	enableSslModeRaw := os.Getenv("DB_ENABLE_SSL_MODE")
	enblSSLMode, err := strconv.ParseBool(enableSslModeRaw)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}

	// 5. Initialize DBConfig struct
	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPass,
		EnableSSLMODE: enblSSLMode,
	}

	

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), // Using the parsed port from screenshots
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
	
}

func GetConfig() *Config {
	if configurations == nil {
loadConfig()
	}
	
	return configurations
}