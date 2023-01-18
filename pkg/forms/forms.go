package forms

import (
	"net/http"
	"net/url"
)

// holds information that is assigned to our forms.
type Form struct {
	url.Values
	Errors errors
}

// This function creates a new form in our codebase.
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// This function checks if a field with given id has content in the form.
func (f *Form) HasValue(tagID string, r *http.Request) bool {
	x := r.Form.Get(tagID)
	if x == "" {
		f.Errors.AddError(tagID, "Field empty") // add error if has no content
		return false
	}
	return true
}

// This function checks if the form is valid.
func (f *Form) IsValid() bool {
	return len(f.Errors) == 0 // form is valid if there are no errors.
}
