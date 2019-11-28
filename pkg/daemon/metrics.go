package daemon

import (
	"github.com/golang/glog"
)

// Message defines log message structure transmit
// via MessageChannel.
type Message struct {
	Name	string
	Type	string
	Content	string
}

// Parser is the main structure for parse instance.
// It contains all the necessary data to parse linuxptp process log messages.
type Parser struct {
	// node name where daemon is running
	nodeName	string
	namespace	string

	ptpUpdate	*LinuxPTPConfUpdate
	// channel ensure Parser.Run() exit when main function exits.
	// stopCh is created by main function and passed by Daemon via NewParser()
	stopCh		<-chan struct{}
	messageChannel	<-chan Message
}

// NewParser is called by daemon to create new parser instance
func NewParser(
	nodeName	string,
	namespace	string,
	ptpUpdate	*LinuxPTPConfUpdate,
	stopCh		<-chan struct{},
	messageChannel	<-chan Message,
) *Parser {
	return &Parser{
		nodeName:	nodeName,
		namespace:	namespace,
		ptpUpdate:	ptpUpdate,
		stopCh:		stopCh,
		messageChannel:	messageChannel,
	}
}

// Run in a for loop to listen for log messages in MessageChannel
func (ps *Parser) Run() {
	for {
		select {
		case msg := <-ps.messageChannel:
			switch msg.Type {
			case "ptp4l":
				glog.V(2).Infof("Parser message(%s) received: %v", msg.Type, msg.Content)
			case "phc2sys":
				glog.V(2).Infof("Parser message(%s) received: %v", msg.Type, msg.Content)
			}
		case <-ps.stopCh:
			glog.Infof("linuxPTP Parser stop signal received, existing..")
			return
		}
	}
	return
}
