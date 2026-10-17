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
     the state is only released, and a human severity override or assignment
     only cleared, when `override_human_edits` is true.
   - Findings in progress re-derive their content and keep both their state
     and any linked ticket in both modes.
   - Resolved, archived, and deleted findings are skipped in both modes.

 Ticket links are never touched. No routing rule derives an assignee, but a
 transformation rule's set_assignee transform does, so a reprocess can
 reassign a finding -- unless a person assigned it, which is sticky in the
 same way a human severity override is, and yields only to
 `override_human_edits`.


## Fields

| Field                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `OverrideHumanEdits`                                                                                                                                                                                                                                                                                 | `*bool`                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                   | When false (the default), a person's parked state, severity override and<br/> assignment survive the re-derivation. When true, reprocessing additionally<br/> releases findings a person parked, clears human severity overrides, and<br/> lets a matching set_assignee rule reassign a person-assigned finding. |
| `RunDispatchers`                                                                                                                                                                                                                                                                                     | `*bool`                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                   | When true, matched rules may re-send notification dispatches for<br/> findings your team may have already seen. Off by default.                                                                                                                                                                      |