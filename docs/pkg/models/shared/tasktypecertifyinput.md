# TaskTypeCertifyInput

The TaskTypeCertify message indicates that a task is a certify task and all related details.

This message contains a oneof named principal. Only a single field of the following list may be set at a time:
  - resource



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Resource`                                                             | [*shared.AppResourceRef](../../../pkg/models/shared/appresourceref.md) | :heavy_minus_sign:                                                     | N/A                                                                    |