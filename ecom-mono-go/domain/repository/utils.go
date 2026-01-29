package repository

import "strings"

// We could just modify this to handle all email duplication error but for future there might be other unique constraints so not good practice
func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "23505")
}