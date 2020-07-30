package repo

import (
	"TH0400/repo/ent"
	"TH0400/repo/ent/topic"
	"TH0400/utils"
	"fmt"
)

//Topicrepo ...
type Topicrepo ent.Topic

//GetTopicsByUserID 获取topics
func (t *Topicrepo) GetTopicsByUserID() ([]*ent.Topic, error) {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if t.CreaterID != 0 {
		tr, err := edb.Topic.
			Query().
			Where(topic.IDEQ(t.CreaterID)).
			All(ctx)
		return tr, err
	}

	err := fmt.Errorf("show me ID")

	return nil, err
}

//GetTopicByTopicID 获取一个话题通过话题ID
func (t *Topicrepo) GetTopicByTopicID() (*Topicrepo, error) {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if t.ID != 0 {
		tr, err := edb.Topic.
			Query().
			Where(topic.IDEQ(t.ID)).
			Only(ctx)
		return (*Topicrepo)(tr), err
	}

	err := fmt.Errorf("show me ID")

	return nil, err
}
