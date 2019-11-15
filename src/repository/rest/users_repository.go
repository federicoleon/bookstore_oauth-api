package rest

import (
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/federicoleon/bookstore_oauth-api/src/domain/users"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
	"encoding/json"
	"errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil
}
