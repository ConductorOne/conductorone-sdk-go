# ExternalOAuthClientInfoClientIDType

The clientIdType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExternalOAuthClientInfoClientIDTypeClientIDTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExternalOAuthClientInfoClientIDType("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ExternalOAuthClientInfoClientIDTypeClientIDTypeUnspecified` | CLIENT_ID_TYPE_UNSPECIFIED                                   |
| `ExternalOAuthClientInfoClientIDTypeClientIDTypeDcr`         | CLIENT_ID_TYPE_DCR                                           |
| `ExternalOAuthClientInfoClientIDTypeClientIDTypeMetadataURL` | CLIENT_ID_TYPE_METADATA_URL                                  |
| `ExternalOAuthClientInfoClientIDTypeClientIDTypeApp`         | CLIENT_ID_TYPE_APP                                           |
| `ExternalOAuthClientInfoClientIDTypeClientIDTypeStatic`      | CLIENT_ID_TYPE_STATIC                                        |