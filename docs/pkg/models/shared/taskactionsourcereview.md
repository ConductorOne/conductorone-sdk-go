# TaskActionSourceReview

The TaskActionSourceReview message tracks which access review decision created the action task.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AccessReviewID`                                                           | `*string`                                                                  | :heavy_minus_sign:                                                         | The ID of the access review associated with the action task.               |
| `CertTicketID`                                                             | `*string`                                                                  | :heavy_minus_sign:                                                         | The ID of the certify ticket that was denied and created this action task. |