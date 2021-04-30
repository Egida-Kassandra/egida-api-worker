package machine

import (
	"fmt"
	"github.com/antonioalfa22/go-utils/command"
	"strconv"
	"strings"
)

func MachineHasGUI() bool {
	comando := "ls /usr/bin/*session"
	output, err := command.RunCommandWithOutput("bash", "-c", comando)
	if err != nil {
		fmt.Println("Error on runing command")
	}
	lines := strings.Split(output, "\n")
	for _, l := range lines {
		if l != "/usr/bin/byobu-select-session" && l != "/usr/bin/dbus-run-session" && l != "" { return true }
	}
	return false
}

func ListOpenPorts() []int64 {
	comando := "sudo lsof -i -P -n | grep LISTEN"
	output, err := command.RunCommandWithOutput("bash", "-c", comando)
	if err != nil {
		fmt.Println("Error on runing command")
	}
	lines := strings.Split(output, "\n")
	var ports []int64
	for _, l := range lines {
		if l != "" {
			num, _ := strconv.ParseInt(strings.Split(strings.Split(l, ":")[1], " ")[0], 10, 64)
			ports = append(ports, num)
		}
	}
	return ports
}