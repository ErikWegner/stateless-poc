package states

import (
	"testing"

	"github.com/qmuntal/stateless"
)

// Transition from _active_ to _createTaskA_
func TestFromActiveToCreateTaskA(t *testing.T) {
	m := machineAtActive()

	err := m.Fire(triggerInitMaintenance)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	finalState := m.MustState()
	if finalState != stateCreateTaskA {
		t.Errorf("want:%s\ngot:%s", stateCreateTaskA, finalState)
	}
}

// Transition from _createTaskA_ to _WaingForTaskACompleted_
func TestFromCreateTaskAToWaitingForTaskACompleted(t *testing.T) {
	m := machineAtCreateTaskA()

	err := m.Fire(triggerWaitingForTaskACompleted)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	finalState := m.MustState()
	if finalState != stateWaitingForTaskACompleted {
		t.Errorf("want:%s\ngot:%s", stateCreateTaskA, finalState)
	}
}

func machineAtActive() *stateless.StateMachine {
	return GetMachine()
}

func machineAtCreateTaskA() *stateless.StateMachine {
	m := machineAtActive()
	m.Fire(triggerInitMaintenance)
	return m
}
