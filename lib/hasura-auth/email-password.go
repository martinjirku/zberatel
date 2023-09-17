package hasuraauth

type RequestEmailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseEmailPassword struct {
	Session Session `json:"session"`
	MFA     *MFA    `json:"mfa"`
}

type MFA struct {
	Ticket string `json:"ticket"`
}

type Session struct {
	AccessToken          string `json:"accessToken"`
	AccessTokenExpiresIn int    `json:"accessTokenExpiresIn"`
	RefreshToken         string `json:"refreshToken"`
	User                 User   `json:"user"`
}
