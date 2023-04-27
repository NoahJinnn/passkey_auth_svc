package mail

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/hellohq/hqservice/ms/auth/config"
)

type Mailer struct {
	apiClient *onesignal.APIClient
	authCtx   context.Context
	cfg       *config.Passcode
}

func NewMailer(cfg *config.Passcode) *Mailer {
	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
	authCtx := context.WithValue(context.Background(), onesignal.AppAuth, cfg.Smtp.OneSignalAppKey)
	return &Mailer{
		apiClient,
		authCtx,
		cfg,
	}
}

func (m *Mailer) Send(email []string, subject string, body string) error {
	notification := *onesignal.NewNotification(m.cfg.Smtp.OneSignalAppID) // Notification |
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
	fmt.Println("send passcode successfully")
	return nil
}
