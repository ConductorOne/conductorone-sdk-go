# XAAScopeDiscoveryStatusState

State of the latest run.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAScopeDiscoveryStatusStateXaaScopeDiscoveryStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAScopeDiscoveryStatusState("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `XAAScopeDiscoveryStatusStateXaaScopeDiscoveryStateUnspecified` | XAA_SCOPE_DISCOVERY_STATE_UNSPECIFIED                           |
| `XAAScopeDiscoveryStatusStateXaaScopeDiscoveryStateRunning`     | XAA_SCOPE_DISCOVERY_STATE_RUNNING                               |
| `XAAScopeDiscoveryStatusStateXaaScopeDiscoveryStateSucceeded`   | XAA_SCOPE_DISCOVERY_STATE_SUCCEEDED                             |
| `XAAScopeDiscoveryStatusStateXaaScopeDiscoveryStateFailed`      | XAA_SCOPE_DISCOVERY_STATE_FAILED                                |