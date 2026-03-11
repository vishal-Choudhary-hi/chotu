package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/vishal-Choudhary-hi/chotu/internal/repository"
	"github.com/vishal-Choudhary-hi/chotu/internal/router"
)

func main() {
	log.Println("Application starting...")
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found")
	}
	repository.InitMySQL()
	repository.InitRedis()

	port := os.Getenv("PORT")

	r := router.SetupRouter()

	log.Println("server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
