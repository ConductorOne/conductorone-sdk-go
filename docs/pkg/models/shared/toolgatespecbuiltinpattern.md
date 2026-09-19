# ToolGateSpecBuiltInPattern

Built-in pattern selector; when set, cel_expression is resolved from the
 registry at instantiation.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ToolGateSpecBuiltInPatternToolGateBuiltInPatternUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ToolGateSpecBuiltInPattern("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternUnspecified`         | TOOL_GATE_BUILT_IN_PATTERN_UNSPECIFIED                                |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternDestructiveAction`   | TOOL_GATE_BUILT_IN_PATTERN_DESTRUCTIVE_ACTION                         |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternExternalSend`        | TOOL_GATE_BUILT_IN_PATTERN_EXTERNAL_SEND                              |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternThirdPartyConnector` | TOOL_GATE_BUILT_IN_PATTERN_THIRD_PARTY_CONNECTOR                      |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternFinanceAction`       | TOOL_GATE_BUILT_IN_PATTERN_FINANCE_ACTION                             |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternInfraMutation`       | TOOL_GATE_BUILT_IN_PATTERN_INFRA_MUTATION                             |
| `ToolGateSpecBuiltInPatternToolGateBuiltInPatternDbWrite`             | TOOL_GATE_BUILT_IN_PATTERN_DB_WRITE                                   |