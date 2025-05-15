package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/controller"
	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/dao"
	"github.com/shin-kanda/curriculum_7_shin_kanda/docker/mysql/db/usecase"
)

var db *sql.DB

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_PWD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	connStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase)

	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func main() {
	// 依存性の注入
	userDAO := dao.NewUserDAO(db)

	searchUsecase := usecase.NewSearchUserUsecase(userDAO)
	registerUsecase := usecase.NewRegisterUserUsecase(userDAO)

	searchController := controller.NewSearchUserController(searchUsecase)
	registerController := controller.NewRegisterUserController(registerUsecase)

	// リクエストハンドラーの設定
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			searchController.Handle(w, r)
		case http.MethodPost:
			registerController.Handle(w, r)
		default:
			log.Printf("fail: HTTP Method is %s\n", r.Method)
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	// Ctrl+CでHTTPサーバー停止時にDBをクローズする
	closeDBWithSysCall()

	// 8000番ポートでリクエストを待ち受ける
	log.Println("Listening on :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
