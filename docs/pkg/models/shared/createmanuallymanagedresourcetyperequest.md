# CreateManuallyManagedResourceTypeRequest

The request message for creating a manually managed resource type.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `DisplayName`                                                     | `string`                                                          | :heavy_check_mark:                                                | The display name for the new resource type.                       |
| `ResourceType`                                                    | [shared.ResourceType](../../../pkg/models/shared/resourcetype.md) | :heavy_check_mark:                                                | The category of the resource type (e.g., ROLE, GROUP, LICENSE).   |