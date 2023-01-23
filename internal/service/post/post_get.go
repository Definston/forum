package post

import "forum/internal/model"

func (p *post) GetAllPosts() (*[]model.Post, error) {
	return p.fillVotes(p.rpost.GetPostsByParentId(0))
}

func (p *post) GetCommById(pid int) (*[]model.Post, error) {
	return p.fillVotes(p.rpost.GetPostsByParentId(pid))
}

func (p *post) GetPostsByTag(tag string) (*[]model.Post, error) {
	return p.fillVotes(p.rpost.GetPostsByTag(tag))
}

func (p *post) GetPostsLikedByUser(uid int) (*[]model.Post, error) {
	return p.fillVotes(p.rpost.GetPostsLikedByUser(uid))
}

func (p *post) GetPostsAddedByUser(uid int) (*[]model.Post, error) {
	return p.fillVotes(p.rpost.GetPostsAddedByUser(uid))
}

func (p *post) fillVotes(posts *[]model.Post, err error) (*[]model.Post, error) {
	if err != nil {
		return nil, err
	}

	po := *posts
	for i := 0; i < len(*posts); i++ {
		if po[i].Likes, po[i].Dislikes, err = p.svote.GetVoteByPostId(po[i].Id); err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (p *post) GetPostById(pid int) (*model.Post, error) {
	post, err := p.rpost.GetPostById(pid)
	if err != nil {
		return nil, err
	}

	if post.Likes, post.Dislikes, err = p.svote.GetVoteByPostId(pid); err != nil {
		return nil, err
	}

	return post, nil
}
