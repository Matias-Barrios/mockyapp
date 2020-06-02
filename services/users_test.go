package services

import "testing"

var mockRequester IRequester = &struct{}{}
var subject = UsersService{}

func TestGetUsersOk(t *testing.T) {
	_requester = mockRequester
	response := subject.GetUsers()
	if response == nil {
		t.Fail()
	}
}
