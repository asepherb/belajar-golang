package problems

import "testing"

func TestShouldReturnTrueWhenThreeNumbersAreEqual(t *testing.T) {
    areThreeNumbersEqual := equalOrNotEqual(1, 2, 3)
    if !areThreeNumbersEqual {
        t.Error("Expected true, got", areThreeNumbersEqual)
    }
}

func equalOrNotEqual(first, second, third int) bool {
    if first == second && second == third {
        return true
    } else {
        return false
    }
  }

  func TestThreeEqualNumbers(t *testing.T) {
    isThreeNumbersEqual := EqualOrNotEqual(1, 1, 1)
    if !isThreeNumbersEqual {
        t.Error("\tExpected true, got ", isThreeNumbersEqual)
    }
}
