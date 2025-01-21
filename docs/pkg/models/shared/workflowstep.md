# WorkflowStep

The WorkflowStep message.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - createAccessReview
  - waitForDuration



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `CreateAccessReview`                                                           | [*shared.CreateAccessReview](../../../pkg/models/shared/createaccessreview.md) | :heavy_minus_sign:                                                             | The CreateAccessReview message.                                                |
| `WaitForDuration`                                                              | [*shared.WaitForDuration](../../../pkg/models/shared/waitforduration.md)       | :heavy_minus_sign:                                                             | The WaitForDuration message.                                                   |