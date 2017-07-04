package main

import (
	"os"

	_ "github.com/micro/go-plugins/registry/kubernetes"
	"github.com/micro/micro/cmd"
)

func init() {
	os.Setenv("MICRO_REGISTRY", "kubernetes")
}

func main() {
	cmd.Init()
}
