package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oklog/ulid/v2"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

// DBリトライ付き初期化関数
func initDB() {
	var err error
	connStr := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") + "?sslmode=disable"
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return // 成功
			}
		}
		log.Printf("DB接続失敗（%d回目）: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("DBに接続できませんでした: ", err)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (this is OK if running in production):", err)
	}

	initDB()
	defer db.Close()

	http.HandleFunc("/user", handler)
	log.Println("Server started at :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			log.Printf("[ERROR] DB Query failed: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
				log.Printf("[ERROR] Row scan failed: %v", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}
		if err := rows.Err(); err != nil {
			log.Printf("[ERROR] Rows iteration error: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(users); err != nil {
			log.Printf("[ERROR] JSON encode failed: %v", err)
			http.Error(w, "JSON encode error", http.StatusInternalServerError)
		}

	case http.MethodPost:
		var req struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("[ERROR] JSON decode failed: %v", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if req.Name == "" || req.Age < 0 || req.Age > 150 {
			log.Printf("[ERROR] Invalid input: name=%v, age=%v", req.Name, req.Age)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		id := ulid.MustNew(ulid.Timestamp(time.Now()), ulid.DefaultEntropy())

		_, err := db.Exec("INSERT INTO users (id, name, age) VALUES ($1, $2, $3)",
			id.String(), req.Name, req.Age)
		if err != nil {
			log.Printf("[ERROR] DB Insert failed: %v", err)
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		user := User{
			Id:   id.String(),
			Name: req.Name,
			Age:  req.Age,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Printf("[ERROR] JSON encode failed: %v", err)
			http.Error(w, "JSON encode error", http.StatusInternalServerError)
		}

	default:
		log.Printf("[ERROR] Method not allowed: %v", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
