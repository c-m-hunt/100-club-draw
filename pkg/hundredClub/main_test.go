package hundredClub

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHundredClub_Draw(t *testing.T) {

	t.Run("test randomness", func(t *testing.T) {
		h := New("entries_test.csv")

		rand.Seed(time.Now().UnixNano())
		results := []DrawResult{}
		for i := 0; i < 100000; i++ {
			results = append(results, h.Draw())
		}

		count := make(map[int]int)

		for _, result := range results {
			count[result.Result[0].Entry.Number]++
		}

		if len(count) != len(h.Entries) {
			t.Errorf("Not all numbers win over a long period. Found %v winning numbers, want %v", len(count), len(h.Entries))
		}

		for no, v := range count {
			fmt.Println(no, v)
			want := 0.8 * 1000
			if v < int(want) {
				t.Errorf("Draw() = %v, want %v", v, want)
			}
		}
	})
}
