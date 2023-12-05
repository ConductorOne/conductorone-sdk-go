package conductoronesdkgo

import (
	"testing"
)

func TestNormalizeTenantInput(t *testing.T) {

	tests := []struct {
		name               string
		input              string
		wantedURL          string
		wantedTenantDomain string
		wantErr            bool
	}{
		{
			name:               "base",
			input:              "insulator.conductor.one",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "tenantOnly",
			input:              "insulator",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "baseWithHTTP",
			input:              "https://insulator.conductor.one",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "differentUrlWithHTTP",
			input:              "https://insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
		{
			name:               "differentUrlWithHTTP",
			input:              "insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
		{
			name:               "differentURL",
			input:              "insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithTenant(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithTenant returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			client := New(got)
			if client.sdkConfiguration.ServerURL != tt.wantedURL {
				t.Errorf("ServerURL Got = '%v', want '%v'", client.sdkConfiguration.ServerURL, tt.wantedURL)
			}
			if client.sdkConfiguration.ServerDefaults[0]["tenantDomain"] != tt.wantedTenantDomain {
				t.Errorf("TenantDomain Got = '%v', want '%v'", client.sdkConfiguration.ServerDefaults[0]["tenantDomain"], tt.wantedTenantDomain)
			}
		})
	}
}

type ClientConfigTest struct {
	name               string
	input              string
	wantedURL          string
	wantedTenantDomain string
	tenantOnly         bool
	wantErr            bool
}

func TestParseClientID(t *testing.T) {
	tests := []ClientConfigTest{
		{name: "base", input: "foobar@insulator.conductor.one/pcc", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "dev", input: "foobar@c1dev.anthony.dev.ductone.com/pcc", wantedURL: "https://c1dev.anthony.dev.ductone.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseClientID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseClientID returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			testClientConfig(t, tt, got)
		})
	}
}

func TestCustomTenantOptions(t *testing.T) {
	tests := []ClientConfigTest{
		{name: "base", input: "insulator.conductor.one", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "tenantOnly", input: "insulator", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "baseWithHTTP", input: "https://insulator.conductor.one", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "differentURL", input: "https://insulator.a.different.com", wantedURL: "https://insulator.a.different.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
		{name: "differentBase", input: "insulator.a.different.com", wantedURL: "https://insulator.a.different.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
		{name: "devURL", input: "https://c1dev.anthony.dev.ductone.com:2443", wantedURL: "https://c1dev.anthony.dev.ductone.com:2443", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithTenantCustom(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithTenant returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			customOptions := &CustomOptions{}
			got(customOptions)
			testClientConfig(t, tt, &customOptions.ClientConfig)
		})
	}
}

func testClientConfig(t *testing.T, tt ClientConfigTest, got *ClientConfig) {
	if got.useWithServer == got.useWithTenant {
		t.Error("useWithServer and useWithTenant should not be the same")
	}
	if got.useWithTenant != tt.tenantOnly {
		t.Errorf("useWithTenant Got = '%v', want '%v'", got.useWithTenant, tt.tenantOnly)
	}
	if got.useWithServer {
		if got.ServerURL != tt.wantedURL {
			t.Errorf("ServerURL Got = '%v', want '%v'", got.ServerURL, tt.wantedURL)
		}
		if got.Tenant != tt.wantedTenantDomain {
			t.Errorf("Tenant Got = '%v', want '%v'", got.Tenant, "")
		}
	}
	if got.useWithTenant {
		if got.Tenant != tt.wantedTenantDomain {
			t.Errorf("Tenant Got = '%v', want '%v'", got.Tenant, tt.wantedTenantDomain)
		}
		if got.ServerURL != "" {
			t.Errorf("ServerURL Got = '%v', want '%v'", got.ServerURL, "")
		}
		if got.GetServerURL() != tt.wantedURL {
			t.Errorf("GetServerURL Got = '%v', want '%v'", got.GetServerURL(), tt.wantedURL)
		}
	}
}
