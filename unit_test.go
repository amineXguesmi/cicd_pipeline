package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"awesomeProject/models"
	"awesomeProject/handlers"
	"awesomeProject/config"
)

// Mock MongoDB Collection
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	config.LoadEnv()
	code := m.Run()

	os.Unsetenv("ENV")

	os.Exit(code)
}


func TestSignup_UserExists(t *testing.T) {
	// Setup
	r := gin.Default()
	r.POST("/signup", handlers.Signup)

	user := models.User{Email: "test@example.com", Password: "password"}
	reqBody, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockCollection := new(MockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"email": user.Email}).Return(mongo.NewSingleResultFromDocument(user, nil, nil))

	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "Username already exists")
}

