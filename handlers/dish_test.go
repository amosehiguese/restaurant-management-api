package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amosehiguese/restaurant-api/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuery is a mock implementation of the store.Query interface
type MockQuery struct {
	mock.Mock
}

func (m *MockQuery) GetAllMenuDishes(ctx context.Context, menuID uuid.UUID) ([]models.Dish, error) {
	args := m.Called(ctx, menuID)
	return args.Get(0).([]models.Dish), args.Error(1)
}

// Add similar methods for other store.Query interface methods...

func TestGetAllMenuDishes(t *testing.T) {
	// Initialize the mock query and handler
	mockQuery := new(MockQuery)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := getField(r, "id")
		dishId, _ := uuid.Parse(id)
		mockQuery.GetAllMenuDishes(context.Background(),dishId )
	})

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/menu/{id}/dishes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Set expectations on the mock query
	mockQuery.On("GetAllMenuDishes", mock.Anything, mock.Anything).Return([]models.Dish{}, nil)

	// Serve the HTTP request and record the response
	handler.ServeHTTP(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	// Add more assertions based on the expected behavior of your handler
}

// Add similar tests for other handlers (CreateMenuDish, RetrieveMenuDish, UpdateMenuDish, DeleteMenuDish)
