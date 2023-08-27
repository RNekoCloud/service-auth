package service

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type orderPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewNotifikasi(p *kafka.Producer, topic string) *orderPlacer {
	return &orderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *orderPlacer) sendNotification(email string) error {
	notificationMsg := fmt.Sprintf(email)
	payload := []byte(notificationMsg)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &op.topic, Partition: kafka.PartitionAny},
		Value:          payload,
	}, op.deliverych)

	if err != nil {
		return err
	}

	<-op.deliverych
	fmt.Printf("Notification sent: %s\n", notificationMsg)

	return nil
}

func KafkaProducer(email string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:9092",
		"client.id":         "something",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return err
	}

	op := NewNotifikasi(p, "mail-topic")
	err = op.sendNotification(email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
