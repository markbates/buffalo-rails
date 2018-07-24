package sessions

import (
	"encoding/json"
	"net/http"

	"github.com/adjust/gorails/session"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

func (s Session) Cookie(c buffalo.Context) (map[string]interface{}, error) {
	cmap := map[string]interface{}{}

	cookie, err := c.Cookies().Get(s.Name)
	if err != nil {
		if errors.Cause(err) == http.ErrNoCookie {
			return cmap, nil
		}
		return cmap, errors.WithStack(err)
	}

	d, err := session.DecryptSignedCookie(cookie, s.SecretKeyBase, s.Salt, s.SignSalt)
	if err != nil {
		return cmap, errors.WithStack(err)
	}

	if err := json.Unmarshal(d, &cmap); err != nil {
		c.Cookies().Delete(s.Name)
		return cmap, nil
	}
	return cmap, nil
}
