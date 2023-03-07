package app

import (
	"bytes"
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

func (app *App) BeginRegistration() (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	// user := datastore.GetUser() // Find or create the new user
	user := &User{}
	return app.wAuthn.BeginRegistration(user)
	// handle errors if present
	// store the sessionData values
	// options.publicKey contain our registration options
}

func (app *App) FinishRegistration(postBody []byte) *webauthn.Credential {
	response, err := protocol.ParseCredentialCreationResponseBody(bytes.NewReader(postBody))
	if err != nil {
		// Handle Error and return.

		return nil
	}

	// user := datastore.GetUser() // Get the user

	// Get the session data stored from the function above
	// session := datastore.GetSession()
	user := &User{}
	session := &UserSession{}
	credential, err := app.wAuthn.CreateCredential(user, *session.Webauthn, response)
	if err != nil {
		// Handle Error and return.

		return nil
	}

	// If creation was successful, store the credential object

	// Pseudocode to add the user credential.
	// user.AddCredential(credential)
	// datastore.SaveUser(user)

	return credential
}
