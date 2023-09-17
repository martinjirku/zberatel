package hasuraauth

type RequestEmailRegistration struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Options  Options `json:"options"`
}
type Options struct {
	Locale       string   `json:"locale"`
	DefaultRole  string   `json:"defaultRole"`
	AllowedRoles []string `json:"allowedRoles"`
	DisplayName  string   `json:"displayName"`
	Metadata     Metadata `json:"metadata"`
	RedirectTo   string   `json:"redirectTo"`
}
