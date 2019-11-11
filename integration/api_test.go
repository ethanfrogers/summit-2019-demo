package integration

import (
	"flag"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	endpoint = flag.String("endpoint", "http://localhost:3000", "endpoint to run tests against")
)

func TestProperHelloMessage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	flag.Parse()
	helloUrl := *endpoint + "/hello"
	resp, err := http.Get(helloUrl)
	if err != nil {
		t.Fatalf("got unexpected error when calling hello endpoint: %s", err.Error())
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	if string(b) != "hello spinnaker summit 2019!" {
		t.Fatalf("response doesn't equal spin sum greeting. got %s", string(b))
	}
}

func TestProperVersion(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	flag.Parse()
	helloUrl := *endpoint + "/version"
	resp, err := http.Get(helloUrl)
	if err != nil {
		t.Fatalf("got unexpected error when calling hello endpoint: %s", err.Error())
	}

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	if string(b) == "development" {
		t.Fatalf("application seems to be in development mode. are you sure you set the right variables!?")
	}
}
