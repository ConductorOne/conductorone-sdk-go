# Migrating from v1 to v2

This guide covers the breaking changes between `conductorone-sdk-go` v1.x and v2.0.0.

The v2.0.0 line was cut to absorb the spec drift accumulated since v1.27.x — specifically the [IGA-1181](https://linear.app/ductone/issue/IGA-1181) collision-fix renames plus six weeks of API additions. The v2.0.0-alpha.1 release is the first prerelease in this line; it is **not yet GA**.

## 1. Module path break

Per [Go's semantic import versioning rules](https://research.swtch.com/vgo-import), v2.0.0 changes the import path:

```diff
- import "github.com/conductorone/conductorone-sdk-go"
+ import "github.com/conductorone/conductorone-sdk-go/v2"
```

Update `go.mod`:

```bash
go get github.com/conductorone/conductorone-sdk-go/v2@v2.0.0-alpha.1
go mod tidy
```

A simple `gofmt`-friendly find-and-replace covers most cases:

```bash
find . -name '*.go' -not -path './vendor/*' \
  -exec sed -i.bak 's|github.com/conductorone/conductorone-sdk-go/|github.com/conductorone/conductorone-sdk-go/v2/|g' {} +
find . -name '*.go.bak' -delete
goimports -w .
```

The v1 module remains available at the unsuffixed path; v1 and v2 can coexist in the same binary if needed during a phased migration.

## 2. Renamed types (IGA-1181 collision fixes)

Several types were renamed to disambiguate cross-package collisions in the OpenAPI spec. The old names existed in v1; the new names are required in v2.

> _Populated from the v2.0.0-alpha.1 regen PR's `**Breaking ⚠️**` list. The placeholders below will be filled in once that PR lands._

### Service splits

Two operations now have separate v1 and v2 service types because the API exposes both endpoints:

| v1 type | v2 type for v1 endpoint | v2 type for v2 endpoint |
|---|---|---|
| `AppOwners.Set` | `AppOwners.Set` (unchanged) | `AppOwnersV2.Set` (new) |
| `AppEntitlementOwners.Set` | `AppEntitlementOwners.Set` (unchanged) | `AppEntitlementOwnersV2.Set` (new) |

Pick the version that matches the API endpoint your code targets. v1 endpoints continue to work; v2 endpoints have additional capabilities (entitlement-scoped owners, etc.).

### Schema renames

| v1 | v2 |
|---|---|
| `NotificationConfig` (in access_conflict context) | `AccessConflictNotificationConfig` |
| `NotificationConfig` (in access_review context) | `AccessReviewNotificationConfig` |
| `TaskAction` | `AutomationsTaskAction` |
| `Webhook` (in automations context) | `AutomationsWebhook` |
| `Form` (in policy step context) | `PolicyForm` |

### Field-level changes

> _To be populated from the regen PR. Likely candidates: `TaskView` shape changes across all `TaskActions.*` and `Task.*` operations, `AccessReview` create/update/get/list response shape changes, `AccessReviewTemplate` similar._

## 3. New services and fields (non-breaking, but newly available)

Six weeks of c1 spec additions are now available in v2:

- **Paper Secret** — `PaperSecret`, `PaperSecretAdmin` services for vault-backed secret sharing.
- **Role Mining Management** — `RoleMiningManagement`, `RoleMiningManagementSearch` for surfaced role suggestions.
- **Principal / Credentials** — `Principal` service with credential CRUD + revocation.
- **Workload Federation** — `WorkloadFederation` (trusts, providers, CEL testing) for non-human identity flows.
- **A2UI** — `A2UI` (surfaces, feedback, actions) for Agent-to-UI integrations.
- **Automations expansion** — additional automation triggers and steps (`schedule_trigger_no_user`, `webhook_listener_auth_capability_url`, etc.).
- **Webhook callback timeout** — `Webhook.callback_timeout` now configurable.
- **Other fields**: `connector_api_version`, `directory_merge_config`, `provisioner_assignment` on entitlement provision policies, `origin` on owners/users, etc.

These are additive and do not require code changes to existing v1 code beyond the import path bump in §1.

## 4. Workflow / CI

If your CI fetches the SDK at a pinned version, update accordingly:

```bash
# replace pinned version
go get github.com/conductorone/conductorone-sdk-go/v2@v2.0.0-alpha.1
```

If you use Dependabot / Renovate, both will detect the new module path as a separate package — you may need to add a config rule pairing v1 → v2 to track upgrades.

## 5. Reporting issues

Issues with v2.0.0-alpha.1 should reference the alpha tag explicitly. File them at https://github.com/ConductorOne/conductorone-sdk-go/issues. The IGA team owns the SDK and TF provider — see the [Terraform Linear project](https://linear.app/ductone/project/423a706e-1c8d-4283-b1e6-43f082f53826) for active work.
