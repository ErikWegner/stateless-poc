package states

import "github.com/qmuntal/stateless"

const (
	triggerInitMaintenance = "InitMaintenance"
)

const (
	stateActive = "Active"
	stateCreateTaskA = "WaitingForTaskA"
)

func GetMachine() *stateless.StateMachine {
	machine := stateless.NewStateMachine(stateActive)

	machine.Configure(stateActive).
	 	Permit(triggerInitMaintenance, stateCreateTaskA)

	return machine
}
