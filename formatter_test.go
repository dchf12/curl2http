package main

import "testing"

func Test_Format(t *testing.T) {
	tests := []struct {
		name    string
		request *Request
		want    string
	}{
		{
			name: "基本的なGETリクエスト",
			request: &Request{
				Method:      "GET",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "",
				Body:        "",
			},
			want: "GET https://example.com\n",
		},
		{
			name: "POSTリクエスト（JSONボディ）",
			request: &Request{
				Method:      "POST",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "application/json",
				Body:        `{"A":"B"}`,
			},
			want: "POST https://example.com\ncontent-type: application/json\n\n{\n  \"A\": \"B\"\n}\n",
		},
		{
			name: "ヘッダーありのリクエスト",
			request: &Request{
				Method:  "GET",
				URL:     "https://example.com",
				Headers: map[string]string{"Authorization": "Bearer token", "Content-Type": "application/json"},
				Body:    "",
			},
			want: "GET https://example.com\nAuthorization: Bearer token\nContent-Type: application/json\n",
		},
		{
			name: "無効なJSONボディ",
			request: &Request{
				Method:      "POST",
				URL:         "https://example.com",
				Headers:     map[string]string{},
				ContentType: "application/json",
				Body:        `{"A":"B"`,
			},
			want: "POST https://example.com\ncontent-type: application/json\n\n{\"A\":\"B\"\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.request); got != tt.want {
				t.Errorf("Format() got =%v\nwant =%v", got, tt.want)
			}
		})
	}
}
