package day2

import "testing"

func TestAnalyzeReports(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	want := 2
	got := analyzeReports(reports)

	if want != got {
		t.Errorf("test failed; want %d, got %d", want, got)
	}
}

func TestApplyProblemDampener(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	want := 4
	got := applyProblemDampener(reports, 2)

	if want != got {
		t.Errorf("test failed; want %d, got %d", want, got)
	}
}
