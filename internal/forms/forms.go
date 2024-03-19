package forms

import (
	"net/http"
	"net/url"
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
