package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

// Работаем с кафкой
func StartKafkaConsumer(broker, topic string, db *sql.DB, cache *Cache) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "order-group",
	})

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Ошибка чтения из Kafka:", err)
				continue
			}

			var order Order
			if err := json.Unmarshal(m.Value, &order); err != nil {
				log.Println("Некорректный JSON:", err)
				continue
			}

			if err := SaveOrder(db, order); err != nil {
				log.Println("Ошибка сохранения в БД:", err)
				continue
			}

			cache.Set(order)
			log.Printf("Заказ %s успешен!", order.OrderUID)
		}
	}()
}
