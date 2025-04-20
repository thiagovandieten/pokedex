package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "IM ALL CAPS And Lets Throw In Some CamelCases",
			expected: []string{"im", "all", "caps", "and", "lets", "throw", "in", "some", "camelcases"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual, _ := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, but got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			t.Logf("word:%v. expected:%v.", word, expectedWord)

			if word != expectedWord {
				t.Errorf("%v and %v are not the same word", word, expectedWord)
			}
		}
	}
}

func TestCmdExit(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		commandExit()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCmdExit")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	if err := cmd.Run(); err == nil {
		return
	} else {
		t.Fatalf("Process ran with err %v, want exit status 0", err)
	}
}
