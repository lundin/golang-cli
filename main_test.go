package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	output := hello()
	if output!="hello" {

t.Errorf("Hello failed, expected %v, got %v","hello",output)

}

}


