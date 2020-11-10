package main

import (
	"log"
	"os/exec"
	"testing"

	"github.com/logrusorgru/aurora"
)

// test-case struct definition
type TestCase struct {
	input  string
	flags  []string
	output string
	name   string
}

func asString(value interface{}) string {
	return aurora.Sprintf(value)
}

// slice of test-cases
var testCases = []TestCase{
	{
		input:  "hello world",
		flags:  nil,
		output: "hello world\n",
		name:   "Test that when no flags are provided, the original string is returned with a newline appended.",
	},
	{
		input:  "hello world",
		flags:  []string{"-n"},
		output: "hello world",
		name:   "Test that when no flags are provided, the original string is returned with a newline appended.",
	},
}

// -- dynamically add a test-case for each flag
func addFlagTests(flags map[string]func(str interface{}) aurora.Value) []TestCase {
	for flag, formatter := range flags {
		var testCase = TestCase{
			input:  "hello world",
			flags:  []string{"-n", flag},
			output: asString(formatter("hello world")),
			name:   "works as expected for flag " + flag,
		}

		testCases = append(testCases, testCase)
	}

	return testCases
}

// Test app-level behaviour via a subprocess.
func TestGecko(test *testing.T) {
	testCases = addFlagTests(formatFlags)
	testCases = addFlagTests(foregroundFlags)
	testCases = addFlagTests(backgroundFlags)

	for _, testData := range testCases {
		test.Run(testData.name, func(test *testing.T) {

			inputSlice := make([]string, 1, 1)
			inputSlice[0] = testData.input

			arguments := append(inputSlice, testData.flags...)

			out, err := exec.Command("./gecko", arguments...).Output()
			if err != nil {
				log.Fatal(err)
			}

			if string(out) != testData.output {
				test.Errorf("expected %v, got %v", string(out), string(out))
			}
		})
	}
}

// Benchmark echo when called from the CLI with various flags, as a
// control for gecko's performance.
func BenchmarkEchoControl(bench *testing.B) {
	for count := 0; count < bench.N; count++ {
		for _, testData := range testCases {
			cmd := exec.Command("./gecko", testData.input)

			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Benchmark Gecko when called from the CLI with various flags
func BenchmarkGecko(bench *testing.B) {
	for count := 0; count < bench.N; count++ {
		for _, testData := range testCases {
			inputSlice := make([]string, 1, 1)
			inputSlice[0] = testData.input

			arguments := append(inputSlice, testData.flags...)

			cmd := exec.Command("./gecko", arguments...)
			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
