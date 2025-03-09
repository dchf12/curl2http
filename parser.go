package main

import (
	"errors"
	"regexp"
	"strings"
)

// Request 解析されたリクエスト情報
type Request struct {
	Method      string
	URL         string
	Headers     map[string]string
	ContentType string
	Body        string
}

// Parse curlコマンドを解析してRequest構造体を返す
func Parse(curlCmd string) (*Request, error) {
	curlCmd = strings.TrimSpace(curlCmd)

	// default request
	req := &Request{
		Method:  "GET",
		Headers: map[string]string{},
	}

	// 複数行のcurlコマンドを1行にする
	curlCmd = cleanCurlCmd(curlCmd)

	// URLを取得
	urlPattern := regexp.MustCompile(`curl\s+(?:-X\s+\w+\s+)?(['"]?)(https?://[^\s'"]+)(['"]?)`)
	urlMatches := urlPattern.FindStringSubmatch(curlCmd)
	if len(urlMatches) >= 3 {
		req.URL = urlMatches[2]
	}
	// メソッドを取得
	methodPattern := regexp.MustCompile(`-X\s+(\w+)`)
	methodMatches := methodPattern.FindStringSubmatch(curlCmd)
	if len(methodMatches) >= 2 {
		req.Method = methodMatches[1]
	}

	// JSONボディを取得
	jsonPattern := regexp.MustCompile(`--json\s+(?:"([^"]+)"|'([^']+)')`)
	jsonMatches := jsonPattern.FindStringSubmatch(curlCmd)
	if len(jsonMatches) >= 3 {
		req.ContentType = "application/json"
		if jsonMatches[1] != "" {
			req.Body = jsonMatches[1]
		} else {
			req.Body = jsonMatches[2]
		}
	}

	// bodyを取得
	dataPattern := regexp.MustCompile(`-d\s+(['"])(.*?)(['"])|--data\s+(['"])(.*?)(['"])`)
	dataMatches := dataPattern.FindStringSubmatch(curlCmd)
	if len(dataMatches) >= 3 && req.Body == "" {
		if req.ContentType == "" {
			req.ContentType = "application/x-www-form-urlencoded"
		}
		if dataMatches[2] != "" {
			req.Body = dataMatches[2]
		} else if len(dataMatches) >= 6 {
			req.Body = dataMatches[5]
		}
	}

	// headerを取得
	headerPattern := regexp.MustCompile(`-H\s+(['"])(.*?)(['"])`)
	headerMatches := headerPattern.FindAllStringSubmatch(curlCmd, -1)
	for _, match := range headerMatches {
		if len(match) >= 3 {
			headerLine := match[2]
			parts := strings.SplitN(headerLine, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				req.Headers[key] = value

				// Content-Typeヘッダーがある場合は更新
				if strings.ToLower(key) == "content-type" {
					req.ContentType = value
				}
			}
		}
	}

	if req.URL == "" {
		return nil, errors.New("URLが見つかりません")
	}
	return req, nil
}

// cleanCurlCmd 複数行のcurlコマンドを1行にする
func cleanCurlCmd(cmd string) string {
	cmd = strings.ReplaceAll(cmd, "\\\n", " ")
	cmd = regexp.MustCompile(`\s+`).ReplaceAllString(cmd, " ")
	return cmd
}
