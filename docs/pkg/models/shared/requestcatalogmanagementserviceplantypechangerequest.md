# RequestCatalogManagementServicePlanTypeChangeRequest

Requests a side-effect-free plan for changing an access profile's type.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `TargetType`                                                               | [*shared.TargetType](../../../pkg/models/shared/targettype.md)             | :heavy_minus_sign:                                                         | Requested type. UNSPECIFIED and the deprecated PROFILE value are rejected. |