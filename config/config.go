package config

import (
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Port       string   `toml:"port"`
	Database   database `toml:"database"`
	Dev        bool     `toml:"dev"`
	PrivateKey string   `toml:"private_key"`
}

type database struct {
	DbUsername   string `toml:"db_username"`
	DbPwd        string `toml:"db_pwd"`
	DbHost       string `toml:"db_host"`
	DbPort       string `toml:"db_port"`
	DbSchemaName string `toml:"db_schema_name"`
	DbArgs       string `toml:"db_args"`
}

var config *Configuration

func InitConfig(configFile string) {
	if strings.Trim(configFile, " ") == "" {
		configFile = "./config/config.toml"
	}
	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}

}

func GetConfig() Configuration {
	if config == nil {
		InitConfig("")
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},
		{"dev"},
		{"private_key"},

		{"database", "db_host"},
		{"database", "db_port"},
		{"database", "db_username"},
		{"database", "db_schema_name"},
		{"database", "db_pwd"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
