package states

import (
	"testing"
)

// Transition from _active_ to _createTaskA_
func TestFromActiveToCreateTaskA(t *testing.T) {
	m := GetMachine()

	err := m.Fire(triggerInitMaintenance)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	finalState := m.MustState()
	if finalState != stateCreateTaskA {
		t.Errorf("want:%s\ngot:%s", stateCreateTaskA, finalState)
	}
}
