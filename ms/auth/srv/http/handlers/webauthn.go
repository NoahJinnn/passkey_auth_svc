package handlers

import (
	"errors"

	"github.com/hellohq/hqservice/api/openapi/model"
	"github.com/hellohq/hqservice/ms/auth/srv/http/server"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type WebauthnHandler struct {
	srv *server.HttpServer
}

func NewWebauthnHandler(srv *server.HttpServer) *WebauthnHandler {
	return &WebauthnHandler{
		srv: srv,
	}
}

func BeginRegistration(c echo.Context) error {
	// fmt.Printf("noah %+v\n", ctx)
	// TODO: Impl Session middleware
	sessionToken, ok := c.Get("session").(jwt.Token)
	if !ok {
		return errors.New("failed to cast session object")
	}
	// uId, err := uuid.FromString(sessionToken.Subject())
	// if err != nil {
	// 	return fmt.Errorf("failed to parse userId from JWT subject:%w", err)
	// }

	// TODO: Impl DAL
	// webauthnUser, user, err := h.getWebauthnUser(h.persister.GetConnection(), uId)
	// _, _, err := srv.app.WebauthnBeginRegistration(ctx)
	// if err != nil {
	// 	return fmt.Errorf("failed to get user: %w", err)
	// }
	// if webauthnUser == nil {
	// 	return dto.NewHTTPError(http.StatusBadRequest, "user not found").SetInternal(errors.New(fmt.Sprintf("user %s not found ", uId)))
	// }

	// TODO: Impl Appl
	options := model.CredentialCreationOptions{}
	// t := true
	// options, sessionData, err := h.webauthn.BeginRegistration(
	// 	webauthnUser,
	// 	webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
	// 		RequireResidentKey: &t,
	// 		ResidentKey:        protocol.ResidentKeyRequirementRequired,
	// 		UserVerification:   protocol.VerificationRequired,
	// 	}),
	// 	webauthn.WithConveyancePreference(protocol.PreferNoAttestation),
	// 	// don't set the excludeCredentials list, so an already registered device can be re-registered
	// )
	// if err != nil {
	// 	return fmt.Errorf("failed to create webauthn creation options: %w", err)
	// }

	// TODO: Impl DAL
	// err = h.persister.GetWebauthnSessionDataPersister().Create(*intern.WebauthnSessionDataToModel(sessionData, models.WebauthnOperationRegistration))
	// if err != nil {
	// 	return fmt.Errorf("failed to store creation options session data: %w", err)
	// }
	// err = h.auditLogger.Create(c, models.AuditLogWebAuthnRegistrationInitSucceeded, user, nil)
	// if err != nil {
	// 	return fmt.Errorf("failed to create audit log: %w", err)
	// }

	// switch {
	// default:
	// 	return errGetUsers(log, err, codeInternal)
	// case err == nil:
	// 	return CustomResponder(func(w http.ResponseWriter, producer runtime.Producer) {
	// 		if err := producer.Produce(w, options); err != nil {
	// 			panic(err) // let the recovery middleware deal with this
	// 		}
	// 	})
	// }

	return nil
}
