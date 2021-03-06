package worker

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	ErrNoneAgents = errors.New("None active agents")
	ErrNoneFuncs  = errors.New("None functions")
	ErrTimeOut    = errors.New("Executing time out")
	ErrUnknown    = errors.New("Unknown error")
	ErrLostConn   = errors.New("Lost connection with Gearmand")
)

// Extract the error message
func getError(data []byte) (err error) {
	rel := bytes.SplitN(data, []byte{'\x00'}, 2)
	if len(rel) != 2 {
		err = fmt.Errorf("Not a error data: %V", data)
		return
	}
	err = errors.New(fmt.Sprintf("%s: %s", rel[0], rel[1]))
	return
}

// An error handler
type ErrorHandler func(error)
