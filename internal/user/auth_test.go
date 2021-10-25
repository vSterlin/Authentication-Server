package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockUsers = []*User{
	{1, "Vladimir", "Sterlin", "vSterlin", "hashedPassword"},
}

type mockRepo struct {
}

func (mr *mockRepo) GetMany() []*User {
	return mockUsers
}

func (mr *mockRepo) GetOne(id int) *User {
	return users[id-1]
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
