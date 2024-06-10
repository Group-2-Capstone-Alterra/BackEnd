package handler

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/user"

	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserService struct {
	CreateCalled     bool
	DeleteCalled     bool
	GetProfileCalled bool
	LoginCalled      bool
	UpdateByIdCalled bool
	mock.Mock
}

func (m *mockUserService) Create(inputCore user.Core) error {
	m.CreateCalled = true
	return nil
}

func (m *mockUserService) Delete(id uint) error {
	m.DeleteCalled = true
	return nil
}

func (m *mockUserService) GetProfile(id uint) (*user.Core, error) {
	args := m.Called(id)
	return args.Get(0).(*user.Core), args.Error(1)
}

func (m *mockUserService) Login(email string, password string) (*user.Core, string, error) {
	m.LoginCalled = true
	return &user.Core{}, "", nil
}

func (m *mockUserService) UpdateById(id uint, inputCore user.Core, file io.Reader, filename string) (string, error) {
	m.UpdateByIdCalled = true
	return "", nil
}

func TestRegister(t *testing.T) {
	e := echo.New()
	reqBody := `{"full_name": "user", "email": "user@gmail.com", "password": "Password"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService := &mockUserService{}
	userHandler := &UserHandler{
		userService: mockUserService,
	}

	if err := userHandler.Register(c); err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, `{"message":"Registration successful! You can now log in to your account."}`, strings.TrimSpace(rec.Body.String()))
}

func TestLogin(t *testing.T) {
	e := echo.New()
	reqBody := `{"email": "user@gmail.com", "password": "Password"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService := &mockUserService{}
	userHandler := &UserHandler{
		userService: mockUserService,
	}

	if err := userHandler.Login(c); err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, mockUserService.LoginCalled)

	expectedResponse := `{"message":"Login successful! You are now logged in.","data":{}}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
}

func TestProfile(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/profile", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService := &mockUserService{}
	userHandler := &UserHandler{
		userService: mockUserService,
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(idToken)))

	fullName := "user test"
	email := "user@gmail.com"
	role := "user"
	profilePicture := "https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/1717984306_almohada.jpeg"
	coordinate := "-8.006553676477932,112.62021458533962"

	coreUser := &user.Core{
		ID:             11,
		FullName:       fullName,
		Email:          email,
		Role:           role,
		ProfilePicture: profilePicture,
		Coordinate:     coordinate,
	}

	mockUserService.On("GetProfile", uint(idToken)).Return(coreUser, nil)

	if err := userHandler.Profile(c); err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	expectedResponse := `{"message":"Profile information loaded successfully.","data":{"id":11,"full_name":"user test","email":"user@gmail.com","role":"user","profile_picture":"https://air-bnb.s3.ap-southeast-2.amazonaws.com/profilepicture/1717984306_almohada.jpeg","coordinate":"-8.006553676477932,112.62021458533962"}}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
}

func TestUpdateUserById(t *testing.T) {
	e := echo.New()
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	_ = writer.WriteField("full_name", "updated user")
	_ = writer.WriteField("email", "updated@example.com")
	_ = writer.WriteField("password", "newpassword")
	_ = writer.WriteField("address", "new address")
	_ = writer.WriteField("number_phone", "08123456789")
	_ = writer.WriteField("coordinate", "44.6463,-49.581")

	// Add a file to the request
	file, _ := os.Open("path/to/file.jpg")
	part, _ := writer.CreateFormFile("profile_picture", "file.jpg")
	_, _ = io.Copy(part, file)

	err := writer.Close()
	if err != nil {
		t.Errorf("Error creating multipart request: %v", err)
	}

	req, err := http.NewRequest(echo.PUT, "/users/:id", reqBody)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService := &mockUserService{}
	userHandler := &UserHandler{
		userService: mockUserService,
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(idToken)))

	if err := userHandler.UpdateUserById(c); err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `{"message":"Profile information updated successfully."}`, strings.TrimSpace(rec.Body.String()))
}

func TestDelete(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/delete", nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserService := &mockUserService{}
	userHandler := &UserHandler{
		userService: mockUserService,
	}

	idToken, _, _ := middlewares.ExtractTokenUserId(c)
	mockUserService.On("Delete", uint(idToken)).Return(nil)

	if err := userHandler.Delete(c); err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	expectedResponse := `{"message":"Your account has been deleted. Thank you for using our service."}`
	assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
}
