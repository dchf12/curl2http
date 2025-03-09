# curl2http

A simple CLI tool written in Go to convert curl commands into human-readable HTTP requests for easier debugging and testing.

## Installation

```shell
go install github.com/dchf12/curl2http@latest
```

## Usage

### From Pipe or Standard Input

```shell
echo 'curl -X POST https://example.com --json "{\"A\":\"B\"}"' | curl2http
```

**Output:**
```
POST https://example.com
content: application/json

{
  "A": "B"
}
```

### From File

```shell
curl2http < request.txt
```

## Error Handling

In case of an invalid curl command, a brief error code will be printed to standard error.

## Testing

Developed using Test-Driven Development (TDD). Aim for a coverage of ~70%.

```shell
go test ./...
```

## License

MIT

