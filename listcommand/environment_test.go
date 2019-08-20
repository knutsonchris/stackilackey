package listcommand

import (
	"testing"

	"github.td.teradata.com/ck250037/stackilackey/add"
)

func TestEnvironment_Environment(t *testing.T) {
	// in order to test list environmnet, we must first add one
	add := add.Add{}
	_, err := add.Environment.Environment("testEnvironment")
	if err != nil {
		t.Fatalf("list environment set up failed. attempting to add testEnvironment failed with error %s", err)
	}
	e := Environment{}
	environments, err := e.Environment("testEnvironment")
	if err != nil {
		t.Fatalf("list environment failed with error %s", err)
	}
	for _, environment := range environments {
		if environment.EnvironmentName == "testEnvironment" {
			// TODO: need to tear this test down with remove environment
			return
		}
	}
	t.Errorf("list environment failed. attempted to list testEnvironment but testEnvironment was not found")
}
