# GrantedStatus

Search entitlements with this granted status for your signed in user.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GrantedStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GrantedStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GrantedStatusUnspecified` | UNSPECIFIED                |
| `GrantedStatusAll`         | ALL                        |
| `GrantedStatusGranted`     | GRANTED                    |
| `GrantedStatusNotGranted`  | NOT_GRANTED                |