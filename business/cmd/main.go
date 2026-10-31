package main

import(
	"fmt"
	"net/http"
	"smikesword/streetcravela.com/business/internal/repository"
	"smikesword/streetcravela.com/business/internal/controller"
	"os"
)

func main(){
	// In your Go code
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	repo, err := repository.New(connStr)
	if err != nil {
		panic(err)
	}
	c := controller.New(repo)

	http.HandleFunc("GET /v1/businesses/{id}", c.HandleGetById)
	http.HandleFunc("GET /v1/businesses", c.HandleGetAll)
	http.HandleFunc("POST /v1/businesses", c.HandleCreateBusiness)

	fmt.Println("listening on port :8881")

	http.ListenAndServe(":8881", nil)

}