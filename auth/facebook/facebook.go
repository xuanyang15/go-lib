package facebook

import (
	"github.com/resure-tech/lib/auth"
)

// Endpoint is Facebook's OAuth endpoint.
var Endpoint = auth.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}
