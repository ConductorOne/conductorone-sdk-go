# AuditVerbosity

How much detail is captured in the audit log for MCP tool calls.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AuditVerbosityAuditVerbosityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AuditVerbosity("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AuditVerbosityAuditVerbosityUnspecified` | AUDIT_VERBOSITY_UNSPECIFIED               |
| `AuditVerbosityAuditVerbosityMinimal`     | AUDIT_VERBOSITY_MINIMAL                   |
| `AuditVerbosityAuditVerbosityStandard`    | AUDIT_VERBOSITY_STANDARD                  |
| `AuditVerbosityAuditVerbosityFull`        | AUDIT_VERBOSITY_FULL                      |