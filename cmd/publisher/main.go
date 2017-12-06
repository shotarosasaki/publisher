package main

import (
	"flag"
	"os"
	"fmt"
)

var (
configPath = flag.String("f", "/etc/publisher/publisher.toml", "specify a path to configuration file")
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Panic occured. %v", err)	// TODO メッセージ見直し
			exitCode = -1	// TODO k8sと絡めた時に、プロセス落とすでよい？
		}
	}()
	return wrappedMain()
}

func wrappedMain() int {
	flag.Parse()

	return 0
}
