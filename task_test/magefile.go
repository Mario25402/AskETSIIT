//go:build mage
// +build mage

package main

import (
	"os"
	"os/exec"
)

// Check verifica la sintaxis de los archivos Go
func Check() error {
	cmd := exec.Command("gofmt", "-e", "../internal")
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.NewFile(0, os.DevNull)
	return cmd.Run()
}

