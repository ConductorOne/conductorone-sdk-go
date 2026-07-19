# Kind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.KindDecoyKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Kind("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `KindDecoyKindUnspecified`          | DECOY_KIND_UNSPECIFIED              |
| `KindDecoyKindUserClientCredential` | DECOY_KIND_USER_CLIENT_CREDENTIAL   |
| `KindDecoyKindConnectorClient`      | DECOY_KIND_CONNECTOR_CLIENT         |
| `KindDecoyKindWorkloadFederation`   | DECOY_KIND_WORKLOAD_FEDERATION      |
| `KindDecoyKindAccessToken`          | DECOY_KIND_ACCESS_TOKEN             |