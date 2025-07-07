package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func StartAPI(cache *Cache) {
	http.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		// ссылка будет вида /order/<order_uid>
		orderID := strings.TrimPrefix(r.URL.Path, "/order/")
		if orderID == "" {
			http.Error(w, "order ID is required", http.StatusBadRequest)
			return
		}

		order, ok := cache.Get(orderID)
		if !ok {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(order); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	})

	http.Handle("/", http.FileServer(http.Dir("./web")))

	log.Println("HTTP сервер запущен на http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
