package i18n

import (
	"embed"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed *.toml
var LocaleFS embed.FS

var (
	bundle = i18n.NewBundle(language.English)
)

func InitializeI18N() error {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(LocaleFS, "en.toml")
	bundle.LoadMessageFileFS(LocaleFS, "sk.toml")
	return nil
}

func Localizer(r *http.Request) *i18n.Localizer {
	accept := r.Header.Get("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, accept, "en")
	return localizer
}
