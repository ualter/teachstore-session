package config

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	"github.com/ualter/teachstore-session/utils"
)

// ServiceConfig
type ServiceConfig struct {
	URL  string `json:"url"`
	Port string `json:"port"`
}

var (
	environment    string
	ServicesConfig = make(map[string]*ServiceConfig)
)

const TEACHSTORE_ENROLLMENT string = "teachstore-enrollment"

// GetString get string values from config
func GetString(path string) string {
	v := GetInterface(path)
	return v.(string)
}

// GetInt get int values from config
func GetInt(path string) int {
	v := GetInterface(path)
	return v.(int)
}

// GetInt get interface{} values from config
func GetInterface(path string) interface{} {
	if environment != "" {
		if viper.IsSet(environment + "." + path) {
			return viper.Get(environment + "." + path)
		}
	}
	if viper.IsSet(path) {
		return viper.Get(path)
	}
	panic(fmt.Sprintf("Path %s was not found!!", path))
}

// GetServiceConfig return the configuration of a service
func GetServiceConfig(serviceName string) *ServiceConfig {
	return ServicesConfig[serviceName]
}

// LoadExternalConfiguration is ...
func LoadExternalConfiguration() {
	var err error
	_, err = utils.MyIP()
	if err != nil {
		panic(err.Error())
	}

	environment = os.Getenv("ENVIRONMENT")

	// TODO Use ConfigMap as Volumes
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Error %s", err.Error())
		panic(err.Error())
	}

	viper.SetDefault("port", "9999")

	if environment == "" {
		services := viper.Get("services")
		loadServicesConfig(services)
	} else {
		services := viper.Get(environment + ".services")
		if services == nil {
			msg := fmt.Sprintf("Please check it!!! It was not found the entry \"%s.services\" at config.yaml file", environment)
			panic(errors.New(msg))
		}
		loadServicesConfig(services)
	}
}

func loadServicesConfig(services interface{}) {
	for _, serviceObj := range services.([]interface{}) {
		for k1, v1 := range serviceObj.(map[interface{}]interface{}) {
			serviceKey := k1.(string)
			ServicesConfig[serviceKey] = &ServiceConfig{}
			for k2, v2 := range v1.(map[interface{}]interface{}) {
				if k2.(string) == "url" {
					ServicesConfig[serviceKey].URL = v2.(string)
				} else if k2.(string) == "port" {
					ServicesConfig[serviceKey].Port = v2.(string)
				}
			}
		}
	}

	for k, _ := range ServicesConfig {
		fmt.Printf("%s=%s\n", k, ServicesConfig[k].URL)
		fmt.Printf("%s=%s\n", k, ServicesConfig[k].Port)
	}
}

func replaceEnvInConfig(vlr string) string {
	body := []byte(vlr)
	r := regexp.MustCompile(`\$\{([^{}]+)\}`)
	replaced := r.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := r.ReplaceAllString(string(b), `$1`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		panic(fmt.Sprintf("Environment variable $%s was not set!!", group1))
	})
	return string(replaced)
}
