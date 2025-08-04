package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// NewConverter creates a new converter instance
func NewConverter() *Converter {
	return &Converter{
		modelMapping: map[string]string{
			"opus":   "anthropic/claude-opus-3",
			"sonnet": "anthropic/claude-sonnet-3.5",
			"haiku":  "anthropic/claude-haiku-3",
		},
		defaultGroups: map[string][]string{
			"full":      {"read", "edit", "browser", "command", "mcp"},
			"review":    {"read", "edit"},
			"architect": {"read", "edit"},
			"web":       {"read", "edit", "browser", "command"},
			"system":    {"read", "edit", "command"},
			"default":   {"read", "edit", "browser", "command"},
		},
		frontmatterRe:   regexp.MustCompile(`(?s)^---\n(.*?)\n---\s*\n(.*)$`),
		slugRe:          regexp.MustCompile(`[^a-z0-9]+`),
		iconSelector:    NewIconSelector(),
		contentAnalyzer: NewContentAnalyzer(),
		yamlSanitizer:   NewYAMLSanitizer(),
	}
}

// parseFrontmatter extracts YAML frontmatter and markdown content
func (c *Converter) parseFrontmatter(content string) (*ClaudeAgent, string, error) {
	// Normalize line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	trimmed := strings.TrimSpace(content)

	lines := strings.Split(trimmed, "\n")
	if len(lines) < 3 || lines[0] != "---" {
		return nil, "", fmt.Errorf("no valid YAML frontmatter found - first line: %q", lines[0])
	}

	// Find the closing ---
	var yamlEnd int
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			yamlEnd = i
			break
		}
	}

	if yamlEnd == 0 {
		return nil, "", fmt.Errorf("no closing --- found")
	}

	yamlContent := strings.Join(lines[1:yamlEnd], "\n")
	markdownContent := strings.Join(lines[yamlEnd+1:], "\n")

	var agent ClaudeAgent
	err := yaml.Unmarshal([]byte(yamlContent), &agent)

	if err != nil {
		// Try sanitization
		sanitizedYAML, sanitizeErr := c.yamlSanitizer.SanitizeFrontmatter(yamlContent)
		if sanitizeErr != nil {
			return nil, "", fmt.Errorf("YAML parsing failed, sanitization also failed: original error: %w, sanitization error: %v", err, sanitizeErr)
		}

		// Retry with sanitized content
		err = yaml.Unmarshal([]byte(sanitizedYAML), &agent)
		if err != nil {
			return nil, "", fmt.Errorf("YAML parsing failed even after sanitization: %w", err)
		}

		// Log successful sanitization
		fmt.Printf("  ⚠ Applied YAML sanitization\n")
	}

	if agent.Name == "" {
		return nil, "", fmt.Errorf("missing required 'name' field")
	}
	if agent.Description == "" {
		return nil, "", fmt.Errorf("missing required 'description' field")
	}

	return &agent, strings.TrimSpace(markdownContent), nil
}

// generateSlug creates a URL-friendly slug from the agent name
func (c *Converter) generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = c.slugRe.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

// determineGroups selects appropriate tool groups based on agent characteristics
func (c *Converter) determineGroups(name, description, content string) []string {
	text := strings.ToLower(fmt.Sprintf("%s %s %s", name, description, content))

	// Review-only agents (code reviewers, auditors)
	if (strings.Contains(text, "review") || strings.Contains(text, "reviewer") || strings.Contains(text, "audit")) &&
		!strings.Contains(text, "architect") {
		return c.defaultGroups["review"]
	}

	// Architect modes typically only edit markdown files
	if strings.Contains(text, "architect") && strings.Contains(text, "review") {
		return c.defaultGroups["architect"]
	}

	// Web/frontend development (needs browser for testing)
	if strings.Contains(text, "frontend") || strings.Contains(text, "react") || strings.Contains(text, "ui") ||
		strings.Contains(text, "css") || strings.Contains(text, "html") || strings.Contains(text, "web") {
		return c.defaultGroups["web"]
	}

	// System/backend development (needs command line tools)
	if strings.Contains(text, "backend") || strings.Contains(text, "api") || strings.Contains(text, "server") ||
		strings.Contains(text, "database") || strings.Contains(text, "devops") || strings.Contains(text, "deploy") ||
		strings.Contains(text, "infrastructure") || strings.Contains(text, "cloud") || strings.Contains(text, "system") {
		return c.defaultGroups["system"]
	}

	// AI/ML engineers and complex development (full access including MCP)
	if strings.Contains(text, "ai") || strings.Contains(text, "llm") || strings.Contains(text, "ml") ||
		strings.Contains(text, "data") || strings.Contains(text, "analytics") || strings.Contains(text, "engineer") ||
		strings.Contains(text, "rag") || strings.Contains(text, "vector") || strings.Contains(text, "embedding") {
		return c.defaultGroups["full"]
	}

	// Default for most other cases
	return c.defaultGroups["default"]
}

// determineFileRestrictions sets file access restrictions based on agent type
func (c *Converter) determineFileRestrictions(name, description, content string) (string, string) {
	text := strings.ToLower(fmt.Sprintf("%s %s %s", name, description, content))

	// Architect modes typically only edit markdown files
	if strings.Contains(text, "architect") && strings.Contains(text, "review") {
		return "\\.md$", "Markdown files only"
	}

	return "", ""
}

// generateWhenToUse creates a description of when to use this mode using intelligent content analysis
func (c *Converter) generateWhenToUse(name, description, content string) string {
	return c.contentAnalyzer.generateWhenToUseStatement(name, description, content)
}

// convertAgent converts a Claude Code sub-agent to Kilo Code mode
func (c *Converter) convertAgent(filePath string) (*KiloMode, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	agent, markdown, err := c.parseFrontmatter(string(content))
	if err != nil {
		return nil, err
	}

	slug := c.generateSlug(agent.Name)
	groups := c.determineGroups(agent.Name, agent.Description, markdown)
	fileRegex, fileDesc := c.determineFileRestrictions(agent.Name, agent.Description, markdown)

	// Generate icon and description
	iconName := c.iconSelector.SelectIcon(agent.Name, agent.Description, markdown)
	shortDescription := generateDescription(agent.Name, agent.Description, markdown)

	// Format name properly (capitalize each word)
	nameParts := strings.Split(agent.Name, "-")
	for i, part := range nameParts {
		if len(part) > 0 {
			nameParts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	formattedName := strings.Join(nameParts, " ")

	// Generate whenToUse description based on agent characteristics
	whenToUse := c.generateWhenToUse(agent.Name, agent.Description, markdown)

	mode := &KiloMode{
		Slug:               slug,
		Name:               formattedName,
		IconName:           iconName,
		RoleDefinition:     agent.Description,
		WhenToUse:          whenToUse,
		Description:        shortDescription,
		Groups:             groups,
		CustomInstructions: markdown,
		Source:             "project", // Default to project, user can change on import
		OriginalModel:      agent.Model,
	}

	if fileRegex != "" {
		mode.FileRegex = fileRegex
		// For file restrictions, override the generated description
		mode.Description = fileDesc
	}

	return mode, nil
}

// convertAgentWithStats converts a Claude Code sub-agent to Kilo Code mode and returns sanitization stats
func (c *Converter) convertAgentWithStats(filePath string) (*KiloMode, bool, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, false, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	agent, markdown, wasSanitized, err := c.parseFrontmatterWithStats(string(content))
	if err != nil {
		return nil, false, err
	}

	slug := c.generateSlug(agent.Name)
	groups := c.determineGroups(agent.Name, agent.Description, markdown)
	fileRegex, fileDesc := c.determineFileRestrictions(agent.Name, agent.Description, markdown)

	// Generate icon and description
	iconName := c.iconSelector.SelectIcon(agent.Name, agent.Description, markdown)
	shortDescription := generateDescription(agent.Name, agent.Description, markdown)

	// Format name properly (capitalize each word)
	nameParts := strings.Split(agent.Name, "-")
	for i, part := range nameParts {
		if len(part) > 0 {
			nameParts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		}
	}
	formattedName := strings.Join(nameParts, " ")

	// Generate whenToUse description based on agent characteristics
	whenToUse := c.generateWhenToUse(agent.Name, agent.Description, markdown)

	mode := &KiloMode{
		Slug:               slug,
		Name:               formattedName,
		IconName:           iconName,
		RoleDefinition:     agent.Description,
		WhenToUse:          whenToUse,
		Description:        shortDescription,
		Groups:             groups,
		CustomInstructions: markdown,
		Source:             "project", // Default to project, user can change on import
		OriginalModel:      agent.Model,
	}

	if fileRegex != "" {
		mode.FileRegex = fileRegex
		// For file restrictions, override the generated description
		mode.Description = fileDesc
	}

	return mode, wasSanitized, nil
}

// parseFrontmatterWithStats extracts YAML frontmatter and markdown content with sanitization tracking
func (c *Converter) parseFrontmatterWithStats(content string) (*ClaudeAgent, string, bool, error) {
	// Normalize line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	trimmed := strings.TrimSpace(content)

	lines := strings.Split(trimmed, "\n")
	if len(lines) < 3 || lines[0] != "---" {
		return nil, "", false, fmt.Errorf("no valid YAML frontmatter found - first line: %q", lines[0])
	}

	// Find the closing ---
	var yamlEnd int
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			yamlEnd = i
			break
		}
	}

	if yamlEnd == 0 {
		return nil, "", false, fmt.Errorf("no closing --- found")
	}

	yamlContent := strings.Join(lines[1:yamlEnd], "\n")
	markdownContent := strings.Join(lines[yamlEnd+1:], "\n")

	var agent ClaudeAgent
	wasSanitized := false
	err := yaml.Unmarshal([]byte(yamlContent), &agent)

	if err != nil {
		// Try sanitization
		sanitizedYAML, sanitizeErr := c.yamlSanitizer.SanitizeFrontmatter(yamlContent)
		if sanitizeErr != nil {
			return nil, "", false, fmt.Errorf("YAML parsing failed, sanitization also failed: original error: %w, sanitization error: %v", err, sanitizeErr)
		}

		// Retry with sanitized content
		err = yaml.Unmarshal([]byte(sanitizedYAML), &agent)
		if err != nil {
			return nil, "", false, fmt.Errorf("YAML parsing failed even after sanitization: %w", err)
		}

		wasSanitized = true
		// Log successful sanitization
		fmt.Printf("  ⚠ Applied YAML sanitization\n")
	}

	if agent.Name == "" {
		return nil, "", wasSanitized, fmt.Errorf("missing required 'name' field")
	}
	if agent.Description == "" {
		return nil, "", wasSanitized, fmt.Errorf("missing required 'description' field")
	}

	return &agent, strings.TrimSpace(markdownContent), wasSanitized, nil
}
