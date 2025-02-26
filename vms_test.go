package go_api_abrha

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVms_ListVms(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"vms": [
				{
					"id": "1"
				},
				{
					"id": "2"
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	vms, resp, err := client.Vms.List(ctx, nil)
	if err != nil {
		t.Errorf("Vms.List returned error: %v", err)
	}

	expectedVms := []Vm{{ID: "1"}, {ID: "2"}}
	if !reflect.DeepEqual(vms, expectedVms) {
		t.Errorf("Vms.List\nVms: got=%#v\nwant=%#v", vms, expectedVms)
	}
	expectedMeta := &Meta{Total: 2}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.List\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
	}
}

func TestVms_ListVmsWithGPUs(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		if r.URL.Query().Get("type") != "gpus" {
			t.Errorf("Vms.ListWithGPUs did not request with a type parameter")
		}
		fmt.Fprint(w, `{
			"vms": [
				{
					"id": "1",
					"size": {
						"gpu_info": {
							"count": 1,
							"vram": {
								"amount": 8,
								"unit": "gib"
							},
							"model": "nvidia_tesla_v100"
						},
						"disk_info": [
							{
								"type": "local",
								"size": {
									"amount": 200,
									"unit": "gib"
								}
							},
							{
								"type": "scratch",
								"size": {
									"amount": 40960,
									"unit": "gib"
								}
							}
						]
					}
				}
			],
			"meta": {
				"total": 1
			}
		}`)
	})

	vms, resp, err := client.Vms.ListWithGPUs(ctx, nil)
	if err != nil {
		t.Errorf("Vms.List returned error: %v", err)
	}

	expectedVms := []Vm{
		{
			ID: "1",
			Size: &Size{
				GPUInfo: &GPUInfo{
					Count: 1,
					VRAM: &VRAM{
						Amount: 8,
						Unit:   "gib",
					},
					Model: "nvidia_tesla_v100",
				},
				DiskInfo: []DiskInfo{
					{
						Type: "local",
						Size: &DiskSize{
							Amount: 200,
							Unit:   "gib",
						},
					},
					{
						Type: "scratch",
						Size: &DiskSize{
							Amount: 40960,
							Unit:   "gib",
						},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(vms, expectedVms) {
		t.Errorf("Vms.List\nVms: got=%#v\nwant=%#v", vms, expectedVms)
	}
	expectedMeta := &Meta{Total: 1}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.List\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
	}
}

func TestVms_ListVmsByTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("Vms.ListByTag did not request with a tag parameter")
		}

		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"vms": [
				{
					"id": "1"
				},
				{
					"id": "2"
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	vms, resp, err := client.Vms.ListByTag(ctx, "testing-1", nil)
	if err != nil {
		t.Errorf("Vms.ListByTag returned error: %v", err)
	}

	expectedVms := []Vm{{ID: "1"}, {ID: "2"}}
	if !reflect.DeepEqual(vms, expectedVms) {
		t.Errorf("Vms.ListByTag returned vms %+v, expected %+v", vms, expectedVms)
	}
	expectedMeta := &Meta{Total: 2}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.ListByTag returned meta %+v, expected %+v", resp.Meta, expectedMeta)
	}
}

func TestVms_ListVmsByName(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		name := "testing"
		if r.URL.Query().Get("name") != name {
			t.Errorf("Vms.ListByName request did not contain the 'name=%s' query parameter", name)
		}

		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"vms": [
				{
					"id": "1",
					"name": "testing"
				},
				{
					"id": "2",
					"name": "testing"
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	vms, _, err := client.Vms.ListByName(ctx, "testing", nil)
	if err != nil {
		t.Errorf("Vms.ListByTag returned error: %v", err)
	}

	expected := []Vm{{ID: "1", Name: "testing"}, {ID: "2", Name: "testing"}}
	if !reflect.DeepEqual(vms, expected) {
		t.Errorf("Vms.ListByTag returned vms %+v, expected %+v", vms, expected)
	}
}

func TestVms_ListVmsMultiplePages(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		dr := vmsRoot{
			Vms: []Vm{
				{ID: "1"},
				{ID: "2"},
			},
			Links: &Links{
				Pages: &Pages{Next: "http://example.com/v1/vms/?page=2"},
			},
		}

		b, err := json.Marshal(dr)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Fprint(w, string(b))
	})

	_, resp, err := client.Vms.List(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 1)
}

func TestVms_RetrievePageByNumber(t *testing.T) {
	setup()
	defer teardown()

	jBlob := `
	{
		"vms": [{"id":"1"},{"id":"2"}],
		"links":{
			"pages":{
				"next":"http://example.com/v1/vms/?page=3",
				"prev":"http://example.com/v1/vms/?page=1",
				"last":"http://example.com/v1/vms/?page=3",
				"first":"http://example.com/v1/vms/?page=1"
			}
		}
	}`

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, jBlob)
	})

	opt := &ListOptions{Page: 2}
	_, resp, err := client.Vms.List(ctx, opt)
	if err != nil {
		t.Fatal(err)
	}

	checkCurrentPage(t, resp, 2)
}

func TestVms_GetVm(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"vm":{"id":"12345"}}`)
	})

	vms, _, err := client.Vms.Get(ctx, "12345")
	if err != nil {
		t.Errorf("Vm.Get returned error: %v", err)
	}

	expected := &Vm{ID: "12345"}
	if !reflect.DeepEqual(vms, expected) {
		t.Errorf("Vms.Get\n got=%#v\nwant=%#v", vms, expected)
	}
}

func TestVms_Create(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &VmCreateRequest{
		Name:   "name",
		Region: "region",
		Size:   "size",
		Image: VmCreateImage{
			ID: 1,
		},
		Volumes: []VmCreateVolume{
			{ID: "hello-im-another-volume"},
			{Name: "should be ignored due to Name", ID: "aaa-111-bbb-222-ccc"},
		},
		Tags:    []string{"one", "two"},
		VPCUUID: "880b7f98-f062-404d-b33c-458d545696f6",
		Backups: true,
		BackupPolicy: &VmBackupPolicyRequest{
			Plan:    "weekly",
			Weekday: "MON",
			Hour:    PtrTo(0),
		},
	}

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name":               "name",
			"region":             "region",
			"size":               "size",
			"image":              float64(1),
			"ssh_keys":           nil,
			"ipv6":               false,
			"private_networking": false,
			"monitoring":         false,
			"volumes": []interface{}{
				map[string]interface{}{"id": "hello-im-another-volume"},
				map[string]interface{}{"id": "aaa-111-bbb-222-ccc"},
			},
			"tags":          []interface{}{"one", "two"},
			"vpc_uuid":      "880b7f98-f062-404d-b33c-458d545696f6",
			"backups":       true,
			"backup_policy": map[string]interface{}{"plan": "weekly", "weekday": "MON", "hour": float64(0)},
		}
		jsonBlob := `
{
  "vm": {
    "id": "1",
    "vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6"
  },
  "links": {
    "actions": [
      {
        "id": 1,
        "href": "http://example.com",
        "rel": "create"
      }
    ]
  }
}
`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprintf(w, jsonBlob)
	})

	vmRoot, resp, err := client.Vms.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Vms.Create returned error: %v", err)
	}

	if id := vmRoot.Vm.ID; id != "1" {
		t.Errorf("expected id '%d', received '%s'", 1, id)
	}

	vpcid := "880b7f98-f062-404d-b33c-458d545696f6"
	if id := vmRoot.Vm.VPCUUID; id != vpcid {
		t.Errorf("expected VPC uuid '%s', received '%s'", vpcid, id)
	}

	if a := resp.Links.Actions[0]; a.ID != 1 {
		t.Errorf("expected action id '%d', received '%d'", 1, a.ID)
	}
}

func TestVms_CreateWithoutVmAgent(t *testing.T) {
	setup()
	defer teardown()

	boolVal := false
	createRequest := &VmCreateRequest{
		Name:   "name",
		Region: "region",
		Size:   "size",
		Image: VmCreateImage{
			ID: 1,
		},
		Volumes: []VmCreateVolume{
			{ID: "hello-im-another-volume"},
			{Name: "should be ignored due to Name", ID: "aaa-111-bbb-222-ccc"},
		},
		Tags:        []string{"one", "two"},
		VPCUUID:     "880b7f98-f062-404d-b33c-458d545696f6",
		WithVmAgent: &boolVal,
	}

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name":               "name",
			"region":             "region",
			"size":               "size",
			"image":              float64(1),
			"ssh_keys":           nil,
			"backups":            false,
			"ipv6":               false,
			"private_networking": false,
			"monitoring":         false,
			"volumes": []interface{}{
				map[string]interface{}{"id": "hello-im-another-volume"},
				map[string]interface{}{"id": "aaa-111-bbb-222-ccc"},
			},
			"tags":          []interface{}{"one", "two"},
			"vpc_uuid":      "880b7f98-f062-404d-b33c-458d545696f6",
			"with_vm_agent": false,
		}
		jsonBlob := `
{
  "vm": {
    "id": "1",
    "vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6"
  },
  "links": {
    "actions": [
      {
        "id": 1,
        "href": "http://example.com",
        "rel": "create"
      }
    ]
  }
}
`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprintf(w, jsonBlob)
	})

	vmRoot, resp, err := client.Vms.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Vms.Create returned error: %v", err)
	}

	if id := vmRoot.Vm.ID; id != "1" {
		t.Errorf("expected id '%d', received '%s'", 1, id)
	}

	vpcid := "880b7f98-f062-404d-b33c-458d545696f6"
	if id := vmRoot.Vm.VPCUUID; id != vpcid {
		t.Errorf("expected VPC uuid '%s', received '%s'", vpcid, id)
	}

	if a := resp.Links.Actions[0]; a.ID != 1 {
		t.Errorf("expected action id '%d', received '%d'", 1, a.ID)
	}
}

func TestVms_WithVmAgentJsonMarshal(t *testing.T) {
	boolF := false
	boolT := true
	tests := []struct {
		in   *VmCreateRequest
		want string
	}{
		{
			in:   &VmCreateRequest{Name: "foo", WithVmAgent: &boolF},
			want: `{"name":"foo","region":"","size":"","image":0,"ssh_keys":null,"backups":false,"ipv6":false,"private_networking":false,"monitoring":false,"tags":null,"with_vm_agent":false}`,
		},
		{
			in:   &VmCreateRequest{Name: "foo", WithVmAgent: &boolT},
			want: `{"name":"foo","region":"","size":"","image":0,"ssh_keys":null,"backups":false,"ipv6":false,"private_networking":false,"monitoring":false,"tags":null,"with_vm_agent":true}`,
		},
		{
			in:   &VmCreateRequest{Name: "foo"},
			want: `{"name":"foo","region":"","size":"","image":0,"ssh_keys":null,"backups":false,"ipv6":false,"private_networking":false,"monitoring":false,"tags":null}`,
		},
	}

	for _, tt := range tests {
		got, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if !reflect.DeepEqual(tt.want, string(got)) {
			t.Errorf("expected: %v, got: %v", tt.want, string(got))
		}
	}
}

func TestVms_CreateWithDisabledPublicNetworking(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &VmCreateRequest{
		Name:   "name",
		Region: "region",
		Size:   "size",
		Image: VmCreateImage{
			ID: 1,
		},
		Volumes: []VmCreateVolume{
			{ID: "hello-im-another-volume"},
			{Name: "should be ignored due to Name", ID: "aaa-111-bbb-222-ccc"},
		},
		Tags:    []string{"one", "two"},
		VPCUUID: "880b7f98-f062-404d-b33c-458d545696f6",
	}

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"name":               "name",
			"region":             "region",
			"size":               "size",
			"image":              float64(1),
			"ssh_keys":           nil,
			"backups":            false,
			"ipv6":               false,
			"private_networking": false,
			"monitoring":         false,
			"volumes": []interface{}{
				map[string]interface{}{"id": "hello-im-another-volume"},
				map[string]interface{}{"id": "aaa-111-bbb-222-ccc"},
			},
			"tags":     []interface{}{"one", "two"},
			"vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6",
		}
		jsonBlob := `
{
  "vm": {
    "id": "1",
    "vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6"
  },
  "links": {
    "actions": [
      {
        "id": 1,
        "href": "http://example.com",
        "rel": "create"
      }
    ]
  }
}
`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body\n got=%#v\nwant=%#v", v, expected)
		}

		fmt.Fprintf(w, jsonBlob)
	})

	vmRoot, _, err := client.Vms.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("Vms.Create returned error: %v", err)
	}

	if id := vmRoot.Vm.ID; id != "1" {
		t.Errorf("expected id '%d', received '%s'", 1, id)
	}
}

func TestVms_CreateMultiple(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &VmMultiCreateRequest{
		Names:  []string{"name1", "name2"},
		Region: "region",
		Size:   "size",
		Image: VmCreateImage{
			ID: 1,
		},
		Tags:    []string{"one", "two"},
		VPCUUID: "880b7f98-f062-404d-b33c-458d545696f6",
	}

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		expected := map[string]interface{}{
			"names":              []interface{}{"name1", "name2"},
			"region":             "region",
			"size":               "size",
			"image":              float64(1),
			"ssh_keys":           nil,
			"backups":            false,
			"ipv6":               false,
			"private_networking": false,
			"monitoring":         false,
			"tags":               []interface{}{"one", "two"},
			"vpc_uuid":           "880b7f98-f062-404d-b33c-458d545696f6",
		}
		jsonBlob := `
{
  "vms": [
    {
      "id": "1",
	  "vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6"
    },
    {
      "id": "2",
	  "vpc_uuid": "880b7f98-f062-404d-b33c-458d545696f6"
    }
  ],
  "links": {
    "actions": [
      {
        "id": 1,
        "href": "http://example.com",
        "rel": "multiple_create"
      }
    ]
  }
}
`

		var v map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		if !reflect.DeepEqual(v, expected) {
			t.Errorf("Request body = %#v, expected %#v", v, expected)
		}

		fmt.Fprintf(w, jsonBlob)
	})

	vms, resp, err := client.Vms.CreateMultiple(ctx, createRequest)
	if err != nil {
		t.Errorf("Vms.CreateMultiple returned error: %v", err)
	}

	if id := vms[0].ID; id != "1" {
		t.Errorf("expected id '%d', received '%s'", 1, id)
	}
	if id := vms[1].ID; id != "2" {
		t.Errorf("expected id '%d', received '%s'", 2, id)
	}

	vpcid := "880b7f98-f062-404d-b33c-458d545696f6"
	if id := vms[0].VPCUUID; id != vpcid {
		t.Errorf("expected VPC uuid '%s', received '%s'", vpcid, id)
	}
	if id := vms[1].VPCUUID; id != vpcid {
		t.Errorf("expected VPC uuid '%s', received '%s'", vpcid, id)
	}

	if a := resp.Links.Actions[0]; a.ID != 1 {
		t.Errorf("expected action id '%d', received '%d'", 1, a.ID)
	}
}

func TestVms_Destroy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Vms.Delete(ctx, "12345")
	if err != nil {
		t.Errorf("Vm.Delete returned error: %v", err)
	}
}

func TestVms_DestroyByTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("tag_name") != "testing-1" {
			t.Errorf("Vms.DeleteByTag did not request with a tag parameter")
		}

		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Vms.DeleteByTag(ctx, "testing-1")
	if err != nil {
		t.Errorf("Vm.Delete returned error: %v", err)
	}
}

//func TestVms_Kernels(t *testing.T) {
//	setup()
//	defer teardown()
//
//	mux.HandleFunc("/api/public/v1/vms/12345/kernels", func(w http.ResponseWriter, r *http.Request) {
//		testMethod(t, r, http.MethodGet)
//		fmt.Fprint(w, `{
//			"kernels": [
//				{
//					"id": "1"
//				},
//				{
//					"id": "2"
//				}
//			],
//			"meta": {
//				"total": 2
//			}
//		}`)
//	})
//
//	opt := &ListOptions{Page: 2}
//	kernels, resp, err := client.Vms.Kernels(ctx, "12345", opt)
//	if err != nil {
//		t.Errorf("Vms.Kernels returned error: %v", err)
//	}
//
//	expectedKernels := []Kernel{{ID: 1}, {ID: 2}}
//	if !reflect.DeepEqual(kernels, expectedKernels) {
//		t.Errorf("Vms.Kernels\nKernels got=%#v\nwant=%#v", kernels, expectedKernels)
//	}
//	expectedMeta := &Meta{Total: 2}
//	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
//		t.Errorf("Vms.Kernels\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
//	}
//}

func TestVms_Snapshots(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345/snapshots", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"snapshots": [
				{
					"id": 1
				},
				{
					"id": 2
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	opt := &ListOptions{Page: 2}
	snapshots, resp, err := client.Vms.Snapshots(ctx, "12345", opt)
	if err != nil {
		t.Errorf("Vms.Snapshots returned error: %v", err)
	}

	expectedSnapshots := []Image{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(snapshots, expectedSnapshots) {
		t.Errorf("Vms.Snapshots\nSnapshots got=%#v\nwant=%#v", snapshots, expectedSnapshots)
	}
	expectedMeta := &Meta{Total: 2}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.Snapshots\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
	}
}

func TestVms_Backups(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345/backups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"backups": [
				{
					"id": 1
				},
				{
					"id": 2
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	opt := &ListOptions{Page: 2}
	backups, resp, err := client.Vms.Backups(ctx, "12345", opt)
	if err != nil {
		t.Errorf("Vms.Backups returned error: %v", err)
	}

	expectedBackups := []Image{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(backups, expectedBackups) {
		t.Errorf("Vms.Backups\nBackups got=%#v\nwant=%#v", backups, expectedBackups)
	}
	expectedMeta := &Meta{Total: 2}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.Backups\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
	}
}

func TestVms_Actions(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345/actions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"actions": [
				{
					"id": 1
				},
				{
					"id": 2
				}
			],
			"meta": {
				"total": 2
			}
		}`)
	})

	opt := &ListOptions{Page: 2}
	actions, resp, err := client.Vms.Actions(ctx, "12345", opt)
	if err != nil {
		t.Errorf("Vms.Actions returned error: %v", err)
	}

	expectedActions := []Action{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(actions, expectedActions) {
		t.Errorf("Vms.Actions\nActions got=%#v\nwant=%#v", actions, expectedActions)
	}
	expectedMeta := &Meta{Total: 2}
	if !reflect.DeepEqual(resp.Meta, expectedMeta) {
		t.Errorf("Vms.Actions\nMeta: got=%#v\nwant=%#v", resp.Meta, expectedMeta)
	}
}

func TestVms_Neighbors(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345/neighbors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"vms": [{"id":"1"},{"id":"2"}]}`)
	})

	neighbors, _, err := client.Vms.Neighbors(ctx, "12345")
	if err != nil {
		t.Errorf("Vms.Neighbors returned error: %v", err)
	}

	expected := []Vm{{ID: "1"}, {ID: "2"}}
	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Vms.Neighbors\n got=%#v\nwant=%#v", neighbors, expected)
	}
}

func TestNetworkV4_String(t *testing.T) {
	network := &NetworkV4{
		IPAddress: "192.168.1.2",
		Netmask:   "255.255.255.0",
		Gateway:   "192.168.1.1",
	}

	stringified := network.String()
	expected := `go_api_abrha.NetworkV4{IPAddress:"192.168.1.2", Netmask:"255.255.255.0", Gateway:"192.168.1.1", Type:""}`
	if expected != stringified {
		t.Errorf("NetworkV4.String\n got=%#v\nwant=%#v", stringified, expected)
	}

}

func TestNetworkV6_String(t *testing.T) {
	network := &NetworkV6{
		IPAddress: "2604:A880:0800:0010:0000:0000:02DD:4001",
		Netmask:   64,
		Gateway:   "2604:A880:0800:0010:0000:0000:0000:0001",
	}
	stringified := network.String()
	expected := `go_api_abrha.NetworkV6{IPAddress:"2604:A880:0800:0010:0000:0000:02DD:4001", Netmask:64, Gateway:"2604:A880:0800:0010:0000:0000:0000:0001", Type:""}`
	if expected != stringified {
		t.Errorf("NetworkV6.String\n got=%#v\nwant=%#v", stringified, expected)
	}
}

func TestVms_IPMethods(t *testing.T) {
	var d Vm

	ipv6 := "1000:1000:1000:1000:0000:0000:004D:B001"

	d.Networks = &Networks{
		V4: []NetworkV4{
			{IPAddress: "192.168.0.1", Type: "public"},
			{IPAddress: "10.0.0.1", Type: "private"},
		},
		V6: []NetworkV6{
			{IPAddress: ipv6, Type: "public"},
		},
	}

	ip, err := d.PublicIPv4()
	if err != nil {
		t.Errorf("unknown error")
	}

	if got, expected := ip, "192.168.0.1"; got != expected {
		t.Errorf("Vm.PublicIPv4 returned %s; expected %s", got, expected)
	}

	ip, err = d.PrivateIPv4()
	if err != nil {
		t.Errorf("unknown error")
	}

	if got, expected := ip, "10.0.0.1"; got != expected {
		t.Errorf("Vm.PrivateIPv4 returned %s; expected %s", got, expected)
	}

	ip, err = d.PublicIPv6()
	if err != nil {
		t.Errorf("unknown error")
	}

	if got, expected := ip, ipv6; got != expected {
		t.Errorf("Vm.PublicIPv6 returned %s; expected %s", got, expected)
	}
}

func TestVms_GetBackupPolicy(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/public/v1/vms/12345/backups/policy", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
				"policy": {
					"vm_id": "12345",
					"backup_enabled": true,
					"backup_policy": {
					"plan": "weekly",
					"weekday": "SUN",
					"hour": 0,
					"monthday": 0,
					"window_length_hours": 4,
					"retention_period_days": 28
					},
					"next_backup_window": {
						"start": "2021-01-01T00:00:00Z",
						"end": "2021-01-01T00:00:00Z"
					}
				}
			}`)
	})

	policy, _, err := client.Vms.GetBackupPolicy(ctx, "12345")
	if err != nil {
		t.Errorf("Vms.GetBackupPolicy returned error: %v", err)
	}

	pt, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expected := &VmBackupPolicy{
		VmID:          "12345",
		BackupEnabled: true,
		BackupPolicy: &VmBackupPolicyConfig{
			Plan:                "weekly",
			Weekday:             "SUN",
			Hour:                0,
			WindowLengthHours:   4,
			RetentionPeriodDays: 28,
		},
		NextBackupWindow: &BackupWindow{
			Start: &Timestamp{Time: pt},
			End:   &Timestamp{Time: pt},
		},
	}
	if !reflect.DeepEqual(policy, expected) {
		t.Errorf("Vms.GetBackupPolicy\n got=%#v\nwant=%#v", policy, expected)
	}
}

func TestVms_ListBackupPolicies(t *testing.T) {
	setup()
	defer teardown()

	ctx := context.Background()
	policyID := 123
	pt, err := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	testBackupPolicy := VmBackupPolicy{
		VmID:          "12345",
		BackupEnabled: true,
		BackupPolicy: &VmBackupPolicyConfig{
			Plan:                "weekly",
			Weekday:             "SUN",
			Hour:                0,
			WindowLengthHours:   4,
			RetentionPeriodDays: 28,
		},
		NextBackupWindow: &BackupWindow{
			Start: &Timestamp{Time: pt},
			End:   &Timestamp{Time: pt},
		},
	}

	mux.HandleFunc("/api/public/v1/vms/backups/policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(&vmBackupPoliciesRoot{
			VmBackupPolicies: map[int]*VmBackupPolicy{policyID: &testBackupPolicy},
			Meta:             &Meta{Total: 1},
			Links:            &Links{},
		})
	})

	policies, _, err := client.Vms.ListBackupPolicies(ctx, &ListOptions{Page: 1})
	require.NoError(t, err)
	assert.Equal(t, map[int]*VmBackupPolicy{policyID: &testBackupPolicy}, policies)
}

func TestVms_ListSupportedBackupPolicies(t *testing.T) {
	setup()
	defer teardown()

	ctx := context.Background()
	testSupportedBackupPolicy := SupportedBackupPolicy{
		Name:                 "weekly",
		PossibleWindowStarts: []int{0, 4, 8, 12, 16, 20},
		WindowLengthHours:    4,
		RetentionPeriodDays:  28,
		PossibleDays:         []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"},
	}

	mux.HandleFunc("/api/public/v1/vms/backups/supported_policies", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(&vmSupportedBackupPoliciesRoot{SupportedBackupPolicies: []*SupportedBackupPolicy{&testSupportedBackupPolicy}})
	})

	policies, _, err := client.Vms.ListSupportedBackupPolicies(ctx)
	require.NoError(t, err)
	assert.Equal(t, []*SupportedBackupPolicy{&testSupportedBackupPolicy}, policies)
}
