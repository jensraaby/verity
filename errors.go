package main

import (
	"fmt"
	"time"
)

type GVError struct {
	Operation string
	When      time.Time
	What      string
}

func (e *GVError) Error() string {
	return fmt.Sprintln("GoVerify Error:", e.Operation, ": ", e.What, " at ", e.When)
}
