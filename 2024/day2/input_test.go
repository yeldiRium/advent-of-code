package day2_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/advent-of-code/2024/day2"
)

func TestParseReports(t *testing.T) {
	t.Run("succeeds for a valid input file.", func(t *testing.T) {
		inputFile := bytes.NewBuffer([]byte("1 2 3\n4 5 6\n"))

		reports, err := day2.ParseReports(bufio.NewReader(inputFile))
		assert.NoError(t, err)
		assert.Equal(t, []day2.Report{
			{1, 2, 3},
			{4, 5, 6},
		}, reports)
	})
}
