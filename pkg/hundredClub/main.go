package hundredClub

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gocarina/gocsv"
	"golang.org/x/exp/slices"
)

type Entry struct {
	Number int    `csv:"number"`
	Name   string `csv:"name"`
}

type HundredClub struct {
	Entries      []*Entry
	PrizesToDraw int
	PrizeNames   []string
}

type DrawResult struct {
	Result []PrizeResult
}

type PrizeResult struct {
	PrizeName string
	Entry     *Entry
}

func New(entriesFilePath string, prizeCount int) *HundredClub {
	entries := generateEntriesFromCSV(entriesFilePath)
	prizeNames := []string{}
	for i := 1; i <= prizeCount; i++ {
		prizeNames = append(prizeNames, humanize.Ordinal(i))
	}

	return &HundredClub{
		Entries:      entries,
		PrizesToDraw: prizeCount,
		PrizeNames:   prizeNames,
	}
}

func (h *HundredClub) DisplayEntries() {
	fmt.Println()
	fmt.Println("Entries:")
	fmt.Println(strings.Repeat("-", 80))
	for _, entry := range h.Entries {
		fmt.Println(entry.Number, entry.Name)
	}

	fmt.Println()
	fmt.Println("Total entries:", len(h.Entries))
}

func (h *HundredClub) DisplayEntriesSummary() {
	fmt.Println()
	fmt.Println("Entry summary:")
	fmt.Println(strings.Repeat("-", 80))
	entriesSummary := make(map[string]int)
	for _, entry := range h.Entries {
		entriesSummary[entry.Name]++
	}

	keys := make([]string, 0, len(entriesSummary))
	for k := range entriesSummary {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, entriesSummary[k])
	}
}

func (h *HundredClub) Draw() DrawResult {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	drawResults := []PrizeResult{}
	entries := make([]*Entry, len(h.Entries))
	copy(entries, h.Entries)
	for i := 0; i < h.PrizesToDraw; i++ {
		prizeIndex := r.Intn(len(entries))
		drawResults = append(drawResults, PrizeResult{PrizeName: h.PrizeNames[i], Entry: entries[prizeIndex]})
		entries = append(entries[:prizeIndex], entries[prizeIndex+1:]...)
	}
	return DrawResult{Result: drawResults}
}

func (h *HundredClub) DrawAndDisplay() {
	drawResult := h.Draw()
	for _, prizeResult := range drawResult.Result {
		fmt.Println(prizeResult.PrizeName, prizeResult.Entry.Number, prizeResult.Entry.Name)
	}
}

func generateEntriesFromCSV(entriesFilePath string) []*Entry {
	in, err := os.Open(entriesFilePath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	entries := []*Entry{}

	if err := gocsv.UnmarshalFile(in, &entries); err != nil {
		panic(err)
	}

	filteredEntries := []*Entry{}
	numbers := []int{}
	for _, entry := range entries {
		if slices.Contains(numbers, entry.Number) {
			panic(fmt.Sprintf("Number %v is defined twice in entries", entry.Number))
		}
		numbers = append(numbers, entry.Number)
		if strings.TrimSpace(entry.Name) != "" {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries
}
