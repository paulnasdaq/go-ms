package main

import "github.com/paulnasdaq/fms-v2/common"

func main() {
	connector, err := common.NewConnector()
	if err != nil {
		panic(err)
	}
	err = connector.PublishEvent("test.hehe", []byte("hehe"))
	if err != nil {
		panic(err)
	}
}
