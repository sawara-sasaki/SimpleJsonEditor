package action

import (
	"testing"
)

func TestGetLinearIntDataSuccess(t *testing.T) {
	result := GetLinearIntData(1, 10)
	correctSlice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range result {
		if v != correctSlice[i] {
			t.Fatal("failed test")
		}
	}
}

func TestGetLinearFloatDataSuccess(t *testing.T) {
	result := GetLinearFloatData(1.0, 2.0, 0.1)
	correctSlice := []interface{}{1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0}
	for i, v := range result {
		if v != correctSlice[i] {
			t.Fatal("failed test")
		}
	}
}

func TestGetLinearFloatDataSuccess2(t *testing.T) {
	result := GetLinearFloatData(1.0, 1.1, 0.01)
	correctSlice := []interface{}{1.00, 1.01, 1.02, 1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09, 1.1}
	for i, v := range result {
		if v != correctSlice[i] {
			t.Fatal("failed test")
		}
	}
}
