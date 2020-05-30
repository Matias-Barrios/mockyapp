package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Matias-Barrios/mockyapp/models"
	"github.com/Matias-Barrios/mockyapp/network"
)

var (
	_requester network.IRequest = network.Request{}
)

type UsersService struct{}

type IUsersService interface {
	GetUsers() models.UsersResponse
}

const (
	url = "https://jsonplaceholder.typicode.com/users"
)

func (s *UsersService) GetUsers() models.UsersResponse {
	statuscode, resp, err := _requester.Execute("GET", url, nil, "")
	if err != nil || statuscode != http.StatusOK {
		log.Fatal(err.Error())
	}
	var usersResponse models.UsersResponse
	err = json.Unmarshal([]byte(resp), &usersResponse)
	if err != nil {
		log.Fatal(err.Error())
	}
	return usersResponse
}
