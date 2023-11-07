package entity

import (
	"context"
	"github.com/maratIbatulin/mongodb/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"reforce.pattern/internal/api/cache"
	"reforce.pattern/pkg/mongodb"
	"time"
)

type User struct {
	ReforceID   string
	PartnerID   string
	TokenExpire int64
}

func (u *User) Get(key string) bool {
	val, ok := cache.Get(key)
	if ok {
		*u = *val.(*User)
	}
	return ok
}
func (u *User) Set(key string) {
	cache.Set(key, u, time.Hour*24)
}
func (u *User) Remove(key string) {
	cache.Remove(key)
}

func (u *User) Partner(mdb *mongodb.Collections) error {
	cursor, err := mdb.ContactPersons.Aggregate(context.TODO(), mongo.Filter().Match(bson.M{"reforce_id": u.ReforceID}).Sort(bson.M{"partner": 1}).Limit(1).Project(bson.M{"_id": 1, "partner": "$partner"}))
	if err != nil {
		return err
	}

	if cursor.Next(context.TODO()) {
		u.PartnerID = cursor.Current.Lookup("partner").StringValue()
	}
	return nil
}
