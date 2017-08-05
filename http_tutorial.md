https://cryptic.io/go-http/

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func main() {
	err := http.ListenAndServe(":9999", helloHandler{})
	log.Fatal(err)
}
```

The http package provides a helper function, `http.HandlerFunc`, which wraps a function which has the signature `func(w http.ResponseWriter, r *http.Request)`, returning an `http.Handler` which will call it in all cases.

The following behaves exactly like the previous example, but uses `http.HandlerFunc` instead of defining a new type.

``` go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

### ServeMux

On their own, the previous examples don’t seem all that useful. If we wanted to have different behavior for different endpoints we would end up with having to parse path strings as well as numerous if or switch statements. Luckily we’re provided with `http.ServeMux`, which does all of that for us. Here’s an example of it being used:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit foo!")
	})

	h.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit bar!")
	})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "You're lost, go home")
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

The `http.ServeMux` is itself an `http.Handler`, so it can be passed into `http.ListenAndServe`. When it receives a request it will check if the request’s path is prefixed by any of its known paths, choosing the longest prefix match it can find. We use the / endpoint as a catch-all to catch any requests to unknown endpoints. 

`http.ServeMux` has both `Handle` and `HandleFunc` methods. These do the same thing, except that Handle takes in an `http.Handler` while `HandleFunc` merely takes in a function, implicitly wrapping it just as `http.HandlerFunc` does.

### Other muxes

There are numerous replacements for `http.ServeMux` like `gorilla/mux` which give you things like automatically pulling variables out of paths, easily asserting what http methods are allowed on an endpoint, and more. Most of these replacements will implement `http.Handler` like `http.ServeMux` does, and accept `http.Handlers` as arguments.

## Composability

When I say that the http package is composable I mean that it is very easy to create re-usable pieces of code and glue them together into a new working application. The `http.Handler` interface is the way all pieces communicate with each other.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	err := http.ListenAndServe(":9999", h)
	log.Fatal(err)
}
```

`numberDumper` implements `http.Handler`, and can be passed into the `http.ServeMux` multiple times to serve multiple endpoints.

## Testing

The httptest package provides a few handy utilities, including `NewRecorder` which implements `http.ResponseWriter` and allows you to effectively make an http request by calling ServeHTTP directly. Here’s an example of a test for our previously implemented `numberDumper`:

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	. "testing"
)

func TestNumberDumper(t *T) {
	// We first create the http.Handler we wish to test
	n := numberDumper(1)

	// We create an http.Request object to test with. The http.Request is
	// totally customizable in every way that a real-life http request is, so
	// even the most intricate behavior can be tested
	r, _ := http.NewRequest("GET", "/one", nil)

	// httptest.Recorder implements the http.ResponseWriter interface, and as
	// such can be passed into ServeHTTP to receive the response. It will act as
	// if all data being given to it is being sent to a real client, when in
	// reality it's being buffered for later observation
	w := httptest.NewRecorder()

	// Pass in our httptest.Recorder and http.Request to our numberDumper. At
	// this point the numberDumper will act just as if it was responding to a
	// real request
	n.ServeHTTP(w, r)

	// httptest.Recorder gives a number of fields and methods which can be used
	// to observe the response made to our request. Here we check the response
	// code
	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	// We can also get the full body out of the httptest.Recorder, and check
	// that its contents are what we expect
	body := w.Body.String()
	if body != fmt.Sprintf("Here's your number: 1\n") {
		t.Fatalf("wrong body returned: %s", body)
	}

}
```