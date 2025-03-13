package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sparsh2/notification-system/src/components/notification-server/types"
)

type KafkaService struct {
	producer sarama.SyncProducer
}

var Kafka *KafkaService

func InitKafka() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{os.Getenv("KAFKA_BOOTSTRAP_SERVERS")}
	if brokers[0] == "" {
		brokers[0] = "notification-system-kafka:9092" // default for k8s
	}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %v", err)
	}

	Kafka = &KafkaService{
		producer: producer,
	}
	return nil
}

func (k *KafkaService) Close() error {
	if k.producer != nil {
		return k.producer.Close()
	}
	return nil
}

func (k *KafkaService) SendNotification(notification *types.NotificationRequest) error {
	// Convert notification to JSON
	notificationBytes, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	// Create Kafka message
	msg := &sarama.ProducerMessage{
		Topic: "notifications",
		Value: sarama.StringEncoder(notificationBytes),
		Key:   sarama.StringEncoder(notification.UserID), // Use UserID as key for partitioning
	}

	// Send message
	partition, offset, err := k.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
