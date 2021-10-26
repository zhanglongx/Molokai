package core

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/zhanglongx/Molokai/tsWrapper"
)

const (
	VERSION = "v1"
	RCFILE  = "~/.molokairc"
)

type Molokai struct {
	Version string `yaml:"version"`

	Token string `yaml:"token"`

	Holdings []Holding `yaml:"holdings"`

	Reminders Reminders `yaml:"reminders"`
}

var (
	errVersion    = errors.New("version not match")
	errRCNotFound = errors.New("molokai rc file not found")
)

func (m *Molokai) Init() error {
	if m.Version != VERSION {
		return errVersion
	}

	if err := initCfg(); err != nil {
		return err
	}

	if err := m.Reminders.Init(); err != nil {
		return err
	}

	tsWrapper.Token = m.Token

	return nil
}

// Runholdings launches runners on each holding
func (m *Molokai) RunHoldings() error {
	if len(m.Holdings) == 0 {
		log.Printf("holding has none, check the input")
		return nil
	}

	for _, h := range m.Holdings {
		if err := h.Run(m.Reminders); err != nil {
			log.Printf("run %s failed", h.Symbol)
			continue
		}
	}

	return nil
}

var smtpCfg struct {
	From struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	Smtp struct {
		Smtp string `yaml:"smtp"`
		Port int    `yaml:"port"`
	}
}

func initCfg() error {
	usr, _ := user.Current()
	dir := usr.HomeDir

	file := RCFILE
	if file == "~" {
		// In case of "~", which won't be caught by the "else if"
		file = dir
	} else if strings.HasPrefix(file, "~/") {
		// Use strings.HasPrefix so we don't match files like
		// "/something/~/something/"
		file = filepath.Join(dir, file[2:])
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return errRCNotFound
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := yaml.UnmarshalStrict(buf, &smtpCfg); err != nil {
		return err
	}

	return nil
}
