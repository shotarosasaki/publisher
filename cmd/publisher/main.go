package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shotarosasaki/publisher"
	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/global"
	"go.uber.org/zap"
)

var (
	// TODO 現行LINEはプログラム引数。環境変数とどちらがよいか再考！
	configPath = flag.String("f", "/etc/publisher/publisher.toml", "specify a path to configuration file")
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Panic occured. %v", err) // TODO メッセージ見直し
			exitCode = -1                        // TODO k8sと絡めた時に、プロセス落とすでよい？
		}
	}()
	return wrappedMain()
}

func wrappedMain() int {
	flag.Parse()

	conf, err := config.New(*configPath)
	if err != nil {
		// TODO ログ出力
		// TODO exitCodeを定数定義
		return -1
	}

	global.InitLogger(conf.Log)

	if err := publisher.Serve(conf); err != nil {
		global.Logger.Error("xxxx", zap.Error(err))
		// TODO exitCodeを定数定義
		return -1
	}

	// TODO exitCodeを定数定義
	return 0
}
