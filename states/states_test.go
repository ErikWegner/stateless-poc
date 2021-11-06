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

// Transition from _active_ to _createTaskA_
func TestFromActiveToCreateTaskA_Denied_Because_UserIsNotAdmin(t *testing.T) {
	m := GetMachine(stateActive, USER_IS_NOT_ADMIN)

	b, err := m.CanFire(triggerInitMaintenance)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	if b {
		t.Errorf("want:%v\ngot:%v", false, true)
	}
}

// Transition from _createTaskA_ to _WaingForTaskACompleted_
func TestFromCreateTaskAToWaitingForTaskACompleted(t *testing.T) {
	m := machineAtCreateTaskA()
	taskId := 42

	err := m.Fire(triggerWaitingForTaskACompleted, taskId)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	finalState := m.MustState()
	if finalState != stateWaitingForTaskACompleted {
		t.Errorf("want:%s\ngot:%s", stateCreateTaskA, finalState)
	}
}

func TestCanFire_triggerWaitingForTaskACompleted_is_false(t *testing.T) {
	m := machineAtActive()

	b, err := m.CanFire(triggerWaitingForTaskACompleted)
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}

	if b {
		t.Errorf("want:%v\ngot:%v", false, true)
	}
}

func machineAtActive() *stateless.StateMachine {
	return GetMachine(stateActive, USER_IS_ADMIN)
}

func machineAtCreateTaskA() *stateless.StateMachine {
	return GetMachine(stateCreateTaskA, USER_IS_ADMIN)
}
