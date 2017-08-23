# Go Tweet Processor

## TODO

* work on api responses, errors, etc.
* log level setup to env variable
* log db conn errors to logger, output internal error to http res
* mongo index optimisation
* unit tests for error cases - table driven tests ?
* content size , chunked transfer ? how to return body size?
* marshal or encode ?
* mongo reconnects
* check how echo handles writing JSON to res
* recover from panic on handler erros ?
* access log
* context
* middleware
* document code
* travis - see gorilla/handlers for example


---

## TODO sequence

1. Unit tests
    * table driven [LATER]
    * test returning errors [DONE]
    * test headers

2. Logger
    * with loglevels [DONE]
    * loglevel to env variable

3. Better way of error handling

4. Middleware - context ?
    * adding headers ?
    * auth - jwt
    * access logging [DONE] - with gorilla/handlers
        * missing request duration

5. MongoDB reconnect / circuit breaker / backoff ?

6. Abstract away configuration from os.Getenv



---
## Read:

### HTTP server general

* https://cryptic.io/go-http/
* http://www.alexedwards.net/blog/organising-database-access - Env{}, interfaces, testing, context !!!
* https://gowebexamples.github.io/
* https://www.rickyanto.com/understanding-go-standard-http-libraries-servemux-handler-handle-and-handlefunc/
* http://modocache.io/restful-go - rest server with many external libs, ex. for testing
* https://medium.com/@jduv/simple-handler-tdd-in-golang-66fe2fa89a64
* https://www.nicolasmerouze.com/how-to-render-json-api-golang-mongodb/ - mongodb, json !!
* https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql - !!! - finish

### Testing

* advanced testing video !

### Context

* 

### Logging

* https://github.com/mash/go-accesslog - access logs example

### Middleware

* https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/ - context, finish reading ! (notes in dev-notes: `terraform-basics.md`)
* http://www.alexedwards.net/blog/making-and-using-middleware
* https://www.nicolasmerouze.com/share-values-between-middlewares-context-golang/ - from series
* https://gist.github.com/cespare/3985516 - apache access logging
* https://stackoverflow.com/questions/20987752/how-to-setup-access-error-log-for-http-listenandserve