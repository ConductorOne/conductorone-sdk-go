# ConnectorActionRefOperation

Which connector RPC this dispatches to.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ConnectorActionRefOperationOperationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ConnectorActionRefOperation("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ConnectorActionRefOperationOperationUnspecified`     | OPERATION_UNSPECIFIED                                 |
| `ConnectorActionRefOperationOperationGrant`           | OPERATION_GRANT                                       |
| `ConnectorActionRefOperationOperationIssueCredential` | OPERATION_ISSUE_CREDENTIAL                            |