package tag

import (
	"regexp"
	"strings"

	"forum/internal/model"
)

type tag struct {
	rtag model.TagRepo
}

func NewServiceTag(rtag model.TagRepo) *tag {
	return &tag{
		rtag: rtag,
	}
}

func (t *tag) GetTagAll() (*[]string, error) {
	tags, err := t.rtag.GetTagAll()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *tag) AddTagByPostId(pid int, tags *map[string]bool) error {
	if len(*tags) == 0 {
		return nil
	}

	for tag := range *tags {
		if err := t.rtag.AddTagByPostId(pid, strings.ToLower(tag)); err != nil {
			return err
		}
	}

	return nil
}

func (t *tag) FindTags(content string) (*map[string]bool, error) {
	regTags := regexp.MustCompile(`#\w+`)
	tags := regTags.FindAllString(content, -1)
	res := map[string]bool{}
	for _, v := range tags {
		res[v[1:]] = true
	}
	return &res, nil
}

// func (t *tag) ReplaceTagsToLink(content string, tags *map[string]bool) (string, error) {
// 	for v := range *tags {
// 		content = strings.ReplaceAll(content, "#"+v, fmt.Sprintf("<a href=\"%s\">#%s</a>", fmt.Sprintf("/?tag=%s", v), v))
// 	}

// 	return content, nil
// }
