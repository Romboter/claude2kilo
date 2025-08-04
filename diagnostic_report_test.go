package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestGenerateDiagnosticReport_CreatesFile(t *testing.T) {
	dir := t.TempDir()
	issues := []FileIssue{{FilePath: "file1.md", IssueType: "Conversion Error", Description: "desc", Suggestion: "fix"}}
	err := GenerateDiagnosticReport(dir, issues, 10, 8, 2)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	path := dir + string(os.PathSeparator) + "conversion-diagnostic-report.md"
	if _, err := os.Stat(path); err != nil {
		t.Errorf("Expected report file to exist, got error: %v", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read report: %v", err)
	}
	if !strings.Contains(string(content), "Diagnostic report saved") && !strings.Contains(string(content), "Summary") {
		t.Error("Expected report content to contain summary or diagnostic report header")
	}
}

func TestGenerateReportContent_ContainsMetrics(t *testing.T) {
	report := DiagnosticReport{
		TotalFiles:      5,
		SuccessfulFiles: 4,
		FailedFiles:     1,
		SanitizedFiles:  2,
		Issues:          []FileIssue{{FilePath: "file.md", IssueType: "Conversion Error", Description: "desc", Suggestion: "fix"}},
		Timestamp:       (func() (tm time.Time) { tm, _ = time.Parse("2006-01-02", "2024-01-01"); return })(),
	}
	content := generateReportContent(report)
	if !strings.Contains(content, "Total files processed") {
		t.Error("Expected summary in report content")
	}
	if !strings.Contains(content, "Conversion Error") {
		t.Error("Expected issue type in report content")
	}
}
