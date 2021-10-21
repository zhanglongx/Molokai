package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/zhanglongx/Molokai/core"
)

func main() {
	if _, err := os.Stat("example.yaml"); os.IsNotExist(err) {
		log.Fatal("file not exists")
	}

	buf, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Fatal("read example.yaml failed")
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
