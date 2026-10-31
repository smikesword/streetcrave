package controller

import (
	// "context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"smikesword/streetcravela.com/business/gen/mock"
	"smikesword/streetcravela.com/business/internal/repository"
	"smikesword/streetcravela.com/business/pkg/model"

	// "smikesword/streetcravela.com/business/pkg/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T){
	tests := []struct{
		name string
		input string
		wantStatus int
	
	}{
		{
			name: "invalid request body with extra field should return 400 bad request",
			input: `{
				"owner_id": 2,
				"rating": 5,
				"business_type": "elote",
				"media_urls": ["s3://bucket"]
			}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "successful creation should return 201 Created",
			input: `{
				"owner_id": 2,
				"business_type": "fruit stand",
				"media_urls": ["s3L://bucket"]
			}`,
			wantStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests{

		t.Run(tt.name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := mock.NewMockbusinessRepository(ctrl)
			
			c := New(repoMock)


			if tt.name == "invalid request body with extra field" {
				// DO NOTHING. 
				// If the handler incorrectly calls the repo, gomock will FAIL the test automatically.
				// This proves the handler correctly stopped at the JSON decoder.
			} else {
				// 🚨 FIX 2: Use gomock.Any() for both arguments
				// We don't care about the exact context or struct, just that it was called.
				repoMock.EXPECT().CreateBusiness(gomock.Any(), gomock.Any()).Return(nil)
			}

			req := httptest.NewRequest("POST", "/v1/businesses", strings.NewReader(tt.input))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c.HandleCreateBusiness(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code, tt.name)
		})
	}
}

func TestGetById(t *testing.T){
	tests := []struct{
		name string
		id string
		expRes *model.Business
		expError error
		wantStatus int
	}{
		{
			name: "non existent should return 404 not found",
			id: "2",
			wantStatus: http.StatusNotFound,
			expError: repository.ErrNotFound,
		},
		{
			name: "existing business should return business",
			id: "1",
			expRes: &model.Business{
				Id: 1,
				OwnerId: 2,
				Type: "food truck",
				MediaUrls: []string{"s3://bucket"},
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repoMock := mock.NewMockbusinessRepository(ctrl)
			
			c := New(repoMock)

			repoMock.EXPECT().GetById(gomock.Any(), tt.id).Return(tt.expRes, tt.expError)

			req := httptest.NewRequest("GET", "/v1/businesses/"+tt.id, nil)
			req.SetPathValue("id", tt.id)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c.HandleGetById(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code, tt.name)

			if tt.expRes != nil {
				expJson, err := json.Marshal(tt.expRes)
				assert.NoError(t, err, "failed to parse expected JSON")
				assert.Contains(t, rec.Body.String(), string(expJson))
			}
			
		})
	}
}