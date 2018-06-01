package wordpress

import (
"github.com/resure-tech/lib/auth"
)

// Endpoint is Wordpress's OAuth endpoint.
var Endpoint = auth.Endpoint{
	AuthURL:  "https://public-api.wordpress.com/oauth2/authorize",
	TokenURL: "https://public-api.wordpress.com/oauth2/token",
}