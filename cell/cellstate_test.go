package cell_test

import (
	"errors"
	"testing"

	"github.com/Irainia/gameoflife-go/cell"
)

func TestNewShouldReturnNilAndErrorForInitialStateNil(t *testing.T) {
	var expectedCellState *cell.CellState = nil
	var expectedError error = errors.New(cell.ArgumentNilError)

	actualCellState, actualError := cell.New(nil)

	if actualCellState != expectedCellState {
		t.Error("expected: nil -- actual: not nil")
		return
	}
	if actualError == nil {
		t.Errorf("expected: %s -- actual: nil", expectedError.Error())
		return
	}
	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected: %s -- actual: %s", expectedError.Error(), actualError.Error())
	}
}
