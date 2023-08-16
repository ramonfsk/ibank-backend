package domain

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ramonfsk/ibank-backend/server/dto"
	"github.com/ramonfsk/ibank-backend/server/errs"
)

type AuthRepository interface {
	IsAuthorized(token string, userID string, fullPath string, httpMethod string) (bool, *errs.AppError)
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, userID string, fullPath string, httpMethod string) (bool, *errs.AppError) {
	request, err := buildRequest(token, userID, fullPath, httpMethod)
	if err != nil {
		return false, err
	}

	if response, err := http.DefaultClient.Do(request); err != nil {
		return false, errs.NewSendRequestError("Error while sending auth request...")
	} else {
		defer response.Body.Close()
		var responseDecoded dto.AuthResponse
		if err := json.NewDecoder(response.Body).Decode(&responseDecoded); err != nil {
			return false, errs.NewParseError("Error while decoding response from auth server")
		}

		return responseDecoded.IsAuthorized, nil
	}
}

func buildRequest(token string, userID string, fullPath string, httpMethod string) (*http.Request, *errs.AppError) {
	body := &dto.AuthRequest{
		Token:     token,
		UserID:    userID,
		RouteName: getRouteName(fullPath, httpMethod),
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, errs.NewEncodeBodyError("Error marshalling")
	}

	if request, err := http.NewRequest(http.MethodPost, "http://localhost:8011/v1/auth/verify", bytes.NewBuffer(bodyJson)); err != nil {
		return nil, errs.NewBuildRequestError("Error to build request")
	} else {
		request.Header.Add("Content-Type", "application/json")

		return request, nil
	}
}

func getRouteName(fullPath string, httpMethod string) string {
	versionApp := "/v1"
	routeName := ""
	if strings.EqualFold(fullPath, "/users") && httpMethod == http.MethodPost {
		routeName = "newUser"
	} else if strings.EqualFold(fullPath, versionApp+"/users/:id") {
		routeName = "getUser"
	} else if strings.EqualFold(fullPath, versionApp+"/users") {
		routeName = "getAllUsers"
	} else if strings.EqualFold(fullPath, versionApp+"/accounts") {
		routeName = "getAllAccounts"
	} else if strings.EqualFold(fullPath, versionApp+"/accounts/:id") {
		routeName = "getAccount"
	} else if strings.EqualFold(fullPath, versionApp+"accounts/:id/transactions") {
		routeName = "getAllTransactionsByAccount"
	} else if strings.EqualFold(fullPath, versionApp+"/accounts/:id/transactionswithperiod") {
		routeName = "getAllTransactionsByAccountWithPeriod"
	} else if strings.EqualFold(fullPath, versionApp+"/transactions/:id") {
		routeName = "getTransaction"
	} else if strings.EqualFold(fullPath, versionApp+"/transactions") && httpMethod == http.MethodPost {
		routeName = "newTransaction"
	} else if strings.EqualFold(fullPath, versionApp+"/transactions") {
		routeName = "getAllTransactions"
	}

	return routeName
}
