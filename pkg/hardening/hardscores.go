package hardening

import (
	"fmt"
	"strings"

	"github.com/antonioalfa22/go-utils/command"
)

func GetLynisScores() []string {
	output, err := command.RunCommandWithOutput("lynis", "audit", "system", "grep", "'Hardening index'")
	if err != nil {
		fmt.Println("Error on runing command")
	}
	return strings.Split(output, "\n")
}
