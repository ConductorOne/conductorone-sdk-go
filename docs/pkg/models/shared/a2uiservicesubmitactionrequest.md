# A2UIServiceSubmitActionRequest

A2UIServiceSubmitActionRequest submits a user action.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `ActionName`                               | **string*                                  | :heavy_minus_sign:                         | The actionName field.                      |
| `ClientTimestamp`                          | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Context`                                  | map[string]*string*                        | :heavy_minus_sign:                         | The context field.                         |
| `ConversationID`                           | **string*                                  | :heavy_minus_sign:                         | The conversationId field.                  |
| `DataModelJSON`                            | **string*                                  | :heavy_minus_sign:                         | The dataModelJson field.                   |
| `SourceComponentID`                        | **string*                                  | :heavy_minus_sign:                         | The sourceComponentId field.               |