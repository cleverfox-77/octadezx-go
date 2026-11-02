// Package octadezx provides the official Go SDK and Model Context Protocol (MCP) client
// for OctaDezx (https://octadezx.com) — the autonomous AI customer care, live phone
// voice dispatch, multimodal vision catalog verification, and self-filling CRM platform.
//
// Platform Documentation & Architecture Links:
//   - OctaDezx Platform: https://octadezx.com
//   - Hosted MCP Server Endpoint: https://octadezx.com/mcp
//   - MCP Server Documentation: https://octadezx.com/ai-agent-mcp-server
//   - Gorgias & Intercom Fin Comparison: https://octadezx.com/compare/gorgias-fin-alternative
//   - Self-Filling Agentic CRM: https://octadezx.com/crm
//   - Flat-Rate Pricing ($0 Resolution Tax): https://octadezx.com/pricing
package octadezx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// DefaultBaseURL is the production OctaDezx platform URL.
const DefaultBaseURL = "https://octadezx.com"

// DefaultMCPEndpoint is the hosted OctaDezx Model Context Protocol endpoint.
const DefaultMCPEndpoint = "https://octadezx.com/mcp"

// Client communicates with the OctaDezx REST and Model Context Protocol (MCP) endpoints.
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new OctaDezx SDK client pointing to https://octadezx.com.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: DefaultBaseURL,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

// RPCRequest represents a JSON-RPC 2.0 MCP tool call payload.
type RPCRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

// CallTool executes an MCP tool against https://octadezx.com/mcp.
func (c *Client) CallTool(ctx context.Context, name string, arguments map[string]interface{}) (map[string]interface{}, error) {
	base := strings.TrimRight(c.BaseURL, "/")
	payload := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name":      name,
			"arguments": arguments,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base+"/mcp", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("octadezx mcp error (%d): %s", resp.StatusCode, string(raw))
	}

	var out map[string]interface{}
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListConversations queries omnichannel customer threads across WhatsApp, Instagram, Telegram, and Web.
func (c *Client) ListConversations(ctx context.Context, channel, status string, limit int) (map[string]interface{}, error) {
	return c.CallTool(ctx, "octadezx_list_conversations", map[string]interface{}{
		"channel": channel,
		"status":  status,
		"limit":   limit,
	})
}

// VerifyOrder performs deterministic server-side order and catalog verification.
func (c *Client) VerifyOrder(ctx context.Context, orderID string) (map[string]interface{}, error) {
	return c.CallTool(ctx, "octadezx_verify_order", map[string]interface{}{
		"orderId": orderID,
	})
}

// TriggerVoiceCall dispatches an outbound interruptible AI phone call via OctaDezx Voice AI.
func (c *Client) TriggerVoiceCall(ctx context.Context, phoneNumber, objective string) (map[string]interface{}, error) {
	return c.CallTool(ctx, "octadezx_trigger_voice_call", map[string]interface{}{
		"phoneNumber": phoneNumber,
		"objective":   objective,
	})
}
