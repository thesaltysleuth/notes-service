package store

import "errors"

type User struct {
	Username string
	Password string // plaintext for now, will hash later
}

type UserStore struct{
	users map[string]User
}

func NewUserStore() *UserStore {
	return &UserStore{users: make(map[string]User)}
}

func (s *UserStore) Add(username, password string) error{
	if _, exists := s.users[username]; exists {
		return errors.New("user exists")
	}
	s.users[username] = User{Username: username, Password: password}
	return nil
}

func (s *UserStore) Validate(username,password string) bool{
	u, exists := s.users[username]
	return exists && u.Password == password
}

