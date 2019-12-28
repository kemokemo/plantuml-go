package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_run(t *testing.T) {
	type test struct {
		name string
		opt  option
		help bool
		args []string
		want int
	}

	tests := []test{
		{"help", option{}, true, []string{}, 0},
		{"test.puml", option{"http://plantuml.com/plantuml", "png", styleTxt}, false, []string{filepath.Clean("./test-files/test.puml")}, 0},
		{"invalid file path", option{"http://plantuml.com/plantuml", "png", styleTxt}, false, []string{filepath.Clean("./test-files/hoge.puml")}, 1},
		{"no input file", option{"http://plantuml.com/plantuml", "png", styleTxt}, false, []string{}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt.server = tt.opt.server
			opt.format = tt.opt.format
			opt.style = tt.opt.style
			help = tt.help
			fileList = tt.args

			if got := run(); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}

	// Open and pass file instead of os.Stdin
	f, err := os.Open(filepath.Clean("./test-files/test.puml"))
	if err != nil {
		t.Errorf("failed to open file:%v", err)
	}
	defer f.Close()
	os.Stdin = f

	tests = []test{
		{"help with stdin", option{}, true, []string{}, 0},
		{"stdin", option{"http://plantuml.com/plantuml", "png", styleTxt}, false, []string{}, 0},
		{"invalid option with stdin", option{"", "", ""}, false, []string{}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt.server = tt.opt.server
			opt.format = tt.opt.format
			opt.style = tt.opt.style
			help = tt.help
			fileList = tt.args

			if got := run(); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
