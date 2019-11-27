package listcommand

import (
	"testing"

	"github.com/knutsonchris/stackilackey/add"
	"github.com/knutsonchris/stackilackey/remove"
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
	found := false
	for _, environment := range environments {
		if environment.EnvironmentName == "testEnvironment" {
			found = true
		}
	}
	if !found {
		t.Fatalf("list environment failed. attempted to list testEnvironment but testEnvironment was not found")

	}

	// to clean up after ourselves, remove the test environment

	remove := remove.Remove{}
	removeEnv := remove.Environment

	_, err = removeEnv.Environment("testEnvironment")
	if err != nil {
		t.Fatalf("list environment tear down failed. Unable to remove test environment %s", err)
	}
}
