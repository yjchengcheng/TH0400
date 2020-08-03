package repo

import (
	"TH0400/logger"
	"TH0400/repo/ent"
	"TH0400/repo/ent/tag"
	"TH0400/repo/ent/tagtopic"
	"TH0400/utils"
	"fmt"
)

// TagTopicRepo ..
type TagTopicRepo ent.TagTopic

// TagRepo ..
type TagRepo ent.Tag

var tagnames []*ent.Tag

//CreateTag 创建tags
func (tt *TagRepo) CreateTag() error {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	_, err := edb.Tag.
		Create().
		SetTagName(tt.TagName).
		SetCreatedAt(tt.CreatedAt).
		SetBaned(tt.Baned).
		Save(ctx)
	if err != nil {
		logger.Errorf("failed create Tag: %s", err.Error())
		return err
	}

	return nil
}

//GetTags 获取tags
func (tt *TagTopicRepo) GetTags() ([]*ent.Tag, error) {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if tt.TopicID != 0 {
		ttr, err := edb.TagTopic.
			Query().
			Where(tagtopic.TagIDEQ(tt.TopicID)).
			All(ctx)

		for i := 0; i < len(ttr); i++ {
			tagr, _ := edb.Tag.Query().Where(tag.IDEQ(ttr[i].TagID)).Only(ctx)

			tagnames[i] = tagr
		}
		return tagnames, err

	}

	err := fmt.Errorf("show me ID")

	return nil, err
}
