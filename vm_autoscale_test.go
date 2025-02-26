package go_api_abrha

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var vmAutoscaleListHistoryJSONResponse = `
{
  "history": [
    {
      "history_event_id": "4344c45f-7574-493b-a96c-df805c65a900",
      "current_instance_count": 0,
      "desired_instance_count": 1,
      "reason": "configuration update",
      "status": "success",
      "created_at": "2024-10-18T19:03:09Z",
      "updated_at": "2024-10-18T19:03:09Z"
    },
    {
      "history_event_id": "9ad436f7-af57-49ff-b416-0043721055b2",
      "current_instance_count": 1,
      "desired_instance_count": 2,
      "reason": "scaling up (desired=2 current=1)",
      "status": "success",
      "created_at": "2024-10-18T19:15:24Z",
      "updated_at": "2024-10-18T19:15:24Z"
    },
    {
      "history_event_id": "45390191-d077-49e9-a3c4-c2eb903bc1a2",
      "current_instance_count": 2,
      "desired_instance_count": 1,
      "reason": "scaling down (desired=1 current=2)",
      "status": "success",
      "created_at": "2024-10-18T19:47:24Z",
      "updated_at": "2024-10-18T19:47:24Z"
    }
  ],
  "links": {},
  "meta": {
    "total": 3
  }
}
`

var vmAutoscaleListMembersJSONResponse = `
{
  "vms": [
    {
      "vm_id": "1677149",
      "created_at": "2024-10-18T19:03:09Z",
      "updated_at": "2024-10-18T19:03:24Z",
      "health_status": "healthy",
      "status": "active",
      "current_utilization": {
        "memory": 0.35,
        "cpu": 0.0012
      }
    },
    {
      "vm_id": "1677150",
      "created_at": "2024-10-18T19:04:09Z",
      "updated_at": "2024-10-18T19:04:24Z",
      "health_status": "healthy",
      "status": "active",
      "current_utilization": {
        "memory": 0.40,
        "cpu": 0.0013
      }
    }
  ],
  "links": {},
  "meta": {
    "total": 2
  }
}
`

var vmAutoscaleListJSONResponse = `
{
  "autoscale_pools": [
    {
      "id": "a4456a02-133d-4fea-8f2d-94dc6a7bf9c9",
      "name": "test-autoscalergroup-03",
      "config": {
        "min_instances": 1,
        "max_instances": 5,
        "target_cpu_utilization": 0.5,
        "cooldown_minutes": 5
      },
      "vm_template": {
        "size": "s-1vcpu-512mb-10gb",
        "region": "s2r1",
        "image": "547864",
        "tags": [
          "test-ag-01"
        ],
        "ssh_keys": [
          "372862",
          "367582",
          "355790"
        ],
        "vpc_uuid": "72b0812c-7535-4388-8507-5ad29b4487b3",
        "with_vm_agent": false,
        "project_id": "",
        "ipv6": true,
        "user_data": "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n"
      },
      "created_at": "2024-10-21T13:05:23Z",
      "updated_at": "2024-10-21T13:05:23Z",
      "current_utilization": {
        "memory": 0.33,
        "cpu": 0.0007
      },
      "status": "active"
    },
    {
      "id": "1044bfca-e490-44a1-aa1c-6f002daf6a13",
      "name": "test-autoscalergroup-01",
      "config": {
        "min_instances": 1,
        "max_instances": 5,
        "target_cpu_utilization": 0.5,
        "cooldown_minutes": 5
      },
      "vm_template": {
        "size": "s-1vcpu-512mb-10gb",
        "region": "s2r1",
        "image": "547864",
        "tags": [
          "test-ag-01"
        ],
        "ssh_keys": [
          "372862",
          "367582",
          "355790"
        ],
        "vpc_uuid": "72b0812c-7535-4388-8507-5ad29b4487b3",
        "with_vm_agent": false,
        "project_id": "",
        "ipv6": true,
        "user_data": "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n"
      },
      "created_at": "2024-10-18T19:03:08Z",
      "updated_at": "2024-10-18T19:03:08Z",
      "current_utilization": {
        "memory": 0.35,
        "cpu": 0.0009
      },
      "status": "active"
    },
    {
      "id": "b92962b5-26a5-4e63-a1d9-a0f5d44b4f23",
      "name": "test-autoscalergroup-02",
      "config": {
        "min_instances": 1,
        "max_instances": 5,
        "target_cpu_utilization": 0.5,
        "cooldown_minutes": 5
      },
      "vm_template": {
        "size": "s-1vcpu-512mb-10gb",
        "region": "s2r1",
        "image": "547864",
        "tags": [
          "test-ag-01"
        ],
        "ssh_keys": [
          "372862",
          "367582",
          "355790"
        ],
        "vpc_uuid": "72b0812c-7535-4388-8507-5ad29b4487b3",
        "with_vm_agent": false,
        "project_id": "",
        "ipv6": true,
        "user_data": "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n"
      },
      "created_at": "2024-10-21T13:05:12Z",
      "updated_at": "2024-10-21T13:05:12Z",
      "current_utilization": {
        "memory": 0.56,
        "cpu": 0.0002
      },
      "status": "active"
    }
  ],
  "links": {},
  "meta": {
    "total": 3
  }
}
`

var vmAutoscaleGetJSONResponse = `
{
  "autoscale_pool": {
    "id": "1044bfca-e490-44a1-aa1c-6f002daf6a13",
    "name": "test-autoscalergroup-01",
    "config": {
      "min_instances": 1,
      "max_instances": 5,
      "target_cpu_utilization": 0.5,
      "cooldown_minutes": 5
    },
    "vm_template": {
      "size": "s-1vcpu-512mb-10gb",
      "region": "s2r1",
      "image": "547864",
      "tags": [
        "test-ag-01"
      ],
      "ssh_keys": [
        "372862",
        "367582",
        "355790"
      ],
      "vpc_uuid": "72b0812c-7535-4388-8507-5ad29b4487b3",
      "with_vm_agent": false,
      "project_id": "",
      "ipv6": true,
      "user_data": "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n"
    },
    "created_at": "2024-10-18T19:03:08Z",
    "updated_at": "2024-10-18T19:03:08Z",
    "current_utilization": {
      "memory": 0.35,
      "cpu": 0.0008
    },
    "status": "active"
  }
}
`

func TestVmAutoscaler_Get(t *testing.T) {
	setup()
	defer teardown()

	autoscalePoolID := "1044bfca-e490-44a1-aa1c-6f002daf6a13"
	mux.HandleFunc(fmt.Sprintf("%s/%s", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, vmAutoscaleGetJSONResponse)
	})

	expectedPoolResp := &VmAutoscalePool{
		ID:   "1044bfca-e490-44a1-aa1c-6f002daf6a13",
		Name: "test-autoscalergroup-01",
		Config: &VmAutoscaleConfiguration{
			MinInstances:         1,
			MaxInstances:         5,
			TargetCPUUtilization: 0.5,
			CooldownMinutes:      5,
		},
		VmTemplate: &VmAutoscaleResourceTemplate{
			Size:     "s-1vcpu-512mb-10gb",
			Region:   "s2r1",
			Image:    "547864",
			Tags:     []string{"test-ag-01"},
			SSHKeys:  []string{"372862", "367582", "355790"},
			VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
			IPV6:     true,
			UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
		},
		CurrentUtilization: &VmAutoscaleResourceUtilization{
			Memory: 0.35,
			CPU:    0.0008,
		},
		Status: "active",
	}

	gotPoolResp, _, err := client.VmAutoscale.Get(ctx, autoscalePoolID)
	require.NoError(t, err)
	require.NotNil(t, gotPoolResp)
	expectedPoolResp.CreatedAt = gotPoolResp.CreatedAt
	expectedPoolResp.UpdatedAt = gotPoolResp.UpdatedAt
	assert.Equal(t, expectedPoolResp, gotPoolResp)
}

func TestVmAutoscaler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(vmAutoscaleBasePath, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, vmAutoscaleListJSONResponse)
	})

	expectedConfig := &VmAutoscaleConfiguration{
		MinInstances:         1,
		MaxInstances:         5,
		TargetCPUUtilization: 0.5,
		CooldownMinutes:      5,
	}
	expectedVmTemplate := &VmAutoscaleResourceTemplate{
		Size:     "s-1vcpu-512mb-10gb",
		Region:   "s2r1",
		Image:    "547864",
		Tags:     []string{"test-ag-01"},
		SSHKeys:  []string{"372862", "367582", "355790"},
		VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
		IPV6:     true,
		UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
	}
	expectedPoolsResp := []*VmAutoscalePool{
		{
			ID:         "1044bfca-e490-44a1-aa1c-6f002daf6a13",
			Name:       "test-autoscalergroup-01",
			Config:     expectedConfig,
			VmTemplate: expectedVmTemplate,
			CurrentUtilization: &VmAutoscaleResourceUtilization{
				Memory: 0.35,
				CPU:    0.0009,
			},
			Status: "active",
		},
		{
			ID:         "b92962b5-26a5-4e63-a1d9-a0f5d44b4f23",
			Name:       "test-autoscalergroup-02",
			Config:     expectedConfig,
			VmTemplate: expectedVmTemplate,
			CurrentUtilization: &VmAutoscaleResourceUtilization{
				Memory: 0.56,
				CPU:    0.0002,
			},
			Status: "active",
		},
		{
			ID:         "a4456a02-133d-4fea-8f2d-94dc6a7bf9c9",
			Name:       "test-autoscalergroup-03",
			Config:     expectedConfig,
			VmTemplate: expectedVmTemplate,
			CurrentUtilization: &VmAutoscaleResourceUtilization{
				Memory: 0.33,
				CPU:    0.0007,
			},
			Status: "active",
		},
	}

	gotPoolsResp, _, err := client.VmAutoscale.List(ctx, nil)
	require.NoError(t, err)
	require.NotEmpty(t, gotPoolsResp)
	sort.SliceStable(gotPoolsResp, func(i, j int) bool {
		return gotPoolsResp[i].Name < gotPoolsResp[j].Name
	})
	for idx := range gotPoolsResp {
		expectedPoolsResp[idx].CreatedAt = gotPoolsResp[idx].CreatedAt
		expectedPoolsResp[idx].UpdatedAt = gotPoolsResp[idx].UpdatedAt
	}
	assert.Equal(t, expectedPoolsResp, gotPoolsResp)
}

func TestVmAutoscaler_ListMembers(t *testing.T) {
	setup()
	defer teardown()

	autoscalePoolID := "1044bfca-e490-44a1-aa1c-6f002daf6a13"
	mux.HandleFunc(fmt.Sprintf("%s/%s/members", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, vmAutoscaleListMembersJSONResponse)
	})

	expectedMembersResp := []*VmAutoscaleResource{
		{
			VmID:         "1677149",
			HealthStatus: "healthy",
			Status:       "active",
			CurrentUtilization: &VmAutoscaleResourceUtilization{
				Memory: 0.35,
				CPU:    0.0012,
			},
		},
		{
			VmID:         "1677150",
			HealthStatus: "healthy",
			Status:       "active",
			CurrentUtilization: &VmAutoscaleResourceUtilization{
				Memory: 0.40,
				CPU:    0.0013,
			},
		},
	}

	gotMembersResp, _, err := client.VmAutoscale.ListMembers(ctx, autoscalePoolID, nil)
	require.NoError(t, err)
	require.NotEmpty(t, gotMembersResp)
	sort.SliceStable(gotMembersResp, func(i, j int) bool {
		return gotMembersResp[i].VmID < gotMembersResp[j].VmID
	})
	for idx := range gotMembersResp {
		expectedMembersResp[idx].CreatedAt = gotMembersResp[idx].CreatedAt
		expectedMembersResp[idx].UpdatedAt = gotMembersResp[idx].UpdatedAt
	}
	assert.Equal(t, expectedMembersResp, gotMembersResp)
}

func TestVmAutoscaler_ListHistory(t *testing.T) {
	setup()
	defer teardown()

	autoscalePoolID := "1044bfca-e490-44a1-aa1c-6f002daf6a13"
	mux.HandleFunc(fmt.Sprintf("%s/%s/history", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, vmAutoscaleListHistoryJSONResponse)
	})

	expectedHistoryResp := []*VmAutoscaleHistoryEvent{
		{
			HistoryEventID:       "4344c45f-7574-493b-a96c-df805c65a900",
			CurrentInstanceCount: 0,
			DesiredInstanceCount: 1,
			Reason:               "configuration update",
			Status:               "success",
		},
		{
			HistoryEventID:       "9ad436f7-af57-49ff-b416-0043721055b2",
			CurrentInstanceCount: 1,
			DesiredInstanceCount: 2,
			Reason:               "scaling up (desired=2 current=1)",
			Status:               "success",
		},
		{
			HistoryEventID:       "45390191-d077-49e9-a3c4-c2eb903bc1a2",
			CurrentInstanceCount: 2,
			DesiredInstanceCount: 1,
			Reason:               "scaling down (desired=1 current=2)",
			Status:               "success",
		},
	}

	gotHistoryResp, _, err := client.VmAutoscale.ListHistory(ctx, autoscalePoolID, nil)
	require.NoError(t, err)
	require.NotEmpty(t, gotHistoryResp)
	sort.SliceStable(gotHistoryResp, func(i, j int) bool {
		return gotHistoryResp[i].CreatedAt.Before(gotHistoryResp[j].CreatedAt)
	})
	for idx := range gotHistoryResp {
		expectedHistoryResp[idx].CreatedAt = gotHistoryResp[idx].CreatedAt
		expectedHistoryResp[idx].UpdatedAt = gotHistoryResp[idx].UpdatedAt
	}
	assert.Equal(t, expectedHistoryResp, gotHistoryResp)
}

func TestVmAutoscaler_Create(t *testing.T) {
	setup()
	defer teardown()

	createReq := &VmAutoscalePoolRequest{
		Name: "test-autoscalergroup-01",
		Config: &VmAutoscaleConfiguration{
			MinInstances:         1,
			MaxInstances:         5,
			TargetCPUUtilization: 0.5,
		},
		VmTemplate: &VmAutoscaleResourceTemplate{
			Size:     "s-1vcpu-512mb-10gb",
			Region:   "s2r1",
			Image:    "547864",
			Tags:     []string{"test-ag-01"},
			SSHKeys:  []string{"372862", "367582", "355790"},
			VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
			IPV6:     true,
			UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
		},
	}

	mux.HandleFunc(vmAutoscaleBasePath, func(w http.ResponseWriter, r *http.Request) {
		req := new(VmAutoscalePoolRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, http.MethodPost)
		assert.Equal(t, createReq, req)
		fmt.Fprintf(w, vmAutoscaleGetJSONResponse)
	})

	expectedPoolResp := &VmAutoscalePool{
		ID:   "1044bfca-e490-44a1-aa1c-6f002daf6a13",
		Name: "test-autoscalergroup-01",
		Config: &VmAutoscaleConfiguration{
			MinInstances:         1,
			MaxInstances:         5,
			TargetCPUUtilization: 0.5,
			CooldownMinutes:      5,
		},
		VmTemplate: &VmAutoscaleResourceTemplate{
			Size:     "s-1vcpu-512mb-10gb",
			Region:   "s2r1",
			Image:    "547864",
			Tags:     []string{"test-ag-01"},
			SSHKeys:  []string{"372862", "367582", "355790"},
			VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
			IPV6:     true,
			UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
		},
		CurrentUtilization: &VmAutoscaleResourceUtilization{
			Memory: 0.35,
			CPU:    0.0008,
		},
		Status: "active",
	}

	createPoolResp, _, err := client.VmAutoscale.Create(ctx, createReq)
	require.NoError(t, err)
	require.NotEmpty(t, createPoolResp)
	expectedPoolResp.CreatedAt = createPoolResp.CreatedAt
	expectedPoolResp.UpdatedAt = createPoolResp.UpdatedAt
	assert.Equal(t, expectedPoolResp, createPoolResp)
}

func TestVmAutoscaler_Update(t *testing.T) {
	setup()
	defer teardown()

	updateReq := &VmAutoscalePoolRequest{
		Name: "test-autoscalergroup-01",
		Config: &VmAutoscaleConfiguration{
			MinInstances:         1,
			MaxInstances:         5,
			TargetCPUUtilization: 0.5,
		},
		VmTemplate: &VmAutoscaleResourceTemplate{
			Size:     "s-1vcpu-512mb-10gb",
			Region:   "s2r1",
			Image:    "547864",
			Tags:     []string{"test-ag-01"},
			SSHKeys:  []string{"372862", "367582", "355790"},
			VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
			IPV6:     true,
			UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
		},
	}

	autoscalePoolID := "d50d8276-ad17-475d-8d2a-26b0acac756c"
	mux.HandleFunc(fmt.Sprintf("%s/%s", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		req := new(VmAutoscalePoolRequest)
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, http.MethodPut)
		assert.Equal(t, updateReq, req)
		fmt.Fprintf(w, vmAutoscaleGetJSONResponse)
	})

	expectedPoolResp := &VmAutoscalePool{
		ID:   "1044bfca-e490-44a1-aa1c-6f002daf6a13",
		Name: "test-autoscalergroup-01",
		Config: &VmAutoscaleConfiguration{
			MinInstances:         1,
			MaxInstances:         5,
			TargetCPUUtilization: 0.5,
			CooldownMinutes:      5,
		},
		VmTemplate: &VmAutoscaleResourceTemplate{
			Size:     "s-1vcpu-512mb-10gb",
			Region:   "s2r1",
			Image:    "547864",
			Tags:     []string{"test-ag-01"},
			SSHKeys:  []string{"372862", "367582", "355790"},
			VpcUUID:  "72b0812c-7535-4388-8507-5ad29b4487b3",
			IPV6:     true,
			UserData: "\n#cloud-config\nruncmd:\n- apt-get update\n- apt-get install -y stress-ng\n",
		},
		CurrentUtilization: &VmAutoscaleResourceUtilization{
			Memory: 0.35,
			CPU:    0.0008,
		},
		Status: "active",
	}

	updatePoolResp, _, err := client.VmAutoscale.Update(ctx, autoscalePoolID, updateReq)
	require.NoError(t, err)
	require.NotEmpty(t, updatePoolResp)
	expectedPoolResp.CreatedAt = updatePoolResp.CreatedAt
	expectedPoolResp.UpdatedAt = updatePoolResp.UpdatedAt
	assert.Equal(t, expectedPoolResp, updatePoolResp)
}

func TestVmAutoscaler_Delete(t *testing.T) {
	setup()
	defer teardown()

	autoscalePoolID := "d50d8276-ad17-475d-8d2a-26b0acac756c"
	mux.HandleFunc(fmt.Sprintf("%s/%s", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.VmAutoscale.Delete(ctx, autoscalePoolID)
	assert.NoError(t, err)
}

func TestVmAutoscaler_DeleteDangerous(t *testing.T) {
	setup()
	defer teardown()

	autoscalePoolID := "d50d8276-ad17-475d-8d2a-26b0acac756c"
	mux.HandleFunc(fmt.Sprintf("%s/%s/dangerous", vmAutoscaleBasePath, autoscalePoolID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		if expectedHeader, err := strconv.ParseBool(r.Header.Get("X-Dangerous")); err != nil {
			t.Fatal(err)
		} else if !expectedHeader {
			t.Errorf("Request header = %v, expected %v", r.Header.Get("X-Dangerous"), true)
		}
	})

	_, err := client.VmAutoscale.DeleteDangerous(ctx, autoscalePoolID)
	assert.NoError(t, err)
}
