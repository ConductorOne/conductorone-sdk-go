# NameIDFormat

Set this when the service provider requires a specific NameID format. This
 also selects the NameID value semantics: EMAIL_ADDRESS uses the user's
 primary email, TRANSIENT creates a new value for each sign-in, and
 PERSISTENT uses the application's pairwise subject. Immutable once set.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.NameIDFormatSamlNameIDFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.NameIDFormat("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `NameIDFormatSamlNameIDFormatUnspecified`    | SAML_NAME_ID_FORMAT_UNSPECIFIED              |
| `NameIDFormatSamlNameIDFormatPersistent`     | SAML_NAME_ID_FORMAT_PERSISTENT               |
| `NameIDFormatSamlNameIDFormatEmailAddress`   | SAML_NAME_ID_FORMAT_EMAIL_ADDRESS            |
| `NameIDFormatSamlNameIDFormatUnspecifiedUrn` | SAML_NAME_ID_FORMAT_UNSPECIFIED_URN          |
| `NameIDFormatSamlNameIDFormatTransient`      | SAML_NAME_ID_FORMAT_TRANSIENT                |