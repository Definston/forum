package post

import (
	"fmt"

	"forum/internal/model"
)

func (p *post) AddPost(post *model.Post) error {
	tags, err := p.stag.FindTags(post.Content)
	if err != nil {
		return nil
	}

	// tagedContent, err := p.stag.ReplaceTagsToLink(post.Content, tags)
	// if err != nil {
	// 	return err
	// }

	// post.Content = tagedContent

	pid, err := p.rpost.AddPost(post)
	if err != nil {
		return err
	}

	if err := p.stag.AddTagByPostId(pid, tags); err != nil {
		return err
	}

	link := fmt.Sprintf("/post/?id=%d", pid)
	if err := p.rpost.AddLinkById(pid, link); err != nil {
		return err
	}

	return nil
}

func (p *post) AddComm(post *model.Post) error {
	pid, err := p.rpost.AddComm(post)
	if err != nil {
		return err
	}

	if post.Resiever != nil {
		if err := p.rpost.AddResieverById(pid, *post.Resiever); err != nil {
			return err
		}
	}

	link := fmt.Sprintf("/post/?id=%d#%d", post.ParentId, pid)
	if err := p.rpost.AddLinkById(pid, link); err != nil {
		return err
	}

	return nil
}
