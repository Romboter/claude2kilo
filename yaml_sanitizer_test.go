package main

import (
	"testing"
)

func TestNewYAMLSanitizer(t *testing.T) {
	s := NewYAMLSanitizer()
	if s == nil {
		t.Fatal("Expected NewYAMLSanitizer to return non-nil")
	}
}

func TestSanitizeFrontmatter_ValidYAML(t *testing.T) {
	s := NewYAMLSanitizer()
	input := "description: A short description.\ntools: [tool1, tool2]"
	output, err := s.SanitizeFrontmatter(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if output == "" {
		t.Error("Expected output, got empty string")
	}
}

func TestSanitizeFrontmatter_InvalidYAML(t *testing.T) {
	s := NewYAMLSanitizer()
	input := "description: \"Unclosed string"
	_, err := s.SanitizeFrontmatter(input)
	if err == nil {
		t.Error("Expected error for invalid YAML, got nil")
	}
}

func TestSanitizeLine_ToolsString(t *testing.T) {
	s := NewYAMLSanitizer()
	line := "tools: tool1, tool2, tool3"
	got := s.sanitizeLine(line)
	want := "tools: [tool1, tool2, tool3]"
	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestSanitizeLine_LongDescription(t *testing.T) {
	s := NewYAMLSanitizer()
	longDesc := "description: " + string(make([]byte, 201))
	line := longDesc
	got := s.sanitizeLine(line)
	if !((len(got) > 0) && (got[:12] == "description:")) {
		t.Errorf("Expected sanitized description, got %q", got)
	}
}

func TestNeedsLiteralBlock(t *testing.T) {
	s := NewYAMLSanitizer()
	if !s.needsLiteralBlock("This has \"quotes\" and <example> and is long enough.") {
		t.Error("Expected needsLiteralBlock to return true")
	}
	if s.needsLiteralBlock("Short description.") {
		t.Error("Expected needsLiteralBlock to return false")
	}
}

func TestSanitizeDescriptionContent(t *testing.T) {
	s := NewYAMLSanitizer()
	input := "This is a description. <example>Example here</example>"
	out := s.sanitizeDescriptionContent(input)
	if out == input {
		t.Error("Expected sanitized content to differ from input")
	}
}

func TestDetectIssues(t *testing.T) {
	s := NewYAMLSanitizer()
	input := "description: " + string(make([]byte, 201)) + "\ntools: tool1, tool2"
	issues := s.DetectIssues(input)
	if len(issues) != 2 {
		t.Errorf("Expected 2 issues, got %d", len(issues))
	}
}
