package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

// Valid checks if the form is valid.
//
// It returns a boolean value.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New creates a new Form object.
//
// It takes a data parameter of type url.Values and returns a pointer to a Form object.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks if the given fields are not empty.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if the given field exists in the request's form data.
//
// Parameters:
// - field: the name of the field to check.
// - r: the HTTP request object.
//
// Returns:
// - bool: true if the field exists and has a non-empty value, false otherwise.
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blanck")
		return false
	}
	return x != ""
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// IsEmail checks if a field is an email address.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
