package config

import (
	"articleproject/api/model/dto"
	"articleproject/constants"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var DatabaseConfig dto.Database
var JWtSecretConfig dto.JWTSecret

func LoadEnv() {
	// godotenv.Load("D:/Article Project New/.config/.env")

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("../.config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&DatabaseConfig); err != nil {
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

//
