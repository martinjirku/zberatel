package partials

import "errors"

type LoginFormInput struct {
	Username  string
	Password  string
	CsfrToken string
}

type LoginForm struct {
	Form         LoginFormInput
	Errors       map[string]error
	GlobalError  error
	RecaptchaKey string
}

func (f *LoginForm) WithGlobalError(errorText string) *LoginForm {
	f.GlobalError = errors.New(errorText)
	return f
}
