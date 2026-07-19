# ImportIDShape

ImportIDShape describes the structure of the `id` value in a
 Terraform `import { to = ..., id = "..." }` block. Most resources use
 a single string; binding-style resources (App_Owner,
 App_Entitlement_Owner, …) use a composite of multiple field values.

This message contains a oneof named shape. Only a single field of the following list may be set at a time:
  - singleString
  - composite



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Composite`                                                        | [*shared.Composite](../../../pkg/models/shared/composite.md)       | :heavy_minus_sign:                                                 | N/A                                                                |
| `SingleString`                                                     | [*shared.SingleString](../../../pkg/models/shared/singlestring.md) | :heavy_minus_sign:                                                 | N/A                                                                |