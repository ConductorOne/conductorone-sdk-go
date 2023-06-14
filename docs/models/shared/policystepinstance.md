# PolicyStepInstance

The PolicyStepInstance message.

This message contains a oneof named instance. Only a single field of the following list may be set at a time:
  - approval
  - provision



## Fields

| Field                                                                                                                                                                                                                   | Type                                                                                                                                                                                                                    | Required                                                                                                                                                                                                                | Description                                                                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Approval`                                                                                                                                                                                                              | [*ApprovalInstance](../../models/shared/approvalinstance.md)                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                      | The ApprovalInstance message.<br/><br/>This message contains a oneof named outcome. Only a single field of the following list may be set at a time:<br/>  - approved<br/>  - denied<br/>  - reassigned<br/>  - restarted<br/>  - reassignedByError<br/> |
| `ID`                                                                                                                                                                                                                    | **string*                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                      | The id field.                                                                                                                                                                                                           |
| `Provision`                                                                                                                                                                                                             | [*ProvisionInstance](../../models/shared/provisioninstance.md)                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                      | The ProvisionInstance message.<br/><br/>This message contains a oneof named outcome. Only a single field of the following list may be set at a time:<br/>  - completed<br/>  - cancelled<br/>  - errored<br/>  - reassignedByError<br/> |
| `State`                                                                                                                                                                                                                 | [*PolicyStepInstanceState](../../models/shared/policystepinstancestate.md)                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                      | The state field.                                                                                                                                                                                                        |