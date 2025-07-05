# CreateAccessReview

The CreateAccessReview message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AccessReviewTemplateCel`                                                     | **string*                                                                     | :heavy_minus_sign:                                                            | The accessReviewTemplateCel field.                                            |
| `AccessReviewTemplateID`                                                      | **string*                                                                     | :heavy_minus_sign:                                                            | The accessReviewTemplateId field.                                             |
| `UseSubjectUser`                                                              | **bool*                                                                       | :heavy_minus_sign:                                                            | If true, the step will use the subject user of the automation as the subject. |
| `UserIdsCel`                                                                  | **string*                                                                     | :heavy_minus_sign:                                                            | The userIdsCel field.                                                         |
| `UserRefs`                                                                    | [][shared.UserRef](../../../pkg/models/shared/userref.md)                     | :heavy_minus_sign:                                                            | The userRefs field.                                                           |