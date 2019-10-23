package main

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

var (
	GlobalConfig Config
)

type Config struct {
	ElasticSearch struct {
		Url string `envconfig:"default=http://localhost:9200"`
	}

	Kafka struct {
		Hosts string `envconfig:"default=localhost:9092"`
		Log   struct {
			Topic string `envconfig:"default=LOG_TOPIC"`
		}
		Event struct {
			Topic  string `envconfig:"default=EVENT_TOPIC"`
			Enable bool   `envconfig:"default=true"`
		}
	}
}

func main() {
	if err := envconfig.Init(&GlobalConfig); err != nil {
		panic("Load config from env error," + err.Error())
	}

	fmt.Printf("Es: %+v", GlobalConfig.ElasticSearch)
	fmt.Printf("Kafka: %+v", GlobalConfig.Kafka)

}
