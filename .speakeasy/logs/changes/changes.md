## Go SDK Changes:
* `ConductoroneApi.AttributeSearch.SearchAttributeValues()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.LocalUserInvitation.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `Invitation` **Added**
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.SearchUserOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.SearchEntitlementOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.GetUserOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppOwnerUser` **Added**
    - `AppOwnerUser` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.GetEntitlementOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppOwnerEntitlement` **Added**
    - `AppOwnerEntitlement` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.DeleteUserOwner()`: `request.Request` **Changed** (Breaking ⚠️)
    - `DeleteAppUserOwnerRequest` **Added**
    - `DeleteUserOwnerRequest` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.DeleteEntitlementOwner()`: `request.Request` **Changed** (Breaking ⚠️)
    - `DeleteAppEntitlementOwnerRequest` **Added**
    - `DeleteEntitlementOwnerRequest` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.CreateUserOwner()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `CreateAppUserOwnerRequest` **Added**
    - `CreateUserOwnerRequest` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AppOwnerUser` **Added**
    - `AppOwnerUser` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppOwnersV2.CreateEntitlementOwner()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `CreateAppEntitlementOwnerRequest` **Added**
    - `CreateEntitlementOwnerRequest` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AppOwnerEntitlement` **Added**
    - `AppOwnerEntitlement` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementOwnersV2.Set()`: `request.Request` **Changed** (Breaking ⚠️)
    - `SetAppEntitlementOwnersRequestV2` **Added**
    - `SetAppEntitlementOwnersV2Request` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementOwnersV2.SearchUserOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppId` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `EntitlementId` **Added**
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementOwnersV2.SearchEntitlementOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `AppId` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `EntitlementId` **Added**
* `ConductoroneApi.ConnectorOwnersV2.SearchUserOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppId` **Added**
    - `ConnectorId` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.ConnectorOwnersV2.SearchEntitlementOwners()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `AppId` **Added**
    - `ConnectorId` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Webhooks.Update()`: 
  * `request.Request.WebhooksServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `Webhook` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `Webhook` **Added**
* `ConductoroneApi.Webhooks.Test()`: `response` **Changed** (Breaking ⚠️)
    - `WebhookInstance` **Removed** (Breaking ⚠️)
    - `Webhook` **Added**
* `ConductoroneApi.Webhooks.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CallbackTimeout` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Webhooks.Get()`: `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `Webhook` **Added**
* `ConductoroneApi.Webhooks.Create()`: 
  *  `request.Request.CallbackTimeout` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `Webhook` **Added**
* `ConductoroneApi.Vault.Update()`: 
  * `request.Request.VaultServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Vault` **Added**
    - `Vault` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Vault` **Added**
    - `Vault` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Vault.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Vault` **Added**
    - `Vault` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Vault.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `GroupAuthzVault` **Added**
    - `GroupAuthzVault` **Removed** (Breaking ⚠️)
    - `MagicVault` **Added**
    - `MagicVault` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Vault` **Added**
    - `Vault` **Removed** (Breaking ⚠️)
* `ConductoroneApi.User.SetExpiringUserDelegationBindingByAdmin()`: 
  * `request.Request.SetExpiringUserDelegationBindingByAdminRequest` **Changed**
    - `DelegationExpireAt` **Changed**
    - `DelegationStartAt` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `ExpiringUserDelegationBinding` **Removed** (Breaking ⚠️)
    - `Item` **Added**
* `ConductoroneApi.User.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
    - `UserId` **Added**
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementUserBinding.RemoveGrantDuration()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBinding` **Removed** (Breaking ⚠️)
    - `Binding` **Added**
* `ConductoroneApi.User.Get()`: `response` **Changed** (Breaking ⚠️)
    - `UserView` **Added**
    - `UserView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.UpdateRequestData()`: 
  * `request.Request.TaskActionsServiceUpdateRequestDataRequest` **Changed** (Breaking ⚠️)
    - `Data` **Changed**
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.UpdateGrantDuration()`: 
  * `request.Request.TaskActionsServiceUpdateGrantDurationRequest` **Changed** (Breaking ⚠️)
    - `Duration` **Changed**
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.SkipStep()`: 
  * `request.Request.TaskActionsServiceSkipStepRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Restart()`: 
  * `request.Request.TaskActionsServiceRestartRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Reassign()`: 
  * `request.Request.TaskActionsServiceReassignRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.ProcessNow()`: 
  * `request.Request.TaskActionsServiceProcessNowRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.HardReset()`: 
  * `request.Request.TaskActionsServiceHardResetRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.EscalateToEmergencyAccess()`: 
  * `request.Request.TaskActionsServiceEscalateToEmergencyAccessRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Deny()`: 
  * `request.Request.TaskActionsServiceDenyRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Comment()`: 
  * `request.Request.TaskActionsServiceCommentRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Close()`: 
  * `request.Request.TaskActionsServiceCloseRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.ApproveWithStepUp()`: 
  * `request.Request.TaskActionsServiceApproveWithStepUpRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskActions.Approve()`: 
  * `request.Request.TaskActionsServiceApproveRequest` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Task.Get()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Task.CreateRevokeTask()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Task.CreateOffboardingTask()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Task.CreateGrantTask()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `GrantDuration` **Changed**
    - `RequestData` **Changed**
    - `Source` **Added**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
    - `TaskGrantSource` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Added**
    - `TaskView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskAudit.List()`: 
  * `request.Request` **Changed**
    - `CommentsOnly` **Added**
    - `NewestFirst` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AccessRequestOutcome` **Added**
    - `AccountLifecycleActionCreated` **Added**
    - `AccountLifecycleActionFailed` **Added**
    - `ActionInstanceCreated` **Added**
    - `ActionInstanceFailed` **Added**
    - `ActionInstanceSucceeded` **Added**
    - `ActionResult` **Added**
    - `ActionSubmitted` **Added**
    - `ApprovalAutoAcceptedByPolicy` **Added**
    - `ApprovalAutoRejectedByPolicy` **Added**
    - `ApprovalInstanceChange` **Added**
    - `ApprovalReassigned` **Added**
    - `ApprovedAutomatically` **Added**
    - `BulkActionError` **Added**
    - `CertifyOutcome` **Added**
    - `Comment` **Added**
    - `ConditionalPolicyExecutionResult` **Added**
    - `ConnectorActionsEnd` **Added**
    - `ConnectorActionsStart` **Added**
    - `CreatedReplacementExtensionGrantTask` **Added**
    - `Created` **Changed** (Breaking ⚠️)
    - `ExpressionPolicyStepError` **Added**
    - `ExternalTicketCreated` **Added**
    - `ExternalTicketError` **Added**
    - `ExternalTicketProvisionStepResolved` **Added**
    - `ExternalTicketTriggered` **Added**
    - `FormInstanceChange` **Added**
    - `GrantDurationUpdated` **Added**
    - `GrantOutcome` **Added**
    - `HardReset` **Added**
    - `Metadata` **Added**
    - `PolicyChanged` **Added**
    - `PolicyEvaluationStep` **Added**
    - `ProvisionCancelled` **Added**
    - `ProvisionError` **Added**
    - `ProvisionReassigned` **Added**
    - `ReassignedToDelegate` **Added**
    - `ReassignmentFallbackToAdmin` **Added**
    - `ReassignmentListError` **Added**
    - `RequestDefaultsApplied` **Added**
    - `RevokeOutcome` **Added**
    - `SlaEscalation` **Added**
    - `StateChange` **Added**
    - `StepSkipped` **Added**
    - `StepUpApproval` **Added**
    - `TaskAuditAccessRequestOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditAccountLifecycleActionCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditAccountLifecycleActionFailed` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceFailed` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceSucceeded` **Removed** (Breaking ⚠️)
    - `TaskAuditActionSubmitted` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalAutoAcceptedByPolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalAutoRejectedByPolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalHappenedAutomatically` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange` **Removed** (Breaking ⚠️)
    - `TaskAuditBulkActionError` **Removed** (Breaking ⚠️)
    - `TaskAuditCertifyOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditComment` **Removed** (Breaking ⚠️)
    - `TaskAuditConditionalPolicyExecutionResult` **Removed** (Breaking ⚠️)
    - `TaskAuditConnectorActionResult` **Removed** (Breaking ⚠️)
    - `TaskAuditCreatedReplacementExtensionGrantTask` **Removed** (Breaking ⚠️)
    - `TaskAuditEscalateToEmergencyAccess` **Removed** (Breaking ⚠️)
    - `TaskAuditExpressionPolicyStepError` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketError` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketProvisionStepResolved` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketTriggered` **Removed** (Breaking ⚠️)
    - `TaskAuditFinishedConnectorActions` **Removed** (Breaking ⚠️)
    - `TaskAuditFormInstanceChange` **Removed** (Breaking ⚠️)
    - `TaskAuditGrantDurationUpdated` **Removed** (Breaking ⚠️)
    - `TaskAuditGrantOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditHardReset` **Removed** (Breaking ⚠️)
    - `TaskAuditMetaData` **Removed** (Breaking ⚠️)
    - `TaskAuditNewTaskCreatedFrom` **Removed** (Breaking ⚠️)
    - `TaskAuditNewTask` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyApprovalReassigned` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyChanged` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyEvaluationStep` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionCancelled` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionError` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionReassigned` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignedToDelegate` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignmentFallbackToAdmin` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignmentListError` **Removed** (Breaking ⚠️)
    - `TaskAuditRestart` **Removed** (Breaking ⚠️)
    - `TaskAuditRevokeOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditSlaEscalation` **Removed** (Breaking ⚠️)
    - `TaskAuditStartedConnectorActions` **Removed** (Breaking ⚠️)
    - `TaskAuditStateChange` **Removed** (Breaking ⚠️)
    - `TaskAuditStepSkipped` **Removed** (Breaking ⚠️)
    - `TaskAuditStepUpApproval` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepTimedOut` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepWaiting` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepTimedOut` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepUntilTime` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepWaiting` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalAttempt` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalBadResponse` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalFatalError` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalTriggered` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookAttempt` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookTriggered` **Removed** (Breaking ⚠️)
    - `TaskCreatedFrom` **Added**
    - `TaskCreated` **Added**
    - `TaskEscalated` **Added**
    - `TaskRestarted` **Added**
    - `WaitStepAnalysisSuccess` **Added**
    - `WaitStepAnalysisTimedOut` **Added**
    - `WaitStepAnalysisWaiting` **Added**
    - `WaitStepSuccess` **Added**
    - `WaitStepTimedOut` **Added**
    - `WaitStepUntilTime` **Added**
    - `WaitStepWaiting` **Added**
    - `WebhookApprovalAttempt` **Added**
    - `WebhookApprovalBadResponse` **Added**
    - `WebhookApprovalFatalError` **Added**
    - `WebhookApprovalSuccess` **Added**
    - `WebhookApprovalTriggered` **Added**
    - `WebhookAttempt` **Added**
    - `WebhookSuccess` **Added**
    - `WebhookTriggered` **Added**
* `ConductoroneApi.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant()`: `response.Bindings[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `DeprovisionAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Export.Update()`: 
  * `request.Request.ExportServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Exporter` **Added**
    - `Exporter` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Added**
    - `Exporter` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Export.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Datasource` **Added**
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Export.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Added**
    - `Exporter` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Export.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Datasource` **Added**
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Added**
    - `Exporter` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverEvent.List()`:  `response.List[].ReceivedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverStream.Update()`: 
  * `request.Request.SsfReceiverStreamServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `SsfReceiverStream` **Added**
    - `SsfReceiverStream` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `SsfReceiverStream` **Added**
    - `SsfReceiverStream` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverStream.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `LastErrorAt` **Changed** (Breaking ⚠️)
    - `LastVerifiedAt` **Changed** (Breaking ⚠️)
    - `OutboundAuthBearer` **Added**
    - `OutboundAuthOauth2` **Added**
    - `PollInterval` **Changed** (Breaking ⚠️)
    - `SsfOutboundAuthBearer` **Removed** (Breaking ⚠️)
    - `SsfOutboundAuthOAuth2` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverStream.GetStats()`: `response` **Changed** (Breaking ⚠️)
    - `SsfReceiverStreamStats` **Removed** (Breaking ⚠️)
    - `Stats` **Added**
* `ConductoroneApi.SsfReceiverStream.Get()`: `response` **Changed** (Breaking ⚠️)
    - `SsfReceiverStream` **Added**
    - `SsfReceiverStream` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverStream.Create()`: 
  *  `request.Request.PollInterval` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `SsfReceiverStream` **Added**
    - `SsfReceiverStream` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SessionSettings.Update()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Added**
    - `SessionSettings` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Added**
    - `SessionSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SessionSettings.TestSourceIp()`: `response` **Changed** (Breaking ⚠️)
    - `Details` **Added**
    - `Status` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SessionSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Added**
    - `SessionSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.OnboardingSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `McpOnboardingGoal` **Added**
    - `McpOnboardingStatus` **Added**
    - `McpOnboardingTargets` **Added**
    - `OnboardingOrgContext` **Removed** (Breaking ⚠️)
    - `OrgContext` **Added**
* `ConductoroneApi.UserNotificationSettings.Update()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ChannelSettings` **Added**
    - `ChannelSettings` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `UserNotificationSettings` **Added**
    - `UserNotificationSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.UserNotificationSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `UserNotificationSettings` **Added**
    - `UserNotificationSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.OrgNotificationSettings.Update()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ChannelSettings` **Added**
    - `ChannelSettings` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `OrgNotificationSettings` **Added**
    - `OrgNotificationSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.OrgNotificationSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `OrgNotificationSettings` **Added**
    - `OrgNotificationSettings` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TenantEmailProvider.Update()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `EmailProvider` **Added**
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `EmailProvider` **Added**
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TenantEmailProvider.Get()`: `response` **Changed** (Breaking ⚠️)
    - `EmailProvider` **Added**
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.OrgDomain.Update()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.OrgDomain.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Contacts.UpdateContacts()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Contacts` **Added**
    - `Contacts` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Contacts` **Added**
    - `Contacts` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Contacts.GetContacts()`: `response` **Changed** (Breaking ⚠️)
    - `Contacts` **Added**
    - `Contacts` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AwsExternalIdSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AwsExternalId` **Added**
    - `AwsExternalId` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.UpdateCredential()`: 
  * `request.Request.ServicePrincipalServiceUpdateCredentialRequest` **Changed** (Breaking ⚠️)
    - `Credential` **Added**
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Credential` **Added**
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.Update()`: 
  * `request.Request.ServicePrincipalServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Added**
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Added**
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.ListCredentials()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `ExpiresAt` **Changed** (Breaking ⚠️)
    - `LastUsedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Principal.ListBindings()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `Subject` **Added**
  * `response.Bindings[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Principal.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.GetCredential()`: `response` **Changed** (Breaking ⚠️)
    - `Credential` **Added**
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.Get()`: `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Added**
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.DeleteBinding()`: `request.Request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `Subject` **Added**
* `ConductoroneApi.Principal.CreateCredential()`: 
  *  `request.Request.ServicePrincipalServiceCreateCredentialRequest.Expires` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `Credential` **Added**
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.Create()`: `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Added**
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Principal.AddBinding()`: `request.Request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `Subject` **Added**
* `ConductoroneApi.WorkloadFederation.UpdateTrust()`: 
  * `request.Request.WorkloadFederationServiceUpdateTrustRequest` **Changed** (Breaking ⚠️)
    - `Trust` **Added**
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Trust` **Added**
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.UpdateProvider()`: 
  * `request.Request.WorkloadFederationServiceUpdateProviderRequest` **Changed** (Breaking ⚠️)
    - `Provider` **Added**
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Provider` **Added**
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.TestToken()`: `response` **Changed** (Breaking ⚠️)
    - `AudienceValidation` **Added**
    - `CelEvaluation` **Added**
    - `CidrCheck` **Added**
    - `IssuerMatch` **Added**
    - `JwtDecode` **Added**
    - `SignatureValidation` **Added**
    - `SubjectValidation` **Added**
    - `TestTokenStepResult1` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult2` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult3` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult4` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult5` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult6` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult` **Removed** (Breaking ⚠️)
    - `TokenFreshness` **Added**
* `ConductoroneApi.WorkloadFederation.SearchTrusts()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.ListTrusts()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.ListProviders()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Oidc` **Added**
    - `Spiffe` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderSpiffe)` **Added**
* `ConductoroneApi.WorkloadFederation.GetTrust()`: `response` **Changed** (Breaking ⚠️)
    - `Trust` **Added**
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.GetProvider()`: `response` **Changed** (Breaking ⚠️)
    - `Provider` **Added**
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.CreateTrust()`: `response` **Changed** (Breaking ⚠️)
    - `Trust` **Added**
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WorkloadFederation.CreateProvider()`: 
  * `request.Request` **Changed**
    - `Oidc` **Added**
    - `Spiffe` **Added**
    - `WellKnownProvider.Enum(wellKnownWorkloadProviderSpiffe)` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Provider` **Added**
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.WebhooksSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CallbackTimeout` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.UserSearch.Search()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `UserExpandMask` **Removed** (Breaking ⚠️)
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
    - `UserId` **Added**
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TaskSearch.Search()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `CreatedAfter` **Changed**
    - `CreatedBefore` **Changed**
    - `ExpandMask` **Added**
    - `IncludeActedAfter` **Changed**
    - `OlderThanDuration` **Changed**
    - `OutcomeAfter` **Changed**
    - `OutcomeBefore` **Changed**
    - `TaskExpandMask` **Removed** (Breaking ⚠️)
    - `TaskTypes[].Action` **Added**
    - `TaskTypes[].Certify` **Added**
    - `TaskTypes[].Finding` **Added**
    - `TaskTypes[].Grant` **Added**
    - `TaskTypes[].Offboarding` **Added**
    - `TaskTypes[].Revoke` **Added**
    - `TaskTypes[].TaskTypeAction` **Removed** (Breaking ⚠️)
    - `TaskTypes[].TaskTypeCertify` **Removed** (Breaking ⚠️)
    - `TaskTypes[].TaskTypeFinding` **Removed** (Breaking ⚠️)
    - `TaskTypes[].TaskTypeGrant` **Removed** (Breaking ⚠️)
    - `TaskTypes[].TaskTypeOffboarding` **Removed** (Breaking ⚠️)
    - `TaskTypes[].TaskTypeRevoke` **Removed** (Breaking ⚠️)
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
    - `PrincipalResourcePath` **Added**
    - `Task` **Added**
    - `Task` **Removed** (Breaking ⚠️)
* `ConductoroneApi.ExportsSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Datasource` **Added**
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.StepUpTransaction.Search()`: 
  * `request.Request` **Changed**
    - `CreatedAfter` **Changed**
    - `CreatedBefore` **Changed**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `ApproveTask` **Added**
    - `Claims` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `ExpiresAt` **Changed** (Breaking ⚠️)
    - `TargetTask` **Removed** (Breaking ⚠️)
    - `TargetTest` **Removed** (Breaking ⚠️)
    - `Test` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.StepUpTransaction.Get()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpTransaction` **Removed** (Breaking ⚠️)
    - `Transaction` **Added**
* `ConductoroneApi.StepUpProvider.UpdateSecret()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Added**
    - `StepUpProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.StepUpProvider.Update()`: 
  * `request.Request.UpdateStepUpProviderRequest` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Added**
    - `StepUpProvider` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Added**
    - `StepUpProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.StepUpProvider.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `LastTestedAt` **Changed** (Breaking ⚠️)
    - `Microsoft` **Added**
    - `Oauth2` **Added**
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.StepUpProvider.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `LastTestedAt` **Changed** (Breaking ⚠️)
    - `Microsoft` **Added**
    - `Oauth2` **Added**
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.StepUpProvider.Get()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Added**
    - `StepUpProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.StepUpProvider.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Microsoft` **Added**
    - `Oauth2` **Added**
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Added**
    - `StepUpProvider` **Removed** (Breaking ⚠️)
* `ConductoroneApi.SsfReceiverEventSearch.Search()`:  `response.List[].ReceivedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PaperSecret.SetTextContent()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecret.SearchMySecrets()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AgeSuite` **Added**
    - `ContentExpiresAt` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PaperSecret.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecret.GetContent()`:  `response.CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PaperSecret.GetByShareCode()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecret.Get()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecret.CreateInternal()`: 
  * `request.Request` **Changed**
    - `ExpiresIn` **Changed**
    - `RequiredAgeSuite` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AgeSuite` **Added**
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecret.CreateExternal()`: 
  * `request.Request` **Changed**
    - `ExpiresIn` **Changed**
    - `RequiredAgeSuite` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AgeSuite` **Added**
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecretAdmin.Search()`: 
  * `request.Request` **Changed**
    - `CreatedAfter` **Changed**
    - `CreatedBefore` **Changed**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AgeSuite` **Added**
    - `ContentExpiresAt` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PaperSecretAdmin.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.PaperSecretAdmin.Get()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `Secret` **Added**
* `ConductoroneApi.RoleMiningManagementSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Entitlements[].RiskLevelValueId` **Added**
    - `LastGeneratedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogSearch.SearchEntitlements()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindings[].CreatedAt` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindings[].DeletedAt` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindings[].DeprovisionAt` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `Entitlement` **Added**
* `ConductoroneApi.PolicySearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PersonalClientSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ExpiresTime` **Changed** (Breaking ⚠️)
    - `LastUsedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.ExternalClientSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `LastUsedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.HooksSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `BuiltinPattern` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Filter` **Added**
    - `Function` **Added**
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.FunctionsSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ProvisionedConcurrency` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AutomationSearch.SearchAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].AccountLifecycleAction` **Added**
    - `AutomationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CallFunction` **Added**
    - `AutomationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorAction` **Added**
    - `AutomationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount` **Added**
    - `AutomationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateAccessReview` **Added**
    - `AutomationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasksV2` **Added**
    - `AutomationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasks` **Added**
    - `AutomationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].EvaluateExpressions` **Added**
    - `AutomationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GeneratePassword` **Added**
    - `AutomationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements` **Added**
    - `AutomationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RemoveFromDelegation` **Added**
    - `AutomationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RunAutomation` **Added**
    - `AutomationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendEmail` **Added**
    - `AutomationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendSlackMessage` **Added**
    - `AutomationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].TaskAction` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Added**
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UpdateUser` **Added**
    - `AutomationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].WaitForDuration` **Added**
    - `AutomationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Added**
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `CircuitBreaker` **Added**
    - `Context` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DisabledReasonCircuitBreaker` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AccessConflict` **Added**
    - `DraftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserCreated` **Added**
    - `DraftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserUpdated` **Added**
    - `DraftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeleted` **Added**
    - `DraftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantFound` **Added**
    - `DraftTriggers[].ScheduleAppUser` **Added**
    - `DraftTriggers[].ScheduleNoUser` **Added**
    - `DraftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Schedule` **Added**
    - `DraftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UsageBasedRevocation` **Added**
    - `DraftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserCreated` **Added**
    - `DraftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserProfileChange` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Webhook` **Added**
    - `LastExecutedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AutomationSearch.SearchAutomationTemplateVersions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationSteps[].AccountLifecycleAction` **Added**
    - `AutomationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CallFunction` **Added**
    - `AutomationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorAction` **Added**
    - `AutomationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount` **Added**
    - `AutomationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateAccessReview` **Added**
    - `AutomationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasksV2` **Added**
    - `AutomationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasks` **Added**
    - `AutomationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].EvaluateExpressions` **Added**
    - `AutomationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GeneratePassword` **Added**
    - `AutomationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements` **Added**
    - `AutomationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RemoveFromDelegation` **Added**
    - `AutomationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RunAutomation` **Added**
    - `AutomationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendEmail` **Added**
    - `AutomationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendSlackMessage` **Added**
    - `AutomationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].TaskAction` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Added**
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UpdateUser` **Added**
    - `AutomationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].WaitForDuration` **Added**
    - `AutomationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Added**
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `Triggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].AccessConflict` **Added**
    - `Triggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].AppUserCreated` **Added**
    - `Triggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].AppUserUpdated` **Added**
    - `Triggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].GrantDeleted` **Added**
    - `Triggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].GrantFound` **Added**
    - `Triggers[].ScheduleAppUser` **Added**
    - `Triggers[].ScheduleNoUser` **Added**
    - `Triggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `Triggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `Triggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].Schedule` **Added**
    - `Triggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].UsageBasedRevocation` **Added**
    - `Triggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].UserCreated` **Added**
    - `Triggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].UserProfileChange` **Added**
    - `Triggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Triggers[].Webhook` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AppOwners[].CreatedAt` **Changed** (Breaking ⚠️)
    - `AppOwners[].DeletedAt` **Changed** (Breaking ⚠️)
    - `AppOwners[].DepartmentSources[].Priority` **Added**
    - `AppOwners[].Profile` **Changed** (Breaking ⚠️)
    - `AppOwners[].UpdatedAt` **Changed** (Breaking ⚠️)
    - `AppUserMapper` **Added**
    - `AppUserMapper` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `RevokeGrantSources` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceSearch.SearchAppResources()`: 
  * `request.Request` **Changed**
    - `AgentStatuses` **Added**
    - `AppIds` **Added**
    - `CredentialTypes` **Added**
    - `Direction` **Added**
    - `ExcludeDeletedApps` **Added**
    - `NhiTypes` **Added**
    - `SecretAging` **Added**
    - `SortField` **Added**
    - `UnownedOnly` **Added**
    - `WithOpenFindings` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppResource` **Added**
    - `AppResource` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppResourceSearch.SearchAppResourceTypes()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AutomationExecutionSearch.SearchAutomationExecutions()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationExecutionExpandMask` **Removed** (Breaking ⚠️)
    - `ExecutionStepStates[].Enum(automationExecutionStatePausedByCircuitBreaker)` **Added**
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationExecution` **Added**
    - `AutomationExecution` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AutomationExecutionSearch.SearchAllAutomationExecutions()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AutomationExecutionExpandMask` **Removed** (Breaking ⚠️)
    - `ExecutionStates[].Enum(automationExecutionStatePausedByCircuitBreaker)` **Added**
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AutomationExecution` **Added**
    - `AutomationExecution` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RoleMiningManagement.UpdateSuggestionState()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementSuggestion` **Removed** (Breaking ⚠️)
    - `Suggestion` **Added**
* `ConductoroneApi.RoleMiningManagement.UpdateRoleMiningConfig()`: `response` **Changed** (Breaking ⚠️)
    - `Config` **Added**
    - `RoleMiningManagementConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RoleMiningManagement.SearchCohortUsers()`: 
  *  `request.Request.SearchCohortUsersRequest.SelectedEntitlements` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `List[].CreatedAt` **Changed** (Breaking ⚠️)
    - `List[].DeletedAt` **Changed** (Breaking ⚠️)
    - `List[].DepartmentSources[].Priority` **Added**
    - `List[].Profile` **Changed** (Breaking ⚠️)
    - `List[].UpdatedAt` **Changed** (Breaking ⚠️)
    - `UsersWithCoverage` **Added**
* `ConductoroneApi.RoleMiningManagement.ListSuggestions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Entitlements[].RiskLevelValueId` **Added**
    - `LastGeneratedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RoleMiningManagement.ListRuns()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CompletedAt` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RoleMiningManagement.GetSuggestion()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementSuggestion` **Removed** (Breaking ⚠️)
    - `Suggestion` **Added**
* `ConductoroneApi.RoleMiningManagement.GetRoleMiningConfig()`: `response` **Changed** (Breaking ⚠️)
    - `Config` **Added**
    - `RoleMiningManagementConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RoleMiningManagement.GetLatestRun()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementRun` **Removed** (Breaking ⚠️)
    - `Run` **Added**
* `ConductoroneApi.RequestSchema.Update()`: 
  * `request.Request.RequestSchemaServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Added**
    - `RequestSchema` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Added**
    - `RequestSchema` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestSchema.RemoveEntitlementBinding()`: `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `EntitlementRef` **Added**
* `ConductoroneApi.RequestSchema.Get()`: `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Added**
    - `RequestSchema` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestSchema.FindBindingForAppEntitlement()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `EntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `EntitlementRef` **Added**
* `ConductoroneApi.RequestSchema.CreateEntitlementBinding()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `EntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `EntitlementRef` **Added**
* `ConductoroneApi.RequestSchema.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `FieldRelationships[].AtLeastOne` **Added**
    - `FieldRelationships[].AtLeastOne` **Removed** (Breaking ⚠️)
    - `FieldRelationships[].DependentOn` **Added**
    - `FieldRelationships[].DependentOn` **Removed** (Breaking ⚠️)
    - `FieldRelationships[].MutuallyExclusive` **Added**
    - `FieldRelationships[].MutuallyExclusive` **Removed** (Breaking ⚠️)
    - `FieldRelationships[].RequiredTogether` **Added**
    - `FieldRelationships[].RequiredTogether` **Removed** (Breaking ⚠️)
    - `Fields[].AdminConfig` **Added**
    - `Fields[].AdminProviderConfig` **Removed** (Breaking ⚠️)
    - `Fields[].BoolField` **Added**
    - `Fields[].BoolField` **Removed** (Breaking ⚠️)
    - `Fields[].FileField` **Added**
    - `Fields[].FileField` **Removed** (Breaking ⚠️)
    - `Fields[].FormStringField` **Removed** (Breaking ⚠️)
    - `Fields[].FormStringMapField` **Removed** (Breaking ⚠️)
    - `Fields[].Int64Field` **Added**
    - `Fields[].Int64Field` **Removed** (Breaking ⚠️)
    - `Fields[].Oauth2Field` **Added**
    - `Fields[].Oauth2Field` **Removed** (Breaking ⚠️)
    - `Fields[].ReadOnly` **Added**
    - `Fields[].SharedConfig` **Added**
    - `Fields[].SharedProviderConfig` **Removed** (Breaking ⚠️)
    - `Fields[].StringField` **Added**
    - `Fields[].StringMapField` **Added**
    - `Fields[].StringSliceField` **Added**
    - `Fields[].StringSliceField` **Removed** (Breaking ⚠️)
    - `Fields[].UserConfig` **Added**
    - `Fields[].UserProviderConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Added**
    - `RequestSchema` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Policies.Update()`: 
  * `request.Request.UpdatePolicyRequest` **Changed** (Breaking ⚠️)
    - `Policy` **Added**
    - `Policy` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Policy` **Added**
    - `Policy` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Policies.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Policies.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Policy` **Added**
    - `Policy` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Policies.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Accept` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Action` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Approval` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Form` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Provision` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Reject` **Removed** (Breaking ⚠️)
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Added**
    - `PolicySteps.Map<PolicySteps>.Steps[].Wait` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Policy` **Added**
    - `Policy` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalUserInvitation.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AcceptedAt` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `ExpiresAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.Update()`: 
  * `request.Request.RequestCatalogManagementServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Catalog` **Added**
    - `ExpandMask` **Added**
    - `RequestCatalogExpandMask` **Removed** (Breaking ⚠️)
    - `RequestCatalog` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Added**
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalUserInvitation.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Invitation` **Added**
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalUserInvitation.Create()`: `response` **Changed** (Breaking ⚠️)
    - `Invitation` **Added**
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalDirectoryConfig.Update()`: 
  * `request.Request.LocalDirectoryConfigServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Added**
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Added**
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalDirectoryConfig.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `InvitationTtl` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.LocalDirectoryConfig.Get()`: `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Added**
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.LocalDirectoryConfig.Create()`: 
  *  `request.Request.InvitationTtl` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Added**
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Roles.Update()`: 
  * `request.Request.UpdateRoleRequest` **Changed** (Breaking ⚠️)
    - `Role` **Added**
    - `Role` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Role` **Added**
    - `Role` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Roles.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Roles.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Role` **Added**
    - `Role` **Removed** (Breaking ⚠️)
* `ConductoroneApi.PersonalClient.Update()`: 
  * `request.Request.PersonalClientServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Client` **Added**
    - `PersonalClient` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Client` **Added**
    - `PersonalClient` **Removed** (Breaking ⚠️)
* `ConductoroneApi.PersonalClient.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ExpiresTime` **Changed** (Breaking ⚠️)
    - `LastUsedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.PersonalClient.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Client` **Added**
    - `PersonalClient` **Removed** (Breaking ⚠️)
* `ConductoroneApi.PersonalClient.Create()`: 
  *  `request.Request.Expires` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `Client` **Added**
    - `PersonalClient` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Hooks.Update()`: 
  * `request.Request.HooksServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Hook` **Added**
    - `Hook` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Hook` **Added**
    - `Hook` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Hooks.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `BuiltinPattern` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Filter` **Added**
    - `Function` **Added**
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Hooks.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Hook` **Added**
    - `Hook` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Hooks.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `BuiltinPattern` **Added**
    - `Filter` **Added**
    - `Function` **Added**
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Hook` **Added**
    - `Hook` **Removed** (Breaking ⚠️)
* `ConductoroneApi.FunctionsInvocationSearch.Search()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Input` **Changed** (Breaking ⚠️)
    - `Output` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.FunctionsInvocation.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `Input` **Changed** (Breaking ⚠️)
    - `Output` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.FunctionsInvocation.Get()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionInvocation` **Removed** (Breaking ⚠️)
    - `Invocation` **Added**
* `ConductoroneApi.Functions.UpdateFunction()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Function` **Added**
    - `Function` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Function` **Added**
    - `Function` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Functions.Test()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionTestResult` **Removed** (Breaking ⚠️)
    - `Result` **Added**
* `ConductoroneApi.Functions.ListTags()`:  `response.Tags.Map<FunctionCommit>.CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementUserBinding.SearchGrantFeed()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `After` **Changed**
    - `AppEntitlementUserBindingExpandHistoryMask` **Removed** (Breaking ⚠️)
    - `Before` **Changed**
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingFeed` **Removed** (Breaking ⚠️)
    - `Feed` **Added**
* `ConductoroneApi.Functions.ListFunctions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `ProvisionedConcurrency` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Functions.ListCommits()`:  `response.List[].CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Functions.GetFunction()`: `response` **Changed** (Breaking ⚠️)
    - `Function` **Added**
    - `Function` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Functions.GetCommitContent()`: `response` **Changed** (Breaking ⚠️)
    - `Commit` **Added**
    - `FunctionCommit` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Functions.CreateFunction()`: `response` **Changed** (Breaking ⚠️)
    - `Commit` **Added**
    - `FunctionCommit` **Removed** (Breaking ⚠️)
    - `Function` **Added**
    - `Function` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Functions.CreateFinalCommit()`: `response` **Changed** (Breaking ⚠️)
    - `Commit` **Added**
    - `FunctionCommit` **Removed** (Breaking ⚠️)
* `ConductoroneApi.FindingSearch.Search()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppResourceIds` **Added**
    - `AppResourceTraitIds` **Added**
    - `AppResourceTypeIds` **Added**
    - `AppUserTypes` **Added**
    - `ConnectorIds` **Added**
    - `CustomSubTypes` **Added**
    - `DecoyIds` **Added**
    - `FindingTypes[]` **Changed** (Breaking ⚠️)
    - `IncludeUnassigned` **Added**
    - `NhiTypes` **Added**
    - `OwnerIdentityUserIds` **Added**
    - `Refs` **Added**
    - `ScopeToAppOwner` **Added**
    - `SourceKinds` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppResourceTarget` **Added**
    - `AppUserTarget` **Added**
    - `AppUserTarget` **Removed** (Breaking ⚠️)
    - `AssignedOwner` **Added**
    - `ComputedOwner` **Added**
    - `ConnectorAnomalyDetectionDisabled` **Added**
    - `ConnectorTarget` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `CustomSubType` **Added**
    - `Custom` **Added**
    - `DecoyCredentialUsed` **Added**
    - `DecoyTarget` **Added**
    - `DedupKeyParts` **Added**
    - `Description` **Added**
    - `FindingOwnerRef1` **Removed** (Breaking ⚠️)
    - `FindingOwnerRef` **Removed** (Breaking ⚠️)
    - `FindingRiskScore` **Removed** (Breaking ⚠️)
    - `FirstObservedAt` **Changed** (Breaking ⚠️)
    - `IdentityUserTarget` **Added**
    - `IdentityUserTarget` **Removed** (Breaking ⚠️)
    - `LastAppearedAt` **Added**
    - `LastObservedAt` **Changed** (Breaking ⚠️)
    - `NhiUnowned` **Added**
    - `ResolvedAt` **Changed** (Breaking ⚠️)
    - `RiskAcceptanceExpiresAt` **Changed** (Breaking ⚠️)
    - `RiskScore` **Added**
    - `ServiceAccountMisclassificationEvidence` **Added**
    - `ServiceAccountMisclassificationEvidence` **Removed** (Breaking ⚠️)
    - `ServiceAccountMisclassificationType` **Removed** (Breaking ⚠️)
    - `ServiceAccountMisclassification` **Added**
    - `ServiceAccountUnowned` **Added**
    - `SimilarUsernameMatchEvidence` **Added**
    - `SimilarUsernameMatchEvidence` **Removed** (Breaking ⚠️)
    - `SimilarUsernameMatchType` **Removed** (Breaking ⚠️)
    - `SimilarUsernameMatch` **Added**
    - `SnoozeUntil` **Changed** (Breaking ⚠️)
    - `SourceKind` **Added**
    - `TenantTarget` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.Request.UpdateFindingRoutingRuleRequest` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `RoutingRule` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `RoutingRule` **Added**
* `ConductoroneApi.FindingRoutingRule.ListFindingRoutingRules()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Action` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `FindingRoutingRuleAction` **Removed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.FindingRoutingRule.GetFindingRoutingRule()`: `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `RoutingRule` **Added**
* `ConductoroneApi.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `RoutingRule` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `RoutingRule` **Added**
* `ConductoroneApi.Finding.UpdateFindingState()`: 
  * `request.Request.UpdateFindingStateRequest` **Changed** (Breaking ⚠️)
    - `AcceptRiskAction` **Removed** (Breaking ⚠️)
    - `AcceptRisk` **Added**
    - `ReopenAction` **Removed** (Breaking ⚠️)
    - `Reopen` **Added**
    - `ResolveAction` **Removed** (Breaking ⚠️)
    - `Resolve` **Added**
    - `SnoozeAction` **Removed** (Breaking ⚠️)
    - `Snooze` **Added**
    - `SuppressStateAction` **Removed** (Breaking ⚠️)
    - `Suppress` **Added**
    - `UnsuppressAction` **Removed** (Breaking ⚠️)
    - `Unsuppress` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Finding` **Added**
    - `Finding` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Finding.GetFinding()`: `response` **Changed** (Breaking ⚠️)
    - `Finding` **Added**
    - `Finding` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Finding.CreateFindingTask()`: `response` **Changed** (Breaking ⚠️)
    - `Finding` **Added**
    - `Finding` **Removed** (Breaking ⚠️)
* `ConductoroneApi.A2Ui.CreateSurfaceFeedback()`: `response` **Changed** (Breaking ⚠️)
    - `A2UiSurfaceFeedback` **Removed** (Breaking ⚠️)
    - `Feedback` **Added**
* `ConductoroneApi.A2Ui.ListSurfaceFeedback()`:  `response.Feedback[].CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.A2Ui.ListSurfaces()`: `response.Surfaces[]` **Changed** (Breaking ⚠️)
    - `Components[].ButtonComponent` **Removed** (Breaking ⚠️)
    - `Components[].Button` **Added**
    - `Components[].C1Chart` **Added**
    - `Components[].C1CodeBlockComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1CodeBlock` **Added**
    - `Components[].C1ConnectorConfigFormComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1ConnectorConfigForm` **Added**
    - `Components[].C1ConnectorSyncDetailComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1ConnectorSyncDetail` **Added**
    - `Components[].C1ConnectorSyncProgressComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1ConnectorSyncProgress` **Added**
    - `Components[].C1DurationPickerComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1DurationPicker` **Added**
    - `Components[].C1MsTeamsNotificationsComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1MsTeamsNotifications` **Added**
    - `Components[].C1OnboardingPlanComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1OnboardingPlan` **Added**
    - `Components[].C1OnboardingWelcomeComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1OnboardingWelcome` **Added**
    - `Components[].C1ResourcePickerComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1ResourcePicker` **Added**
    - `Components[].C1SlackNotificationsComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1SlackNotifications` **Added**
    - `Components[].C1StatusIndicatorComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1StatusIndicator` **Added**
    - `Components[].C1TodoListComponent` **Removed** (Breaking ⚠️)
    - `Components[].C1TodoList` **Added**
    - `Components[].CardComponent` **Removed** (Breaking ⚠️)
    - `Components[].Card` **Added**
    - `Components[].CheckBoxComponent` **Removed** (Breaking ⚠️)
    - `Components[].CheckBox` **Added**
    - `Components[].ChoicePickerComponent` **Removed** (Breaking ⚠️)
    - `Components[].ChoicePicker` **Added**
    - `Components[].ColumnComponent` **Removed** (Breaking ⚠️)
    - `Components[].Column` **Added**
    - `Components[].DateTimeInputComponent` **Removed** (Breaking ⚠️)
    - `Components[].DateTimeInput` **Added**
    - `Components[].DividerComponent` **Removed** (Breaking ⚠️)
    - `Components[].Divider` **Added**
    - `Components[].ProgressBarComponent` **Removed** (Breaking ⚠️)
    - `Components[].ProgressBar` **Added**
    - `Components[].RowComponent` **Removed** (Breaking ⚠️)
    - `Components[].Row` **Added**
    - `Components[].SliderComponent` **Removed** (Breaking ⚠️)
    - `Components[].Slider` **Added**
    - `Components[].TextComponent` **Removed** (Breaking ⚠️)
    - `Components[].TextFieldComponent` **Removed** (Breaking ⚠️)
    - `Components[].TextField` **Added**
    - `Components[].Text` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `Role` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Finding.BulkUpdateFindingState()`: `request.Request` **Changed** (Breaking ⚠️)
    - `AcceptRisk` **Added**
    - `AssignOwner` **Added**
    - `BulkAcceptRiskAction` **Removed** (Breaking ⚠️)
    - `BulkAssignOwnerAction` **Removed** (Breaking ⚠️)
    - `BulkReopenAction` **Removed** (Breaking ⚠️)
    - `BulkSnoozeAction` **Removed** (Breaking ⚠️)
    - `BulkSuppressAction` **Removed** (Breaking ⚠️)
    - `BulkUnsuppressAction` **Removed** (Breaking ⚠️)
    - `FindingSearchRequest` **Removed** (Breaking ⚠️)
    - `Reopen` **Added**
    - `SearchRequest` **Added**
    - `Snooze` **Added**
    - `Suppress` **Added**
    - `Unsuppress` **Added**
* `ConductoroneApi.AccessReview.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessReviewExpandMask` **Removed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `CompletionDate` **Changed**
    - `ExpandMask` **Added**
    - `NotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `ScopeType.Enum(accessReviewScopeTypeByUsers)` **Added**
    - `ScopeV2` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `AccessReview` **Added**
* `ConductoroneApi.AccessReview.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `AccessReview` **Added**
* `ConductoroneApi.AccessReview.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AccessReview` **Added**
    - `AccessReview` **Removed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AccessReview.Update()`: 
  * `request.Request.AccessReviewServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewExpandMask` **Removed** (Breaking ⚠️)
    - `AccessReview` **Added**
    - `AccessReview` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `AccessReview` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `List[].AccessReviewEntitlement` **Added**
    - `List[].AccessReviewSetupEntitlement` **Removed** (Breaking ⚠️)
    - `ScopeV2` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.Request.AccessReviewSetupEntitlementAndScopeServiceSetRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `AccessReviewSetupEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `ScopeV2` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `List[].AccessReviewEntitlement` **Added**
    - `List[].AccessReviewSetupEntitlement` **Removed** (Breaking ⚠️)
    - `ScopeV2` **Added**
* `ConductoroneApi.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.Request.AccessReviewSetScopeByResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `ScopeV2` **Added**
* `ConductoroneApi.AccessReviewTemplate.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Removed** (Breaking ⚠️)
    - `AccessReviewDuration` **Changed**
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `ColumnConfig` **Added**
    - `NotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `RecurrenceRule` **Added**
    - `RecurrenceRule` **Removed** (Breaking ⚠️)
    - `ReviewSignatureConfig` **Removed** (Breaking ⚠️)
    - `ReviewerAttributeConfig` **Added**
    - `ScopeType.Enum(accessReviewScopeTypeByUsers)` **Added**
    - `Scope` **Added**
    - `SignatureConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Added**
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReviewTemplate.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Added**
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReviewTemplate.Update()`: 
  * `request.Request.AccessReviewTemplateServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Added**
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Added**
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `List[].AccessReviewTemplateEntitlement` **Added**
    - `List[].AccessReviewTemplateSetupEntitlement` **Removed** (Breaking ⚠️)
    - `Scope` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.Request.AccessReviewTemplateSetupEntitlementServiceSetRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `AccessReviewTemplateSetupEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `Scope` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `List[].AccessReviewTemplateEntitlement` **Added**
    - `List[].AccessReviewTemplateSetupEntitlement` **Removed** (Breaking ⚠️)
    - `Scope` **Added**
* `ConductoroneApi.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.Request.AccessReviewTemplateSetScopeByResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `Scope` **Added**
* `ConductoroneApi.AccessConflict.CreateMonitor()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `NotificationConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `NegateGroupB` **Added**
    - `NotificationConfig` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AccessConflict.GetMonitor()`: `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `NegateGroupB` **Added**
    - `NotificationConfig` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AccessConflict.UpdateMonitor()`: 
  * `request.Request.ConflictMonitorUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `NegateGroupB` **Added**
    - `NotificationConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `NegateGroupB` **Added**
    - `NotificationConfig` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementMonitorBinding.CreateAppEntitlementMonitorBinding()`: `response` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementMonitorBinding.GetAppEntitlementMonitorBinding()`: `response` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Apps.Create()`: 
  *  `request.Request.Annotations` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `App` **Added**
    - `App` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Apps.Get()`: `response` **Changed** (Breaking ⚠️)
    - `App` **Added**
    - `App` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Apps.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AppOwners[].CreatedAt` **Changed** (Breaking ⚠️)
    - `AppOwners[].DeletedAt` **Changed** (Breaking ⚠️)
    - `AppOwners[].DepartmentSources[].Priority` **Added**
    - `AppOwners[].Profile` **Changed** (Breaking ⚠️)
    - `AppOwners[].UpdatedAt` **Changed** (Breaking ⚠️)
    - `AppUserMapper` **Added**
    - `AppUserMapper` **Removed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `RevokeGrantSources` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Apps.Update()`: 
  * `request.Request.UpdateAppRequest` **Changed** (Breaking ⚠️)
    - `App` **Added**
    - `App` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `App` **Added**
    - `App` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.Create()`: 
  * `request.Request.ConnectorServiceCreateRequest` **Changed** (Breaking ⚠️)
    - `Config` **Changed**
    - `ConnectorExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Added**
    - `ConnectorView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.CreateDelegated()`: 
  * `request.Request.ConnectorServiceCreateDelegatedRequest` **Changed** (Breaking ⚠️)
    - `AppManagedStateBindingRef` **Added**
    - `AppManagedStateBindingRef` **Removed** (Breaking ⚠️)
    - `ConnectorExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Added**
    - `ConnectorView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.Get()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Added**
    - `ConnectorView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.GetCredentials()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorCredential` **Removed** (Breaking ⚠️)
    - `Credential` **Added**
* `ConductoroneApi.Connector.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Connector` **Added**
    - `Connector` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.RotateCredential()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorCredential` **Removed** (Breaking ⚠️)
    - `Credential` **Added**
* `ConductoroneApi.Connector.Update()`: 
  * `request.Request.ConnectorServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `ConnectorExpandMask` **Removed** (Breaking ⚠️)
    - `Connector` **Added**
    - `Connector` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Added**
    - `ConnectorView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Connector.UpdateConnectorSchedule()`: 
  * `request.Request.UpdateConnectorScheduleRequest` **Changed** (Breaking ⚠️)
    - `ConnectorScheduleCron` **Removed** (Breaking ⚠️)
    - `Cron` **Added**
* `ConductoroneApi.Connector.UpdateDelegated()`: 
  * `request.Request.ConnectorServiceUpdateDelegatedRequest` **Changed** (Breaking ⚠️)
    - `ConnectorExpandMask` **Removed** (Breaking ⚠️)
    - `Connector` **Added**
    - `Connector` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Added**
    - `ConnectorView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.DurationGrant` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.Request.AppAccessRequestDefaults.DurationGrant` **Changed**
  *  `response.DurationGrant` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.DurationGrant` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppUser.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Added**
    - `AppUser` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppUser.ListAppUserCredentials()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `EncryptedData` **Added**
    - `EncryptedData` **Removed** (Breaking ⚠️)
    - `ExpiresAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppUser.ListAppUsersForUser()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Added**
    - `AppUser` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppUser.Search()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppIds` **Added**
    - `AppUserExpandMask` **Removed** (Breaking ⚠️)
    - `ExcludeDeletedApps` **Added**
    - `ExpandMask` **Added**
    - `SortBy` **Added**
    - `WithOpenFindings` **Added**
    - `WithoutResponsibleParty` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Added**
    - `AppUser` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppUser.Update()`: 
  * `request.Request.AppUserServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AppUserExpandMask` **Removed** (Breaking ⚠️)
    - `AppUser` **Added**
    - `AppUser` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppUserView` **Added**
    - `AppUserView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.Create()`: 
  * `request.Request.CreateAppEntitlementRequest` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AppEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `DurationGrant` **Changed**
    - `ExpandMask` **Added**
    - `ProvisionPolicy` **Added**
    - `ProvisionPolicy` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Added**
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.CreateAutomation()`: 
  * `request.Request.CreateAutomationRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomation` **Removed** (Breaking ⚠️)
    - `Automation` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomation` **Removed** (Breaking ⚠️)
    - `Automation` **Added**
* `ConductoroneApi.AppEntitlements.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Added**
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.GetAutomation()`: `response.AppEntitlementAutomation` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationLastRunStatus` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `Basic` **Added**
    - `Cel` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `Entitlements` **Added**
    - `LastRunStatus` **Added**
    - `None` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppEntitlements.ListAutomationExclusions()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.ListForAppResource()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppEntitlements.ListForAppUser()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppEntitlements.ListUsers()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementId` **Added**
    - `AppEntitlementUserBindingCreatedAt` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingDeprovisionAt` **Changed** (Breaking ⚠️)
    - `AppId` **Added**
    - `AppUserId` **Added**
    - `AppUserView` **Removed** (Breaking ⚠️)
    - `AppUser` **Added**
* `ConductoroneApi.AppEntitlements.Update()`: 
  * `request.Request.UpdateAppEntitlementRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `Entitlement` **Added**
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Added**
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlements.UpdateAutomation()`: 
  * `request.Request.AppEntitlementServiceUpdateAutomationRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `Basic` **Added**
    - `Cel` **Added**
    - `Entitlements` **Added**
    - `None` **Added**
  * `response.AppEntitlementAutomation` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationLastRunStatus` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `Basic` **Added**
    - `Cel` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `Entitlements` **Added**
    - `LastRunStatus` **Added**
    - `None` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementSearch.Search()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `RequestSchemaIds` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Facets` **Added**
    - `Facets` **Removed** (Breaking ⚠️)
    - `List[].ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `List[].AppEntitlement` **Added**
    - `List[].AppEntitlement` **Removed** (Breaking ⚠️)
    - `List[].ObjectPermissions` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Added**
    - `AppUser` **Removed** (Breaking ⚠️)
    - `Discovered` **Changed** (Breaking ⚠️)
    - `Expired` **Changed** (Breaking ⚠️)
    - `GrantReasons[].CreatedAt` **Changed** (Breaking ⚠️)
    - `GrantReasons[].DeletedAt` **Changed** (Breaking ⚠️)
    - `GrantReasons[].ReasonExpiresAt` **Changed** (Breaking ⚠️)
    - `GrantReasons[].UpdatedAt` **Changed** (Breaking ⚠️)
    - `User` **Added**
    - `User` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementSearch.SearchGrants()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBinding` **Added**
    - `AppEntitlementUserView` **Removed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `Entitlement` **Added**
* `ConductoroneApi.Finding.BulkCreateFindingTasks()`: `request.Request` **Changed** (Breaking ⚠️)
    - `FindingSearchRequest` **Removed** (Breaking ⚠️)
    - `SearchRequest` **Added**
* `ConductoroneApi.Directory.Update()`: 
  * `request.Request.DirectoryServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `All` **Added**
    - `CelExpression` **Added**
    - `DirectoryAccountFilterAll` **Removed** (Breaking ⚠️)
    - `DirectoryAccountFilterCel` **Removed** (Breaking ⚠️)
    - `DirectoryExpandMask` **Removed** (Breaking ⚠️)
    - `DirectoryMergeConfig` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `MergeConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Added**
    - `DirectoryView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Directory.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Directory` **Added**
    - `Directory` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementUserBinding.SearchPastGrants()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingExpandHistoryMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response.List[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingHistory` **Removed** (Breaking ⚠️)
    - `History` **Added**
* `ConductoroneApi.AppEntitlementUserBinding.UpdateGrantDuration()`: 
  *  `request.Request.UpdateGrantDurationRequest.NewDeprovisionAt` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBinding` **Removed** (Breaking ⚠️)
    - `Binding` **Added**
* `ConductoroneApi.AppEntitlementOwners.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `DepartmentSources[].Priority` **Added**
    - `Profile` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppOwners.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `DepartmentSources[].Priority` **Added**
    - `Profile` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppReport.List()`:  `response.List[].CreatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceType.CreateManuallyManagedResourceType()`: 
  *  `request.Request.CreateManuallyManagedResourceTypeRequest.ResourceType.Enum(sessionPolicy)` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Added**
    - `AppResourceType` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceType.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppResourceTypeView` **Added**
    - `AppResourceTypeView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceType.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Added**
    - `AppResourceType` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceType.UpdateManuallyManagedResourceType()`: 
  * `request.Request.UpdateManuallyManagedResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Added**
    - `AppResourceType` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Added**
    - `AppResourceType` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResource.CreateManuallyManagedAppResource()`: 
  *  `request.Request.CreateManuallyManagedAppResourceRequest.Annotations` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResource` **Added**
    - `AppResource` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResource.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppResourceView` **Added**
    - `AppResourceView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResource.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppResource` **Added**
    - `AppResource` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.AppResource.Update()`: 
  * `request.Request.AppResourceServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AppResourceExpandMask` **Removed** (Breaking ⚠️)
    - `AppResource` **Added**
    - `AppResource` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceView` **Added**
    - `AppResourceView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppResourceOwners.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `DepartmentSources[].Priority` **Added**
    - `Profile` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.AppUsageControls.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Added**
    - `AppUsageControls` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppUsageControls.Update()`: 
  * `request.Request.UpdateAppUsageControlsRequest` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Added**
    - `AppUsageControls` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Added**
    - `AppUsageControls` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AppEntitlementsProxy.Create()`: 
  * `request.Request.CreateAppEntitlementProxyRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementProxyExpandMask` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementProxyView` **Removed** (Breaking ⚠️)
    - `AppProxyEntitlementView` **Added**
* `ConductoroneApi.AppEntitlementsProxy.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementProxyView` **Removed** (Breaking ⚠️)
    - `AppProxyEntitlementView` **Added**
* `ConductoroneApi.Attributes.CreateAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.CreateComplianceFrameworkAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.CreateRiskLevelAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.GetAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.GetComplianceFrameworkAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.GetRiskLevelAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `Value` **Added**
* `ConductoroneApi.Attributes.ListAttributeValues()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Attributes.ListComplianceFrameworks()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Attributes.ListRiskLevels()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.TenantAuthConfig.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `AuthConfigC1Local` **Removed** (Breaking ⚠️)
    - `AuthConfigGoogle` **Removed** (Breaking ⚠️)
    - `AuthConfigJumpCloud` **Removed** (Breaking ⚠️)
    - `AuthConfigMicrosoft` **Removed** (Breaking ⚠️)
    - `AuthConfigOidc` **Removed** (Breaking ⚠️)
    - `AuthConfigOkta` **Removed** (Breaking ⚠️)
    - `AuthConfigOneLogin` **Removed** (Breaking ⚠️)
    - `AuthConfigPingOne` **Removed** (Breaking ⚠️)
    - `C1Local` **Added**
    - `DeprecationDeadline` **Changed**
    - `Google` **Added**
    - `Jumpcloud` **Added**
    - `Microsoft` **Added**
    - `Oidc` **Added**
    - `Okta` **Added**
    - `Onelogin` **Added**
    - `Pingone` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AuthConfig` **Added**
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TenantAuthConfig.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AuthConfig` **Added**
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.TenantAuthConfig.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `AuthConfigC1Local` **Removed** (Breaking ⚠️)
    - `AuthConfigGoogle` **Removed** (Breaking ⚠️)
    - `AuthConfigJumpCloud` **Removed** (Breaking ⚠️)
    - `AuthConfigMicrosoft` **Removed** (Breaking ⚠️)
    - `AuthConfigOidc` **Removed** (Breaking ⚠️)
    - `AuthConfigOkta` **Removed** (Breaking ⚠️)
    - `AuthConfigOneLogin` **Removed** (Breaking ⚠️)
    - `AuthConfigPingOne` **Removed** (Breaking ⚠️)
    - `C1Local` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeprecationDeadline` **Changed** (Breaking ⚠️)
    - `Google` **Added**
    - `Jumpcloud` **Added**
    - `Microsoft` **Added**
    - `Oidc` **Added**
    - `Okta` **Added**
    - `Onelogin` **Added**
    - `Pingone` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.TenantAuthConfig.Update()`: 
  * `request.Request.TenantAuthConfigServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AuthConfig` **Added**
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AuthConfig` **Added**
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Directory.Get()`: `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Added**
    - `DirectoryView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.AutomationExecution.GetAutomationExecution()`: `response` **Changed** (Breaking ⚠️)
    - `AutomationExecutionView` **Removed** (Breaking ⚠️)
    - `AutomationExecution` **Added**
    - `AutomationExecution` **Removed** (Breaking ⚠️)
    - `View` **Added**
* `ConductoroneApi.AutomationExecution.ListAutomationExecutions()`: `response.AutomationExecutions[]` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `CompletedAt` **Changed** (Breaking ⚠️)
    - `Context` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `Duration` **Changed** (Breaking ⚠️)
    - `State.Enum(automationExecutionStatePausedByCircuitBreaker)` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.Directory.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `All` **Added**
    - `CelExpression` **Added**
    - `DirectoryAccountFilterAll` **Removed** (Breaking ⚠️)
    - `DirectoryAccountFilterCel` **Removed** (Breaking ⚠️)
    - `DirectoryExpandMask` **Removed** (Breaking ⚠️)
    - `DirectoryMergeConfig` **Removed** (Breaking ⚠️)
    - `ExpandMask` **Added**
    - `MergeConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Added**
    - `DirectoryView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.CreateAutomation()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].AccountLifecycleAction` **Added**
    - `AutomationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CallFunction` **Added**
    - `AutomationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorAction` **Added**
    - `AutomationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount` **Added**
    - `AutomationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateAccessReview` **Added**
    - `AutomationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasksV2` **Added**
    - `AutomationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasks` **Added**
    - `AutomationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].EvaluateExpressions` **Added**
    - `AutomationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GeneratePassword` **Added**
    - `AutomationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements` **Added**
    - `AutomationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RemoveFromDelegation` **Added**
    - `AutomationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RunAutomation` **Added**
    - `AutomationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendEmail` **Added**
    - `AutomationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendSlackMessage` **Added**
    - `AutomationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].TaskAction` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Added**
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UpdateUser` **Added**
    - `AutomationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].WaitForDuration` **Added**
    - `AutomationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Added**
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Context` **Added**
    - `DraftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AccessConflict` **Added**
    - `DraftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserCreated` **Added**
    - `DraftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserUpdated` **Added**
    - `DraftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeleted` **Added**
    - `DraftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantFound` **Added**
    - `DraftTriggers[].ScheduleAppUser` **Added**
    - `DraftTriggers[].ScheduleNoUser` **Added**
    - `DraftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Schedule` **Added**
    - `DraftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UsageBasedRevocation` **Added**
    - `DraftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserCreated` **Added**
    - `DraftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserProfileChange` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Webhook` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation` **Added**
    - `Automation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.ExecuteAutomation()`: 
  * `request.Request.ExecuteAutomationRequest` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `Context` **Added**
* `ConductoroneApi.Automation.GetAutomation()`: `response` **Changed** (Breaking ⚠️)
    - `Automation` **Added**
    - `Automation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.Automation.ListAutomations()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].AccountLifecycleAction` **Added**
    - `AutomationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CallFunction` **Added**
    - `AutomationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorAction` **Added**
    - `AutomationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].ConnectorCreateAccount` **Added**
    - `AutomationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateAccessReview` **Added**
    - `AutomationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasksV2` **Added**
    - `AutomationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].CreateRevokeTasks` **Added**
    - `AutomationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].EvaluateExpressions` **Added**
    - `AutomationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GeneratePassword` **Added**
    - `AutomationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].GrantEntitlements` **Added**
    - `AutomationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RemoveFromDelegation` **Added**
    - `AutomationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].RunAutomation` **Added**
    - `AutomationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendEmail` **Added**
    - `AutomationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SendSlackMessage` **Added**
    - `AutomationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].SetCredential` **Added**
    - `AutomationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].StoreCredential` **Added**
    - `AutomationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].TaskAction` **Added**
    - `AutomationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Added**
    - `AutomationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].UpdateUser` **Added**
    - `AutomationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].WaitForDuration` **Added**
    - `AutomationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `AutomationSteps[].Webhook` **Added**
    - `AutomationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `CircuitBreaker` **Added**
    - `Context` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DisabledReasonCircuitBreaker` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AccessConflict` **Added**
    - `DraftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserCreated` **Added**
    - `DraftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].AppUserUpdated` **Added**
    - `DraftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantDeleted` **Added**
    - `DraftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].GrantFound` **Added**
    - `DraftTriggers[].ScheduleAppUser` **Added**
    - `DraftTriggers[].ScheduleNoUser` **Added**
    - `DraftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Schedule` **Added**
    - `DraftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UsageBasedRevocation` **Added**
    - `DraftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserCreated` **Added**
    - `DraftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].UserProfileChange` **Added**
    - `DraftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `DraftTriggers[].Webhook` **Added**
    - `LastExecutedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed** (Breaking ⚠️)
    - `ConfigSchema` **Removed** (Breaking ⚠️)
    - `FormSchema` **Added**
    - `RequestSchemaForm` **Removed** (Breaking ⚠️)
    - `Schema` **Added**
* `ConductoroneApi.Automation.UpdateAutomation()`: 
  * `request.Request.UpdateAutomationRequest` **Changed** (Breaking ⚠️)
    - `Automation` **Added**
    - `Automation` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `Automation` **Added**
    - `Automation` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.Create()`: 
  * `request.Request` **Changed** (Breaking ⚠️)
    - `Annotations` **Added**
    - `ExpandMask` **Added**
    - `GrantPolicyId` **Removed** (Breaking ⚠️)
    - `RequestCatalogExpandMask` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Added**
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.CreateBundleAutomation()`: 
  * `request.Request.CreateBundleAutomationRequest` **Changed** (Breaking ⚠️)
    - `BundleAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `Cel` **Added**
    - `EnforceOnSmallProfiles` **Added**
    - `Entitlements` **Added**
    - `RemovedMembersThresholdPercent` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `Cel` **Added**
    - `CircuitBreaker` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `EnforceOnSmallProfiles` **Added**
    - `Entitlements` **Added**
    - `RemovedMembersThresholdPercent` **Added**
    - `State` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.CreateRequestableEntry()`: `response` **Changed** (Breaking ⚠️)
    - `RequestableEntry` **Added**
    - `RequestableEntry` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.Get()`: `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Added**
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.GetBundleAutomation()`: `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `Cel` **Added**
    - `CircuitBreaker` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `EnforceOnSmallProfiles` **Added**
    - `Entitlements` **Added**
    - `RemovedMembersThresholdPercent` **Added**
    - `State` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.GetRequestableEntry()`: `response` **Changed** (Breaking ⚠️)
    - `RequestableEntry` **Added**
    - `RequestableEntry` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.List()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `RequestCatalog` **Added**
    - `RequestCatalog` **Removed** (Breaking ⚠️)
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.List[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Added**
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `ObjectPermissions` **Added**
* `ConductoroneApi.RequestCatalogManagement.SetBundleAutomation()`: 
  * `request.Request.SetBundleAutomationRequest` **Changed** (Breaking ⚠️)
    - `BundleAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `Cel` **Added**
    - `EnforceOnSmallProfiles` **Added**
    - `Entitlements` **Added**
    - `RemovedMembersThresholdPercent` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCel` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `Cel` **Added**
    - `CircuitBreaker` **Added**
    - `CreatedAt` **Changed** (Breaking ⚠️)
    - `DeletedAt` **Changed** (Breaking ⚠️)
    - `EnforceOnSmallProfiles` **Added**
    - `Entitlements` **Added**
    - `RemovedMembersThresholdPercent` **Added**
    - `State` **Added**
    - `UpdatedAt` **Changed** (Breaking ⚠️)
* `ConductoroneApi.SignInPolicy.Delete()`: **Added**
* `ConductoroneApi.Automation.ResolvePausedAutomationExecutions()`:  `response.BulkActionId` **Added**
* `ConductoroneApi.Automation.ClearAutomationCircuitBreaker()`: 
  * `request.Request.ClearAutomationCircuitBreakerRequest` **Changed**
    - `Decision` **Added**
    - `Reason` **Added**
  *  `response.BulkActionId` **Added**
* `ConductoroneApi.Auth.Introspect()`: `response` **Changed**
    - `DeviceClientId` **Added**
    - `TenantId` **Added**
* `ConductoroneApi.AppResourceOwnersV2.Set()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.ReorderAppEntitlementRoutingRules()`: **Added**
* `ConductoroneApi.XaaAccessProfile.GetByAppEntitlementId()`: **Added**
* `ConductoroneApi.A2Ui.SubmitAction()`: 
  *  `request.Request.A2UiServiceSubmitActionRequest.ClientTimestamp` **Changed**
* `ConductoroneApi.UserOwnersV2.Set()`: **Added**
* `ConductoroneApi.UserOwnersV2.SearchUserOwners()`: **Added**
* `ConductoroneApi.UserOwnersV2.SearchEntitlementOwners()`: **Added**
* `ConductoroneApi.UserOwnersV2.DeleteUserOwner()`: **Added**
* `ConductoroneApi.UserOwnersV2.DeleteEntitlementOwner()`: **Added**
* `ConductoroneApi.UserOwnersV2.CreateUserOwner()`: **Added**
* `ConductoroneApi.UserOwnersV2.CreateEntitlementOwner()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.Set()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.SearchUserOwners()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.SearchEntitlementOwners()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.DeleteUserOwner()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.DeleteEntitlementOwner()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.CreateUserOwner()`: **Added**
* `ConductoroneApi.AppUserOwnersV2.CreateEntitlementOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.SearchUserOwners()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.SearchEntitlementOwners()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.GetUserOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.GetEntitlementOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.DeleteUserOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.DeleteEntitlementOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.CreateUserOwner()`: **Added**
* `ConductoroneApi.AppResourceOwnersV2.CreateEntitlementOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.GetUserOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.GetEntitlementOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.DeleteUserOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.DeleteEntitlementOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.CreateUserOwner()`: **Added**
* `ConductoroneApi.AppEntitlementOwnersV2.CreateEntitlementOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.GetUserOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.GetEntitlementOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.DeleteUserOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.DeleteEntitlementOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.CreateUserOwner()`: **Added**
* `ConductoroneApi.ConnectorOwnersV2.CreateEntitlementOwner()`: **Added**
* `ConductoroneApi.User.Introspect()`: **Added**
* `ConductoroneApi.TerraformExport.GetSchema()`: **Added**
* `ConductoroneApi.Task.CreateResourceActionTask()`: **Added**
* `ConductoroneApi.Task.CreateActionTask()`: **Added**
* `ConductoroneApi.RequestSettings.Update()`: **Added**
* `ConductoroneApi.RequestSettings.Get()`: **Added**
* `ConductoroneApi.IdentityPolicyTenantDefaults.Update()`: **Added**
* `ConductoroneApi.IdentityPolicyTenantDefaults.Get()`: **Added**
* `ConductoroneApi.UserDeveloperPreferences.Update()`: **Added**
* `ConductoroneApi.UserDeveloperPreferences.Get()`: **Added**
* `ConductoroneApi.XaaSettings.Update()`: **Added**
* `ConductoroneApi.XaaSettings.ListHistory()`: **Added**
* `ConductoroneApi.XaaSettings.Get()`: **Added**
* `ConductoroneApi.AiGovernanceSettings.Update()`: **Added**
* `ConductoroneApi.AiGovernanceSettings.ListHistory()`: **Added**
* `ConductoroneApi.RoleMiningManagement.CreateAccessProfileFromCohort()`: `request.Request` **Changed**
    - `CelExpression` **Added**
    - `Entitlements[].RiskLevelValueId` **Added**
* `ConductoroneApi.RoleMiningManagement.GetCustomAnalysisResult()`:  `response.Clusters[].Entitlements[].RiskLevelValueId` **Added**
* `ConductoroneApi.AiGovernanceSettings.GetTenantDefaults()`: **Added**
* `ConductoroneApi.AiGovernanceSettings.Get()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.Update()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.Search()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.ListHistory()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.List()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.Get()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.Delete()`: **Added**
* `ConductoroneApi.XaaClientAudienceMapping.Create()`: **Added**
* `ConductoroneApi.SignInPolicy.Update()`: **Added**
* `ConductoroneApi.SignInPolicy.Search()`: **Added**
* `ConductoroneApi.SignInPolicy.List()`: **Added**
* `ConductoroneApi.SignInPolicy.Get()`: **Added**
* `ConductoroneApi.AppUser.ListOwnedServiceAccounts()`: **Added**
* `ConductoroneApi.McpTool.Delete()`: **Added**
* `ConductoroneApi.SignInPolicy.Create()`: **Added**
* `ConductoroneApi.SessionPolicy.Update()`: **Added**
* `ConductoroneApi.SessionPolicy.UnassignUser()`: **Added**
* `ConductoroneApi.SessionPolicy.UnassignGroup()`: **Added**
* `ConductoroneApi.SessionPolicy.Search()`: **Added**
* `ConductoroneApi.SessionPolicy.ListAssignments()`: **Added**
* `ConductoroneApi.SessionPolicy.List()`: **Added**
* `ConductoroneApi.SessionPolicy.Get()`: **Added**
* `ConductoroneApi.SessionPolicy.Delete()`: **Added**
* `ConductoroneApi.SessionPolicy.Create()`: **Added**
* `ConductoroneApi.SessionPolicy.AssignUser()`: **Added**
* `ConductoroneApi.SessionPolicy.AssignGroup()`: **Added**
* `ConductoroneApi.FindingAudit.Search()`: **Added**
* `ConductoroneApi.RoleMiningManagement.ListCustomAnalysisResults()`: **Added**
* `ConductoroneApi.RecoveryPolicy.Update()`: **Added**
* `ConductoroneApi.RecoveryPolicy.Search()`: **Added**
* `ConductoroneApi.RecoveryPolicy.List()`: **Added**
* `ConductoroneApi.RecoveryPolicy.Get()`: **Added**
* `ConductoroneApi.RecoveryPolicy.Delete()`: **Added**
* `ConductoroneApi.RecoveryPolicy.Create()`: **Added**
* `ConductoroneApi.TunnelCredentials.UpdateBridge()`: **Added**
* `ConductoroneApi.TunnelCredentials.RevokeBridgeCredential()`: **Added**
* `ConductoroneApi.TunnelCredentials.ListBridges()`: **Added**
* `ConductoroneApi.TunnelCredentials.ListBridgeCredentials()`: **Added**
* `ConductoroneApi.TunnelCredentials.ListBridgeAnnouncedServices()`: **Added**
* `ConductoroneApi.TunnelCredentials.GetBridge()`: **Added**
* `ConductoroneApi.TunnelCredentials.DeleteBridge()`: **Added**
* `ConductoroneApi.TunnelCredentials.CreateBridgeCredential()`: **Added**
* `ConductoroneApi.TunnelCredentials.CreateBridge()`: **Added**
* `ConductoroneApi.PersonalDevice.UpdateDevice()`: **Added**
* `ConductoroneApi.PersonalDevice.Search()`: **Added**
* `ConductoroneApi.PersonalDevice.RevokeDeviceClient()`: **Added**
* `ConductoroneApi.PersonalDevice.RevokeDevice()`: **Added**
* `ConductoroneApi.PersonalDevice.ListDeviceClients()`: **Added**
* `ConductoroneApi.PersonalDevice.GetDevice()`: **Added**
* `ConductoroneApi.FindingTransformationRule.UpdateFindingTransformationRule()`: **Added**
* `ConductoroneApi.FindingTransformationRule.ListFindingTransformationRules()`: **Added**
* `ConductoroneApi.FindingTransformationRule.GetFindingTransformationRule()`: **Added**
* `ConductoroneApi.FindingTransformationRule.DeleteFindingTransformationRule()`: **Added**
* `ConductoroneApi.FindingTransformationRule.CreateFindingTransformationRule()`: **Added**
* `ConductoroneApi.Finding.CreateFinding()`: **Added**
* `ConductoroneApi.DecoySearch.Search()`: **Added**
* `ConductoroneApi.Decoy.Update()`: **Added**
* `ConductoroneApi.Decoy.Rotate()`: **Added**
* `ConductoroneApi.Decoy.List()`: **Added**
* `ConductoroneApi.Decoy.Get()`: **Added**
* `ConductoroneApi.Decoy.Delete()`: **Added**
* `ConductoroneApi.Decoy.Create()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.Update()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.Search()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.List()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.Get()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.Delete()`: **Added**
* `ConductoroneApi.CredentialInventoryPolicy.Create()`: **Added**
* `ConductoroneApi.ConnectorAuthoringActivation.RollbackRevision()`: **Added**
* `ConductoroneApi.ConnectorAuthoringActivation.ActivateRevision()`: **Added**
* `ConductoroneApi.XaaScope.Update()`: **Added**
* `ConductoroneApi.XaaScope.Search()`: **Added**
* `ConductoroneApi.XaaScope.ListHistory()`: **Added**
* `ConductoroneApi.XaaScope.List()`: **Added**
* `ConductoroneApi.XaaScope.Get()`: **Added**
* `ConductoroneApi.XaaScope.Delete()`: **Added**
* `ConductoroneApi.XaaScope.Create()`: **Added**
* `ConductoroneApi.XaaResourceServer.Update()`: **Added**
* `ConductoroneApi.XaaResourceServer.Search()`: **Added**
* `ConductoroneApi.OnboardingSettings.Update()`: 
  * `request.Request` **Changed**
    - `McpOnboardingGoal` **Added**
    - `McpOnboardingStatus` **Added**
    - `McpOnboardingTargets` **Added**
  * `response` **Changed**
    - `McpOnboardingGoal` **Added**
    - `McpOnboardingStatus` **Added**
    - `McpOnboardingTargets` **Added**
* `ConductoroneApi.XaaResourceServer.ListHistory()`: **Added**
* `ConductoroneApi.XaaResourceServer.List()`: **Added**
* `ConductoroneApi.XaaResourceServer.Get()`: **Added**
* `ConductoroneApi.XaaResourceServer.Delete()`: **Added**
* `ConductoroneApi.XaaResourceServer.Create()`: **Added**
* `ConductoroneApi.XaaAccessProfileScopeBinding.Search()`: **Added**
* `ConductoroneApi.XaaAccessProfileScopeBinding.List()`: **Added**
* `ConductoroneApi.XaaAccessProfileScopeBinding.DeleteBindings()`: **Added**
* `ConductoroneApi.XaaAccessProfileScopeBinding.CreateBindings()`: **Added**
* `ConductoroneApi.SystemLog.ListEvents()`: `request.Request` **Changed**
    - `Since` **Changed**
    - `Until` **Changed**
* `ConductoroneApi.XaaAccessProfile.Update()`: **Added**
* `ConductoroneApi.XaaAccessProfile.Search()`: **Added**
* `ConductoroneApi.XaaAccessProfile.ListHistory()`: **Added**
* `ConductoroneApi.XaaAccessProfile.List()`: **Added**
* `ConductoroneApi.XaaAccessProfile.Get()`: **Added**
* `ConductoroneApi.XaaAccessProfile.Delete()`: **Added**
* `ConductoroneApi.XaaAccessProfile.Create()`: **Added**
* `ConductoroneApi.McpServer.UpdateCredentials()`: **Added**
* `ConductoroneApi.McpServer.Update()`: **Added**
* `ConductoroneApi.McpServer.TestConnection()`: **Added**
* `ConductoroneApi.McpServer.SearchWithToolCount()`: **Added**
* `ConductoroneApi.McpServer.ResyncTools()`: **Added**
* `ConductoroneApi.McpServer.Register()`: **Added**
* `ConductoroneApi.McpServer.ListConnections()`: **Added**
* `ConductoroneApi.McpServer.ListCatalog()`: **Added**
* `ConductoroneApi.McpServer.List()`: **Added**
* `ConductoroneApi.McpServer.GetCatalog()`: **Added**
* `ConductoroneApi.McpServer.Get()`: **Added**
* `ConductoroneApi.McpServer.DiscoverOidcEndpoints()`: **Added**
* `ConductoroneApi.McpServer.Delete()`: **Added**
* `ConductoroneApi.AppEntitlementSearch.SearchGraph()`: **Added**
* `ConductoroneApi.AppEntitlementSearch.CountGrantsForUserByApp()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.UpdateAppEntitlementRoutingRule()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.ListAppEntitlementRoutingRules()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.GetAppEntitlementRoutingRule()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.DeleteAppEntitlementRoutingRule()`: **Added**
* `ConductoroneApi.AppEntitlementRoutingRule.CreateAppEntitlementRoutingRule()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.ListToolsByProfileHistory()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.ListProfilesByToolHistory()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.List()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.GetAccessProfilesForTools()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.DeleteBindings()`: **Added**
* `ConductoroneApi.McpAccessProfileToolBinding.CreateBindings()`: **Added**
* `ConductoroneApi.McpAccessProfile.Update()`: **Added**
* `ConductoroneApi.McpAccessProfile.SearchRequestableConnectors()`: **Added**
* `ConductoroneApi.McpAccessProfile.ListRequestableConnectors()`: **Added**
* `ConductoroneApi.McpAccessProfile.List()`: **Added**
* `ConductoroneApi.McpAccessProfile.GetByAppEntitlementId()`: **Added**
* `ConductoroneApi.McpAccessProfile.Get()`: **Added**
* `ConductoroneApi.McpAccessProfile.Delete()`: **Added**
* `ConductoroneApi.McpAccessProfile.Create()`: **Added**
* `ConductoroneApi.McpTool.Update()`: **Added**
* `ConductoroneApi.McpTool.Search()`: **Added**
* `ConductoroneApi.McpTool.ListHistory()`: **Added**
* `ConductoroneApi.McpTool.List()`: **Added**
* `ConductoroneApi.McpTool.Get()`: **Added**
