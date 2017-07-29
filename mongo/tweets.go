package mongo

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

// type DBgetter interface {
// 	CountTweets() (int, error)
// }

type MongoDataStore struct {
	*mgo.Session
}

func NewMongoStore() *MongoDataStore {
	mongoURI := fmt.Sprintf("%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	log.Printf("Connecting to mongodb at: %s\n", mongoURI)
	session, err := mgo.Dial(mongoURI)
	if err != nil {
		log.Fatalf("Error connecting to the mongodb at %s\n", mongoURI)
	}
	return &MongoDataStore{Session: session}
}

func (ms *MongoDataStore) getCollection() *mgo.Collection {
	return ms.Session.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MONGO_COLLECTION"))
}

// @TODO separate file

// @TODO DBgetter interface as receiver?
func (ms *MongoDataStore) CountTweets() (int, error) {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()
	return coll.Count()
}

func (ms *MongoDataStore) GetAuthors() []bson.M {
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

	results := []bson.M{}
	err1 := pipe.All(&results)
	// @TODO to return - bubble up
	if err1 != nil {
		log.Fatalf("ERROR : %s\n", err1.Error())
	}
	return results
}

func (ms *MongoDataStore) CountTags() []bson.M {
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

	results := []bson.M{}
	err1 := pipe.All(&results)
	// @TODO to return - bubble up
	if err1 != nil {
		log.Fatalf("ERROR : %s\n", err1.Error())
	}
	return results
}

//@TODO return else than interface{}
func (ms *MongoDataStore) GetAuthorTweets(name string) interface{} {
	session := ms.Copy()
	defer session.Close()
	coll := ms.getCollection()

	//@TODO change to Twit type
	names := []struct {
		Name string `bson:"text"`
	}{}

	err := coll.Find(bson.M{"user.name": name}).Select(bson.M{"text": 1}).Sort("-_id").All(&names)
	// @TODO return and bubble error
	if err != nil {
		log.Fatal(err)
	}
	return names
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
