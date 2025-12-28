package testequal

import (
	"fmt"
	"reflect"
)

func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if equal(expected, actual) {
		return true
	}

	msg := fmt.Sprintf("not equal:\nexpected: %v\nactual  : %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if format, ok := msgAndArgs[0].(string); ok {
			msg += "\nmessage : " + fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}

	t.Errorf(msg)
	return false
}

func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	t.Helper()
	if !equal(expected, actual) {
		return true
	}

	msg := fmt.Sprintf("equal:\nexpected: %v\nactual  : %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if format, ok := msgAndArgs[0].(string); ok {
			msg += "\nmessage : " + fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}

	t.Errorf(msg)
	return false
}

func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if equal(expected, actual) {
		return
	}

	msg := fmt.Sprintf("not equal:\nexpected: %v\nactual  : %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if format, ok := msgAndArgs[0].(string); ok {
			msg += "\nmessage : " + fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}

	t.Errorf(msg)
	t.FailNow()
}

func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	t.Helper()
	if !equal(expected, actual) {
		return
	}

	msg := fmt.Sprintf("equal:\nexpected: %v\nactual  : %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if format, ok := msgAndArgs[0].(string); ok {
			msg += "\nmessage : " + fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}

	t.Errorf(msg)
	t.FailNow()
}

func equal(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	at, bt := reflect.TypeOf(a), reflect.TypeOf(b)
	if at != bt {
		return false
	}

	switch a.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		string:
		return a == b
	case []int, []byte, map[string]string:
		return reflect.DeepEqual(a, b)
	}
	return false
}
