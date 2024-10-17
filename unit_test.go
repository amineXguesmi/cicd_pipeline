package main

import (
	"awesomeProject/models"
	"awesomeProject/utils"
	"github.com/golang-jwt/jwt/v4"
	"testing"

	"awesomeProject/config"
)

func TestGenerateJWTToken(t *testing.T) {
	// Define a mock secret key
	mockSecretKey := "secret"
	config.JWTSecret = mockSecretKey
	
	// Define a test user ID
	testUserID := "user123"

	// Call the function to generate the token
	tokenString, _ := utils.GenerateJWTToken(testUserID)

	jwtClaim := models.JWTToken{}

	_, _ = jwt.ParseWithClaims(tokenString, &jwtClaim, func(token *jwt.Token) (interface{}, error) {
		return "secret", nil
	})

	if jwtClaim.ID != "user123" {
		t.Errorf("token mismatch")
	}

}
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

// Mock MongoDB client
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	args := m.Called(name, opts)
	return args.Get(0).(*mongo.Database)
}

func TestSignup_UserAlreadyExists(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Define the user
	user := models.User{
		Email:    "test@example.com",
		Password: "password",
	}

	// Create a request to pass to the handler
	reqBody, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up response recorder
	rr := httptest.NewRecorder()

	// Set up mock Mongo collection
	mockCollection := new(MockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"email": user.Email}).Return(mongo.NewSingleResultFromDocument(user, nil, nil))

	// Call handler
	r.POST("/signup", handlers.Signup)
	r.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "Username already exists")
}

func TestSignup_HashesPassword(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Define the user
	user := models.User{
		Email:    "test@example.com",
		Password: "password",
	}

	// Create a request to pass to the handler
	reqBody, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up response recorder
	rr := httptest.NewRecorder()

	// Mock bcrypt password generation
	bcrypt.GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
		return []byte("$2a$10$hashedPassword"), nil
	}

	// Set up mock Mongo collection
	mockCollection := new(MockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"email": user.Email}).Return(nil, mongo.ErrNoDocuments)
	mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("models.User")).Return(&mongo.InsertOneResult{}, nil)

	// Call handler
	r.POST("/signup", handlers.Signup)
	r.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "User created successfully")
}

func TestSignup_FailedToHashPassword(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Define the user
	user := models.User{
		Email:    "test@example.com",
		Password: "password",
	}

	// Create a request to pass to the handler
	reqBody, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up response recorder
	rr := httptest.NewRecorder()

	// Mock bcrypt to return an error
	bcrypt.GenerateFromPassword = func(password []byte, cost int) ([]byte, error) {
		return nil, errors.New("bcrypt error")
	}

	// Call handler
	r.POST("/signup", handlers.Signup)
	r.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Failed to hash password")
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Define the login input
	loginInput := models.LoginInput{
		Email:    "wrong@example.com",
		Password: "wrongpassword",
	}

	// Create a request to pass to the handler
	reqBody, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up response recorder
	rr := httptest.NewRecorder()

	// Mock Mongo collection to return no documents
	mockCollection := new(MockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"email": loginInput.Email}).Return(mongo.NewSingleResultFromError(mongo.ErrNoDocuments))

	// Call handler
	r.POST("/login", handlers.Login)
	r.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid credentials")
}

func TestLogin_ValidCredentials(t *testing.T) {
	// Set up Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Define the login input
	loginInput := models.LoginInput{
		Email:    "test@example.com",
		Password: "password",
	}

	// Define the user that exists in the database
	user := models.User{
		Email:    loginInput.Email,
		Password: "$2a$10$hashedPassword", // This should be a hashed password
	}

	// Create a request to pass to the handler
	reqBody, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Set up response recorder
	rr := httptest.NewRecorder()

	// Mock Mongo collection to return a valid user
	mockCollection := new(MockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"email": loginInput.Email}).Return(mongo.NewSingleResultFromDocument(user, nil, nil))

	// Mock bcrypt password comparison to return nil (indicating valid password)
	bcrypt.CompareHashAndPassword = func(hashedPassword, password []byte) error {
		return nil
	}

	// Call handler
	r.POST("/login", handlers.Login)
	r.ServeHTTP(rr, req)

	// Check the response
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "access_token")
}
