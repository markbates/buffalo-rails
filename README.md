# github.com/markbates/buffalo-rails

A set of tools to make it easy to work with legacy Rails applications from a [https://gobuffalo.io](https://gobuffalo.io) application.

## Installation

```bash
$ go get -u -v github.com/markbates/buffalo-rails
```

## Sessions

The `github.com/markbates/buffalo-rails/rails/sessions` package provides tool for reading from a Rails session.

Currently only cookie sessions are supported. PR's are welcome for others.

### Reading a Rails Session Cookie

```go
app := buffalo.New(buffalo.Options{})
app.GET("/", func(c buffalo.Context) error {
  s := sessions.New("test")

  // set the secret key based used to create the cookie
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
```
