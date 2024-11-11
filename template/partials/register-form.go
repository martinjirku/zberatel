package partials

import "errors"

type RegisterFormInput struct {
	Username             string
	Password             string
	PasswordConfirmation string
	Email                string
	CsfrToken            string
}

type RegisterForm struct {
	Form         RegisterFormInput
	Errors       map[string][]error
	GlobalError  error
	RecaptchaKey string
}

func (f *RegisterForm) WithGlobalError(errorText string) *RegisterForm {
	f.GlobalError = errors.New(errorText)
	return f
}

func (f *RegisterForm) SetError(fieldName, message string) {
	if _, ok := f.Errors[fieldName]; !ok {
		f.Errors[fieldName] = []error{}
	}
	f.Errors[fieldName] = append(f.Errors[fieldName], errors.New(message))
}

func (f *RegisterForm) HasError() bool {
	return f.GlobalError != nil || len(f.Errors) > 0
}
