package mongodb

import "github.com/maratIbatulin/mongodb/mongo"

// Collections структура для коллекций, которые используются сервисом в mongo
type Collections struct {
	ContactPersons mongo.Collection
}
