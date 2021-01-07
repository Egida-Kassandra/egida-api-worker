package packages

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/command"
	"strings"
)


func GetAllPackages() []string {
	output, err := command.RunCommandWithOutput("apt", "list", "--installed")
	if err != nil {
		fmt.Println("Error on runing command")
	}
	return strings.Split(output, "\n")
}