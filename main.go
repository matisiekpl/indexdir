package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	filepath.WalkDir(".", walk)
}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.IsDir() {
		generate(s)
	}
	return nil
}

func generate(wd string) {
	logrus.Info("working directory: ", wd)
	entries, err := os.ReadDir(wd)
	if err != nil {
		logrus.Panic(errors.New("failed to read directory"))
	}

	buf := bytes.NewBufferString("")
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") || entry.Name() == "index.html" {
			continue
		}
		buf.WriteString(fmt.Sprintf("<a href='/%s/%s'>%s</a><br/>\n", wd, entry.Name(), entry.Name()))
	}
	err = os.WriteFile(wd+"/index.html", buf.Bytes(), 0644)
	if err != nil {
		logrus.Panic(errors.New("failed to write file"))
	}
	logrus.Info("index.html created")
}
