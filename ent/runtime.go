// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/ent/email"
	"github.com/hellohq/hqservice/ent/identity"
	"github.com/hellohq/hqservice/ent/jwk"
	"github.com/hellohq/hqservice/ent/passcode"
	"github.com/hellohq/hqservice/ent/primaryemail"
	"github.com/hellohq/hqservice/ent/schema"
	"github.com/hellohq/hqservice/ent/user"
	"github.com/hellohq/hqservice/ent/webauthncredential"
	"github.com/hellohq/hqservice/ent/webauthncredentialtransport"
	"github.com/hellohq/hqservice/ent/webauthnsessiondata"
	"github.com/hellohq/hqservice/ent/webauthnsessiondataallowedcredential"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	emailFields := schema.Email{}.Fields()
	_ = emailFields
	// emailDescCreatedAt is the schema descriptor for created_at field.
	emailDescCreatedAt := emailFields[3].Descriptor()
	// email.DefaultCreatedAt holds the default value on creation for the created_at field.
	email.DefaultCreatedAt = emailDescCreatedAt.Default.(func() time.Time)
	// emailDescUpdatedAt is the schema descriptor for updated_at field.
	emailDescUpdatedAt := emailFields[4].Descriptor()
	// email.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	email.DefaultUpdatedAt = emailDescUpdatedAt.Default.(func() time.Time)
	// email.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	email.UpdateDefaultUpdatedAt = emailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// emailDescID is the schema descriptor for id field.
	emailDescID := emailFields[0].Descriptor()
	// email.DefaultID holds the default value on creation for the id field.
	email.DefaultID = emailDescID.Default.(func() uuid.UUID)
	identityFields := schema.Identity{}.Fields()
	_ = identityFields
	// identityDescCreatedAt is the schema descriptor for created_at field.
	identityDescCreatedAt := identityFields[5].Descriptor()
	// identity.DefaultCreatedAt holds the default value on creation for the created_at field.
	identity.DefaultCreatedAt = identityDescCreatedAt.Default.(func() time.Time)
	// identityDescUpdatedAt is the schema descriptor for updated_at field.
	identityDescUpdatedAt := identityFields[6].Descriptor()
	// identity.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	identity.DefaultUpdatedAt = identityDescUpdatedAt.Default.(func() time.Time)
	// identity.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	identity.UpdateDefaultUpdatedAt = identityDescUpdatedAt.UpdateDefault.(func() time.Time)
	// identityDescID is the schema descriptor for id field.
	identityDescID := identityFields[0].Descriptor()
	// identity.DefaultID holds the default value on creation for the id field.
	identity.DefaultID = identityDescID.Default.(func() uuid.UUID)
	jwkFields := schema.Jwk{}.Fields()
	_ = jwkFields
	// jwkDescCreatedAt is the schema descriptor for created_at field.
	jwkDescCreatedAt := jwkFields[2].Descriptor()
	// jwk.DefaultCreatedAt holds the default value on creation for the created_at field.
	jwk.DefaultCreatedAt = jwkDescCreatedAt.Default.(func() time.Time)
	passcodeFields := schema.Passcode{}.Fields()
	_ = passcodeFields
	// passcodeDescCreatedAt is the schema descriptor for created_at field.
	passcodeDescCreatedAt := passcodeFields[5].Descriptor()
	// passcode.DefaultCreatedAt holds the default value on creation for the created_at field.
	passcode.DefaultCreatedAt = passcodeDescCreatedAt.Default.(func() time.Time)
	// passcodeDescUpdatedAt is the schema descriptor for updated_at field.
	passcodeDescUpdatedAt := passcodeFields[6].Descriptor()
	// passcode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	passcode.DefaultUpdatedAt = passcodeDescUpdatedAt.Default.(func() time.Time)
	// passcode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	passcode.UpdateDefaultUpdatedAt = passcodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// passcodeDescID is the schema descriptor for id field.
	passcodeDescID := passcodeFields[0].Descriptor()
	// passcode.DefaultID holds the default value on creation for the id field.
	passcode.DefaultID = passcodeDescID.Default.(func() uuid.UUID)
	primaryemailFields := schema.PrimaryEmail{}.Fields()
	_ = primaryemailFields
	// primaryemailDescCreatedAt is the schema descriptor for created_at field.
	primaryemailDescCreatedAt := primaryemailFields[3].Descriptor()
	// primaryemail.DefaultCreatedAt holds the default value on creation for the created_at field.
	primaryemail.DefaultCreatedAt = primaryemailDescCreatedAt.Default.(func() time.Time)
	// primaryemailDescUpdatedAt is the schema descriptor for updated_at field.
	primaryemailDescUpdatedAt := primaryemailFields[4].Descriptor()
	// primaryemail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	primaryemail.DefaultUpdatedAt = primaryemailDescUpdatedAt.Default.(func() time.Time)
	// primaryemail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	primaryemail.UpdateDefaultUpdatedAt = primaryemailDescUpdatedAt.UpdateDefault.(func() time.Time)
	// primaryemailDescID is the schema descriptor for id field.
	primaryemailDescID := primaryemailFields[0].Descriptor()
	// primaryemail.DefaultID holds the default value on creation for the id field.
	primaryemail.DefaultID = primaryemailDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	webauthncredentialFields := schema.WebauthnCredential{}.Fields()
	_ = webauthncredentialFields
	// webauthncredentialDescCreatedAt is the schema descriptor for created_at field.
	webauthncredentialDescCreatedAt := webauthncredentialFields[6].Descriptor()
	// webauthncredential.DefaultCreatedAt holds the default value on creation for the created_at field.
	webauthncredential.DefaultCreatedAt = webauthncredentialDescCreatedAt.Default.(func() time.Time)
	// webauthncredentialDescUpdatedAt is the schema descriptor for updated_at field.
	webauthncredentialDescUpdatedAt := webauthncredentialFields[7].Descriptor()
	// webauthncredential.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	webauthncredential.DefaultUpdatedAt = webauthncredentialDescUpdatedAt.Default.(func() time.Time)
	// webauthncredential.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	webauthncredential.UpdateDefaultUpdatedAt = webauthncredentialDescUpdatedAt.UpdateDefault.(func() time.Time)
	webauthncredentialtransportFields := schema.WebauthnCredentialTransport{}.Fields()
	_ = webauthncredentialtransportFields
	// webauthncredentialtransportDescID is the schema descriptor for id field.
	webauthncredentialtransportDescID := webauthncredentialtransportFields[0].Descriptor()
	// webauthncredentialtransport.DefaultID holds the default value on creation for the id field.
	webauthncredentialtransport.DefaultID = webauthncredentialtransportDescID.Default.(func() uuid.UUID)
	webauthnsessiondataFields := schema.WebauthnSessionData{}.Fields()
	_ = webauthnsessiondataFields
	// webauthnsessiondataDescCreatedAt is the schema descriptor for created_at field.
	webauthnsessiondataDescCreatedAt := webauthnsessiondataFields[5].Descriptor()
	// webauthnsessiondata.DefaultCreatedAt holds the default value on creation for the created_at field.
	webauthnsessiondata.DefaultCreatedAt = webauthnsessiondataDescCreatedAt.Default.(func() time.Time)
	// webauthnsessiondataDescUpdatedAt is the schema descriptor for updated_at field.
	webauthnsessiondataDescUpdatedAt := webauthnsessiondataFields[6].Descriptor()
	// webauthnsessiondata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	webauthnsessiondata.DefaultUpdatedAt = webauthnsessiondataDescUpdatedAt.Default.(func() time.Time)
	// webauthnsessiondata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	webauthnsessiondata.UpdateDefaultUpdatedAt = webauthnsessiondataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// webauthnsessiondataDescID is the schema descriptor for id field.
	webauthnsessiondataDescID := webauthnsessiondataFields[0].Descriptor()
	// webauthnsessiondata.DefaultID holds the default value on creation for the id field.
	webauthnsessiondata.DefaultID = webauthnsessiondataDescID.Default.(func() uuid.UUID)
	webauthnsessiondataallowedcredentialFields := schema.WebauthnSessionDataAllowedCredential{}.Fields()
	_ = webauthnsessiondataallowedcredentialFields
	// webauthnsessiondataallowedcredentialDescCreatedAt is the schema descriptor for created_at field.
	webauthnsessiondataallowedcredentialDescCreatedAt := webauthnsessiondataallowedcredentialFields[3].Descriptor()
	// webauthnsessiondataallowedcredential.DefaultCreatedAt holds the default value on creation for the created_at field.
	webauthnsessiondataallowedcredential.DefaultCreatedAt = webauthnsessiondataallowedcredentialDescCreatedAt.Default.(func() time.Time)
	// webauthnsessiondataallowedcredentialDescUpdatedAt is the schema descriptor for updated_at field.
	webauthnsessiondataallowedcredentialDescUpdatedAt := webauthnsessiondataallowedcredentialFields[4].Descriptor()
	// webauthnsessiondataallowedcredential.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	webauthnsessiondataallowedcredential.DefaultUpdatedAt = webauthnsessiondataallowedcredentialDescUpdatedAt.Default.(func() time.Time)
	// webauthnsessiondataallowedcredential.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	webauthnsessiondataallowedcredential.UpdateDefaultUpdatedAt = webauthnsessiondataallowedcredentialDescUpdatedAt.UpdateDefault.(func() time.Time)
	// webauthnsessiondataallowedcredentialDescID is the schema descriptor for id field.
	webauthnsessiondataallowedcredentialDescID := webauthnsessiondataallowedcredentialFields[0].Descriptor()
	// webauthnsessiondataallowedcredential.DefaultID holds the default value on creation for the id field.
	webauthnsessiondataallowedcredential.DefaultID = webauthnsessiondataallowedcredentialDescID.Default.(func() uuid.UUID)
}
