package titlecase

import "testing"

func TestToTitleCase(t *testing.T) {
	// I want to test the ToTitleCase function through multiple test cases in a matrix
	// The matrix will contain the input string and the expected output
	testCases := []struct {
		input    string
		expected string
	}{
		{"this is a test", "This Is a Test"},
		{"I cannot believe the audacity OF some people", "I Cannot Believe the Audacity of Some People"},
		{"a river ruins through it", "A River Ruins Through It"},
		{"via vs vs. vs", "Via vs vs. vs"},
		{"a tale of two cItIeS", "A Tale of Two Cities"},
	}

	for _, tc := range testCases {

		result := ToTitleCase(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s, but got %s", tc.expected, result)
		}
	}
}
