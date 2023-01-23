package reg

import (
	"strings"

	"forum/internal/model"
)

type reg struct {
	repo model.UserRepo
}

func NewServiceReg(repo model.UserRepo) *reg {
	return &reg{
		repo: repo,
	}
}

func (r *reg) AddUser(user *model.User) (bool, error) {
	if err := r.repo.AddUser(user); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
