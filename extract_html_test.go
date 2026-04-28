package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "h1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "h2 fallback when no h1",
			inputBody: "<html><body><h2>Test Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "h1 takes priority over h2",
			inputBody: "<html><body><h1>First</h1><h2>Second</h2></body></html>",
			expected:  "First",
		},
		{
			name:      "no heading",
			inputBody: "<html><body><p>No heading here</p></body></html>",
			expected:  "",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name: "extracts content from main tag",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "multiple paragraphs in main",
			inputBody: `<html><body>
				<main>
					<p>First.</p>
					<p>Second.</p>
				</main>
			</body></html>`,
			expected: "First.",
		},
		{
			name: "empty main tag",
			inputBody: `<html><body>
				<main></main>
			</body></html>`,
			expected: "",
		},
		{
			name:      "empty body",
			inputBody: `<html><body></body></html>`,
			expected:  "",
		},
		{
			name:      "empty string input",
			inputBody: ``,
			expected:  "",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}
