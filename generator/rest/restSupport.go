package rest

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func ExtractCredentials(language string, r *http.Request) Credentials {
	username, password, err := decodeBasicAuthHeader(r)
	if err == nil {
		return Credentials{
			Language:      language,
			RequestUID:    r.Header.Get("X-request-uid"),
			SessionUID:    "",
			EndUserAccess: "",
			EndUserRole:   "supplier",
			EndUserUID:    username,
			Password:      password,
		}
	}
	return Credentials{
		Language:      language,
		RequestUID:    r.Header.Get("X-request-uid"),
		SessionUID:    r.Header.Get("X-session-uid"),
		EndUserAccess: r.Header.Get("X-enduser-access"),
		EndUserRole:   r.Header.Get("X-enduser-role"),
		EndUserUID:    r.Header.Get("X-enduser-uid"),
	}
}

type Credentials struct {
	Language      string
	RequestUID    string
	SessionUID    string
	EndUserAccess string
	EndUserRole   string
	EndUserUID    string
	Password      string
}

func decodeBasicAuthHeader(r *http.Request) (string, string, error) {
	auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return "", "", fmt.Errorf("Invalid/missing header")
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return "", "", fmt.Errorf("Invalid/missing header-values")
	}

	return pair[0], pair[1], nil

}
