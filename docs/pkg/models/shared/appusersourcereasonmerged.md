# AppUserSourceReasonMerged

The merge that moved a record onto its account.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `ActorUserID`                                      | `*string`                                          | :heavy_minus_sign:                                 | The ID of the user who confirmed the merge.        |
| `MergeRecordID`                                    | `*string`                                          | :heavy_minus_sign:                                 | The ID of the merge's record in the merge history. |
| `PlanID`                                           | `*string`                                          | :heavy_minus_sign:                                 | The ID of the merge batch the merge belonged to.   |
| `PreviousAppUserID`                                | `*string`                                          | :heavy_minus_sign:                                 | The ID of the account the record was moved off.    |