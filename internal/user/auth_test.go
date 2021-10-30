package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockUsers = []*User{
	{1, "Vladimir", "Sterlin", "vSterlin", "v@v.com", "hashedPassword"},
}

type mockRepo struct {
}

func mockCurrentUserMiddleware(next *http.Request) *http.Request {
	return next.WithContext(context.WithValue(next.Context(), UserContext, mockUsers[0]))
}

func (mr *mockRepo) GetMany() []*User {
	return mockUsers
}

func (mr *mockRepo) GetOne(id int) *User {
	return users[id-1]
}

func (ur *mockRepo) GetOneByEmail(email string) *User {
	for _, u := range users {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func (mr *mockRepo) InsertOne(u *User) *User {
	u.Id = len(users) + 1
	users = append(users, u)
	return u
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
	u := &User{FirstName: "t", LastName: "t", Username: "t", Password: "t"}
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
