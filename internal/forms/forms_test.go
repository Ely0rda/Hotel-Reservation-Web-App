package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {

	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}

}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required foelds missing")
	}
	posteData := url.Values{}
	posteData.Add("a", "a")
	posteData.Add("b", "a")
	posteData.Add("c", "a")
	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = posteData
	form = New(r.PostForm)
	form.Required("a", "c", "b")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {

	posteData := url.Values{}
	posteData.Add("a", "a")
	r, _ := http.NewRequest("POST", "/whatever", nil)
	r.PostForm = posteData
	form := New(r.PostForm)
	if !form.Has("a") {
		t.Error("show field does not exist when it is ")
	}
	r, _ = http.NewRequest("POST", "/whatever", nil)
	form = New(r.PostForm)
	if form.Has("a") {
		t.Error("show exist, when it is not")
	}

}

func TestForm_MinLength(t *testing.T) {
	//The field doesn't exist
	postedValues := url.Values{}
	form := New(postedValues)
	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form show min length for non-existent field")
	}
	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}
	//data exist but shorted then the length given to
	//the function
	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)
	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}
	//valid : data exist and longer then the length
	//given to the function
	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}
	isError = form.Errors.Get("x")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("form show valid email for non-existing field")
	}
	postedValues = url.Values{}
	postedValues.Add("email", "a@a.com")
	form = New(postedValues)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}
	postedValues = url.Values{}
	postedValues.Add("email", "x")
	form = New(postedValues)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid meail when we should not have one")
	}

}
