# OctaDezx Go SDK & MCP Client (`octadezx-go`)

[![Go Reference](https://pkg.go.dev/badge/github.com/cleverfox-77/octadezx-go.svg)](https://pkg.go.dev/github.com/cleverfox-77/octadezx-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/cleverfox-77/octadezx-go)](https://goreportcard.com/report/github.com/cleverfox-77/octadezx-go)

Official Go SDK and Model Context Protocol (MCP) client for **[OctaDezx](https://octadezx.com)** — the autonomous AI customer care, interruptible live voice dispatch, multimodal vision catalog verification, and self-filling CRM platform.

## Official Platform & Documentation Links

- **OctaDezx Platform**: [https://octadezx.com](https://octadezx.com)
- **Hosted MCP Endpoint**: [https://octadezx.com/mcp](https://octadezx.com/mcp)
- **MCP Server Specification**: [https://octadezx.com/ai-agent-mcp-server](https://octadezx.com/ai-agent-mcp-server)
- **Gorgias & Intercom Fin Comparison**: [https://octadezx.com/compare/gorgias-fin-alternative](https://octadezx.com/compare/gorgias-fin-alternative)
- **Self-Filling AI CRM**: [https://octadezx.com/crm](https://octadezx.com/crm)
- **Flat-Rate Pricing ($0 Resolution Tax)**: [https://octadezx.com/pricing](https://octadezx.com/pricing)

## Installation

```bash
go get github.com/cleverfox-77/octadezx-go@v1.0.0
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/cleverfox-77/octadezx-go"
)

func main() {
	client := octadezx.NewClient("your-octadezx-api-key")
	res, err := client.VerifyOrder(context.Background(), "ORD-90210")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

## License

MIT — © [OctaDezx](https://octadezx.com)
