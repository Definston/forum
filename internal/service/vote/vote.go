package vote

import (
	"database/sql"
	"errors"
	"strconv"

	"forum/internal/model"
)

type vote struct {
	repo model.VoteRepo
}

func NewServiceVote(repo model.VoteRepo) *vote {
	return &vote{
		repo: repo,
	}
}

func (v *vote) Vote(uid int, postId, vote string) error {
	pid, err := strconv.Atoi(postId)
	if err != nil {
		return err
	}

	var voteRequest int
	if vote == "1" {
		voteRequest = 1
	} else if vote == "0" {
		voteRequest = 0
	} else {
		return errors.New("")
	}

	voteDatabase, err := v.repo.GetVote(pid, uid)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		return v.repo.AddVote(pid, uid, voteRequest)
	}

	if voteDatabase == true && voteRequest == 1 || voteDatabase == false && voteRequest == 0 {
		return v.repo.DelVote(pid, uid)
	}

	if err := v.repo.DelVote(pid, uid); err != nil {
		return err
	}
	return v.repo.AddVote(pid, uid, voteRequest)
}

func (v *vote) GetVoteByPostId(pid int) (likes, dislikes int, err error) {
	if l, d, err := v.repo.GetVoteByPostId(pid); err != nil {
		return 0, 0, nil
	} else {
		return int(l), int(d), nil
	}
}
