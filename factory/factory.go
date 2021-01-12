package factory

import (
	"cache-service/configuration"
	"cache-service/constants"
	"github.com/patrickmn/go-cache"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var mgoSession *mgo.Session

func GetDbSession(config *configuration.ConfigData) (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{config.Database.Host},
			Database: config.Database.Name,
			Username: config.Database.Username,
			Password: config.Database.Password,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatal("Failed to get mongo mgoSession ", err)
			return nil, err
		}
	}
	return mgoSession, nil
}

func GetCache() (*cache.Cache, error) {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	return cache.New(constants.DefaultInterval*time.Minute, constants.CleanupInterval*time.Minute), nil
}

func GetMongoCollection(config *configuration.ConfigData) *mgo.Collection {
	session, err := GetDbSession(config)
	if err != nil {
		log.Fatal("unable to connect to db", err)
		return nil
	}
	return session.DB(config.Database.Name).C(config.Database.MongoCollectionName)
}
