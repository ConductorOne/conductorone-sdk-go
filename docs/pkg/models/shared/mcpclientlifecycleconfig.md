# MCPClientLifecycleConfig

MCPClientLifecycleConfig controls how long inactive MCP clients remain
 visible, when their access is closed, and when their records are removed.
 Durations are measured from the client's last activity. Any duration left at
 zero disables that transition.


## Fields

| Field                  | Type                   | Required               | Description            |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `InactivityCloseAfter` | `*string`              | :heavy_minus_sign:     | N/A                    |
| `InactivityHideAfter`  | `*string`              | :heavy_minus_sign:     | N/A                    |
| `RetentionDeleteAfter` | `*string`              | :heavy_minus_sign:     | N/A                    |