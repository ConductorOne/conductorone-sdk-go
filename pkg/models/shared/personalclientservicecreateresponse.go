// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PersonalClientServiceCreateResponse message contains the created personal client and client secret.
type PersonalClientServiceCreateResponse struct {
	// The PersonalClient message contains information about a presonal client credential.
	PersonalClient *PersonalClient `json:"client,omitempty"`
	// The client secret that corresponds to the personal client. Make sure to save this, because it cannot be returned or queried again.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

func (o *PersonalClientServiceCreateResponse) GetPersonalClient() *PersonalClient {
	if o == nil {
		return nil
	}
	return o.PersonalClient
}

func (o *PersonalClientServiceCreateResponse) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}
