package singleton

import "fmt"

type notificationServiceImpl struct {
}

func (impl *notificationServiceImpl) Notify(message string) error {
	fmt.Printf("Notication service - notify: %s\n", message)
	return nil
}

var impl = &notificationServiceImpl{}

func GetNotificationService() NotificationService {
	return impl
}
