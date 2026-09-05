# ClientIDType

How the client_id was established.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClientIDTypeClientIDTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClientIDType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ClientIDTypeClientIDTypeUnspecified` | CLIENT_ID_TYPE_UNSPECIFIED            |
| `ClientIDTypeClientIDTypeDcr`         | CLIENT_ID_TYPE_DCR                    |
| `ClientIDTypeClientIDTypeMetadataURL` | CLIENT_ID_TYPE_METADATA_URL           |
| `ClientIDTypeClientIDTypeApp`         | CLIENT_ID_TYPE_APP                    |