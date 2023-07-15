package main

import "testing"

func ignore_Example_Main() {
	goMain([]string{})
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./crssy", "Gifu"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
