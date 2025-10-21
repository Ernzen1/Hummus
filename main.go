package main

import (
	"fmt"
	"log"
	"os"
	"study/service/svcHandler"
)

func main() {

	f, err := os.OpenFile("C:\\Program Files\\placeholder\\service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln(fmt.Errorf("error opening file: %v", err))
	}
	defer f.Close()

	log.SetOutput(f)
	svcHandler.RunService("myservice1", false)

}
