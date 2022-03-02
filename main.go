package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/jessevdk/go-flags"
	"github.com/zhanglongx/Molokai/core"
)

const (
	VERSION = "1.0.2"
)

var opt struct {
	Version bool `long:"version"`
}

func main() {
	args, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		log.Fatal(err)
	}

	if opt.Version {
		fmt.Println("Molokai " + VERSION)
		os.Exit(0)
	}

	var cfgFile string
	if len(args) > 2 {
		log.Fatal("more than 1 input file")
	} else if len(args) == 1 {
		log.Printf("no cfg file given, using default: molokai.yaml")
		cfgFile = "molokai.yaml"
	} else {
		cfgFile = args[1]
	}

	log.Printf("using cfg file: %s", cfgFile)

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		log.Fatalf("file %s not exists", cfgFile)
	}

	buf, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("read %s failed", cfgFile)
	}

	var Molokai core.Molokai
	if err := yaml.Unmarshal(buf, &Molokai); err != nil {
		log.Fatal("parse example.yaml failed")
	}

	if err := Molokai.Init(); err != nil {
		log.Fatal(err)
	}

	if err := Molokai.RunHoldings(); err != nil {
		log.Fatal(err)
	}
}
