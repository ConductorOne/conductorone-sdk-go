# UpdateSuggestionStateRequest

The UpdateSuggestionStateRequest message.


## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `CreatedCatalogID`                                                                                           | `*string`                                                                                                    | :heavy_minus_sign:                                                                                           | The ID of the access profile created from this suggestion, set when accepting.                               |
| `State`                                                                                                      | [*shared.UpdateSuggestionStateRequestState](../../../pkg/models/shared/updatesuggestionstaterequeststate.md) | :heavy_minus_sign:                                                                                           | The new state to transition the suggestion to.                                                               |