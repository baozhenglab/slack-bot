package slackbot

type SlackbotService interface {
	SendMessage(form map[string]string) error
}
