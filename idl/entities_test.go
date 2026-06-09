package idl

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseExistingMatterElements(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "matter-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	matterFileContent := `
// Some comments
/* More comments */
cluster DoorLock = 257 {
  revision 1;

  enum DoorStateEnum : enum8 {
    kSecured = 0;
    kUnsecured = 1;
  }

  bitmap Feature : bitmap32 {
    kPinGeneration = 0x1;
    kRfidCredential = 0x2;
  }

  struct AccessControlEntryStruct {
    fabric_idx fabricIndex = 1;
    privilege privilege = 2;
    auth_mode authMode = 3;
  }

  info event event_id = 1 {
    nullable int32u attributeId = 0;
    int32u eventId = 1;
  }

  info event optional opt_event = 2 {
    int32u someField = 0;
  }

  readonly attribute DoorStateEnum doorState = 0;
  attribute int32u pinCodeLength = 1;

  request command LockDoor() = 0;
  response command LockDoorResponse() = 1;
  command optional OptCommand() = 2;
  timed command access(invoke: manage) UpdatePIN(UpdatePINRequest): DefaultSuccess = 3;
  command access(invoke: administer) ResetPIN() = 4;
}

cluster BasicInformation = 40 {
  enum ProductTypeEnum : enum8 {
    kItem = 0;
    kService = 1;
  }
}
`

	filePath := filepath.Join(tmpDir, "test.matter")
	err = os.WriteFile(filePath, []byte(matterFileContent), 0644)
	if err != nil {
		t.Fatalf("failed to write test matter file: %v", err)
	}

	got, err := parseExistingMatterElements(filePath)
	if err != nil {
		t.Fatalf("parseExistingMatterElements failed: %v", err)
	}

	want := map[string]bool{
		"doorlock":                                             true,
		"doorlock.doorstateenum":                               true,
		"doorlock.doorstateenum.secured":                       true,
		"doorlock.doorstateenum.unsecured":                     true,
		"doorlock.feature":                                     true,
		"doorlock.feature.pingeneration":                       true,
		"doorlock.feature.rfidcredential":                      true,
		"doorlock.accesscontrolentrystruct":                    true,
		"doorlock.accesscontrolentrystruct.fabricindex":        true,
		"doorlock.accesscontrolentrystruct.privilege":          true,
		"doorlock.accesscontrolentrystruct.authmode":           true,
		"doorlock.event.event_id":                              true,
		"doorlock.event.event_id.attributeid":                  true,
		"doorlock.event.event_id.eventid":                      true,
		"doorlock.event.opt_event":                             true,
		"doorlock.event.opt_event.somefield":                   true,
		"doorlock.attribute.doorstate":                         true,
		"doorlock.attribute.pincodelength":                     true,
		"doorlock.command.lockdoor":                            true,
		"doorlock.command.lockdoorresponse":                    true,
		"doorlock.command.optcommand":                          true,
		"doorlock.command.updatepin":                           true,
		"doorlock.command.resetpin":                            true,
		"basicinformation":                                     true,
		"basicinformation.producttypeenum":                     true,
		"basicinformation.producttypeenum.item":                true,
		"basicinformation.producttypeenum.service":             true,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseExistingMatterElements returned mismatch.\nGot:\n")
		for k := range got {
			t.Logf("  %q: true,", k)
		}
		t.Errorf("\nWant:\n")
		for k := range want {
			t.Logf("  %q: true,", k)
		}
	}
}
