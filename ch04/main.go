package main

import "fmt"

//ch03 -Xjre "F:\soft\java\jdk\jdk8" java.lang.String
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
		//	F:\soft\java\jdk\jdk8
	}
}
