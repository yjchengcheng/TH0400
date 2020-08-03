package service

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTopic(t *testing.T) {

	userregister := CreateTopic{
		Title:      "怎么学习英语?",
		Content:    "如何有效的学习英语....................",
		IsReleased: false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CreateID:   1,
	}

	err := userregister.CreateTopic()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(err)
}

func TestGetTopic(t *testing.T) {
	userregister := GetTopicCondition{
		TopicID: 1,
	}

	res, err := userregister.GetTopic()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(res)
}

func TestGetTopics(t *testing.T) {
	userregister := GetTopicsCondition{
		UserID:   1,
		PageSize: 1,
		PageNum:  2,
	}

	res, err := userregister.GetTopics()
	if err != nil {
		t.Error(err.Error())
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].TopicID)
		fmt.Println(res[i].TopicName)
		fmt.Println(res[i].TopicContent)
	}
	t.Log(res)
}
