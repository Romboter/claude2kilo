package main

import (
	"testing"
)

func TestNewContentAnalyzer(t *testing.T) {
	ca := NewContentAnalyzer()
	if ca == nil {
		t.Fatal("Expected NewContentAnalyzer to return non-nil")
	}
}

func TestFindBestMatch(t *testing.T) {
	ca := NewContentAnalyzer()
	patterns := map[string]string{"foo|bar": "desc1", "baz": "desc2"}
	text := "this is a foo bar test"
	desc, score := ca.findBestMatch(text, patterns)
	if desc != "desc1" || score == 0 {
		t.Errorf("Expected desc1 and nonzero score, got %q, %d", desc, score)
	}
}

func TestExtractFromProactiveStatement(t *testing.T) {
	ca := NewContentAnalyzer()
	desc := ca.extractFromProactiveStatement("Use PROACTIVELY for code review and analysis.")
	if desc == "" {
		t.Error("Expected to extract proactive statement")
	}
}

func TestGenerateWhenToUseStatement_Proactive(t *testing.T) {
	ca := NewContentAnalyzer()
	out := ca.generateWhenToUseStatement("", "Use PROACTIVELY for code review.", "")
	if out == "" || out == ca.fallbackPattern+"." {
		t.Error("Expected a specific when-to-use statement, got fallback")
	}
}

func TestGenerateWhenToUseStatement_RolePattern(t *testing.T) {
	ca := NewContentAnalyzer()
	out := ca.generateWhenToUseStatement("", "This is for debugging and troubleshooting.", "")
	if out == "" || out == ca.fallbackPattern+"." {
		t.Error("Expected a specific when-to-use statement, got fallback")
	}
}

func TestGenerateDescription(t *testing.T) {
	desc := generateDescription("AI Engineer", "", "")
	if desc == "" || desc == "Development specialist" {
		t.Error("Expected a specific short description, got fallback")
	}
}
