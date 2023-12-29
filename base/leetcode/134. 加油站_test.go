package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompleteCircuit(t *testing.T) {
	tests := []struct {
		cost   []int
		gas    []int
		expect int
	}{
		{
			expect: 3,
			gas:    []int{1, 2, 3, 4, 5},
			cost:   []int{3, 4, 5, 1, 2},
		}, {
			expect: -1,
			gas:    []int{2, 3, 4},
			cost:   []int{3, 4, 5},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, canCompleteCircuit(test.gas, test.cost))
	}
}

func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumGas, sumCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumGas += gas[j]
			sumCost += cost[j]
			if sumCost > sumGas {
				break
			}
			cnt++
		}

		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}
