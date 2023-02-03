package utils

import (
	"github.com/google/uuid"
)

/*
Returns a newly created random UUID
*/
func UUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
