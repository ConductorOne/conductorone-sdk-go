# ToolGateFilterBuiltInPattern

Built-in pattern selector: a named canned CEL filter maintained by
 ConductorOne. When set, cel_expression must be empty and is resolved from
 the registry at create time; the kind is retained for provenance.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ToolGateFilterBuiltInPatternToolGateBuiltInPatternUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ToolGateFilterBuiltInPattern("custom_value")
```


## Values

| Name                                                                    | Value                                                                   |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternUnspecified`         | TOOL_GATE_BUILT_IN_PATTERN_UNSPECIFIED                                  |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternDestructiveAction`   | TOOL_GATE_BUILT_IN_PATTERN_DESTRUCTIVE_ACTION                           |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternExternalSend`        | TOOL_GATE_BUILT_IN_PATTERN_EXTERNAL_SEND                                |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternThirdPartyConnector` | TOOL_GATE_BUILT_IN_PATTERN_THIRD_PARTY_CONNECTOR                        |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternFinanceAction`       | TOOL_GATE_BUILT_IN_PATTERN_FINANCE_ACTION                               |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternInfraMutation`       | TOOL_GATE_BUILT_IN_PATTERN_INFRA_MUTATION                               |
| `ToolGateFilterBuiltInPatternToolGateBuiltInPatternDbWrite`             | TOOL_GATE_BUILT_IN_PATTERN_DB_WRITE                                     |