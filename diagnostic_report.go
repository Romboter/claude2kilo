package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DiagnosticReport generates detailed reports about conversion issues
type DiagnosticReport struct {
	TotalFiles      int
	SuccessfulFiles int
	FailedFiles     int
	SanitizedFiles  int
	Issues          []FileIssue
	Timestamp       time.Time
}

// FileIssue represents a specific issue with a file
type FileIssue struct {
	FilePath    string
	IssueType   string
	Description string
	Suggestion  string
}

// GenerateDiagnosticReport creates a comprehensive report of conversion results
func GenerateDiagnosticReport(inputDir string, issues []FileIssue, totalFiles, successfulFiles, sanitizedFiles int) error {
	report := DiagnosticReport{
		TotalFiles:      totalFiles,
		SuccessfulFiles: successfulFiles,
		FailedFiles:     totalFiles - successfulFiles,
		SanitizedFiles:  sanitizedFiles,
		Issues:          issues,
		Timestamp:       time.Now(),
	}

	// Generate report content
	content := generateReportContent(report)

	// Save report to file
	reportPath := filepath.Join(inputDir, "conversion-diagnostic-report.md")
	if err := os.WriteFile(reportPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write diagnostic report: %w", err)
	}

	fmt.Printf("\n📊 Diagnostic report saved to: %s\n", reportPath)
	return nil
}

// generateReportContent creates the markdown content for the diagnostic report
func generateReportContent(report DiagnosticReport) string {
	var content strings.Builder

	content.WriteString("# Claude Agent Conversion Diagnostic Report\n\n")
	content.WriteString(fmt.Sprintf("Generated: %s\n\n", report.Timestamp.Format("2006-01-02 15:04:05")))

	// Summary section
	content.WriteString("## Summary\n\n")
	content.WriteString(fmt.Sprintf("- **Total files processed**: %d\n", report.TotalFiles))
	content.WriteString(fmt.Sprintf("- **Successfully converted**: %d (%.1f%%)\n",
		report.SuccessfulFiles, float64(report.SuccessfulFiles)/float64(report.TotalFiles)*100))
	content.WriteString(fmt.Sprintf("- **Failed conversions**: %d (%.1f%%)\n",
		report.FailedFiles, float64(report.FailedFiles)/float64(report.TotalFiles)*100))
	content.WriteString(fmt.Sprintf("- **Files requiring sanitization**: %d (%.1f%%)\n\n",
		report.SanitizedFiles, float64(report.SanitizedFiles)/float64(report.TotalFiles)*100))

	// Success rate analysis
	if report.FailedFiles == 0 {
		content.WriteString("🎉 **Perfect conversion rate!** All files were successfully converted.\n\n")
	} else if report.FailedFiles <= 5 {
		content.WriteString("✅ **Excellent conversion rate!** Only a few files need manual attention.\n\n")
	} else if report.FailedFiles <= 15 {
		content.WriteString("⚠️ **Good conversion rate** with some files needing manual fixes.\n\n")
	} else {
		content.WriteString("❌ **Multiple files need attention** - consider reviewing the patterns below.\n\n")
	}

	// Issues section
	if len(report.Issues) > 0 {
		content.WriteString("## Issues Found\n\n")

		// Group issues by type
		issueGroups := make(map[string][]FileIssue)
		for _, issue := range report.Issues {
			issueGroups[issue.IssueType] = append(issueGroups[issue.IssueType], issue)
		}

		for issueType, issues := range issueGroups {
			content.WriteString(fmt.Sprintf("### %s (%d files)\n\n", issueType, len(issues)))

			for _, issue := range issues {
				content.WriteString(fmt.Sprintf("**File**: `%s`\n", issue.FilePath))
				content.WriteString(fmt.Sprintf("**Issue**: %s\n", issue.Description))
				if issue.Suggestion != "" {
					content.WriteString(fmt.Sprintf("**Suggestion**: %s\n", issue.Suggestion))
				}
				content.WriteString("\n")
			}
		}
	}

	// Recommendations section
	content.WriteString("## Recommendations\n\n")

	if report.SanitizedFiles > 0 {
		content.WriteString(fmt.Sprintf("### YAML Sanitization Applied (%d files)\n\n", report.SanitizedFiles))
		content.WriteString("The converter automatically fixed common YAML issues in these files:\n")
		content.WriteString("- Long descriptions with unescaped quotes → Converted to YAML literal blocks\n")
		content.WriteString("- Tools field as comma-separated string → Converted to YAML array format\n")
		content.WriteString("- Embedded examples and special characters → Properly formatted\n\n")
		content.WriteString("✅ **No action needed** - these files were automatically fixed during conversion.\n\n")
	}

	if report.FailedFiles > 0 {
		content.WriteString("### Manual Fixes Required\n\n")
		content.WriteString("For the remaining failed files, consider these approaches:\n\n")

		content.WriteString("#### Option 1: Manual Fix\n")
		content.WriteString("Edit the files directly to fix the YAML frontmatter issues.\n\n")

		content.WriteString("#### Option 2: LLM-Assisted Fix\n")
		content.WriteString("Use this prompt with an LLM to fix the files:\n\n")
		content.WriteString("```\n")
		content.WriteString("Please fix the YAML frontmatter in this Claude agent file. The issues are:\n")
		content.WriteString("1. Ensure proper YAML syntax with correct quoting\n")
		content.WriteString("2. Convert long descriptions to YAML literal blocks (|) if they contain quotes or examples\n")
		content.WriteString("3. Convert tools field from comma-separated string to YAML array format\n")
		content.WriteString("4. Preserve all content meaning while making it valid YAML\n\n")
		content.WriteString("[Paste the problematic file content here]\n")
		content.WriteString("```\n\n")
	}

	// Best practices section
	content.WriteString("## Best Practices for Future Claude Agent Files\n\n")
	content.WriteString("To avoid conversion issues in the future:\n\n")
	content.WriteString("### YAML Frontmatter Guidelines\n")
	content.WriteString("1. **Use literal blocks for long descriptions**:\n")
	content.WriteString("   ```yaml\n")
	content.WriteString("   description: |\n")
	content.WriteString("     Your long description here\n")
	content.WriteString("     with multiple lines and examples\n")
	content.WriteString("   ```\n\n")
	content.WriteString("2. **Format tools as YAML arrays**:\n")
	content.WriteString("   ```yaml\n")
	content.WriteString("   tools: [Read, Write, Bash, Grep]\n")
	content.WriteString("   ```\n\n")
	content.WriteString("3. **Escape quotes in descriptions**:\n")
	content.WriteString("   ```yaml\n")
	content.WriteString("   description: \"Use this agent when you need to 'quote' something\"\n")
	content.WriteString("   ```\n\n")
	content.WriteString("4. **Required fields**:\n")
	content.WriteString("   - `name`: Agent identifier (kebab-case recommended)\n")
	content.WriteString("   - `description`: What the agent does and when to use it\n\n")

	content.WriteString("## Conversion Statistics\n\n")
	content.WriteString("| Metric | Count | Percentage |\n")
	content.WriteString("|--------|-------|------------|\n")
	content.WriteString(fmt.Sprintf("| Total Files | %d | 100%% |\n", report.TotalFiles))
	content.WriteString(fmt.Sprintf("| Successful | %d | %.1f%% |\n",
		report.SuccessfulFiles, float64(report.SuccessfulFiles)/float64(report.TotalFiles)*100))
	content.WriteString(fmt.Sprintf("| Failed | %d | %.1f%% |\n",
		report.FailedFiles, float64(report.FailedFiles)/float64(report.TotalFiles)*100))
	content.WriteString(fmt.Sprintf("| Auto-Fixed | %d | %.1f%% |\n",
		report.SanitizedFiles, float64(report.SanitizedFiles)/float64(report.TotalFiles)*100))

	return content.String()
}
