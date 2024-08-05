package modules

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

var Enable = make(chan bool, 1)
var Stop = make(chan struct{})

var CurrentConfig Config

func TaskActuator(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("停止任务")
			return
		default:
			if CurrentConfig.Message == "" || CurrentConfig.ConcurrencyLimit < 0 {
				continue
			}
			var wg sync.WaitGroup
			for i := 0; i < CurrentConfig.ConcurrencyLimit; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					CurrentConfig.MQTTConfig.ClientID = uuid.New().String()
					client, err := NewMQTTClient(CurrentConfig.MQTTConfig)
					if err != nil {
						fmt.Printf("Error creating MQTT client: %v\n", err)
						return
					}
					err = client.PublishMessage(CurrentConfig.Message)
					if err != nil {
						fmt.Printf("Error publishing message: %v\n", err)
					} else {
						fmt.Println("Message published successfully")
					}
				}()
			}
			wg.Wait()
			time.Sleep(1 * time.Second) // 控制消息发布的频率
			fmt.Println("我这里触发了一次并发......")
		}
	}

}

func StartTaskActuator() bool {
	fmt.Println("启动")
	go TaskActuator(Stop) // 在启动任务时启动 TaskActuator 协程
	Enable <- true
	fmt.Println("启动完成")
	return true
}

func StopTaskActuator() bool {
	fmt.Println("停止")
	close(Stop)                // 通知停止
	Stop = make(chan struct{}) // 重新创建 Stop 通道，以便后续继续使用
	fmt.Println("停止完成")
	return true
}
