package go_api_abrha

import (
	"testing"
)

func TestStringify_Map(t *testing.T) {
	result := Stringify(map[int]*VmBackupPolicy{
		1: {VmID: "1", BackupEnabled: true},
		2: {VmID: "2"},
		3: {VmID: "3", BackupEnabled: true},
	})

	expected := `map[1:go_api_abrha.VmBackupPolicy{VmID:"1", BackupEnabled:true}, 2:go_api_abrha.VmBackupPolicy{VmID:"2", BackupEnabled:false}, 3:go_api_abrha.VmBackupPolicy{VmID:"3", BackupEnabled:true}]`
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
