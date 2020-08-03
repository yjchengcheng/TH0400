package service

import (
	"TH0400/logger"
	"TH0400/repo"
	"time"
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
	TopicName    string `json:"topic_name"`
	TopicContent string `json:"topic_content"`
}

// CreateTopic ...
type CreateTopic struct {
	Title      string    `json:"Title"`
	Content    string    `json:"content"`
	IsReleased bool      `json:"is_released"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreateID   int       `json:"create_id"`
}

// CreateTopic ...
func (ct *CreateTopic) CreateTopic() error {

	etopic := repo.Topicrepo{
		Title:      ct.Title,
		Content:    ct.Content,
		IsReleased: ct.IsReleased,
		CreatedAt:  ct.CreatedAt,
		UpdatedAt:  ct.UpdatedAt,
		CreaterID:  ct.CreateID,
	}

	err := etopic.CreateTopic()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return err
	}

	return nil
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
		top := new(GetTopics)
		top.TopicID = tr[i].ID
		top.TopicName = tr[i].Title
		top.TopicContent = tr[i].Content
		gettopics = append(gettopics, top)
	}

	return
}

// GetTopic ...
func (ts *GetTopicCondition) GetTopic() (gettopic *GetTopic, err error) {
	gettopic = new(GetTopic)

	etopic := repo.Topicrepo{
		ID: ts.TopicID,
	}

	tr, err := etopic.GetTopicByTopicID()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	gettopic.TopicName = tr.Title
	gettopic.TopicContent = tr.Content

	return
}
