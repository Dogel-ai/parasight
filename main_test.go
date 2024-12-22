package main

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"testing"
)

var dataMasks = []struct {
	input    string
	expected [][]uint8
}{
	{input: "Hello", expected: [][]uint8{{0, 64, 0, 0, 8, 0, 0, 0},
		{0, 64, 32, 0, 0, 4, 0, 1},
		{0, 64, 32, 0, 8, 4, 0, 0},
		{0, 64, 32, 0, 8, 4, 0, 0},
		{0, 64, 32, 0, 8, 4, 2, 1}}},

	{input: "bFeeHspqxh", expected: [][]uint8{{0, 64, 32, 0, 0, 0, 2, 0},
		{0, 64, 0, 0, 0, 4, 2, 0},
		{0, 64, 32, 0, 0, 4, 0, 1},
		{0, 64, 32, 0, 0, 4, 0, 1},
		{0, 64, 0, 0, 8, 0, 0, 0},
		{0, 64, 32, 16, 0, 0, 2, 1},
		{0, 64, 32, 16, 0, 0, 0, 0},
		{0, 64, 32, 16, 0, 0, 0, 1},
		{0, 64, 32, 16, 8, 0, 0, 0},
		{0, 64, 32, 0, 8, 0, 0, 0}}},

	{input: "iXvdqmTLOb", expected: [][]uint8{{0, 64, 32, 0, 8, 0, 0, 1},
		{0, 64, 0, 16, 8, 0, 0, 0},
		{0, 64, 32, 16, 0, 4, 2, 0},
		{0, 64, 32, 0, 0, 4, 0, 0},
		{0, 64, 32, 16, 0, 0, 0, 1},
		{0, 64, 32, 0, 8, 4, 0, 1},
		{0, 64, 0, 16, 0, 4, 0, 0},
		{0, 64, 0, 0, 8, 4, 0, 0},
		{0, 64, 0, 0, 8, 4, 2, 1},
		{0, 64, 32, 0, 0, 0, 2, 0}}},

	{input: "kWebuFmSZG", expected: [][]uint8{{0, 64, 32, 0, 8, 0, 2, 1},
		{0, 64, 0, 16, 0, 4, 2, 1},
		{0, 64, 32, 0, 0, 4, 0, 1},
		{0, 64, 32, 0, 0, 0, 2, 0},
		{0, 64, 32, 16, 0, 4, 0, 1},
		{0, 64, 0, 0, 0, 4, 2, 0},
		{0, 64, 32, 0, 8, 4, 0, 1},
		{0, 64, 0, 16, 0, 0, 2, 1},
		{0, 64, 0, 16, 8, 0, 2, 0},
		{0, 64, 0, 0, 0, 4, 2, 1}}},

	//{input: "Z15VUkgTvrXYR6A", expected: [][]uint8{{0, 1, 0, 1, 1, 0, 1, 0}, {0, 0, 1, 1, 0, 0, 0, 1}, {0, 0, 1, 1, 0, 1, 0, 1}, {0, 1, 0, 1, 0, 1, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 1}, {0, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 0, 1, 0, 0}, {0, 1, 1, 1, 0, 1, 1, 0}, {0, 1, 1, 1, 0, 0, 1, 0}, {0, 1, 0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0, 0, 1}, {0, 1, 0, 1, 0, 0, 1, 0}, {0, 0, 1, 1, 0, 1, 1, 0}, {0, 1, 0, 0, 0, 0, 0, 1}}},
	/*{input: "TqxuQjNm3ReaoM6", expected: [][]uint8{}},
	{input: "UdNhYw7AMu", expected: [][]uint8{}},
	{input: "jUbdwjYozxsXnLw", expected: [][]uint8{}},
	{input: "a9FWVdvj367m8XzyR72QALYlkQehAb94", expected: [][]uint8{}},
	{input: "6rdIHjbyezTObLS9bPMrZ9kostUODQDJ", expected: [][]uint8{}},
	{input: "ON3CVRhbgQf3HidZ0KzMPuxw6dYCBxQB", expected: [][]uint8{}},
	{input: "w5Q5rc6vhEcgQUsqtNMl9Qu54PqKvPgI", expected: [][]uint8{}},
	{input: "ilSD52D9l7bvORKKzMqPpW4FrwOFRmWi", expected: [][]uint8{}},*/
	//{input: "According to all known laws of aviation, there is no way a bee should be able to fly. Its wings are too small to get its fat little body off the ground. The bee, of course, flies anyway because bees don't care what humans think is impossible. Yellow, black. Yellow, black. Yellow, black. Yellow, black. Ooh, black and yellow! Let's shake it up a little."},
}

func BenchmarkGetDataMask(b *testing.B) {
	for _, v := range dataMasks {
		b.Run(fmt.Sprintf("input_%s", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getDataMask(stringToUint8(v.input))
			}
		})
	}
}

func TestGetDataMask(t *testing.T) {
	for _, v := range dataMasks {
		t.Run(fmt.Sprintf("input_%s", v.input), func(t *testing.T) {
			output, _ := getDataMask(stringToUint8(v.input))
			for i := range output {
				if !(bytes.Equal(output[i], v.expected[i])) {
					t.Errorf("Output %q not equal to expected %q", output, v.expected)
				}
			}

		})
	}
}

var powersTable = []struct {
	base, power int
}{
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(100), power: rand.IntN(1024)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
	{base: rand.IntN(1024), power: rand.IntN(20480)},
}

func BenchmarkGetIntPower(b *testing.B) {
	for _, v := range powersTable {
		b.Run(fmt.Sprintf("input_%d^%d", v.base, v.power), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getIntPower(v.base, v.power)
			}
		})
	}
}
