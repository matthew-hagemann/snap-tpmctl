package main_test

import "testing"

type AppMock struct{}

func (a AppMock) Run() error {
	return nil
}

func TestRun(t *testing.T) {}
