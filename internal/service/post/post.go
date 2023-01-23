package post

import "forum/internal/model"

type post struct {
	rpost model.PostRepo
	stag  model.TagService
	svote model.VoteService
}

func NewServicePost(repo model.PostRepo, tag model.TagService, svote model.VoteService) *post {
	return &post{
		rpost: repo,
		stag:  tag,
		svote: svote,
	}
}

func (p *post) IsValidContent(content string) bool {
	for i := 0; i < len(content); {
		if content[i] == ' ' {
			content = content[1:]
		} else {
			return true
		}
	}

	if len(content) == 0 || content == "" {
		return false
	}

	return true
}
