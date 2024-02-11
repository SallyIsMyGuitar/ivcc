package polestar

import (
	"context"
	"fmt"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/oauth"
	"github.com/evcc-io/evcc/util/request"
	"github.com/hasura/go-graphql-client"
	"github.com/samber/lo"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/oauth2"
)

// https://github.com/TA2k/ioBroker.polestar

const (
	OAuthURI  = "https://polestarid.eu.polestar.com"
	basicAuth = "cG9seHBsb3JlOlhhaUtvb0hlaXJlaXNvb3NhaDBFdjZxdW9oczhjb2hGZUtvaHdpZTFhZTdraWV3b2hkb295ZWk5QWVZZWlXb2g"
)

// https://polestarid.eu.polestar.com/.well-known/openid-configuration
var OAuth2Config = &oauth2.Config{
	// ClientID:    "polxplore",
	// RedirectURL: "polestar-explore://explore.polestar.com",
	ClientID:    "polmystar",
	RedirectURL: "https://www.polestar.com/sign-in-callback",
	Endpoint: oauth2.Endpoint{
		AuthURL:  OAuthURI + "/as/authorization.oauth2",
		TokenURL: OAuthURI + "/as/token.oauth2",
	},
	Scopes: []string{
		"openid", "profile", "email", "customer:attributes",
		// "conve:recharge_status", "conve:fuel_status", "conve:odometer_status",
		// "energy:charging_connection_status", "energy:electric_range", "energy:estimated_charging_time", "energy:recharge_status",
		// "energy:battery_charge_level", "energy:charging_system_status", "energy:charging_timer", "energy:electric_range", "energy:recharge_status",
		// "energy:battery_charge_level",
	},
}

type Identity struct {
	*request.Helper
	oauth2.TokenSource
}

// NewIdentity creates Polestar identity
func NewIdentity(log *util.Logger) *Identity {
	return &Identity{
		Helper: request.NewHelper(log),
	}
}

func (v *Identity) Login(user, password string) error {
	if v.Client.Jar == nil {
		var err error
		v.Client.Jar, err = cookiejar.New(&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		})
		if err != nil {
			return err
		}
	}

	state := lo.RandomString(16, lo.AlphanumericCharset)
	uri := OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	var param request.InterceptResult
	v.Client.CheckRedirect, param = request.InterceptRedirect("resumePath", true)
	defer func() { v.Client.CheckRedirect = nil }()

	if _, err := v.Get(uri); err != nil {
		return err
	}

	resume, err := param()
	if err != nil {
		return err
	}

	params := url.Values{
		"pf.username": []string{user},
		"pf.pass":     []string{password},
	}

	uri = fmt.Sprintf("%s/as/%s/resume/as/authorization.ping?client_id=%s", OAuthURI, resume, OAuth2Config.ClientID)
	v.Client.CheckRedirect, param = request.InterceptRedirect("code", true)
	defer func() { v.Client.CheckRedirect = nil }()

	var code string
	if _, err = v.Post(uri, request.FormContent, strings.NewReader(params.Encode())); err == nil {
		code, err = param()
	}

	if err != nil {
		return err
	}

	gqlClient := graphql.NewClient("https://pc-api.polestar.com/eu-north-1/auth", v.Client)

	type Token struct {
		AccessToken  string `graphql:"access_token"`
		RefreshToken string `graphql:"refresh_token"`
		ExpiresIn    int    `graphql:"expires_in"`
	}

	var res struct {
		Token `graphql:"getAuthToken(code: $code)"`
	}

	if err := gqlClient.Query(context.Background(), &res, map[string]any{
		"code": code,
	}, graphql.OperationName("getAuthToken")); err != nil {
		return err
	}

	v.TokenSource = oauth.RefreshTokenSource(&oauth2.Token{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(res.ExpiresIn) * time.Second),
		// Expiry:       time.Now(), //.Add(time.Duration(res.ExpiresIn) * time.Second),
	}, v)

	return err
}

func (v *Identity) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	gqlClient := graphql.NewClient("https://pc-api.polestar.com/eu-north-1/auth", v.Client)

	type Token struct {
		AccessToken  string `graphql:"access_token"`
		RefreshToken string `graphql:"refresh_token"`
		ExpiresIn    int    `graphql:"expires_in"`
	}

	var res struct {
		Token `graphql:"refreshAuthToken(token: $token)"`
	}

	err := gqlClient.Query(context.Background(), &res, map[string]any{
		"token": token.RefreshToken,
	}, graphql.OperationName("refreshAuthToken"))

	return &oauth2.Token{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(res.ExpiresIn) * time.Second),
	}, err
}
