package main

import (
	"log"
	"orderservice/internal"
)

func main() {
	// Инит подключения к бд
	db, err := internal.ConnectDB()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	// Инит кеша
	cache := internal.NewCache()

	// подгружаем из бд в кеш
	orders, err := internal.LoadOrders(db)
	if err == nil {
		cache.Load(orders)
		log.Printf("Loaded %d orders into cache", len(orders))
	}

	// запуск консюма
	internal.StartKafkaConsumer(
		internal.GetEnv("KAFKA_BROKER", "localhost:9092"),
		internal.GetEnv("KAFKA_TOPIC", "order-topic"),
		db,
		cache,
	)

	// запуск api
	internal.StartAPI(cache)
}
