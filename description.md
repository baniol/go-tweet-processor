## web/handlers

Controllers.

`requestHandler` struct has a nested dbConn object

```go
type requestHandler struct {
	dbConn db.DBLayer
}
```

`initHandlers()`

Instance of `requestHandler` struct.

```go
h := new(requestHandler)
```

Db connection:

```
err := h.connect("MONGO", mongoURI)
```

The `connect` function sets the dbConn for the `requestHandler` receiver.

For each handler, the receiver is `requestHandler` stuct.
The receiver contains db methods:

```
count, _ := rh.dbConn.CountTweets()
```

----

## Db package

```
dblayer, err := db.ConnectDatabase(o, conn)
```

### db/dblayer

`DBLayer` interface with db methods.

`ConnectDatabase` returns a handler of a desired db.

### db/mongo

```
type MongoDataStore struct {
	*mgo.Session
}
```

`func NewMongoStore() (*MongoDataStore, error)` constructor.


---

## Mocking mongo - interfaces


---

## Links to testing in Go

http://thylong.com/golang/2016/mocking-mongo-in-golang/

    All of our functions will then deal with interfaces instead of the underlying structs (that can be either a real database or a mock).

http://relistan.com/writing-testable-apps-in-go/ - good and simple explanation of DI and interfaces for testing

---

https://elithrar.github.io/article/testing-http-handlers-go/ - testing handlers

Tip: make strings like application/json or Content-Type package-level constants, so you don’t have to type (or typo) them over and over. A typo in your tests can cause unintended behaviour, becasue you’re not testing what you think you are.

You should also make sure to test not just for success, but for failure too: test that your handlers return errors when they should (e.g. a HTTP 403, or a HTTP 500).



---
https://gist.github.com/nmerouze/2e26a02d23c4c62173fd - example of how to apply Collection + modify Header with next + reflect.TypeOf

https://github.com/thylong/regexrace !! - tests + deploy with kubectl

https://github.com/thylong/mongo_mock_go_example

https://www.youtube.com/watch?v=yszygk1cpEc - advanced testing with Go

http://www.unixstickers.com/image/cache/data/stickers/golang/golang.sh-600x600.png

https://medium.com/@mvmaasakkers/writing-integration-tests-with-mongodb-support-231580a566cd - integration testing


https://developers.almamedia.fi/painless-mongodb-testing-with-docker-and-golang/ - testing with docker

https://github.com/modocache/signatures/blob/master/signatures/server_test.go - tests, before, after, describe

https://blog.envimate.com/2016/04/21/gomongo-part-2/ - microservice tutorial, docker ?