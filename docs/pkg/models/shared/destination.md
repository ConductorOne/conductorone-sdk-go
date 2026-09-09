# Destination

Where the claim is released.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DestinationOidcClaimDestinationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Destination("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `DestinationOidcClaimDestinationUnspecified`  | OIDC_CLAIM_DESTINATION_UNSPECIFIED            |
| `DestinationOidcClaimDestinationIDTokenOnly`  | OIDC_CLAIM_DESTINATION_ID_TOKEN_ONLY          |
| `DestinationOidcClaimDestinationUserinfoOnly` | OIDC_CLAIM_DESTINATION_USERINFO_ONLY          |