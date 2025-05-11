package database

import (
	"context"

	"github.com/sahay-shashank/mongodb-server/internal/core/details"
)

func CloseDatabase() details.APIDetails {
	if err := mongoClient.Disconnect(context.Background()); err != nil {
		return details.APIDetails{
			Error:      true,
			StatusCode: details.MongoDBConnectionCloseFailure,
			Message:    details.GetMessage(details.MongoDBConnectionCloseFailure),
		}
	}
	return details.APIDetails{
		StatusCode: details.MongoDBConnectionCloseSuccessful,
		Message:    details.GetMessage(details.MongoDBConnectionCloseSuccessful),
	}
}
