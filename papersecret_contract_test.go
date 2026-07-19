package conductoronesdkgo

import (
	"testing"

	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

func TestPaperSecretMLKEMSuiteContract(t *testing.T) {
	internalSuite := shared.PaperSecretServiceCreateInternalRequestRequiredAgeSuiteAgeSuiteMlkem768X25519
	externalSuite := shared.RequiredAgeSuiteAgeSuiteMlkem768X25519
	responseSuite := shared.PaperSecretServiceCreateResponseAgeSuiteAgeSuiteMlkem768X25519

	internal := shared.PaperSecretServiceCreateInternalRequest{RequiredAgeSuite: internalSuite.ToPointer()}
	external := shared.PaperSecretServiceCreateExternalRequest{RequiredAgeSuite: externalSuite.ToPointer()}
	response := shared.PaperSecretServiceCreateResponse{AgeSuite: responseSuite.ToPointer()}

	if got := internal.GetRequiredAgeSuite(); got == nil || string(*got) != "AGE_SUITE_MLKEM768X25519" {
		t.Fatalf("internal required age suite = %v", got)
	}
	if got := external.GetRequiredAgeSuite(); got == nil || string(*got) != "AGE_SUITE_MLKEM768X25519" {
		t.Fatalf("external required age suite = %v", got)
	}
	if got := response.GetAgeSuite(); got == nil || string(*got) != "AGE_SUITE_MLKEM768X25519" {
		t.Fatalf("response age suite = %v", got)
	}
}
