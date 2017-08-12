package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBLayer interface {
	CountTweets() (int, error)
	GetAuthors() (interface{}, error)
	GetTags() (interface{}, error)
	GetAuthorTweets(string) (interface{}, error)
}

type MongoDataStore struct {
	*mgo.Session
}

// @TODO data store as a receiver ?
func ConnectMongo() (*MongoDataStore, error) {
	mongoURI := fmt.Sprintf("%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	log.Printf("Connecting to mongodb at: %s\n", mongoURI)
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatalf("Error connecting to the mongodb at %s\n", mongoURI)
	}
	return &MongoDataStore{Session: session}, nil
}

func (ms *MongoDataStore) getCollection() *mgo.Collection {
	return ms.Session.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_COLLECTION"))
}

func (ms *MongoDataStore) CountTweets() (int, error) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()
	return coll.Count()
}

func (ms *MongoDataStore) GetAuthors() (interface{}, error) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()

	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id":    "$user.name",
				"tweets": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"tweets": -1,
			},
		},
	}

	pipe := coll.Pipe(pipeline)

	var results []interface{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoDataStore) GetTags() (interface{}, error) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()

	pipeline := []bson.M{
		bson.M{
			"$project": bson.M{
				"tag": "$entities.hashtags.text",
				"_id": 0,
			},
		},
		bson.M{
			"$unwind": "$tag",
		},
		bson.M{
			"$project": bson.M{"tag": bson.M{"$toLower": "$tag"}},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$tag",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
	}

	pipe := coll.Pipe(pipeline)

	var results []interface{}
	err := pipe.All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

//@TODO return else than interface{}
func (ms *MongoDataStore) GetAuthorTweets(name string) (interface{}, error) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()

	//@TODO change to Tweet type
	names := []struct {
		Name string `bson:"text"`
	}{}

	err := coll.Find(bson.M{"user.name": name}).Select(bson.M{"text": 1}).Sort("-_id").All(&names)
	if err != nil {
		return nil, err
	}
	return names, nil
}

func (ms *MongoDataStore) InsertTweet(tweet []byte) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()
	// @TODO tweet struct
	var tweetDecoded interface{}
	json.Unmarshal(tweet, &tweetDecoded)
	err := coll.Insert(tweetDecoded)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[√] Tweet inserted")
}
