package auth

import "context"

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"User": a.User, "password": a.Password}, nil
}

func (c *Authentication) RequireTransportSecurity() bool {
	return false
}
