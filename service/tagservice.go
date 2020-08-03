package service

import (
	"TH0400/logger"
	"TH0400/repo"
	"time"
)

// GetTagsCondition ...
type GetTagsCondition struct {
	TopicID int `json:"topic"`
}

// GetTag ...
type GetTag struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

// CreateTopic ...
type CreateTag struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Baned     bool      `json:"baned"`
}

// CreateTag ...
func (ct *CreateTag) CreateTag() error {

	etag := repo.TagRepo{
		TagName:   ct.Name,
		CreatedAt: ct.CreatedAt,
		Baned:     ct.Baned,
	}

	err := etag.CreateTag()
	if err != nil {
		logger.Errorf("not found Tag: %s", err.Error())
		return err
	}

	return nil
}

// GetTags ...
func (ts *GetTagsCondition) GetTags() (gettags []*GetTag, err error) {

	etag := repo.TagTopicRepo{
		TopicID: ts.TopicID,
	}

	tr, err := etag.GetTags()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	for i := 0; i < len(tr); i++ {
		tag := new(GetTag)
		tag.TagID = tr[i].ID
		tag.TagName = tr[i].TagName
		gettags = append(gettags, tag)

	}

	return
}
