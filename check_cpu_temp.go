package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	var (
		w = flag.Int("w", 0, "int flag")
		c = flag.Int("c", 0, "int flag")
		f = flag.String("f", "", "string flag")
	)
	flag.Parse()

	// cmdResult, err := exec.Command("cat", "/sys/devices/platform/coretemp.0/hwmon/hwmon1/temp3_input").Output()
	cmdResult, err := exec.Command("cat", *f).Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	value, err := strconv.Atoi(strings.ReplaceAll(string(cmdResult), "\n", ""))
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	temp := value / 1000

	msg := fmt.Sprintf("CPU Temperature is %dC\n", temp)
	fmt.Printf(msg)

	if temp >= *c {
		os.Exit(2)
	} else if temp >= *w {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
