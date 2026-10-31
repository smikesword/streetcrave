package controller

import (
	"encoding/json"
	// "errors"
	"net/http"
	// "smikesword/streetcravela.com/review/internal/repository"
	"smikesword/streetcravela.com/review/pkg/model"
	"context"
)

type reviewRepository interface{
	GetByBusinessId(ctx context.Context, businessId string) ([]model.Reviews, error)
	GetRatingAvgByBusinessId(ctx context.Context, businessId string) (*float32, error)
	CreateReview(ctx context.Context, businessId string, review model.Reviews)(error)
}

type reviewController struct {
	repository reviewRepository
}

func New(repo reviewRepository) *reviewController{
	return &reviewController{repo}
}

func (c *reviewController) HandleGetBusinessReviews(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	ctx := r.Context()
	reviews, err := c.repository.GetByBusinessId(ctx, id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(reviews)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *reviewController) HandleCreateRating(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	ctx := r.Context()

	var newReview model.Reviews
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields to prevent unexpected data
	err := decoder.Decode(&newReview)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = c.repository.CreateReview(ctx, id, newReview)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *reviewController) HandleGetAvgBusinessRating(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	ctx := r.Context()
	rating, err := c.repository.GetRatingAvgByBusinessId(ctx, id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(*rating)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
}
