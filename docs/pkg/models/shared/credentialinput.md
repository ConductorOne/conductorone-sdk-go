# CredentialInput

A credential input described using ConductorOne's shared form field contract.
 Discovery includes labels, rendering hints and connector restrictions. When
 configuring a requester input, provide the field name, type and supported
 rules: string choices, unique string-list choices with item-count bounds,
 integer choices or inclusive bounds, or a boolean. Maps are fixed-only.
 Provider configuration, including CEL, is not supported for credential inputs.
 Names contain 1–256 bytes. Strings contain at most 4096 bytes, allowed
 sets and lists at most 1024 entries, and integers must be within the exact JSON
 integer range (-9007199254740991 through 9007199254740991). Saved arguments
 and requester policies together are limited to 64 KiB.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `DefaultValue`                                                     | [*shared.DefaultValue](../../../pkg/models/shared/defaultvalue.md) | :heavy_minus_sign:                                                 | N/A                                                                |
| `Field`                                                            | [*shared.FormField](../../../pkg/models/shared/formfield.md)       | :heavy_check_mark:                                                 | N/A                                                                |