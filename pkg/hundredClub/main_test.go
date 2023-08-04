package hundredClub

import (
	"math/rand"
	"testing"
	"time"
)

var bigNumber = 100000

func TestHundredClub_Draw(t *testing.T) {

	// Test that it fails when a number is defined twice in entries
	t.Run("test entries fail when duplucate numbers", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		_ = New("entries_test_bad.csv", 3)
	})

	h := New("entries_test.csv", 3)

	rand.Seed(time.Now().UnixNano())
	results := []DrawResult{}
	for i := 0; i < bigNumber; i++ {
		results = append(results, h.Draw())
	}

	// Test the correct amount of entries are loaded
	t.Run("test entries loaded", func(t *testing.T) {
		if len(h.Entries) != 100 {
			t.Errorf("Draw() = %v, want %v", len(h.Entries), 100)
		}
	})

	t.Run("test randomness", func(t *testing.T) {

		count := make(map[int]int)

		for _, result := range results {
			count[result.Result[0].Entry.Number]++

			if len(result.Result) != 3 {
				t.Errorf("Draw() = %v, want %v", len(result.Result), 3)
			}
		}

		if len(count) != len(h.Entries) {
			t.Errorf("Not all numbers win over a long period. Found %v winning numbers, want %v", len(count), len(h.Entries))
		}

		// Ensure all numbers win roughly the same amount. Should be within 10% of the average as long as the number of draws is large enough
		for _, v := range count {
			want := 0.90 * float32(bigNumber/len(h.Entries))
			if v < int(want) {
				t.Errorf("Draw() = %v, want %v", v, want)
			}
		}
	})

	t.Run("test entry cannot win twice in draw", func(t *testing.T) {
		for _, result := range results {
			if result.Result[0].Entry.Number == result.Result[1].Entry.Number || result.Result[0].Entry.Number == result.Result[2].Entry.Number || result.Result[1].Entry.Number == result.Result[2].Entry.Number {
				t.Errorf("Entries cannot win more than one prize")
			}
		}
	})
}
