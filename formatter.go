package main

import (
	"encoding/json"
	"sort"
	"strings"
)

func Format(req *Request) string {
	var output strings.Builder

	output.WriteString(req.Method)
	output.WriteString(" ")
	output.WriteString(req.URL)
	output.WriteString("\n")

	var keys []string
	for k := range req.Headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		output.WriteString(k)
		output.WriteString(": ")
		output.WriteString(req.Headers[k])
		output.WriteString("\n")
	}

	if req.ContentType != "" && len(req.Headers) == 0 {
		output.WriteString("content-type: ")
		output.WriteString(req.ContentType)
		output.WriteString("\n")
	}

	if req.Body == "" {
		return output.String()
	}
	output.WriteString("\n")

	if req.ContentType == "application/json" {
		var jsonObj any
		if err := json.Unmarshal([]byte(req.Body), &jsonObj); err == nil {
			if formatted, err := json.MarshalIndent(jsonObj, "", "  "); err == nil {
				output.Write(formatted)
				output.WriteString("\n")
				return output.String()
			}
		}
	}

	output.WriteString(req.Body)
	output.WriteString("\n")
	return output.String()
}
