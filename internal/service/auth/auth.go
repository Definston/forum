package auth

import (
	"errors"

	"forum/internal/model"
)

type auth struct {
	repo model.UserRepo
}

func NewServiceAuth(repo model.UserRepo) *auth {
	return &auth{
		repo: repo,
	}
}

func (a *auth) GetUserById(uid int) (*model.User, error) {
	if uid > 0 {
		return a.repo.GetUserById(uid)
	} else {
		return nil, nil
	}
}

func (a *auth) AuthUser(email, pass string) (*model.User, error) {
	user, err := a.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if pass != user.Pass {
		return nil, errors.New("worng pass")
	}

	return user, nil
}

func (a *auth) IsValidUser(email, nick, pass string) bool {
	for i := 0; i < len(email); {
		if email[i] == ' ' {
			email = email[1:]
		} else {
			break
		}
	}

	for i := 0; i < len(nick); {
		if nick[i] == ' ' {
			nick = nick[1:]
		} else {
			break
		}
	}

	for i := 0; i < len(pass); {
		if pass[i] == ' ' {
			pass = pass[1:]
		} else {
			break
		}
	}

	if email == "" || nick == "" || pass == "" || len(email) == 0 || len(nick) == 0 || len(pass) == 0 {
		return false
	}

	return true
}
