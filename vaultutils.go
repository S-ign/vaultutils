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

// Data
type Data struct {
	Keymap map[string]string `json:"keymap"`
}

// Auth Authenticates user request
func Auth(req handler.Request, vaultEngine, functionURL string) error {
	email := req.Header.Get("email")
	token := req.Header.Get("apitoken")

	var vd VaultData
	vd.AccessToken = "mC9Ucju63Z7%&O07GQvzvf@o"
	vd.Action = "listSecretData"
	vd.Path = fmt.Sprintf("%v/data/%v", vaultEngine, email)

	postHeaders := make(map[string]string)
	postHeaders["email"] = email
	postHeaders["apitoken"] = token

	b, err := httputils.PostRequest(vd, functionURL, postHeaders)
	if err != nil {
		return err
	}

	//m := make(map[string]interface{})
	var m Data
	err = json.Unmarshal(b, &m)

	if m.Keymap[email] != token {
		return fmt.Errorf(fmt.Sprintf("vault auth: Unauthorized Access\nb:%v\nm:%v~~token:%v\n\n%v", string(b), m, token, err))
	}
	return nil
}
