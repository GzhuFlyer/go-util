package main

import (
	"mixtest/mqtt"
)

func main() {
	// for i := 0; i < 1; i++ {
	// 	myqueue.QueueT4()
	// 	// myqueue.QueueT3()
	// // }
	// mqttTwo.M1()
	// mqttTwo.M2()
	// s := mqtt.NewAdaptorWithAuth("tcp://localhost:1883", "client123", "fzw", "12345678")
	// mqtt config : {tcp://10.240.34.35:30010 C2@1727584984128937984@c2app 1727584984128937984 e3bc38a4faa625d074664d572d810c1e}
	// s := mqtt.NewAdaptorWithAuth("tcp://10.240.34.35:30010", "C2@1727884703555842048@1334", "admin", "e3bc38a4faa625d074664d572d810c1e")

	// ret := s.Connect()
	// fmt.Println("connect ret = ", ret)
	// s.SetQoS(1)
	// mqtt.HandlerSub(s)
	// time.Sleep(500000 * time.Second)
	// mqtt.HandlerPub(s)
	mqtt.TestCode()
}
