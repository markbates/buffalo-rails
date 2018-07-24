package sessions

import "github.com/gobuffalo/envy"

const DefaultSalt = "encrypted cookie"            // default value for Rails 4 app
const DefaultSignSalt = "signed encrypted cookie" // default value for Rails 4 app

type Session struct {
	Name          string
	SecretKeyBase string // can be found in config/initializers/secret_token.rb or config/secrets.yml
	Salt          string
	SignSalt      string
}

func New(name string) Session {
	s := Session{
		Name:          name,
		SecretKeyBase: envy.Get("SECRET_KEY_BASE", ""),
		Salt:          envy.Get("RAILS_SALT", DefaultSalt),
		SignSalt:      envy.Get("RAILS_SIGN_SALT", DefaultSignSalt),
	}
	return s
}
