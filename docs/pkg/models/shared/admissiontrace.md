# AdmissionTrace

AdmissionTrace records the bounded authority decision made during admission.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Allocations`                                                             | [][shared.TraceAllocation](../../../pkg/models/shared/traceallocation.md) | :heavy_minus_sign:                                                        | Per-scope limits resolved during admission.                               |
| `AuthorityRefs`                                                           | [][shared.AuthorityRef](../../../pkg/models/shared/authorityref.md)       | :heavy_minus_sign:                                                        | Configurations read during admission.                                     |