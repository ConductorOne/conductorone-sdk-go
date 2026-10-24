# EdgeServiceListInferenceRoutesResponse

EdgeServiceListInferenceRoutesResponse contains configured route metadata
 for an application.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `CapabilitySummary`                                                                                    | [*shared.EdgeCapabilitySummary](../../../pkg/models/shared/edgecapabilitysummary.md)                   | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `Routes`                                                                                               | [][shared.EdgeInferenceRouteOption](../../../pkg/models/shared/edgeinferencerouteoption.md)            | :heavy_minus_sign:                                                                                     | The routes field.                                                                                      |
| `UnservedTargets`                                                                                      | [][shared.EdgeInferenceUnservedTarget](../../../pkg/models/shared/edgeinferenceunservedtarget.md)      | :heavy_minus_sign:                                                                                     | Targets of this app's routes that connected Edge processes cannot serve. Empty when none has reported. |