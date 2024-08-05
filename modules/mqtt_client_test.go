package modules

import (
	"fmt"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	go TaskActuator()
	fmt.Println("check1")
	Enable <- true
	time.Sleep(1 * time.Second)
	time.Sleep(3 * time.Second)
	Enable <- false
}
