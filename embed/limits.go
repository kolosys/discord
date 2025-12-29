package embed

import "errors"

// Discord embed limits
const (
	MaxTitleLength       = 256
	MaxDescriptionLength = 4096
	MaxFields            = 25
	MaxFieldNameLength   = 256
	MaxFieldValueLength  = 1024
	MaxFooterLength      = 2048
	MaxAuthorNameLength  = 256
	MaxTotalLength       = 6000
)

// Sentinel errors for embed validation
var (
	ErrTitleTooLong       = errors.New("embed: title exceeds 256 characters")
	ErrDescriptionTooLong = errors.New("embed: description exceeds 4096 characters")
	ErrTooManyFields      = errors.New("embed: exceeds 25 fields maximum")
	ErrFieldNameTooLong   = errors.New("embed: field name exceeds 256 characters")
	ErrFieldValueTooLong  = errors.New("embed: field value exceeds 1024 characters")
	ErrFooterTooLong      = errors.New("embed: footer text exceeds 2048 characters")
	ErrAuthorNameTooLong  = errors.New("embed: author name exceeds 256 characters")
	ErrEmbedTooLarge      = errors.New("embed: total length exceeds 6000 characters")
)
