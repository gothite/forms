# gothite/forms #
*Form handling and data validation for Go*

## Installation ##

`go get github.com/gothite/forms`


## Usage ##
```go
package main

import (
        "fmt"

        "github.com/gothite/forms"
        "github.com/gothite/forms/fields"
        "github.com/gothite/forms/validators"
)

func authorize(form *forms.Form) error {
        // check user data
        // form.Fields["email"].Value
        // form.Fields["password"].Value
        return nil
}

// LoginForm handles user login.
var LoginForm = forms.NewFormFactory(
        authorize, // or nil
        &fields.Email{
                Name:   "email",
                Errors: map[string]string{"Invalid": "Please, set a valid email."},
        },
        &fields.String{Name: "password", MinLength: validators.MinLength(5)},
)

func main() {
        data := map[string]interface{}{"email": "me@pyvimcom", "password": "pass"}
        form := LoginForm(data)

        if !form.IsValid() {
                fmt.Printf("Form error: %v\n", form.Error)

                for name, field := range form.Fields {
                        fmt.Printf("%v error: %v\n", name, field.Error)
                }
        }
}
```
