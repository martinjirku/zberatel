package hasuraauth

import "time"

type User struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"createdAt"`
	DisplayName         string    `json:"displayName"`
	AvatarURL           string    `json:"avatarUrl"`
	Locale              string    `json:"locale"`
	Email               string    `json:"email"`
	IsAnonymous         bool      `json:"isAnonymous"`
	DefaultRole         string    `json:"defaultRole"`
	Metadata            Metadata  `json:"metadata"`
	ActiveMfaType       string    `json:"activeMfaType"`
	EmailVerified       bool      `json:"emailVerified"`
	PhoneNumber         string    `json:"phoneNumber"`
	PhoneNumberVerified bool      `json:"phoneNumberVerified"`
	Roles               []string  `json:"roles"`
}
type Metadata struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
