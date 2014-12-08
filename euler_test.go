package euler

import "testing"

func TestProblemOne(t *testing.T) {
    const out = 233168
    if x := ProblemOne(); x != out {
        t.Error("ProblemOne() = %v, want %v", x, out)
    }
}
