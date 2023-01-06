package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTest = []struct {
	name               string     //the name of the test which correspond to the handler
	url                string     //url the call the handler
	method             string     //type of method the handler deal with
	params             []postData // if its is a post request there will be some data there
	expectedStatusCode int        //we will be using status code to check the corectness of our handler
}{
	//Test cases for all of our GeT handlers
	//A test case corectnesse is determined by the status returned
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"rs", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-resrvation-post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "JOhn"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555-555-55"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	//This is a server provided by golang for testing
	//we pass to it the handler routes that has all the routes configured
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTest {
		//There are two types of tet we will be dealing with
		//get requests
		//and post requests
		if e.method == "GET" {
			//The server provide us with a client so making request now is easy
			//client is like if someone called the web server and made requests from his browser
			//we can get the response from the server directly
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			//if the response status code differnet from the expected one
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {

			//if it is a post request
			//valuse for storing our form inputs and their values
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			//making post request with data(keys, values)
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
