package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"
	"flag"

	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaUsername string
	kafkaPassword string
	brokerList string
	brokers []string
)

func main() {
	flag.StringVar(&kafkaUsername, "user", "admin", "kafka auth username")
	flag.StringVar(&kafkaPassword, "password", "admin", "kafka auth password")
	flag.StringVar(&brokerList, "broker-list", "cpaas-kafka:9092", "kafka broker list, seperated with comma")

	flag.Parse()

	brokers = strings.Split(brokerList, ",")

	writer := getWriter()
	reader := getReader()
}

func getWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Dialer: &kafka.Dialer{
			ClientID: "testwriter",
			SASLMechanism: plain.Mechanism{
				Username: kafkaUsername,
				Password: kafkaPassword,
			},
		},
		Brokers:    brokers,
		Topic:      "test",
		BatchBytes: 65535600,
	})
}

func getReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Dialer: &kafka.Dialer{
			ClientID: "testreader",
			SASLMechanism: plain.Mechanism{
				Username: kafkaUsername,
				Password: kafkaPassword,
			},
		},
		Brokers:    brokers,
		Topic:      "test",
		GroupID:  "test",
		MinBytes: 10,   // 10B
		MaxBytes: 10e8, // 1GB
	})
}
