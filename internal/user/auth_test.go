package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockUsers = []*UserWithPassword{
	{User: User{Id: 1, FirstName: "Vladimir", LastName: "Sterlin", Username: "vSterlin", Email: "v@v.com"}, Password: "password"},
}

type mockRepo struct{}

func mockCurrentUserMiddleware(next *http.Request) *http.Request {
	return next.WithContext(context.WithValue(next.Context(), UserContext, &mockUsers[0].User))
}

func (mr *mockRepo) GetMany() []*User {
	us := []*User{}
	for _, u := range mockUsers {
		us = append(us, &u.User)
	}
	return us
}

func (mr *mockRepo) GetOne(id int) *User {
	return &mockUsers[id-1].User
}

func (ur *mockRepo) GetOneByEmail(email string) *UserWithPassword {
	for _, u := range mockUsers {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func (mr *mockRepo) InsertOne(u *UserWithPassword) *User {
	u.Id = len(mockUsers) + 1
	mockUsers = append(mockUsers, u)
	return &u.User
}

func setupController() *UserController {
	ur := &mockRepo{}
	us := NewUserService(ur)
	as := NewAuthService(us)

	uc := NewUserController(us, as)

	return uc
}

func TestSignUp(t *testing.T) {
	uc := setupController()

	w := httptest.NewRecorder()
	u := &UserWithPassword{User: User{FirstName: "t", LastName: "t", Username: "t", Email: "e"}, Password: "t"}
	json.NewEncoder(w).Encode(u)

	initialP := u.Password
	req := httptest.NewRequest(http.MethodPost, "/users", w.Body)
	uc.SignUp(w, req)
	res := w.Result()
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&u)

	if u.Password == "password" {
		t.Errorf("Expected password not to equal \"%s\", got \"%s\" instead from %+v\n", initialP, u.Password, u)
	}

}

func TestCurrentUser(t *testing.T) {
	uc := setupController()

	w := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/users/current-user", w.Body)

	var u *User

	req = mockCurrentUserMiddleware(req)

	uc.GetCurrentUser(w, req)
	res := w.Result()
	defer res.Body.Close()
	json.NewDecoder(req.Body).Decode(&u)

	if u.Id != mockUsers[0].Id {
		t.Errorf("Expected id not to equal \"%d\", got \"%d\" instead from %+v\n", mockUsers[0].Id, u.Id, u)
	}

}
