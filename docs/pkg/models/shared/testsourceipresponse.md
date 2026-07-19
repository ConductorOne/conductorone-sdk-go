# TestSourceIPResponse

The TestSourceIPResponse message.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Allowed`                                                                             | `*bool`                                                                               | :heavy_minus_sign:                                                                    | Whether the tested IP address is allowed by the CIDR rules.                           |
| `CheckedIP`                                                                           | `*string`                                                                             | :heavy_minus_sign:                                                                    | The IP address that was checked, either from the request or inferred from the caller. |
| `Details`                                                                             | [*shared.Status](../../../pkg/models/shared/status.md)                                | :heavy_minus_sign:                                                                    | N/A                                                                                   |