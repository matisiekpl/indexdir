package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		logrus.Panic(errors.New("failed to get working directory"))
	}
	logrus.Info("working directory: ", wd)
	entries, err := os.ReadDir("./")
	if err != nil {
		logrus.Panic(errors.New("failed to read directory"))
	}

	buf := bytes.NewBufferString("")
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		buf.WriteString(fmt.Sprintf("<a href='%s'>%s</a><br/>\n", entry.Name(), entry.Name()))
	}
	err = os.WriteFile("index.html", buf.Bytes(), 0644)
	if err != nil {
		logrus.Panic(errors.New("failed to write file"))
	}
	logrus.Info("index.html created")
}
