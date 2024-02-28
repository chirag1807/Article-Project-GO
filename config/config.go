package config

import (
	"articleproject/api/model/dto"
	"articleproject/constants"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config dto.Config
var JWtSecretConfig dto.JWTSecret

func LoadEnv(envFilePath string) {
	// godotenv.Load("D:/Article Project New/.config/.env")

	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(envFilePath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	viper.BindEnv("Config.Database.Username", "DATABASE_USERNAME")

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("Error While Decoding .env File")
	}

	fileContent, err := os.ReadFile(constants.SECRET_JSON_FILE_PATH)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(fileContent, &JWtSecretConfig)

	if err != nil {
		fmt.Println(err)
	}

}
