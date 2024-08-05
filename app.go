package main

import (
	"context"
	"fmt"
	"mqtt-test-2/modules"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type StartInput struct {
	Host          string `json:"host"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Date1         string `json:"date1"`
	Date2         string `json:"date2"`
	Concurrencies int    `json:"concurrencies"`
	Topic         string `json:"topic"`
	Message       string `json:"message"`
	Port          string `json:"port"`
}

func (a *App) Start(input StartInput) bool {

	fmt.Println("check.....myStruct", input.Username)
	fmt.Println("check.....Concurrencies", input.Concurrencies)
	fmt.Println("check.....myStruct", input.Username)
	mqConfig := modules.MQTTConfig{
		ClientID: "mqttx_61862496",
		Username: input.Username,
		Password: input.Password,
		Topic:    input.Topic,
		Broker:   fmt.Sprintf("tcp://%s:%s", input.Host, input.Port),
	}
	source := modules.Config{
		MQTTConfig:       mqConfig,
		ConcurrencyLimit: input.Concurrencies,
		Message:          input.Message,
	}
	return source.Start()
}

func (a *App) Stop() bool {
	source := modules.Config{}
	return source.Stop()
}
