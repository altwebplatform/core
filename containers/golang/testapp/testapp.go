package testapp

import (
	"github.com/willf/pad"
	"strconv"
)

func ComputeBusinessLogic() int {
	return 37
}

func PaddedLogic() string {
	return pad.Right(strconv.Itoa(ComputeBusinessLogic()), 10, "!")
}
