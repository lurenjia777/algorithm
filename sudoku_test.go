package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestSudoku(t *testing.T) {
	rand.Seed(time.Now().Unix())
	GenerateSudoku()
	SettleSudoku(initdata)
}
