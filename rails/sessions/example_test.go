package sessions_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/markbates/buffalo-rails/rails/sessions"
	"github.com/pkg/errors"
)

const value = `bXF4TUI3RzdkSVlVeElRUk9jZGZ5YkhnTFFHUStzdit6V2k4bnJWVVd5bFF5ZHZWWmQ2T1J0M2NvRkp0ajlWcmFyNnUxVm13TFBNYkdtY29lekJKMk1EdTRYZ2lpSEtvVjgvRjA2RHk0SWZEQis5TUdyL1ZBUUw0Z2hESEhscFdlaFQvOW1uczAzV1ZhNnczWHVUNTBET0loejF1Y0hoWmFmVkorUEpoUGlGYjJHQTFYSkhSSU5LbEFIQ0V4bkJxemp2TmJORjQzOVNTanRrVVFQRDlZSHI0U1pleldHdEhpbXBvMFNEL0JBMEtZTXFZeUoyc2VuZjlQdUphckVMRGFzTC9FLytGamJFOWlrSFhQc3V1QTlqbllnTmFpWHRuR3RTQ0VoaitTU0ZnN0JNZWNuQkZyMWJjTmsrQnBQa1pzaFI0U0d6SFJNWU5sZzJnKzY1QllOcXVYODZhdGl3eGhyUG9vQUlsSm5VaGJ6L1dXTjZQZGNmMzZmUG5FL3E1clFocE5zL0pQU3ZWdHUreGpMRzNQMlFkVWI1NXRvWERGWW5JKzA5MHBuRT0tLWJSUWM5UXFhMkVNSUtqVjNwQnFRaHc9PQ%3D%3D--37989ccb1eb59c4cd7da99eb496a95dd94519f9e`

func Example_Session_Cookie() {
	app := buffalo.New(buffalo.Options{})
	app.GET("/", func(c buffalo.Context) error {
		s := sessions.New("test")
		s.SecretKeyBase = "development"

		m, err := s.Cookie(c)
		if err != nil {
			return errors.WithStack(err)
		}

		// find a value from the rails cookie, in this case a value named "token"
		tok, ok := m["token"].(string)
		if !ok {
			return errors.New("could not find token")
		}
		return c.Render(200, render.String(tok))
	})

	res := httptest.NewRecorder()
	cook := &http.Cookie{Name: "test", Value: value}
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(cook)

	app.ServeHTTP(res, req)

	body := res.Body.String()
	fmt.Println(body)

	// output: 3b67f127e9a82e5b9b81d68e64d69249
}
