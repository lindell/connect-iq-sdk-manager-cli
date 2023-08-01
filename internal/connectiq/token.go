package connectiq

import "time"

// Token contains information about the token used for requests to the Garmin API
type Token struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}
