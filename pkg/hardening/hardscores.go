package hardening

import (
	"fmt"
	"strings"

	"github.com/antonioalfa22/go-utils/command"
)

func GetLynisScores() []string {
	comando := "lynis audit system --no-colors | grep 'Hardening index'"
	output, err := command.RunCommandWithOutput("bash", "-c", comando)
	if err != nil {
		fmt.Println("Error on runing command")
	}
	return strings.Split(output, "\n")
}
