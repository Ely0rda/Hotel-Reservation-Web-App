package forms

import (
	"errors"

	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form a struct for embeding
// url.Values => form values
// Errors => Errors that gets generated when the user
// enters the wrong values to the form

type Form struct {
	//type Values map[string][]string
	//for form values
	url.Values
	Errors Errors
}

// Valid return true if there are no errors
// ,otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struuct
func New(data url.Values) *Form {
	return &Form{
		data,
		Errors(map[string][]string{}),
		//errors(map[string][]string{}),
	}
}

// Required check for all the required fields that has
// been passed to it
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//Has checks if form fields is in post and not empty

func (f *Form) Has(field string) bool {
	x := f.Get(field)

	return x != ""
}

// MinLength checking the length of the fields

func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checking the corecteness of an email
func (f *Form) IsEmail(field string) error {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
		return errors.New("can't get template from cache")
	}
	return nil
}
