package service

import (
	"TH0400/logger"
	"TH0400/repo"
)

// GetTopicsCondition 获取一堆话题
type GetTopicsCondition struct {
	UserID   int `json:"user_id"`
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

// GetTopics 获取一堆话题
type GetTopics struct {
	TopicID      int    `json:"topic_id"`
	TopicName    string `json:"topic_name"`
	TopicContent string `json:"topic_content"`
}

// GetTopicCondition ...
type GetTopicCondition struct {
	TopicID int `json:"topic"`
}

// GetTopic ...
type GetTopic struct {
	TopicID      int    `json:"topic_id"`
	TopicName    string `json:"topic_name"`
	TopicContent string `json:"topic_content"`
}

// GetTopics ...
func (ts *GetTopicsCondition) GetTopics() (gettopics []*GetTopics, err error) {

	etopic := repo.Topicrepo{
		CreaterID: ts.UserID,
	}

	tr, err := etopic.GetTopicsByUserID()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	for i := 0; i < ts.PageSize; i++ {
		gettopics[i].TopicID = tr[i].ID
		gettopics[i].TopicName = tr[i].Title
		gettopics[i].TopicContent = tr[i].Content
	}

	return
}

// GetTopic ...
func (ts *GetTopicCondition) GetTopic() (gettopic *GetTopic, err error) {

	etopic := repo.Topicrepo{
		ID: ts.TopicID,
	}

	tr, err := etopic.GetTopicByTopicID()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	gettopic.TopicID = tr.ID
	gettopic.TopicName = tr.Title
	gettopic.TopicContent = tr.Content

	return
}
