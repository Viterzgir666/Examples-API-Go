package generics

import (
	"../utils"
)

var (
	RandomId          = utils.GenerateIntegerInRange(100, 200)
	RandomUserId      = utils.GenerateIntegerInRange(1, 10)
	RandomTitle       = utils.GenerateTitle()
	RandomDescription = utils.GenerateDescription()
)
