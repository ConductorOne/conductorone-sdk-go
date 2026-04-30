# BusinessHours

BusinessHours defines a weekly time window in a specific timezone.


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `Days`                     | []`int`                    | :heavy_minus_sign:         | 0=Sun, 1=Mon, ..., 6=Sat.  |
| `End`                      | `*string`                  | :heavy_minus_sign:         | "HH:MM" in 24-hour format. |
| `Start`                    | `*string`                  | :heavy_minus_sign:         | "HH:MM" in 24-hour format. |
| `Timezone`                 | `*string`                  | :heavy_minus_sign:         | The timezone field.        |