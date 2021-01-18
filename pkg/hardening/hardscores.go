package hardening

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/command"
	"strings"
)

func GetLynisScores() []string {
	output, err := command.RunCommandWithOutput("lynis", "audit", "system")
	if err != nil {
		fmt.Println("Error on runing command")
	}
	return strings.Split(output, "\n")
}
