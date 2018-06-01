package auth

import (
	"bytes"
	"net/url"
	"strings"
)

type Endpoint struct {
	AuthURL  string
	TokenURL string
}

type Config struct {
	// Application ID.
	ClientID string

	// Application secret.
	ClientSecret string

	// Endpoint contains the resource server's token endpoint
	// URLs. These are constants specific to each server and are
	// often available via site-specific packages, such as
	// google.Endpoint or github.Endpoint.
	Endpoint Endpoint

	// RedirectURL is the URL to redirect users going through
	// the OAuth flow, after the resource owner's URLs.
	RedirectURL string

	// Scope specifies optional requested permissions.
	Scopes []string
}

func CondVal(v string) []string {
	if v == "" {
		return nil
	}
	return []string{v}
}

func (c *Config) AuthCodeURL(state string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.AuthURL)
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {c.ClientID},
		"redirect_uri":  {c.RedirectURL},
		"scope":         CondVal(strings.Join(c.Scopes, " ")),
		"state":         CondVal(state),
	}

	if strings.Contains(c.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

func (c *Config) TokenURL(code string) string {
	var buf bytes.Buffer
	buf.WriteString(c.Endpoint.TokenURL)
	v := url.Values {
		"client_id":     	{c.ClientID},
		"redirect_uri":  	{c.RedirectURL},
		"client_secret": 	{c.ClientSecret},
		"code": 			{code},
	}

	if strings.Contains(c.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}