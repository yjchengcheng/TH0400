package service

import (
	"TH0400/logger"
	"TH0400/repo"
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
		gettags[i].TagID = tr[i].ID
		gettags[i].TagName = tr[i].TagName

	}

	return
}
