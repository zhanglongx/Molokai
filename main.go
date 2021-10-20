package main

import (
	"fmt"
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

	var holdings core.Holding
	if err := yaml.Unmarshal(buf, &holdings); err != nil {
		log.Fatal("parse example.yaml failed")
	}

	fmt.Println(holdings)
}
