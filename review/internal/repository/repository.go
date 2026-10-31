package repository

import(
	"database/sql"
	_ "github.com/lib/pq"
	"smikesword/streetcravela.com/review/pkg/model"
	"context"
)

type Repository struct {
	db *sql.DB
}

func New(connString string) (*Repository, error){
	db, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	r := Repository{db}

	if err := r.initialize(); err != nil {
		return nil,err
	}

	return &r, nil
}

func (r *Repository) initialize() error{
	if _, err := r.db.Exec("CREATE TABLE IF NOT EXISTS reviews (id SERIAL PRIMARY KEY, business_id INTEGER, user_id INTEGER, rating NUMERIC(3,2), comment TEXT, posted_at TIMESTAMP DEFAULT NOW())"); err != nil {
		return err
	}

	if _, err := r.db.Exec("INSERT INTO reviews(business_id, user_id, rating, comment) values('1', '2','4.5', 'the carnitas are goated')"); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByBusinessId(ctx context.Context, businessId string) ([]model.Reviews, error){
	stmt, err := r.db.Prepare("SELECT * FROM reviews WHERE business_id = $1")
	
	if err != nil {
		return nil, err
	}
	
	defer stmt.Close()

	rows, err := stmt.Query(businessId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reviews := []model.Reviews{}

	for rows.Next(){
		var review model.Reviews

		if err := rows.Scan(&review.Id, &review.BusinessId, &review.UserId, &review.Rating, &review.Comment, &review.PostedAt); err != nil{
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *Repository) GetRatingAvgByBusinessId(ctx context.Context, businessId string)(*float32, error){
	stmt, err := r.db.Prepare("SELECT ROUND(AVG(rating),2) FROM reviews WHERE business_id = $1")

	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	var avg float32

	err = stmt.QueryRow(businessId).Scan(&avg)

	if err != nil {
		return nil, err
	}

	return &avg, nil
}

func (r *Repository) CreateReview(ctx context.Context, businessId string, review model.Reviews)(error){
	stmt, err := r.db.Prepare("INSERT INTO reviews (business_id, user_id, rating, comment) VALUES ($1, $2, $3, $4) RETURNING id")

	if err != nil {
		return err
	}

	var id int

	err = stmt.QueryRow(businessId, review.UserId, review.Rating, review.Comment).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}