// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSchool holds the string denoting the school field in the database.
	FieldSchool = "school"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldLikes holds the string denoting the likes field in the database.
	FieldLikes = "likes"
	// FieldViews holds the string denoting the views field in the database.
	FieldViews = "views"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserName,
	FieldPassword,
	FieldSchool,
	FieldLevel,
	FieldLikes,
	FieldViews,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldIsDeleted,
}

var (
	// UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultLevel holds the default value on creation for the level field.
	DefaultLevel int
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int) error
	// DefaultLikes holds the default value on creation for the likes field.
	DefaultLikes int
	// LikesValidator is a validator for the "likes" field. It is called by the builders before save.
	LikesValidator func(int) error
	// DefaultViews holds the default value on creation for the views field.
	DefaultViews int
	// ViewsValidator is a validator for the "views" field. It is called by the builders before save.
	ViewsValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// DefaultIsDeleted holds the default value on creation for the is_deleted field.
	DefaultIsDeleted bool
)
