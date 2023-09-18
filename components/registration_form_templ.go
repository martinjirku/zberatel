// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func RegistrationForm(username string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<form action=\"/auth/login\" class=\"is-3\"><h1 class=\"is-full title is-spaced\">")
		if err != nil {
			return err
		}
		var_2 := `Registrácia`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><div class=\"field\"><label class=\"control\" for=\"firstname\">")
		if err != nil {
			return err
		}
		var_3 := `Krstné meno`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input class=\"input\" type=\"text\" id=\"username\" placeholder=\"Username\"></div><div class=\"field\"><label class=\"control\" for=\"lastname\">")
		if err != nil {
			return err
		}
		var_4 := `Priezvisko`
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input class=\"input\" type=\"text\" id=\"username\" placeholder=\"Username\"></div><div class=\"field\"><label class=\"control\" for=\"username\">")
		if err != nil {
			return err
		}
		var_5 := `Používateľské meno`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input class=\"input\" type=\"text\" id=\"username\" placeholder=\"Username\"></div><div class=\"field\"><label class=\"control\" for=\"password\">")
		if err != nil {
			return err
		}
		var_6 := `Heslo`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><input class=\"input\" type=\"password\" id=\"password\" placeholder=\"Password\"></div><div class=\"field is-grouped\"><div class=\"control\"><button type=\"submit\" class=\"button is-link\">")
		if err != nil {
			return err
		}
		var_7 := `Registrovať`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button></div><div class=\"control\"><a href=\"/auth/login\" class=\"button is-link is-light\">")
		if err != nil {
			return err
		}
		var_8 := `Prihlásenie`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div></div></form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
