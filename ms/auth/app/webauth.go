package app

import (
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

func NewWebAuthn(cfg config) *webauthn.WebAuthn {
	var w *webauthn.WebAuthn
	var err error
	wconfig := &webauthn.Config{
		RPDisplayName: "Go Webauthn",                               // Display Name for your site
		RPID:          "go-webauthn.local",                         // Generally the FQDN for your site
		RPOrigins:     []string{"https://login.go-webauthn.local"}, // The origin URLs allowed for WebAuthn requests
	}

	if w, err = webauthn.New(wconfig); err != nil {
		fmt.Println(err)
	}
	return w
}

func (app *App) WebauthnBeginRegistration(ctx Ctx) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	// user := datastore.GetUser() // Find or create the new user
	user := &User{}
	return app.wAuthn.BeginRegistration(user)
	// handle errors if present
	// store the sessionData values
	// options.publicKey contain our registration options
}

// func (app *App) WebauthnFinishRegistration(postBody []byte) *webauthn.Credential {

// 	return credential
// }
