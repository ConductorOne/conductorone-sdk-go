# ServicePrincipalServiceListBindingsResponse

The ServicePrincipalServiceListBindingsResponse message.


## Fields

| Field                                                                                                       | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `Bindings`                                                                                                  | [][shared.ServicePrincipalBinding](../../../pkg/models/shared/serviceprincipalbinding.md)                   | :heavy_minus_sign:                                                                                          | Active bindings held by the subject in this page. Empty when the<br/> subject is unbound. Order is unspecified. |
| `NextPageToken`                                                                                             | `*string`                                                                                                   | :heavy_minus_sign:                                                                                          | The nextPageToken field.                                                                                    |