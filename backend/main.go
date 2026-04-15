package main

import (
	"database/sql"
	"log"
	"task-manager/database"
	"task-manager/handlers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "host=localhost user=postgres password=Lucaide23 dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Banco não respondeu", err)
	}
	log.Println("Banco conectado")
}

func main() {
	database.Connect()
	initDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	tasks := r.Group("/tasks")
	{

		tasks.GET("", handlers.GetTask)

		tasks.POST("", handlers.CreateTask)

		tasks.PUT("/:id", handlers.UpdateTask)

		tasks.DELETE("/:id", handlers.DeleteTask)
	}
	r.Run(":8080")
}
