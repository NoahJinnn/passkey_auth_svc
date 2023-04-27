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
	cfg       config.SMTP
}

func NewMailer(cfg config.SMTP) *Mailer {

	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
	authCtx := context.WithValue(context.Background(), onesignal.AppAuth, cfg.OneSignalAppKey)
	return &Mailer{
		apiClient,
		authCtx,
		cfg,
	}
}

func (m *Mailer) Send(body string) error {
	notification := *onesignal.NewNotification(m.cfg.OneSignalAppID) // Notification |
	notification.SetEmailSubject("Testing")
	notification.SetEmailFromName("HelloHQ Pte. Ltd.")
	notification.SetEmailFromAddress("noah@hellohq.com")
	notification.SetEmailBody(body)
	notification.SetIncludeEmailTokens([]string{"trannguyenhcmut@gmail.com"})
	resp, r, err := m.apiClient.DefaultApi.CreateNotification(m.authCtx).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return err
	}
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)
	return nil
}
