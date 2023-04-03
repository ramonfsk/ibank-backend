package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/ramonfsk/ibank/server/dto"
	"github.com/ramonfsk/ibank/server/errs"
	"github.com/ramonfsk/ibank/server/mocks/service"
)

var router *gin.Engine
var uh UserHandler
var mockService *service.MockUserService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService = service.NewMockUserService(ctrl)
	uh = UserHandler{mockService}
	router = gin.Default()
	router.GET("/users", uh.getAllUsers)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestShouldReturnUsersWithStatusCode200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.UserResponse{
		{
			ID:        "1",
			Name:      "Ramon Ferreira",
			Birthdate: "1990-01-01",
			Password:  "123",
			Email:     "rfnascimento@ibm.com",
			Document:  "12196183067",
			Phone:     "5561999991111",
			Status:    "1",
			IsAdmin:   true},
	}
	mockService.EXPECT().GetAllUsers("").Return(dummyCustomers, nil)

	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func TestShouldReturnStatusCode500WithErrorMessage(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllUsers("").Return(nil, errs.NewUnexpectedError("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/users", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Assert
	if recorder.Code != http.StatusBadGateway {
		t.Error("Failed while testing the status code")
	}
}
