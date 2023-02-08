package main

import (
	"encoding/json"
	"github.com/calmera/go-nuts/engine"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configFile string
	var cfg engine.Config
	if len(os.Args) == 2 {
		configFile, _ = homedir.Expand(os.Args[1])
		b, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(b, &cfg); err != nil {
			panic(err)
		}
	}

	n, err := engine.NewEngine(cfg)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(n *engine.Engine) {
		<-c
		n.Close()
		os.Exit(1)
	}(n)

	if err := n.Start(); err != nil {
		panic(err)
	}

	n.Wait()
}
