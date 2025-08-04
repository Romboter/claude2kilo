package main

import (
	"regexp"
)

// ClaudeAgent represents the frontmatter of a Claude Code sub-agent
type ClaudeAgent struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Model       string   `yaml:"model"`
	Tools       []string `yaml:"tools,omitempty"`
}

// KiloMode represents a Kilo Code mode configuration
type KiloMode struct {
	Slug               string   `yaml:"slug"`
	Name               string   `yaml:"name"`
	IconName           string   `yaml:"iconName"` // NEW FIELD
	RoleDefinition     string   `yaml:"roleDefinition"`
	WhenToUse          string   `yaml:"whenToUse,omitempty"`
	Description        string   `yaml:"description"` // NEW FIELD (now included in YAML)
	Groups             []string `yaml:"groups"`
	CustomInstructions string   `yaml:"customInstructions"`
	Source             string   `yaml:"source"`
	FileRegex          string   `yaml:"-"` // Not included in YAML output
	OriginalModel      string   `yaml:"-"` // Not included in YAML output
}

// CustomModesFile represents the root structure for Kilo Code custom modes
type CustomModesFile struct {
	CustomModes []KiloMode `yaml:"customModes"`
}

// IconSelector handles intelligent icon selection
type IconSelector struct {
	exactRoleMap           map[string]string
	domainKeywords         map[string]string
	characteristicKeywords map[string]string
	fallbackMap            map[string]string
	validIcons             map[string]bool
}

// ContentAnalyzer handles intelligent content analysis for generating "when to use" statements
type ContentAnalyzer struct {
	rolePatterns    map[string]string
	domainPatterns  map[string]string
	actionPatterns  map[string]string
	fallbackPattern string
}

// Converter handles the conversion logic
type Converter struct {
	modelMapping    map[string]string
	defaultGroups   map[string][]string
	frontmatterRe   *regexp.Regexp
	slugRe          *regexp.Regexp
	iconSelector    *IconSelector
	contentAnalyzer *ContentAnalyzer
	yamlSanitizer   *YAMLSanitizer
}
