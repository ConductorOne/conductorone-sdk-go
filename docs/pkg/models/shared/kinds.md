# Kinds

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.KindsDecoyKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Kinds("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `KindsDecoyKindUnspecified`          | DECOY_KIND_UNSPECIFIED               |
| `KindsDecoyKindUserClientCredential` | DECOY_KIND_USER_CLIENT_CREDENTIAL    |
| `KindsDecoyKindConnectorClient`      | DECOY_KIND_CONNECTOR_CLIENT          |
| `KindsDecoyKindWorkloadFederation`   | DECOY_KIND_WORKLOAD_FEDERATION       |
| `KindsDecoyKindAccessToken`          | DECOY_KIND_ACCESS_TOKEN              |