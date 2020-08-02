// Code generated by entc, DO NOT EDIT.

package ent

import (
	"TH0400/repo/ent/answer"
	"TH0400/repo/ent/schema"
	"TH0400/repo/ent/tag"
	"TH0400/repo/ent/topic"
	"TH0400/repo/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescContent is the schema descriptor for content field.
	answerDescContent := answerFields[1].Descriptor()
	// answer.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	answer.ContentValidator = answerDescContent.Validators[0].(func(string) error)
	// answerDescCreatedAt is the schema descriptor for created_at field.
	answerDescCreatedAt := answerFields[2].Descriptor()
	// answer.DefaultCreatedAt holds the default value on creation for the created_at field.
	answer.DefaultCreatedAt = answerDescCreatedAt.Default.(func() time.Time)
	// answerDescUpdatedAt is the schema descriptor for updated_at field.
	answerDescUpdatedAt := answerFields[3].Descriptor()
	// answer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	answer.DefaultUpdatedAt = answerDescUpdatedAt.Default.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescTagName is the schema descriptor for tag_name field.
	tagDescTagName := tagFields[1].Descriptor()
	// tag.TagNameValidator is a validator for the "tag_name" field. It is called by the builders before save.
	tag.TagNameValidator = tagDescTagName.Validators[0].(func(string) error)
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[2].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescBaned is the schema descriptor for baned field.
	tagDescBaned := tagFields[3].Descriptor()
	// tag.DefaultBaned holds the default value on creation for the baned field.
	tag.DefaultBaned = tagDescBaned.Default.(bool)
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescTitle is the schema descriptor for title field.
	topicDescTitle := topicFields[1].Descriptor()
	// topic.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	topic.TitleValidator = topicDescTitle.Validators[0].(func(string) error)
	// topicDescContent is the schema descriptor for content field.
	topicDescContent := topicFields[2].Descriptor()
	// topic.DefaultContent holds the default value on creation for the content field.
	topic.DefaultContent = topicDescContent.Default.(string)
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicFields[4].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(func() time.Time)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicFields[5].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[1].Descriptor()
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = func() func(string) error {
		validators := userDescUserName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(user_name string) error {
			for _, fn := range fns {
				if err := fn(user_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescLevel is the schema descriptor for level field.
	userDescLevel := userFields[4].Descriptor()
	// user.DefaultLevel holds the default value on creation for the level field.
	user.DefaultLevel = userDescLevel.Default.(int)
	// user.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	user.LevelValidator = userDescLevel.Validators[0].(func(int) error)
	// userDescLikes is the schema descriptor for likes field.
	userDescLikes := userFields[5].Descriptor()
	// user.DefaultLikes holds the default value on creation for the likes field.
	user.DefaultLikes = userDescLikes.Default.(int)
	// userDescViews is the schema descriptor for views field.
	userDescViews := userFields[6].Descriptor()
	// user.DefaultViews holds the default value on creation for the views field.
	user.DefaultViews = userDescViews.Default.(int)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[7].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[8].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescIsDeleted is the schema descriptor for is_deleted field.
	userDescIsDeleted := userFields[9].Descriptor()
	// user.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	user.DefaultIsDeleted = userDescIsDeleted.Default.(bool)
}
