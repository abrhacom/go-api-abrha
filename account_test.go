package go_api_abrha

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{ "account": {
			"vm_limit": 25,
			"floating_ip_limit": 25,
			"reserved_ip_limit": 25,
			"volume_limit": 22,
			"email": "sammy@abrha.com",
			"name": "Sammy the Shark",
			"uuid": "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
			"email_verified": true
			}
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Account.Get(ctx)
	if err != nil {
		t.Errorf("Account.Get returned error: %v", err)
	}

	expected := &Account{
		VmLimit:         25,
		FloatingIPLimit: 25,
		ReservedIPLimit: 25,
		Email:           "sammy@abrha.com",
		Name:            "Sammy the Shark",
		UUID:            "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		EmailVerified:   true,
		VolumeLimit:     22,
	}
	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Account.Get returned %+v, expected %+v", acct, expected)
	}
}

func TestAccountString(t *testing.T) {
	acct := &Account{
		VmLimit:         25,
		FloatingIPLimit: 25,
		ReservedIPLimit: 25,
		VolumeLimit:     22,
		Email:           "sammy@abrha.com",
		Name:            "Sammy the Shark",
		UUID:            "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		EmailVerified:   true,
		Status:          "active",
		StatusMessage:   "message",
	}

	stringified := acct.String()
	expected := `go_api_abrha.Account{VmLimit:25, FloatingIPLimit:25, ReservedIPLimit:25, VolumeLimit:22, Email:"sammy@abrha.com", Name:"Sammy the Shark", UUID:"b6fr89dbf6d9156cace5f3c78dc9851d957381ef", EmailVerified:true, Status:"active", StatusMessage:"message"}`
	if expected != stringified {
		t.Errorf("\n     got %+v\nexpected %+v", stringified, expected)
	}

}

func TestAccountGetWithTeam(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{ "account": {
			"vm_limit": 25,
			"floating_ip_limit": 25,
			"volume_limit": 22,
			"email": "sammy@abrha.com",
			"name": "Sammy the Shark",
			"uuid": "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
			"email_verified": true,
			"team": {
				"name": "My Team",
				"uuid": "b6fr89dbf6d9156cace5f3c78dc9851d957381ef"
			}
			}
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Account.Get(ctx)
	if err != nil {
		t.Errorf("Account.Get returned error: %v", err)
	}

	expected := &Account{
		VmLimit:         25,
		FloatingIPLimit: 25,
		Email:           "sammy@abrha.com",
		Name:            "Sammy the Shark",
		UUID:            "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		EmailVerified:   true,
		VolumeLimit:     22,
		Team: &TeamInfo{
			Name: "My Team",
			UUID: "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		},
	}
	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("Account.Get returned %+v, expected %+v", acct, expected)
	}
}

func TestAccountStringWithTeam(t *testing.T) {
	acct := &Account{
		VmLimit:         25,
		FloatingIPLimit: 25,
		ReservedIPLimit: 25,
		VolumeLimit:     22,
		Email:           "sammy@abrha.com",
		Name:            "Sammy the Shark",
		UUID:            "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		EmailVerified:   true,
		Status:          "active",
		StatusMessage:   "message",
		Team: &TeamInfo{
			Name: "My Team",
			UUID: "b6fr89dbf6d9156cace5f3c78dc9851d957381ef",
		},
	}

	stringified := acct.String()
	expected := `go_api_abrha.Account{VmLimit:25, FloatingIPLimit:25, ReservedIPLimit:25, VolumeLimit:22, Email:"sammy@abrha.com", Name:"Sammy the Shark", UUID:"b6fr89dbf6d9156cace5f3c78dc9851d957381ef", EmailVerified:true, Status:"active", StatusMessage:"message", Team:go_api_abrha.TeamInfo{Name:"My Team", UUID:"b6fr89dbf6d9156cace5f3c78dc9851d957381ef"}}`
	if expected != stringified {
		t.Errorf("\n     got %+v\nexpected %+v", stringified, expected)
	}

}
