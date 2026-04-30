## Go SDK Changes:
* `ConductoroneApi.AccessReview.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `ErrorState` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `ErrorState` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[].AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `ErrorState` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `ErrorState` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScheduledStartDate` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoGenerateReport` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView` **Added**
    - `ExemptCertifiedAccessConflicts` **Added**
    - `IsCampaignScheduleEnabled` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `RecurrenceRule` **Added**
    - `ReviewInstructions` **Added**
    - `ReviewSignatureConfig` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
    - `UsePolicyOverride` **Added**
  * `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.SkipStep()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Restart()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Reassign()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ProcessNow()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.HardReset()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Deny()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Comment()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Close()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskActions.Approve()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.Get()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateRevokeTask()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateOffboardingTask()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.Task.CreateGrantTask()`: 
  *  `request.Request.TaskGrantSource.IsExtension` **Added**
  * `response.TaskView` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.TaskAudit.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated.PolicyActionInstance` **Added**
    - `TaskAuditActionInstanceFailed.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceFailed.PolicyActionInstance` **Added**
    - `TaskAuditActionInstanceSucceeded.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceSucceeded.PolicyActionInstance` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `TaskAuditCreatedReplacementExtensionGrantTask` **Added**
    - `TaskAuditMetaData.User.Origin` **Added**
    - `TaskAuditNewTaskCreatedFrom` **Added**
    - `TaskAuditReassignmentFallbackToAdmin` **Added**
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed**
    - `ExcludeApplicationIds` **Added**
    - `PendingActionFilter` **Added**
    - `RequireApprovalReason` **Added**
    - `RequireDenialReason` **Added**
    - `TaskTypes[].TaskTypeAction.ScopeRole` **Added**
    - `TaskTypes[].TaskTypeAction.TaskActionInstance` **Added**
    - `TaskTypes[].TaskTypeFinding` **Added**
    - `TaskTypes[].TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `ApproversPath` **Added**
    - `ResourceBindingsPath` **Added**
    - `RoleResourcePath` **Added**
    - `ScopeResourcePath` **Added**
    - `Task.ApproverIds` **Added**
    - `Task.Form.FieldRelationships[].DependentOn` **Added**
    - `Task.Form.Fields[].Required` **Added**
    - `Task.Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Task.Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Task.Form.Fields[].StringMapField` **Added**
    - `Task.Origin.Enum(taskOriginCascadeDelete)` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Action.ActionTargetBatonResourceAction` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].PolicyForm` **Added**
    - `Task.PolicyInstance.Policy.PolicySteps.Map<PolicySteps>.Steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.RevocationTargets` **Added**
    - `Task.TaskType.TaskTypeAction.DisplayName` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.Type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.TaskType.TaskTypeGrant.TaskGrantSource.IsExtension` **Added**
* `ConductoroneApi.PolicySearch.Search()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.FunctionsSearch.Search()`: 
  *  `request.Request.FunctionTypes[].Enum(functionTypeCodeMode)` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
    - `UseSpn` **Added**
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.List()`: `response.List[].PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Get()`: `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request.PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.PolicySteps.Map<PolicySteps>.Steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetBatonResourceAction` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
    - `UseSpn` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
    - `UseSpn` **Added**
* `ConductoroneApi.Functions.GetFunction()`: `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
    - `UseSpn` **Added**
* `ConductoroneApi.Functions.CreateFunction()`: 
  *  `request.Request.FunctionType.Enum(functionTypeCodeMode)` **Added**
  * `response.Function` **Changed** (Breaking ⚠️)
    - `EncryptedValues` **Removed** (Breaking ⚠️)
    - `FunctionType.Enum(functionTypeCodeMode)` **Added**
    - `Secret` **Added**
    - `UseSpn` **Added**
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed** (Breaking ⚠️)
    - `ConfigSchema.FieldGroups[].FieldNames` **Added**
    - `ConfigSchema.FieldGroups[].Fields` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].CheckboxField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].ConnectorCheckboxField` **Added**
    - `ConfigSchema.Fields[].ConnectorSelectField` **Added**
    - `ConfigSchema.Fields[].ConnectorStringField` **Added**
    - `ConfigSchema.Fields[].ConnectorStringMapField` **Added**
    - `ConfigSchema.Fields[].ConnectorTextField` **Added**
    - `ConfigSchema.Fields[].KeyValueField.SupportsFileUpload` **Added**
    - `ConfigSchema.Fields[].SelectField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].StringField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].StringMapField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.Fields[].TextField` **Removed** (Breaking ⚠️)
    - `Form.FieldRelationships[].DependentOn` **Added**
    - `Form.Fields[].Required` **Added**
    - `Form.Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Form.Fields[].StringField.TextField.Suffix` **Added**
    - `Form.Fields[].StringMapField` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].AutomationsTaskAction` **Added**
    - `Automation.AutomationSteps[].AutomationsWebhook` **Added**
    - `Automation.AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `Automation.AutomationSteps[].SendEmail.EmailCel` **Added**
    - `Automation.AutomationSteps[].SendEmail.Email` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `Automation.AutomationSteps[].SetCredential` **Added**
    - `Automation.AutomationSteps[].StoreCredential` **Added**
    - `Automation.AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `Automation.DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `Automation.PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
    - `WebhookCapabilityUrl` **Added**
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response.Automation` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
* `ConductoroneApi.Automation.DeleteAutomation()`: `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationsDeleteAutomationRequest` **Added**
    - `DeleteAutomationRequest` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].AutomationsTaskAction` **Added**
    - `Automation.AutomationSteps[].AutomationsWebhook` **Added**
    - `Automation.AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `Automation.AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `Automation.AutomationSteps[].SendEmail.EmailCel` **Added**
    - `Automation.AutomationSteps[].SendEmail.Email` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `Automation.AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `Automation.AutomationSteps[].SetCredential` **Added**
    - `Automation.AutomationSteps[].StoreCredential` **Added**
    - `Automation.AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `Automation.AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `Automation.DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `Automation.PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
    - `WebhookCapabilityUrl` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: 
  * `request.Request` **Changed**
    - `Direction` **Added**
    - `SortField` **Added**
    - `TriggerTypes[].Enum(triggerTypeScheduleNoUser)` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
    - `PrimaryTriggerType.Enum(triggerTypeScheduleNoUser)` **Added**
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AutomationsTaskAction` **Added**
    - `AutomationSteps[].AutomationsWebhook` **Added**
    - `AutomationSteps[].ConnectorCreateAccount.PasswordCel` **Added**
    - `AutomationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `AutomationSteps[].GeneratePassword.PasswordPolicyId` **Added**
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.AppEntitlementRefs` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionList` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementExclusionNone` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionCriteria` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionListCel` **Added**
    - `AutomationSteps[].GrantEntitlements.GrantEntitlementInclusionList` **Added**
    - `AutomationSteps[].SendEmail.EmailCel` **Added**
    - `AutomationSteps[].SendEmail.Email` **Added**
    - `AutomationSteps[].SendSlackMessage.UseSubjectUser` **Added**
    - `AutomationSteps[].SendSlackMessage.UserIdsCel` **Added**
    - `AutomationSteps[].SendSlackMessage.UserRefs` **Added**
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Triggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Triggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].ScheduleTriggerNoUser` **Added**
    - `Triggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityUrl` **Added**
* `ConductoroneApi.AppEntitlements.DeleteAutomation()`: `request.Request` **Changed** (Breaking ⚠️)
    - `AppDeleteAutomationRequest` **Added**
    - `DeleteAutomationRequest` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.CreateAutomation()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppCreateAutomationRequest` **Added**
    - `CreateAutomationRequest` **Removed** (Breaking ⚠️)
  *  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
* `ConductoroneApi.AccessConflict.UpdateMonitor()`: 
  * `request.Request.ConflictMonitorUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessConflict.GetMonitor()`: `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessConflict.CreateMonitor()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
  * `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.ComplianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.RiskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `AccuracyIssueAction` **Added**
    - `AutoCloseCampaign` **Added**
    - `AutoCloseDecision` **Added**
    - `AutoStartCampaign` **Added**
    - `DefaultView.Enum(accessReviewViewTypeByResource)` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByInheritance)` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByResource)` **Added**
* `ConductoroneApi.SsfReceiverStream.Test()`: **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  * `request.Request.ConnectorServiceUpdateDelegatedRequest.Connector` **Changed**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.RoleMiningManagement.CreateAccessProfileFromCohort()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetCustomAnalysisResult()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetLatestRun()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetRoleMiningConfig()`: **Added**
* `ConductoroneApi.RoleMiningManagement.GetSuggestion()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListRuns()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListSuggestions()`: **Added**
* `ConductoroneApi.RoleMiningManagement.SearchCohortUsers()`: **Added**
* `ConductoroneApi.RoleMiningManagement.TriggerAnalysis()`: **Added**
* `ConductoroneApi.RoleMiningManagement.TriggerCustomAnalysis()`: **Added**
* `ConductoroneApi.RoleMiningManagement.UpdateRoleMiningConfig()`: **Added**
* `ConductoroneApi.RoleMiningManagement.UpdateSuggestionState()`: **Added**
* `ConductoroneApi.AutomationExecutionSearch.SearchAllAutomationExecutions()`: **Added**
* `ConductoroneApi.AppSearch.SearchUserOwnership()`: **Added**
* `ConductoroneApi.HooksSearch.Search()`: **Added**
* `ConductoroneApi.ExternalClientSearch.Search()`: **Added**
* `ConductoroneApi.RoleMiningManagementSearch.Search()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Get()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.Search()`: **Added**
* `ConductoroneApi.PaperSecretAdmin.SearchAuditEvents()`: **Added**
* `ConductoroneApi.PaperSecret.CreateExternal()`: **Added**
* `ConductoroneApi.PaperSecret.CreateInternal()`: **Added**
* `ConductoroneApi.PaperSecret.Get()`: **Added**
* `ConductoroneApi.PaperSecret.GetByShareCode()`: **Added**
* `ConductoroneApi.PaperSecret.GetContent()`: **Added**
* `ConductoroneApi.PaperSecret.Revoke()`: **Added**
* `ConductoroneApi.PaperSecret.SearchAuditEvents()`: **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: **Added**
* `ConductoroneApi.PaperSecret.SetTextContent()`: **Added**
* `ConductoroneApi.SsfReceiverEventSearch.Search()`: **Added**
* `ConductoroneApi.WorkloadFederation.CreateProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.CreateTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.DeleteProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.DeleteTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.GetProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.GetTrust()`: **Added**
* `ConductoroneApi.WorkloadFederation.ListProviders()`: **Added**
* `ConductoroneApi.WorkloadFederation.ListTrusts()`: **Added**
* `ConductoroneApi.WorkloadFederation.SearchTrusts()`: **Added**
* `ConductoroneApi.WorkloadFederation.TestCel()`: **Added**
* `ConductoroneApi.WorkloadFederation.TestToken()`: **Added**
* `ConductoroneApi.WorkloadFederation.UpdateProvider()`: **Added**
* `ConductoroneApi.WorkloadFederation.UpdateTrust()`: **Added**
* `ConductoroneApi.Principal.AddBinding()`: **Added**
* `ConductoroneApi.Principal.Create()`: **Added**
* `ConductoroneApi.Principal.CreateCredential()`: **Added**
* `ConductoroneApi.Principal.Delete()`: **Added**
* `ConductoroneApi.Principal.DeleteBinding()`: **Added**
* `ConductoroneApi.Principal.Get()`: **Added**
* `ConductoroneApi.Principal.GetCredential()`: **Added**
* `ConductoroneApi.Principal.List()`: **Added**
* `ConductoroneApi.Principal.ListBindings()`: **Added**
* `ConductoroneApi.Principal.ListCredentials()`: **Added**
* `ConductoroneApi.Principal.RevokeCredential()`: **Added**
* `ConductoroneApi.Principal.Update()`: **Added**
* `ConductoroneApi.Principal.UpdateCredential()`: **Added**
* `ConductoroneApi.Contacts.GetContacts()`: **Added**
* `ConductoroneApi.Contacts.UpdateContacts()`: **Added**
* `ConductoroneApi.TenantEmailProvider.Get()`: **Added**
* `ConductoroneApi.TenantEmailProvider.GetEmailCapabilities()`: **Added**
* `ConductoroneApi.TenantEmailProvider.SearchAuditEvents()`: **Added**
* `ConductoroneApi.TenantEmailProvider.Test()`: **Added**
* `ConductoroneApi.TenantEmailProvider.Update()`: **Added**
* `ConductoroneApi.OrgNotificationSettings.Get()`: **Added**
* `ConductoroneApi.OrgNotificationSettings.Update()`: **Added**
* `ConductoroneApi.UserNotificationSettings.Get()`: **Added**
* `ConductoroneApi.UserNotificationSettings.Update()`: **Added**
* `ConductoroneApi.OnboardingSettings.Get()`: **Added**
* `ConductoroneApi.OnboardingSettings.Update()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Create()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Delete()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Get()`: **Added**
* `ConductoroneApi.SsfReceiverStream.GetStats()`: **Added**
* `ConductoroneApi.SsfReceiverStream.List()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Revoke()`: **Added**
* `ConductoroneApi.SsfReceiverStream.Update()`: **Added**
* `ConductoroneApi.SsfReceiverEvent.List()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.SearchEntitlementOwners()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.SearchUserOwners()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Get()`: **Added**
* `ConductoroneApi.LocalUserInvitation.Create()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Update()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.List()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Get()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.Set()`: **Added**
* `ConductoroneApi.AppOwnersV2.SearchEntitlementOwners()`: **Added**
* `ConductoroneApi.AppOwnersV2.SearchUserOwners()`: **Added**
* `ConductoroneApi.AppOwnersV2.Set()`: **Added**
* `ConductoroneApi.Functions.GetFunctionSecretEncryptionKey()`: **Removed** (Breaking ⚠️)
* `ConductoroneApi.Apps.Create()`: 
  * `request.Request` **Changed**
    - `AppEntitlementOwnerRefs` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
  * `response.App` **Changed**
    - `AccessModel` **Added**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Apps.Get()`: `response.App` **Changed**
    - `AccessModel` **Added**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Apps.List()`: `response.List[]` **Changed**
    - `AccessModel` **Added**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Apps.Update()`: 
  * `request.Request.UpdateAppRequest.App` **Changed**
    - `AccessModel` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
  * `response.App` **Changed**
    - `AccessModel` **Added**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.Connector.Create()`: `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.Connector.CreateDelegated()`: 
  *  `request.Request.ConnectorServiceCreateDelegatedRequest.AppEntitlementOwnerRefs` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.Connector.Get()`: `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.Connector.List()`: `response.List[].Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.Connector.Update()`: 
  * `request.Request.ConnectorServiceUpdateRequest.Connector` **Changed**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `ConfigUpdatedAt` **Added**
    - `ConnectorApiVersion` **Added**
    - `ConnectorSyncCronSchedule` **Added**
    - `ParallelSyncWorkerCount` **Added**
* `ConductoroneApi.LocalUserInvitation.Search()`: **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.Request.AppAccessRequestDefaults.RequestSchemaId` **Added**
  *  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.RequestSchemaId` **Added**
* `ConductoroneApi.AppEntitlements.AddManuallyManagedMembers()`:  `response.BulkActionId` **Added**
* `ConductoroneApi.AppEntitlements.Create()`: 
  * `request.Request.CreateAppEntitlementRequest` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.LocalDirectoryConfig.Delete()`: **Added**
* `ConductoroneApi.LocalDirectoryConfig.Create()`: **Added**
* `ConductoroneApi.AppEntitlements.Get()`: `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.GetAutomation()`:  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
* `ConductoroneApi.AppEntitlements.List()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListAutomationExclusions()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppResource()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.ListUsers()`:  `response.List[].OriginatingTicketId` **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  * `request.Request.UpdateAppEntitlementRequest.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
  * `response.AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlements.UpdateAutomation()`:  `response.AppEntitlementAutomation.ManagedByRequestCatalogId` **Added**
* `ConductoroneApi.AppEntitlementSearch.Search()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: 
  * `request.Request` **Changed**
    - `EntitlementSlugs` **Added**
    - `Purpose` **Added**
  * `response.List[]` **Changed**
    - `AppEntitlementUserView.OriginatingTicketId` **Added**
    - `AppEntitlementView.ActorObjectPermissions` **Added**
    - `AppEntitlementView.AppEntitlement.ExternalId` **Added**
    - `AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlementView.AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.AppEntitlementOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.AppOwners.List()`:  `response.List[].Origin` **Added**
* `ConductoroneApi.AppResource.CreateManuallyManagedAppResource()`: `response.AppResource` **Changed**
    - `AccessConfigId` **Added**
    - `ExternalId` **Added**
    - `Profile` **Added**
* `ConductoroneApi.AppResource.Get()`: `response.AppResourceView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.ExternalId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResource.List()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.ExternalId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResource.Update()`: 
  *  `request.Request.AppResourceServiceUpdateRequest.AppResource.AccessConfigId` **Added**
  * `response.AppResourceView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.ExternalId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppResourceOwners.List()`: `response` **Changed**
    - `ImmutableUserIds` **Added**
    - `List[].Origin` **Added**
* `ConductoroneApi.Hooks.Update()`: **Added**
* `ConductoroneApi.Hooks.List()`: **Added**
* `ConductoroneApi.Hooks.Get()`: **Added**
* `ConductoroneApi.Hooks.Delete()`: **Added**
* `ConductoroneApi.Hooks.Create()`: **Added**
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: **Added**
* `ConductoroneApi.Functions.Test()`: **Added**
* `ConductoroneApi.RequestCatalogManagement.Create()`: 
  *  `request.Request.GrantPolicyId` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ExternalId` **Added**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.RequestCatalogManagement.CreateBundleAutomation()`: 
  *  `request.Request.CreateBundleAutomationRequest.BundleAutomationRuleCel` **Added**
  * `response` **Changed**
    - `BundleAutomationCircuitBreaker.State.Enum(circuitBreakerStateSupportDisabled)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.RequestCatalogManagement.Get()`: `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ExternalId` **Added**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.RequestCatalogManagement.GetBundleAutomation()`: `response` **Changed**
    - `BundleAutomationCircuitBreaker.State.Enum(circuitBreakerStateSupportDisabled)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.RequestCatalogManagement.List()`: `response.List[].RequestCatalog` **Changed**
    - `AccessEntitlements[].ExternalId` **Added**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.RequestCatalogManagement.SetBundleAutomation()`: 
  *  `request.Request.SetBundleAutomationRequest.BundleAutomationRuleCel` **Added**
  * `response` **Changed**
    - `BundleAutomationCircuitBreaker.State.Enum(circuitBreakerStateSupportDisabled)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCel` **Added**
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  * `request.Request.RequestCatalogManagementServiceUpdateRequest.RequestCatalog` **Changed**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
  * `response.RequestCatalogView.RequestCatalog` **Changed**
    - `AccessEntitlements[].ExternalId` **Added**
    - `AccessEntitlements[].ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AccessEntitlements[].Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
    - `GrantPolicyId` **Added**
* `ConductoroneApi.Functions.GetLockFile()`: **Added**
* `ConductoroneApi.Directory.Create()`: 
  *  `request.Request.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.Get()`:  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.List()`:  `response.List[].Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Directory.Update()`: 
  *  `request.Request.DirectoryServiceUpdateRequest.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `ConductoroneApi.Functions.GetCommitContent()`: **Added**
* `ConductoroneApi.Functions.CreateInitialCommit()`: **Added**
* `ConductoroneApi.Functions.Invoke()`: 
  *  `request.Request.FunctionsServiceInvokeRequest.VfsId` **Added**
* `ConductoroneApi.Functions.CreateFinalCommit()`: **Added**
* `ConductoroneApi.FindingSearch.Search()`: **Added**
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: **Added**
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: **Added**
* `ConductoroneApi.FindingRoutingRule.DeleteFindingRoutingRule()`: **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  * `request.Request` **Changed**
    - `FieldGroups` **Added**
    - `FieldRelationships` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
    - `JustificationVisibility` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
  * `response.RequestSchema.Form` **Changed**
    - `FieldRelationships[].DependentOn` **Added**
    - `Fields[].Required` **Added**
    - `Fields[].StringField.PickerField.C1UserFilter` **Added**
    - `Fields[].StringField.TextField.Suffix` **Added**
    - `Fields[].StringMapField` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`: `response.List[]` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppResource.AccessConfigId` **Added**
    - `AppResource.ExternalId` **Added**
    - `AppResource.Profile` **Added**
* `ConductoroneApi.AppSearch.Search()`: `response.List[]` **Changed**
    - `AccessModel` **Added**
    - `AppOwners[].Origin` **Added**
    - `AppUserMapper` **Added**
    - `EnableConnectorSourcedOwnership` **Added**
    - `IdentityMatching.Enum(appUserIdentityMatchingCustom)` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: **Added**
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`: `response.List[].AppEntitlementView` **Changed**
    - `ActorObjectPermissions` **Added**
    - `AppEntitlement.ExternalId` **Added**
    - `AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `AppEntitlement.Purpose.Enum(appEntitlementPurposeValueOwnership)` **Added**
* `ConductoroneApi.StepUpProvider.Create()`: 
  *  `request.Request.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Get()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.List()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Search()`:  `response.List[].StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.Update()`: 
  *  `request.Request.UpdateStepUpProviderRequest.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
  *  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.StepUpProvider.UpdateSecret()`:  `response.StepUpProvider.StepUpMicrosoftSettings.ValidationMode` **Added**
* `ConductoroneApi.Finding.GetFinding()`: **Added**
* `ConductoroneApi.UserSearch.Search()`: 
  * `request.Request` **Changed**
    - `DelegateStatus` **Added**
    - `DelegatedUserIds` **Added**
    - `ExcludeOrigins` **Added**
    - `IsDelegate` **Added**
    - `Origins` **Added**
  *  `response.List[].User.Origin` **Added**
* `ConductoroneApi.WebhooksSearch.Search()`:  `response.List[].CallbackTimeout` **Added**
* `ConductoroneApi.SessionSettings.Get()`: `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
    - `ExternalClientsEnabled` **Added**
* `ConductoroneApi.SessionSettings.Update()`: 
  * `request.Request.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
    - `ExternalClientsEnabled` **Added**
  * `response.SessionSettings` **Changed**
    - `CidrRestriction1` **Changed**
    - `CidrRestriction2` **Changed**
    - `CidrRestriction3` **Changed**
    - `CidrRestriction4` **Changed**
    - `CidrRestriction5` **Added**
    - `ClientIdApprovalRequestPolicyId` **Added**
    - `ClientIdMetadataDocumentPolicy` **Added**
    - `ExternalClientsEnabled` **Added**
* `ConductoroneApi.SystemLog.ListEvents()`: 
  *  `request.Request.UntilEventUid` **Added**
* `ConductoroneApi.Finding.CreateFindingTask()`: **Added**
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Update()`: **Added**
* `ConductoroneApi.TenantAuthConfig.List()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Get()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Delete()`: **Added**
* `ConductoroneApi.TenantAuthConfig.Create()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: **Added**
* `ConductoroneApi.A2Ui.SubmitAction()`: **Added**
* `ConductoroneApi.A2Ui.ListSurfaces()`: **Added**
* `ConductoroneApi.A2Ui.ListSurfaceFeedback()`: **Added**
* `ConductoroneApi.A2Ui.CreateSurfaceFeedback()`: **Added**
* `ConductoroneApi.User.Get()`:  `response.UserView.User.Origin` **Added**
* `ConductoroneApi.User.List()`:  `response.List[].User.Origin` **Added**
* `ConductoroneApi.Webhooks.Create()`: 
  *  `request.Request.CallbackTimeout` **Added**
  *  `response.Webhook.CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.Get()`:  `response.Webhook.CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.List()`:  `response.List[].CallbackTimeout` **Added**
* `ConductoroneApi.Webhooks.Update()`: 
  *  `request.Request.WebhooksServiceUpdateRequest.Webhook.CallbackTimeout` **Added**
  *  `response.Webhook.CallbackTimeout` **Added**
