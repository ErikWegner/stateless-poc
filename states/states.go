package states

import "github.com/qmuntal/stateless"

const (
	triggerInitMaintenance          = "InitMaintenance"
	triggerWaitingForTaskACompleted = "WaitingForTaskACompleted"
)

const (
	stateActive                   = "Active"
	stateCreateTaskA              = "WaitingForTaskA"
	stateWaitingForTaskACompleted = "WaitingForTaskACompleted"
)

func GetMachine() *stateless.StateMachine {
	machine := stateless.NewStateMachine(stateActive)

	machine.Configure(stateActive).
		Permit(triggerInitMaintenance, stateCreateTaskA)

	machine.Configure(stateCreateTaskA).
		Permit(triggerWaitingForTaskACompleted, stateWaitingForTaskACompleted)

	return machine
}
