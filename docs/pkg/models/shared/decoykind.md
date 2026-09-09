# DecoyKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DecoyKindDecoyKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DecoyKind("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `DecoyKindDecoyKindUnspecified`          | DECOY_KIND_UNSPECIFIED                   |
| `DecoyKindDecoyKindUserClientCredential` | DECOY_KIND_USER_CLIENT_CREDENTIAL        |
| `DecoyKindDecoyKindConnectorClient`      | DECOY_KIND_CONNECTOR_CLIENT              |
| `DecoyKindDecoyKindWorkloadFederation`   | DECOY_KIND_WORKLOAD_FEDERATION           |
| `DecoyKindDecoyKindAccessToken`          | DECOY_KIND_ACCESS_TOKEN                  |