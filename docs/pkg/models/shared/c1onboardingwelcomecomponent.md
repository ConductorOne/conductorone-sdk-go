# C1OnboardingWelcomeComponent

C1OnboardingWelcomeComponent renders the onboarding welcome screen with org context and intent collection.
 Backend pre-populates recommended_catalog_id / recommended_display_name from detected IDP.
 Frontend detects auth backend via introspect for contextual UI text.


## Fields

| Field                             | Type                              | Required                          | Description                       |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `RecommendedCatalogID`            | `*string`                         | :heavy_minus_sign:                | The recommendedCatalogId field.   |
| `RecommendedDisplayName`          | `*string`                         | :heavy_minus_sign:                | The recommendedDisplayName field. |