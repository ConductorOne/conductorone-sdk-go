# PolicyStepInstance

The policy step instance includes a reference to an instance of a policy step that tracks state and has a unique ID.

This message contains a oneof named instance. Only a single field of the following list may be set at a time:
  - approval
  - provision
  - accept
  - reject



## Fields

| Field                                                                                                                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AcceptInstance`                                                                                                                                                                                                                                                                                                          | [*shared.AcceptInstance](../../../pkg/models/shared/acceptinstance.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.<br/> The instance is just a marker for it being copied into an active policy.                                                                     |
| `ApprovalInstance`                                                                                                                                                                                                                                                                                                        | [*shared.ApprovalInstance](../../../pkg/models/shared/approvalinstance.md)                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | The approval instance object describes the way a policy step should be approved as well as its outcomes and state.<br/><br/>This message contains a oneof named outcome. Only a single field of the following list may be set at a time:<br/>  - approved<br/>  - denied<br/>  - reassigned<br/>  - restarted<br/>  - reassignedByError<br/> |
| `ProvisionInstance`                                                                                                                                                                                                                                                                                                       | [*shared.ProvisionInstance](../../../pkg/models/shared/provisioninstance.md)                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | A provision instance describes the specific configuration of an executing provision policy step including actions taken and notification id.<br/><br/>This message contains a oneof named outcome. Only a single field of the following list may be set at a time:<br/>  - completed<br/>  - cancelled<br/>  - errored<br/>  - reassignedByError<br/> |
| `RejectInstance`                                                                                                                                                                                                                                                                                                          | [*shared.RejectInstance](../../../pkg/models/shared/rejectinstance.md)                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.<br/> The instance is just a marker for it being copied into an active policy.                                                                        |
| `ID`                                                                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | The ID of the PolicyStepInstance. This is required by many action submission endpoints to indicate what step you're approving.                                                                                                                                                                                            |
| `PolicyGenerationID`                                                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | The policy generation id refers to the version of the policy that this step was created from.                                                                                                                                                                                                                             |
| `State`                                                                                                                                                                                                                                                                                                                   | [*shared.PolicyStepInstanceState](../../../pkg/models/shared/policystepinstancestate.md)                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                        | The state of the step, which is either active or done.                                                                                                                                                                                                                                                                    |