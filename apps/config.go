package apps

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type LinePayConfig struct {
	ChannelID     string
	ChannelSecret string
	APIUrl        string
}

var (
	RLPConfig LinePayConfig
	GinMode   = "debug"
)

func init() {
	CurrentEnvironment, ok := os.LookupEnv("GO_ENV")

	var err error

	if !ok {
		CurrentEnvironment = "development"
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Environment: " + CurrentEnvironment)

	RLPConfig = setLinePayConfig()
}

func setLinePayConfig() LinePayConfig {
	return LinePayConfig{
		ChannelID:     os.Getenv("LINE_CHANNEL_ID"),
		ChannelSecret: os.Getenv("LINE_CHANNEL_SECRET_KEY"),
		APIUrl:        os.Getenv("SANDBOX_URL"),
	}
}
