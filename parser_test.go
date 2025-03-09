package main

import (
	"reflect"
	"testing"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name    string
		curlCmd string
		want    *Request
		wantErr bool
	}{
		{
			name:    "基本的なGETリクエスト",
			curlCmd: "curl https://example.com",
			want: &Request{
				Method:      "GET",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "",
				Body:        "",
			},
			wantErr: false,
		},
		{
			name:    "メソッド指定あり",
			curlCmd: "curl -X POST https://example.com",
			want: &Request{
				Method:      "POST",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "",
				Body:        "",
			},
			wantErr: false,
		},
		{
			name:    "JSONボディあり",
			curlCmd: `curl -X POST https://example.com --json '{"A":"B"}'`,
			want: &Request{
				Method:      "POST",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "application/json",
				Body:        `{"A":"B"}`,
			},
			wantErr: false,
		},
		{
			name: "複数行のコマンド",
			curlCmd: `curl -X POST https://example.com \
  --json '{"A":"B"}'`,
			want: &Request{
				Method:      "POST",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "application/json",
				Body:        `{"A":"B"}`,
			},
			wantErr: false,
		},
		{
			name:    "ヘッダーあり",
			curlCmd: `curl -X GET https://example.com -H "Authorization: Bearer token"`,
			want: &Request{
				Method:      "GET",
				URL:         "https://example.com",
				Headers:     map[string]string{"Authorization": "Bearer token"},
				ContentType: "",
				Body:        "",
			},
			wantErr: false,
		},
		{
			name:    "不正なコマンド",
			curlCmd: "invalid command",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.curlCmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CleanCurlCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		want string
	}{
		{
			name: "1行",
			cmd:  "curl https://example.com",
			want: "curl https://example.com",
		},
		{
			name: "複数行",
			cmd: `curl -X POST \
					https://example.com \
					-H "Content-Type: application/json" \
					-json '{"key": "value"}'`,
			want: "curl -X POST https://example.com -H \"Content-Type: application/json\" -json '{\"key\": \"value\"}'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanCurlCmd(tt.cmd)
			if got != tt.want {
				t.Errorf("cleanCurlCmd() = %v, want %v", got, tt.want)
			}
		})
	}

}
