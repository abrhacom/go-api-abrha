package go_api_abrha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestVmActions_Shutdown(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "shutdown",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.Shutdown(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.Shutdown returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Shutdown returned %+v, expected %+v", action, expected)
	}
}

func TestVmActions_ShutdownByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "shutdown",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.ShutdownByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.ShutdownByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.ShutdownByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.ShutdownByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerOff(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_off",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.PowerOff(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.PowerOff returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Poweroff returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerOffByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_off",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.PowerOffByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.PowerOffByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.PowerOffByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PoweroffByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerOn(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_on",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.PowerOn(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.PowerOn returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PowerOn returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerOnByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_on",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.PowerOnByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.PowerOnByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.PowerOnByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PowerOnByTag returned %+v, expected %+v", action, expected)
	}
}
func TestVmAction_Reboot(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "reboot",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)

	})

	action, _, err := client.VmActions.Reboot(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.Reboot returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Reboot returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_Restore(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type":  "restore",
		"image": float64(1),
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)

	})

	action, _, err := client.VmActions.Restore(ctx, "1", 1)
	if err != nil {
		t.Errorf("VmActions.Restore returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Restore returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_Resize(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "resize",
		"size": "1024mb",
		"disk": true,
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)

	})

	action, _, err := client.VmActions.Resize(ctx, "1", "1024mb", true)
	if err != nil {
		t.Errorf("VmActions.Resize returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Resize returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_Rename(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "rename",
		"name": "Vm-Name",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.Rename(ctx, "1", "Vm-Name")
	if err != nil {
		t.Errorf("VmActions.Rename returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Rename returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerCycle(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_cycle",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)

	})

	action, _, err := client.VmActions.PowerCycle(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.PowerCycle returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PowerCycle returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PowerCycleByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "power_cycle",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.PowerCycleByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.PowerCycleByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.PowerCycleByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PowerCycleByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_Snapshot(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "snapshot",
		"name": "Image-Name",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.Snapshot(ctx, "1", "Image-Name")
	if err != nil {
		t.Errorf("VmActions.Snapshot returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Snapshot returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_SnapshotByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "snapshot",
		"name": "Image-Name",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.SnapshotByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.SnapshotByTag(ctx, "testing-1", "Image-Name")
	if err != nil {
		t.Errorf("VmActions.SnapshotByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.SnapshotByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnableBackups(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_backups",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.EnableBackups(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.EnableBackups returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableBackups returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnableBackupsByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_backups",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.EnableBackupByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.EnableBackupsByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.EnableBackupsByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableBackupsByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnableBackupsWithPolicy(t *testing.T) {
	setup()
	defer teardown()

	policyRequest := &VmBackupPolicyRequest{
		Plan:    "weekly",
		Weekday: "TUE",
		Hour:    PtrTo(20),
	}

	policy := map[string]interface{}{
		"hour":     float64(20),
		"plan":     "weekly",
		"weekday":  "TUE",
		"monthday": float64(0),
	}

	request := &ActionRequest{
		"type":          "enable_backups",
		"backup_policy": policy,
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.EnableBackupsWithPolicy(ctx, "1", policyRequest)
	if err != nil {
		t.Errorf("VmActions.EnableBackups returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableBackups returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_ChangeBackupPolicy(t *testing.T) {
	setup()
	defer teardown()

	policyRequest := &VmBackupPolicyRequest{
		Plan:    "weekly",
		Weekday: "SUN",
		Hour:    PtrTo(0),
	}

	policy := map[string]interface{}{
		"hour":     float64(0),
		"plan":     "weekly",
		"weekday":  "SUN",
		"monthday": float64(0),
	}

	request := &ActionRequest{
		"type":          "change_backup_policy",
		"backup_policy": policy,
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.ChangeBackupPolicy(ctx, "1", policyRequest)
	if err != nil {
		t.Errorf("VmActions.EnableBackups returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableBackups returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_DisableBackups(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "disable_backups",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.DisableBackups(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.DisableBackups returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.DisableBackups returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_DisableBackupsByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "disable_backups",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.DisableBackupsByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.DisableBackupsByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.DisableBackupsByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.DisableBackupsByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_PasswordReset(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "password_reset",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.PasswordReset(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.PasswordReset returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.PasswordReset returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_RebuildByImageID(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type":  "rebuild",
		"image": float64(2),
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = \n%#v, expected \n%#v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.RebuildByImageID(ctx, "1", 2)
	if err != nil {
		t.Errorf("VmActions.RebuildByImageID returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.RebuildByImageID returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_RebuildByImageSlug(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type":  "rebuild",
		"image": "Image-Name",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.RebuildByImageSlug(ctx, "1", "Image-Name")
	if err != nil {
		t.Errorf("VmActions.RebuildByImageSlug returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.RebuildByImageSlug returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_ChangeKernel(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type":   "change_kernel",
		"kernel": float64(2),
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.ChangeKernel(ctx, "1", 2)
	if err != nil {
		t.Errorf("VmActions.ChangeKernel returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.ChangeKernel returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnableIPv6(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_ipv6",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.EnableIPv6(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.EnableIPv6 returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableIPv6 returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnableIPv6ByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_ipv6",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.EnableIPv6ByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.EnableIPv6ByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.EnableIPv6ByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnableIPv6byTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnablePrivateNetworking(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_private_networking",
	}

	mux.HandleFunc("/api/public/v1/vms/1/actions", func(w http.ResponseWriter, r *http.Request) {
		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.EnablePrivateNetworking(ctx, "1")
	if err != nil {
		t.Errorf("VmActions.EnablePrivateNetworking returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnablePrivateNetworking returned %+v, expected %+v", action, expected)
	}
}

func TestVmAction_EnablePrivateNetworkingByTag(t *testing.T) {
	setup()
	defer teardown()

	request := &ActionRequest{
		"type": "enable_private_networking",
	}

	mux.HandleFunc("/api/public/v1/vms/actions", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("VmActions.EnablePrivateNetworkingByTag did not request with a tag parameter")
		}

		v := new(ActionRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, http.MethodPost)

		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}

		fmt.Fprint(w, `{"actions": [{"status":"in-progress"},{"status":"in-progress"}]}`)
	})

	action, _, err := client.VmActions.EnablePrivateNetworkingByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("VmActions.EnablePrivateNetworkingByTag returned error: %v", err)
	}

	expected := []Action{{Status: "in-progress"}, {Status: "in-progress"}}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.EnablePrivateNetworkingByTag returned %+v, expected %+v", action, expected)
	}
}

func TestVmActions_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/123/actions/456", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `{"action":{"status":"in-progress"}}`)
	})

	action, _, err := client.VmActions.Get(ctx, "123", 456)
	if err != nil {
		t.Errorf("VmActions.Get returned error: %v", err)
	}

	expected := &Action{Status: "in-progress"}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("VmActions.Get returned %+v, expected %+v", action, expected)
	}
}
