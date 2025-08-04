package main

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// YAMLSanitizer handles sanitization of malformed YAML frontmatter
type YAMLSanitizer struct {
	// Patterns for detecting problematic content
	longDescPattern    *regexp.Regexp
	toolsStringPattern *regexp.Regexp
	yamlKeyPattern     *regexp.Regexp
}

// NewYAMLSanitizer creates a new YAML sanitizer with predefined patterns
func NewYAMLSanitizer() *YAMLSanitizer {
	return &YAMLSanitizer{
		// Detect long descriptions that likely contain problematic content
		longDescPattern: regexp.MustCompile(`^(\s*description:\s*)(.{200,}.*)$`),
		// Detect tools field as comma-separated string instead of array
		toolsStringPattern: regexp.MustCompile(`^(\s*tools:\s*)([^[\]]+(?:,\s*[^[\]]+)+)\s*$`),
		// Pattern to match YAML key-value pairs
		yamlKeyPattern: regexp.MustCompile(`^(\s*)(\w+):\s*(.*)$`),
	}
}

// SanitizeFrontmatter attempts to fix malformed YAML frontmatter
func (ys *YAMLSanitizer) SanitizeFrontmatter(content string) (string, error) {
	lines := strings.Split(content, "\n")
	var sanitizedLines []string

	for _, line := range lines {
		sanitizedLine := ys.sanitizeLine(line)
		sanitizedLines = append(sanitizedLines, sanitizedLine)
	}

	result := strings.Join(sanitizedLines, "\n")

	// Validate the sanitized YAML
	var testAgent ClaudeAgent
	if err := yaml.Unmarshal([]byte(result), &testAgent); err != nil {
		return "", fmt.Errorf("sanitization failed to produce valid YAML: %w", err)
	}

	return result, nil
}

// sanitizeLine handles sanitization of individual YAML lines
func (ys *YAMLSanitizer) sanitizeLine(line string) string {
	// Handle tools field as comma-separated string
	if match := ys.toolsStringPattern.FindStringSubmatch(line); match != nil {
		key := match[1]
		value := strings.TrimSpace(match[2])

		// Convert comma-separated string to YAML array format
		tools := strings.Split(value, ",")
		var cleanTools []string
		for _, tool := range tools {
			cleanTools = append(cleanTools, strings.TrimSpace(tool))
		}

		// Format as YAML array
		return key + "[" + strings.Join(cleanTools, ", ") + "]"
	}

	// Handle long descriptions that likely contain problematic content
	if match := ys.longDescPattern.FindStringSubmatch(line); match != nil {
		key := match[1]
		value := strings.TrimSpace(match[2])

		// Check if this description needs literal block formatting
		if ys.needsLiteralBlock(value) {
			// Convert to literal block format
			sanitizedDesc := ys.sanitizeDescriptionContent(value)
			result := key + "|\n"

			// Add description lines with proper indentation
			descLines := strings.Split(sanitizedDesc, "\n")
			for _, descLine := range descLines {
				result += "  " + descLine + "\n"
			}

			// Remove trailing newline
			return strings.TrimSuffix(result, "\n")
		} else {
			// For shorter descriptions, just escape quotes properly
			return key + `"` + strings.ReplaceAll(value, `"`, `\"`) + `"`
		}
	}

	return line
}

// needsLiteralBlock checks if a description value needs literal block formatting
func (ys *YAMLSanitizer) needsLiteralBlock(value string) bool {
	return strings.Contains(value, `"`) &&
		(strings.Contains(value, "<example>") ||
			strings.Contains(value, "Context:") ||
			len(value) > 300)
}

// sanitizeDescriptionContent handles complex description content
func (ys *YAMLSanitizer) sanitizeDescriptionContent(desc string) string {
	// Split on common patterns to create proper line breaks
	desc = strings.ReplaceAll(desc, " <example>", "\n\n<example>")
	desc = strings.ReplaceAll(desc, "</example> ", "</example>\n\n")
	desc = strings.ReplaceAll(desc, " Examples:", "\n\nExamples:")
	desc = strings.ReplaceAll(desc, "Context: ", "\nContext: ")
	desc = strings.ReplaceAll(desc, "user: ", "\nuser: ")
	desc = strings.ReplaceAll(desc, "assistant: ", "\nassistant: ")
	desc = strings.ReplaceAll(desc, "<commentary>", "\n<commentary>")
	desc = strings.ReplaceAll(desc, "</commentary>", "\n</commentary>")

	// Clean up multiple consecutive newlines
	for strings.Contains(desc, "\n\n\n") {
		desc = strings.ReplaceAll(desc, "\n\n\n", "\n\n")
	}

	return strings.TrimSpace(desc)
}

// DetectIssues analyzes YAML content and returns a list of detected issues
func (ys *YAMLSanitizer) DetectIssues(yamlContent string) []string {
	var issues []string

	lines := strings.Split(yamlContent, "\n")
	for i, line := range lines {
		if ys.longDescPattern.MatchString(line) {
			issues = append(issues, fmt.Sprintf("Line %d: Long description with potential quote issues", i+1))
		}

		if ys.toolsStringPattern.MatchString(line) {
			issues = append(issues, fmt.Sprintf("Line %d: Tools field as string instead of array", i+1))
		}
	}

	return issues
}
