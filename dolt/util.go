package main

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

func Execute(cmd *exec.Cmd) (string, error) {
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	outBuf, err := ioutil.ReadAll(outPipe)
	if err != nil {
		return "", err
	}

	errBuf, err := ioutil.ReadAll(errPipe)
	if err != nil {
		return "", err
	}

	if len(errBuf) != 0 {
		return "", errors.New(string(errBuf))
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return string(outBuf), nil
}
