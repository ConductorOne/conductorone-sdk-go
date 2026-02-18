# StepUpMicrosoftSettings

StepUpMicrosoftSettings configures a Microsoft Entra step-up provider using Conditional Access.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ConditionalAccessIds`                                                              | []*string*                                                                          | :heavy_minus_sign:                                                                  | Authentication context IDs (C1-C99). Required for ACRS mode; ignored for OIDC mode. |
| `Tenant`                                                                            | **string*                                                                           | :heavy_minus_sign:                                                                  | Microsoft Entra tenant ID (GUID or domain). Used for response validation.           |
| `ValidationMode`                                                                    | [*shared.ValidationMode](../../../pkg/models/shared/validationmode.md)              | :heavy_minus_sign:                                                                  | Validation approach. See MicrosoftValidationMode for details on each mode.          |