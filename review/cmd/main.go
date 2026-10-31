package main

import(
	"fmt"
	"net/http"
	"smikesword/streetcravela.com/review/internal/repository"
	"smikesword/streetcravela.com/review/internal/controller"
	"os"
)

func main(){
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

		
	repo, err := repository.New(connStr)
	if err != nil {
		panic(err)
	}
	c := controller.New(repo)

	http.HandleFunc("GET /v1/reviews/{id}", c.HandleGetBusinessReviews)
	http.HandleFunc("GET /v1/reviews/{id}/avg", c.HandleGetAvgBusinessRating)
	http.HandleFunc("POST /v1/reviews/{id}", c.HandleCreateRating)

	fmt.Println("listening on port :8888")

	http.ListenAndServe(":8888", nil)

}