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
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "net/http/httptest"
    "testing"
    "awesomeProject/models"
    "awesomeProject/handlers"
)

var bcryptGenerateFromPassword = bcrypt.GenerateFromPassword

// Mock MongoDB Collection
type MockCollection struct {
    mock.Mock
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
    args := m.Called(ctx, filter)
    return args.Get(0).(*mongo.SingleResult)
}

// Test 1: User Already Exists
func TestSignup_UserExists(t *testing.T) {
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

    assert.Equal(t, http.StatusConflict, rr.Code)
    assert.Contains(t, rr.Body.String(), "Username already exists")
}

// Test 2: Successful Signup
func TestSignup_Success(t *testing.T) {
    r := gin.Default()
    r.POST("/signup", handlers.Signup)

    user := models.User{Email: "newuser@example.com", Password: "password"}
    reqBody, _ := json.Marshal(user)
    req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    mockCollection := new(MockCollection)
    mockCollection.On("FindOne", mock.Anything, bson.M{"email": user.Email}).Return(nil)

    bcryptGenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
        return []byte("$2a$10$hashedPassword"), nil
    }

    r.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
}

// Test 3: Invalid Password
func TestSignup_InvalidPassword(t *testing.T) {
    r := gin.Default()
    r.POST("/signup", handlers.Signup)

    user := models.User{Email: "test@example.com", Password: ""}
    reqBody, _ := json.Marshal(user)
    req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()

    r.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
}
