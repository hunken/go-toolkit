package kafka

import (
	logger "cws-proxy/pkg/log"
	"cws-proxy/pkg/util"
	"cws-proxy/pkg/worker"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"time"
)

var log = logger.Logger{}
var batchSize = 100

type Consumer interface {
	GetMessage(worker worker.Worker, kafkaName string)
}

type BaseConsumer struct {
}

func (consumer BaseConsumer) GetMessage(worker worker.Worker, kafkaName string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Debug("OK OK")

	kafkaName = strings.ToUpper(string(kafkaName))

	bootstrapServers := os.Getenv(kafkaName + "_KAFKA_SERVERS")
	kafkaGroup := os.Getenv(kafkaName + "_KAFKA_GROUP")
	kafkaTopic := os.Getenv(kafkaName + "_KAFKA_TOPIC")

	log.Info("bootstrapServers -> " + string(bootstrapServers))
	log.Info("kafkaGroup -> " + string(kafkaGroup))
	log.Info("kafkaTopic -> " + string(kafkaTopic))

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          kafkaGroup,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{kafkaTopic}, nil)

	var messages []string

	start := time.Now()
	var timeLimit = 1.0

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			// worker.Run(string(msg.Value), *db)
			messages = append(messages, string(msg.Value))
			elapsed := time.Since(start)

			if int(len(messages)) == batchSize || util.Greater(elapsed.Seconds(), timeLimit) {
				fmt.Println("LENGTH " + strconv.Itoa(len(messages)))
				process(messages, worker)
				messages = nil
				fmt.Println("LENGTH " + strconv.Itoa(len(messages)))
				start = time.Now()
			}
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			log.Fatal("Consumer error: " + string(msg.Value))
		}
	}

	c.Close()
}

func process(messages []string, worker worker.Worker) {
	fmt.Println("Length batch " + strconv.Itoa(len(messages)))
	for _, value := range messages {
		worker.Run(string(value))
		// fmt.Println("key i -> " + strconv.Itoa(i) + " | val -> " + string(v))
		// log.Info("Process element " + strconv.Itoa(i))
	}
}
