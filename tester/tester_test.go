package tester

import (
	"fmt"
	"testing"
	"time"
)

func Test_Wait(t *testing.T) {
	wait := NewWait()
	start := time.Now()
	wait.WaitForMilliSeconds(100)
	elapsed := time.Since(start)

	if elapsed < 100*time.Millisecond {
		t.Errorf("Expected to wait at least 100 milliseconds, but waited %v", elapsed)
	}
}

func Test_Assert(t *testing.T) {
	assert := NewAssert(t, false)

	var err error = nil
	assert.IsErrNil(err, "Error should be nil")

	assert.IsTrue(1+1 == 2, "1 + 1 should equal 2")

	assert.AreEqual("hello", "hello", "Strings should be equal")

	var obj interface{} = struct{}{}
	assert.IsNotNil(obj, "Object should not be nil")

	assert.IsNotEmpty("not empty", "String should not be empty")
}

func Test_AssertFailingCases(t *testing.T) {
	assert := NewAssert(t, false)
	var err error = fmt.Errorf("some error")

	// The following assertions are expected to fail.
	assert.IsErrNil(err, "Error should be nil")
	assert.IsTrue(false, "Condition should be true")
	assert.AreEqual(1, 2, "Values should be equal")
	assert.IsNotNil(nil, "Object should not be nil")
	assert.IsNotEmpty("", "String should not be empty")
}
