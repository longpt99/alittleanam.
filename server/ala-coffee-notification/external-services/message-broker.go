package externalservices

import (
	"ala-coffee-notification/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

const (
	ConsumerGroup             = "notifications-group"
	ConsumerTopic             = "notifications"
	KafkaServerAddress string = "localhost:9092"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Notification struct {
	From    User   `json:"from"`
	To      User   `json:"to"`
	Message string `json:"message"`
}

type UserNotifications map[string][]Notification

type NotificationStore struct {
	data UserNotifications
	mu   sync.RWMutex
}

func (ns *NotificationStore) Add(userID string,
	notification Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.data[userID] = append(ns.data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	return ns.data[userID]
}

type Consumer struct {
	store *NotificationStore
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		fmt.Println(userID)

		var notification Notification

		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}

		sess.MarkMessage(msg, "")

		s := services.InitEmailService()
		s.SendEmail(notification.From.Email, notification.To.Email, fmt.Sprintf(`%s %s`, notification.To.FirstName, notification.To.LastName))
	}

	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress}, ConsumerGroup, config)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return group, nil
}

func SetupConsumerGroup(ctx context.Context) error {
	group, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
		return err
	}
	defer group.Close()

	consumer := &Consumer{
		store: &NotificationStore{
			data: make(UserNotifications),
		},
	}

	for {
		err = group.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil {
			return err
		}
	}
}
