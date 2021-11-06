package states

import "github.com/qmuntal/stateless"

type WorkflowContext struct {
	state       stateless.State
	userIsAdmin bool
}

const USER_IS_ADMIN = true
const USER_IS_NOT_ADMIN = false
