# README

Example project for the MCP Server

## Disclaimer

We strongly advise against downloading and testing the binary. This is a demonstration of implementation and is not intended for practical use.

If you choose to do so, please proceed at your own risk. We shall not be held liable for any issues arising from the use of this source code or its distributed binary.

## Cline Configuration

```json
{
  "mcpServers": {
    "mcp-server-start": {
      "name": "mcp-server-start",
      "description": "this is a test server",
      "command": "${ABSOLUTE_PATH_TO_BINARY_DIR}/mcp_start",
      "args": []
    }
  }
}
```
