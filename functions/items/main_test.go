package main

import (
	"net/http"
	"testing"

	httpdelivery "github.com/EwanValentine/invoicely/pkg/delivery/http"
	"github.com/EwanValentine/invoicely/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type MockItemRepository struct{}

func (r *MockItemRepository) Get(id string) (*domain.Item, error) {
	return &domain.Item{}, nil
}

func (r *MockItemRepository) List() (*[]domain.Item, error) {
	return &[]domain.Item{}, nil
}

func (r *MockItemRepository) Store(item *domain.Item) error {
	return nil
}

func TestCanFetchClient(t *testing.T) {
	request := httpdelivery.Req{
		PathParameters: map[string]string{"id": "123"},
		HTTPMethod:     "GET",
	}
	h := &Handler{&MockItemRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCanCreateClient(t *testing.T) {
	request := httpdelivery.Req{
		Body:       `{ "name": "test client", "description": "some test", "rate": 40 }`,
		HTTPMethod: "POST",
	}
	h := &Handler{&MockItemRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestCanFetchClients(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "GET",
	}
	h := &Handler{&MockItemRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestHandlerInvalidJSON(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body:       "",
	}
	h := &Handler{&MockItemRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}