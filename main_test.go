package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("5\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}

func Test_checkNumbers(t *testing.T) {
	checkNumbersTests := []struct {
		name     string
		number   *strings.Reader
		expected bool
		msg      string
	}{
		{"prime", strings.NewReader("3"), false, "3 is a prime number!"},
		{"not prime", strings.NewReader("4"), false, "4 is not a prime number because it is divisible by 2!"},
		{"not number", strings.NewReader("st"), false, "Please enter a whole number!"},
		{"leave", strings.NewReader("q"), true, ""},
	}

	for _, test := range checkNumbersTests {
		msg, result := checkNumbers(bufio.NewScanner(test.number))
		if test.expected && !result {
			t.Errorf("%s: expected true but got false", test.name)
		}

		if !test.expected && result {
			t.Errorf("%s: expected false but got true", test.name)
		}

		if test.msg != msg {
			t.Errorf("%s: expected %s but got %s", test.name, test.msg, msg)
		}
	}
}

func Test_intro(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = old

	out, _ := io.ReadAll(r)

	text := `Is it Prime?
------------
Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.
-> `
	if string(out) != text {
		t.Errorf("Is wrong text of intro %s", string(out))
	}

}

func Test_prompt(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()
	_ = w.Close()

	os.Stdout = old

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("Is wrong output of prompt %s", string(out))
	}

}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
