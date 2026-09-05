# BulkReprocessAction

BulkReprocessAction re-evaluates eligible findings against transformation
 and routing rules using each finding's original detector-created state
 (original severity, original annotations) rather than any rule-mutated
 current state.

 `override_human_edits` chooses how far re-derivation goes for
 human-attributed edits:

   - Open findings are re-derived and re-routed in both modes.
   - Findings parked by a rule (snoozed, suppressed, or risk-accepted by a
     routing rule with no subsequent human action) are released to open,
     re-derived, and re-routed in both modes.
   - Findings parked by a person re-derive their content in both modes, but
     the state is only released, and a human severity override only cleared,
     when `override_human_edits` is true.
   - Findings in progress re-derive their content and keep both their state
     and any linked ticket in both modes.
   - Resolved, archived, and deleted findings are skipped in both modes.

 Assigned owners and ticket links are never touched; rules do not derive them.


## Fields

| Field                                                                                                                                                                                                            | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `OverrideHumanEdits`                                                                                                                                                                                             | `*bool`                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                               | When false (the default), a person's parked state and severity override<br/> survive the re-derivation. When true, reprocessing additionally releases<br/> findings a person parked and clears human severity overrides. |
| `RunDispatchers`                                                                                                                                                                                                 | `*bool`                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                               | When true, matched rules may re-send notification dispatches for<br/> findings your team may have already seen. Off by default.                                                                                  |