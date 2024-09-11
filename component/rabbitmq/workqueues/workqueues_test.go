package workqueues

import (
	"testing"
)

func Test_Send(t *testing.T) {
	Publish("sms", "hihi hahaa")
}

func Test_Receiver(t *testing.T) {
	StartConsumer("sms")
}
