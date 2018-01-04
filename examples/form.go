package main

import (
	"errors"
	"fmt"

	"github.com/gothite/forms"
	"github.com/gothite/forms/codes"
	"github.com/gothite/forms/fields"
	"github.com/gothite/forms/validators"
)

type LoginFormData struct {
	Email    string `forms:"email"`
	Password string `forms:"password"`
}

func (data *LoginFormData) Clean(form *forms.Form) error {
	// Check data
	return nil
}

// LoginForm handles user login.
var LoginForm = forms.NewForm(
	map[uint]error{codes.Invalid: errors.New("Please, check data.")},
	&fields.Email{
		Name:   "email",
		Errors: map[uint]string{codes.Invalid: "Please, set a valid email."},
	},
	&fields.String{
		Name: "password",
		Validators: []validators.StringValidator{
			validators.StringMinLength(5),
		},
	},
)

func main() {
	var form LoginFormData
	data := map[string]interface{}{"email": "hello@binlockme", "password": "pass"}

	if err, errors := LoginForm.Validate(&form, data); err != nil {
		fmt.Printf("Form error: %v\n", err)

		for field, err := range errors {
			fmt.Printf("%v error: %v\n", field, err)
		}
	} else {
		fmt.Printf("Email: %s\n", form.Email)
		fmt.Printf("Password: %s\n", form.Password)
	}
}
