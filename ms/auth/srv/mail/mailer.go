package mail

import (
	"context"
	"fmt"
	"os"

	"github.com/NoahJinnn/passkey_auth_svc/ms/auth/config"
	"github.com/OneSignal/onesignal-go-api"
)

type IMailer interface {
	Send(email []string, subject string, body string) error
}

type Mailer struct {
	apiClient *onesignal.APIClient
	authCtx   context.Context
	cfg       *config.Passcode
}

func NewMailer(cfg *config.Passcode) IMailer {
	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
	authCtx := context.WithValue(context.Background(), onesignal.AppAuth, cfg.OneSignalAppKey)
	return &Mailer{
		apiClient,
		authCtx,
		cfg,
	}
}

func (m *Mailer) Send(email []string, subject string, body string) error {
	notification := *onesignal.NewNotification(m.cfg.OneSignalAppID) // Notification |
	notification.SetEmailSubject(subject)
	notification.SetEmailFromName(m.cfg.Email.FromName)
	notification.SetEmailFromAddress(m.cfg.Email.FromAddress)
	notification.SetEmailBody(body)
	notification.SetIncludeEmailTokens(email)
	_, r, err := m.apiClient.DefaultApi.CreateNotification(m.authCtx).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return err
	}
	return nil
}
