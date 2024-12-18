package day1

import "testing"

func TestCalculateTotalDistances(t *testing.T) {
	IDsLeft := []int{3, 4, 2, 1, 3, 3}
	IDsRight := []int{4, 3, 5, 3, 9, 3}

	want := 11
	got, _ := calculateTotalDistances(IDsLeft, IDsRight)

	if want != got {
		t.Errorf("test failed; want %d, got %d", want, got)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	IDsLeft := []int{3, 4, 2, 1, 3, 3}
	IDsRight := []int{4, 3, 5, 3, 9, 3}

	want := 31
	got, _ := calculateSimilarityScore(IDsLeft, IDsRight)

	if want != got {
		t.Errorf("test failed; want %d, got %d", want, got)
	}
}
