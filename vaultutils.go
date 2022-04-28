package vaultutils

import (
	"encoding/json"
	"fmt"

	"github.com/S-ign/httputils"
	handler "github.com/openfaas/templates-sdk/go-http"
)

// VaultData .
type VaultData struct {
	AccessToken string `json:"access_token"`
	Action      string `json:"action"`
	Path        string `json:"path"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

// Auth Authenticates user request
func Auth(req handler.Request, vaultEngine, functionURL string) error {
	email := req.Header.Get("email")
	token := req.Header.Get("apitoken")

	var vd VaultData
	vd.AccessToken = "mC9Ucju63Z7%&O07GQvzvf@o"
	vd.Action = "listSecretData"
	vd.Path = fmt.Sprintf("%v/%v", vaultEngine, email)

	postHeaders := make(map[string]string)
	postHeaders["email"] = email
	postHeaders["apitoken"] = token

	b, err := httputils.PostRequest(vd, functionURL, postHeaders)
	if err != nil {
		return err
	}

	m := make(map[string]string)
	err = json.Unmarshal(b, &m)

	if m[email] != token {
		return fmt.Errorf(fmt.Sprintf("vault auth: Unauthorized Access\n%v~~%v", string(b), token))
	}
	return nil
}
