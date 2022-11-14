package ch07

import "fmt"

func UsingInterface() {
	way1()
	fmt.Sprintln("------------------------------------------------------------")
	way2()
}

// -- way 1 ---------------------------------------

func way1() {
	client := &ThirdPartyMessageClient{
		name:          "UNICON",
		serverAddress: "https://www.faxxooxxoxoxox.com",
		serverPort:    8443,
	}
	provider := &MessageServiceProvider{
		messageClient: client,
	}

	var messageService MessageService = provider
	messageService.Send("Kut Zhang", "Lana Chang", "Show me the money")
}

type MessageService interface {
	// send message
	Send(sender string, receiver string, message string) error
}

// -- define interface provider
type MessageServiceProvider struct {
	messageClient *ThirdPartyMessageClient
}

func (provider *MessageServiceProvider) Send(sender string, receiver string, message string) error {
	fmt.Printf("Message provider - sending message: %s - %s - %s\n", sender, receiver, message)
	err := provider.messageClient.send(sender, receiver, message)
	if err != nil {
		return err
	} else {
		fmt.Println("Message provider - message has been sent.")
		return nil
	}
}

// -- define third party message client
type ThirdPartyMessageClient struct {
	name          string
	serverAddress string
	serverPort    int
}

func (client *ThirdPartyMessageClient) send(sender string, receiver string, message string) error {
	serverEndpoint := fmt.Sprintf("%s:%d", client.serverAddress, client.serverPort)
	fmt.Printf("Requesting %s service to send message: %s\n", client.name, serverEndpoint)
	fmt.Printf("Sending message from %s to %s: %s\n", sender, receiver, message)
	fmt.Printf("Message has been sent")
	return nil
}

// -- way 2 ------------------------------------------------------

func way2() {
	var engineCleaner = new(EngineWorker)
	engineCleaner.kill()
	engineCleaner.clean()
}

type EngineCleaner interface {
	kill()
	clean()
}

type Killer struct {
}

func (killer *Killer) kill() {
	fmt.Println("Killer killing.............")
}

type Cleaner struct {
}

func (cleaner *Cleaner) clean() {
	fmt.Println("Cleaner cleaning............")
}

type EngineWorker struct {
	Killer
	Cleaner
}
