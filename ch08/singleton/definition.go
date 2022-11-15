package singleton

type NotificationService interface {
	Notify(message string) error
}
