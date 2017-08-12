# Go Tweet Processor

## TODO

* work on api responses, errors, etc.
* log levels (logrus, echo logger ?)
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


---

## TODO sequence

1. Unit tests
    * table driven [LATER]
    * test returning errors [DONE]
    * test headers

2. Logger with loglevels

3. Better way of error handling

4. Middleware - context ?
    * adding headers ?
    * auth - jwt
    * access logging

5. MongoDB reconnect / circuit breaker / backoff ?

6. Abstract away configuration from os.Getenv



---
## Read:

### HTTP server general

* https://cryptic.io/go-http/
* http://www.alexedwards.net/blog/organising-database-access - Env{}, interfaces, testing, context !!
* https://gowebexamples.github.io/
* https://www.rickyanto.com/understanding-go-standard-http-libraries-servemux-handler-handle-and-handlefunc/

### Testing

* advanced testing video

### Context

* https://joeshaw.org/revisiting-context-and-http-handler-for-go-17/