# VirtualMCPPresentation

VirtualMCPPresentation selects one presentation mode and provides a stable
 location for mode-specific settings.

This message contains a oneof named mode. Only a single field of the following list may be set at a time:
  - rawTools
  - codeMode
  - dynamicTools



## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `CodeMode`                                                                                                     | [*shared.VirtualMCPCodeModePresentation](../../../pkg/models/shared/virtualmcpcodemodepresentation.md)         | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `DynamicTools`                                                                                                 | [*shared.VirtualMCPDynamicToolsPresentation](../../../pkg/models/shared/virtualmcpdynamictoolspresentation.md) | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |
| `RawTools`                                                                                                     | [*shared.VirtualMCPRawToolsPresentation](../../../pkg/models/shared/virtualmcprawtoolspresentation.md)         | :heavy_minus_sign:                                                                                             | N/A                                                                                                            |