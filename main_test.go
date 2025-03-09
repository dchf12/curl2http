package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "シンプルな入力",
			input:   "curl https://example.com",
			want:    "curl https://example.com",
			wantErr: false,
		},
		{
			name: "複数行の入力",
			input: `curl -X POST https://example.com \
--json '{"A":"B"}'`,
			want: `curl -X POST https://example.com \
--json '{"A":"B"}'`,
			wantErr: false,
		},
		{
			name:    "空の入力",
			input:   "",
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got, err := readInput(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Run(t *testing.T) {
	reader := strings.NewReader("curl -X POST https://example.com --json '{\"A\":\"B\"}'")
	var writer bytes.Buffer

	err := run(reader, &writer)
	if err != nil {
		t.Errorf("run() error = %v", err)
	}

	expected := "POST https://example.com\ncontent-type: application/json\n\n{\n  \"A\": \"B\"\n}\n"
	if writer.String() != expected {
		t.Errorf("run() got = %v, want %v", writer.String(), expected)
	}
}
