package main

import (
	"fmt"

	"github.com/gothite/forms"
	"github.com/gothite/forms/fields"
	"github.com/gothite/forms/validators"
)

type LoginFormData struct {
	Email    string `forms:"email"`
	Password string `forms:"password"`
}

func (form *LoginFormData) Clean() error {
	// Check data
	return nil
}

// LoginForm handles user login.
var LoginForm = forms.NewForm(
	&fields.Email{
		Name:   "email",
		Errors: map[string]string{"Invalid": "Please, set a valid email."},
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
	data := map[string]interface{}{"email": "me@pyvimcom", "password": "pass"}

	if ok, errors := LoginForm.Validate(&form, data); !ok {
		for field, err := range errors {
			fmt.Printf("%v error: %v\n", field, err)
		}
	} else {
		fmt.Printf("Email: %s\n", form.Email)
		fmt.Printf("Password: %s\n", form.Password)

		if err := form.Clean(); err != nil {
			fmt.Printf("Form error: %v\n", err)
		}
	}
}
