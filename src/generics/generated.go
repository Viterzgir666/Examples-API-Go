package generics

import (
	"github.com/Viterzgir666/Examples-API-Go/src/utils"
)

var (
	RandomId          = utils.GenerateIntegerInRange(100, 200)
	RandomUserId      = utils.GenerateIntegerInRange(1, 10)
	RandomTitle       = utils.GenerateTitle()
	RandomDescription = utils.GenerateDescription()
)
