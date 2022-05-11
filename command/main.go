package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage =  func() {
		fmt.Fprintf(os.Stderr, "命令提示,异常退出\n")
		flag.PrintDefaults()
	}
}

func main() {
	// 1.正式接收参数
	name := flag.String("name", "everyone", "username")
	flag.Parse()
	fmt.Println("hello: ", *name)
}
