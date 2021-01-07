package services

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/command"
	"strings"
)

func GetAllServices() []string {
	output, err := command.RunCommandWithOutput("service", "--status-all")
	if err != nil {
		fmt.Println("Error on runing command")
	}
	return strings.Split(output, "\n")
}