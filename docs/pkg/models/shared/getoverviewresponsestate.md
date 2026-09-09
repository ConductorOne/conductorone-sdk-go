# GetOverviewResponseState

Whether spend governance is configured for the organization.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GetOverviewResponseStateOverviewStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GetOverviewResponseState("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `GetOverviewResponseStateOverviewStateUnspecified` | OVERVIEW_STATE_UNSPECIFIED                         |
| `GetOverviewResponseStateOverviewStateResolved`    | OVERVIEW_STATE_RESOLVED                            |
| `GetOverviewResponseStateOverviewStateNoPolicy`    | OVERVIEW_STATE_NO_POLICY                           |
| `GetOverviewResponseStateOverviewStateNoCeiling`   | OVERVIEW_STATE_NO_CEILING                          |