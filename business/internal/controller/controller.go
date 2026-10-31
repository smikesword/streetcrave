package controller

import(
	"net/http"
	"context"
	"smikesword/streetcravela.com/business/pkg/model"
	"encoding/json"
)

type businessRepository interface{
	GetById(ctx context.Context, businessId string) (*model.Business, error)
	CreateBusiness(ctx context.Context, business model.Business) (error)
	GetAll(ctx context.Context) ([]model.Business, error)
}

type businessController struct{
	repository businessRepository
}


func New(repo businessRepository) *businessController{
	return &businessController{repo}
}

func (c *businessController) HandleGetById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	ctx := r.Context()
	business, err := c.repository.GetById(ctx, id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(business)

	if err != nil {
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

func (c *businessController) HandleGetAll(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	businesses, err := c.repository.GetAll(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(businesses)

	if err != nil {
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

func (c *businessController) HandleCreateBusiness(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	var newReview model.Business
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Disallow unknown fields to prevent unexpected data
	
	err := decoder.Decode(&newReview)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.repository.CreateBusiness(ctx, newReview)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
