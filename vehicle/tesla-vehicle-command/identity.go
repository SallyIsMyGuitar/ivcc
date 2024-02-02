package vc

import (
	"context"
	"sync"

	"github.com/evcc-io/evcc/server/db/settings"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/oauth"
	"github.com/evcc-io/evcc/util/request"
	"github.com/teslamotors/vehicle-command/pkg/account"
	"golang.org/x/oauth2"
)

// https://auth.tesla.com/oauth2/v3/.well-known/openid-configuration

// OAuth2Config is the OAuth2 configuration for authenticating with the Tesla API.
var OAuth2Config = &oauth2.Config{
	RedirectURL: "https://auth.tesla.com/void/callback",
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://auth.tesla.com/en_us/oauth2/v3/authorize",
		TokenURL:  "https://auth.tesla.com/oauth2/v3/token",
		AuthStyle: oauth2.AuthStyleInParams,
	},
	Scopes: []string{"openid", "email", "offline_access"},
}

const userAgent = "evcc/evcc-io"

var TESLA_CLIENT_ID, TESLA_CLIENT_SECRET string

func init() {
	if TESLA_CLIENT_ID != "" {
		OAuth2Config.ClientID = TESLA_CLIENT_ID
	}
	if TESLA_CLIENT_SECRET != "" {
		OAuth2Config.ClientSecret = TESLA_CLIENT_SECRET
	}
}

type Identity struct {
	oauth2.TokenSource
	mu    sync.Mutex
	log   *util.Logger
	token *oauth2.Token
	acct  *account.Account
}

func NewIdentity(log *util.Logger, ts oauth2.TokenSource) (*Identity, error) {
	token, err := ts.Token()
	if err != nil {
		return nil, err
	}

	// acct, err := account.New(token.AccessToken, userAgent)
	// if err != nil {
	// 	return nil, err
	// }

	v := &Identity{
		token: token,
		// acct:        acct,
	}

	v.TokenSource = oauth.RefreshTokenSource(token, v)

	return v, nil
}

func (v *Identity) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, request.NewClient(v.log))
	ts := OAuth2Config.TokenSource(ctx, token)

	token, err := ts.Token()
	if err != nil {
		return nil, err
	}

	claims, err := TokenClaims(token)
	if err != nil {
		return nil, err
	}

	subject, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}

	err = settings.SetJson(SettingsKey(subject), token)

	return token, err
}

// func (v *Identity) Account() *account.Account {
// 	token, err := v.Token()
// 	if err != nil {
// 		v.log.ERROR.Println(err)
// 		return v.acct
// 	}

// 	if token.AccessToken != v.token.AccessToken {
// 		acct, err := account.New(token.AccessToken, userAgent)
// 		if err != nil {
// 			v.log.ERROR.Println(err)
// 			return v.acct
// 		}

// 		v.acct = acct
// 	}

// 	return v.acct
// }
