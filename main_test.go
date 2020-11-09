package main

import (
	"log"
	"os/exec"
	"testing"
)

// Test app-level behaviour via a subprocess.
func Test(test *testing.T) {
	testStrings := []string{"hello world", ""}

	for _, testString := range testStrings {
		test.Run("returns original string when called without flags set, plus a terminal newline", func(test *testing.T) {
			out, err := exec.Command("./gecko", testString).Output()
			if err != nil {
				log.Fatal(err)
			}

			if string(out) != testString+"\n" {
				test.Errorf("expected %v, got %v", string(out), string(out))
			}
		})
	}

	for _, testString := range testStrings {
		test.Run("returns original string when called only -n is set", func(test *testing.T) {
			out, err := exec.Command("./gecko", testString, "-n").Output()
			if err != nil {
				log.Fatal(err)
			}

			if string(out) != testString {
				test.Errorf("expected %v, got %v", string(out), string(out))
			}
		})
	}

}
