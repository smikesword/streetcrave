package repository

import(
	"database/sql"
	_ "github.com/lib/pq"
	"smikesword/streetcravela.com/business/pkg/model"
	"context"
	"strings"
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
	if _, err := r.db.Exec("CREATE TABLE IF NOT EXISTS businesses (id SERIAL PRIMARY KEY, owner_id INTEGER, business_type VARCHAR(256), media_urls TEXT)"); err != nil {
		return err
	}

	if _, err := r.db.Exec("INSERT INTO businesses(owner_id, business_type, media_urls) values('1', 'food truck', 'http://s3:bucket/carnitas,http://s3bucket/alpastor')"); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetById(ctx context.Context, businessId string) (*model.Business, error){
	stmt, err := r.db.Prepare("SELECT * FROM businesses WHERE id = $1")
	
	if err != nil {
		return nil, err
	}
	
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, businessId)

	if row == nil {
		return nil, ErrNotFound
	}

	business := model.Business{}

	var media string

	if err := row.Scan(&business.Id, &business.OwnerId, &business.Type, &media); err != nil{
		return nil, err
	}

	if media != "" {
		business.MediaUrls = strings.Split(media, ",")
	}


	return &business, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]model.Business, error){
	stmt, err := r.db.Prepare("SELECT * FROM businesses")
	
	if err != nil{
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	businesses := []model.Business{}

	for rows.Next(){
		var business model.Business
		var media string
		
		if err := rows.Scan(&business.Id, &business.OwnerId, &business.Type, &media); err != nil{
			return nil, err
		}

		if media != "" {
			business.MediaUrls = strings.Split(media, ",")
		}

		businesses = append(businesses, business)
	}

	return businesses, nil

}

func (r *Repository) CreateBusiness(ctx context.Context, business model.Business) (error){
	stmt, err := r.db.Prepare("INSERT INTO businesses(owner_id, business_type, media_urls) values ($1, $2, $3) RETURNING ID")

	if err != nil{
		return err
	}

	defer stmt.Close()

	var id int

	err = stmt.QueryRow(business.OwnerId, business.Type, business.MediaUrls).Scan(&id)

	if err != nil {
		return err
	}

	return nil
}

