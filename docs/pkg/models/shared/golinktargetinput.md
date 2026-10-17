# GoLinkTargetInput

The GoLinkTarget message.

This message contains a oneof named target. Only a single field of the following list may be set at a time:
  - redirect
  - action
  - chooser



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Action`                                                                                   | [*shared.GoLinkActionTarget](../../../pkg/models/shared/golinkactiontarget.md)             | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Chooser`                                                                                  | [*shared.GoLinkChooserTargetInput](../../../pkg/models/shared/golinkchoosertargetinput.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Redirect`                                                                                 | [*shared.GoLinkRedirectTarget](../../../pkg/models/shared/golinkredirecttarget.md)         | :heavy_minus_sign:                                                                         | N/A                                                                                        |