package common_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/advent-of-code/2024/day1/common"
)

func TestParseInputLists(t *testing.T) {
	t.Run("succeeds for a valid input file.", func(t *testing.T) {
		inputFile := bytes.NewBuffer([]byte("1   2\n3   4\n"))

		leftList, rightList, err := common.ParseInputLists(bufio.NewReader(inputFile))
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 3}, leftList)
		assert.Equal(t, []int{2, 4}, rightList)
	})
}
