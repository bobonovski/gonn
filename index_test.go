package gonn

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndexBuild(t *testing.T) {
	index := NewIndex(10, 5)
	rand.Seed(time.Now().Unix())
	// add items randomly
	for k := 0; k < 100; k += 1 {
		var item []float32
		id := strconv.FormatInt(int64(k), 10)
		for d := 0; d < 5; d += 1 {
			item = append(item, rand.Float32())
		}
		index.AddItem(id, item)
	}
	// build index
	err := index.Build()
	assert.Equal(t, err, nil)
}
