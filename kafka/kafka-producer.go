package kafka

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type book struct {
	Name     string   `json:"name"`
	Path     string   `json:"path"`
	Version  int      `json:"version"`
	Keywords []string `json:"keywords"`
}

//Producer export function
func Producer(bpath string, keywords []string) {

	result, err := json.Marshal(book{
		Name:     path.Base(bpath),
		Path:     bpath,
		Version:  1,
		Keywords: keywords,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "topicme"
	// Produce messages to topic (asynchronously)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          result,
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
