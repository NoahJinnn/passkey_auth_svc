package mail

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
	"github.com/hellohq/hqservice/ms/auth/config"
)

type Mailer struct {
}

func NewMailer(cfg config.SMTP) {

	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, cfg.OneSignalAppKey)

	notification := *onesignal.NewNotification(cfg.OneSignalAppID) // Notification |
	notification.SetEmailSubject("Testing")
	notification.SetEmailBody("<html><head>Welcome to Cat Facts</head><body><h1>Welcome to Cat Facts<h1><h4>Learn more about everyone's favorite furry companions!</h4><hr/><p>Hi Nick,</p><p>Thanks for subscribing to Cat Facts! We can't wait to surprise you with funny details about your favorite animal.</p><h5>Today's Cat Fact (March 27)</h5><p>In tigers and tabbies, the middle of the tongue is covered in backward-pointing spines, used for breaking off and gripping meat.</p><a href='https://catfac.ts/welcome'>Show me more Cat Facts</a><hr/><p><small>(c) 2018 Cat Facts, inc</small></p><p><small><a href='[unsubscribe_url]'>Unsubscribe</a></small></p></body></html>")
	notification.SetIncludeEmailTokens([]string{"noah@hellohq.com"})
	// notification.SetIncludedSegments([]string{"Active Users"})
	resp, r, err := apiClient.DefaultApi.CreateNotification(appAuth).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotification`: CreateNotificationSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)

	// player := *onesignal.NewPlayer("abc", 11)
	// player.SetIdentifier("noah@hellohq.com")
	// resp, r, err := apiClient.DefaultApi.CreatePlayer(appAuth).Player(player).Execute()
	// if err != nil {
	//     fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePlayer``: %v\n", err)
	//     fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	// }
	// // response from `CreatePlayer`: CreatePlayerSuccessResponse
	// fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePlayer`: %v\n", resp)

}
