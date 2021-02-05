package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func IsDuplKeyError(err error) bool {
	if writeErr, ok := err.(mongo.WriteException); ok {
		if len(writeErr.WriteErrors) > 0 {
			if writeErr.WriteErrors[0].Code == Err11000DupKeyError {
				return true
			}
		}
	}

	return false
}
