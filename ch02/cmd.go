package main

import (
	"ch02/classpath"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	XjreOption  string
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.versionFlag, "v", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%v args:%v\n", cmd.class, cmd.args)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	data, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Cannot found or load mian class %s\n", cmd.class)
	}
	fmt.Printf("class data : %v\n", data)

	//fmt.Printf("classpath:%s class:%s args:%v\n",
	//	cmd.cpOption, cmd.class, cmd.args)
}
