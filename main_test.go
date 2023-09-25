package main

import (
	"strings"
	"testing"
)

func TestGenerateName(t *testing.T) {
	name := GenerateName()
	ok := strings.Contains(name, ".mp3")
	if !ok {
		t.Fail()
		t.Log("harus berformat.mp3")
	}
}
