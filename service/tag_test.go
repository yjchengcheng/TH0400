package service

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTag(t *testing.T) {
	tag := CreateTag{
		Name:      "学习",
		CreatedAt: time.Now(),
		Baned:     false,
	}

	err := tag.CreateTag()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(err)
}

func TestGetTag(t *testing.T) {
	tag := GetTagsCondition{
		TopicID: 1,
	}

	res, err := tag.GetTags()
	if err != nil {
		t.Error(err.Error())
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].TagID)
		fmt.Println(res[i].TagName)
	}

}
