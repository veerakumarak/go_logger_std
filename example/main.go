package main

import std "github.com/veerakumarak/go_logger_std"

func main() {
	logger := std.NewLogger()
	logger.Info("Hello World")
}
