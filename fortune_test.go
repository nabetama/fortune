package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func createDummyData(t *testing.T) {
	t.Helper()
	fp, err := os.Create("./dummy")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fp.WriteString("A secure family is a nation's strength.\n%\nDon't rely on motivation for anything - it is fleeting and unreliable.\nDiscipline, however, is unyielding - force youself to follow through.\n%")
}

func deleteDummyData(t *testing.T) {
	t.Helper()
	os.Remove("./dummy")
}

func TestRandomOne(t *testing.T) {
	var tests = []struct{
		name string
		ss []string
		isAvailable bool
	}{
		{"if slice is not empty, can get string", []string{"a"}, true},
		{"if slice is empty, can't get string", []string{}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := randomOne(tc.ss); tc.isAvailable == true && got == "Not matched"{
				t.Errorf("expect なんらか取得できる, but %s", got)
			}
		})
	}
}

func TestGetFortune(t *testing.T) {
	var tests = []struct{
		name     string
		filepath string
		pattern string
		readFile bool
		fortunes int
	}{
		{"ファイルから表示候補のスライスを作れる", "./dummy","", true, 2},
		{"ファイルから表示候補のスライスを作れる", "./dummy", "Discipline, however, is unyielding", true, 1},
	}

	createDummyData(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.readFile == true {
				ss, _ := getFortunes(tc.filepath, tc.pattern)
				if len(ss) != tc.fortunes {
					t.Errorf("候補が作れない")
				}
			}
		})
	}
	deleteDummyData(t)
}

func TestCLI_Run(t *testing.T) {
	type fields struct {
		output io.Writer
	}

	tests := []struct{
		name string
		fields fields
		args []string
		expected int
	}{
		{"ランダムに文字列が表示される", fields{output: os.Stdout}, []string{"", ""}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			cli := &CLI{
				outStream: &buf,
				errStream: os.Stderr,
			}

			cli.Run(tc.args)

			if buf.String() == NotMatchedMessage {
				t.Errorf("want: %s, but %s", "any strings", NotMatchedMessage)
			}
		})
	}
}