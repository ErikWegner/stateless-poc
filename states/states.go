package states

import (
	"context"
	"reflect"

	"github.com/qmuntal/stateless"
)

const (
	triggerInitMaintenance          = "InitMaintenance"
	triggerWaitingForTaskACompleted = "WaitingForTaskACompleted"
)

const (
	stateActive                   = "Active"
	stateCreateTaskA              = "WaitingForTaskA"
	stateWaitingForTaskACompleted = "WaitingForTaskACompleted"
)

func GetMachine(state stateless.State, userIsAdmin bool) *stateless.StateMachine {
	myState := WorkflowContext{
		state:       state,
		userIsAdmin: userIsAdmin,
	}
	machine := stateless.NewStateMachineWithExternalStorage(func(_ context.Context) (stateless.State, error) {
		return myState.state, nil
	}, func(_ context.Context, state stateless.State) error {
		myState.state = state
		return nil
	}, stateless.FiringImmediate)

	machine.SetTriggerParameters(triggerWaitingForTaskACompleted, reflect.TypeOf(0))

	machine.Configure(stateActive).
		Permit(triggerInitMaintenance, stateCreateTaskA, func(_ context.Context, _ ...interface{}) bool {
			return myState.userIsAdmin
		})

	machine.Configure(stateCreateTaskA).
		Permit(triggerWaitingForTaskACompleted, stateWaitingForTaskACompleted)

	return machine
}
