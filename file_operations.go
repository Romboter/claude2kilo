package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// saveModeConfig saves the Kilo Code mode configuration as YAML
func (c *Converter) saveModeConfig(modes []KiloMode, outputDir, filename string) (string, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create the custom modes file structure
	customModesFile := CustomModesFile{
		CustomModes: modes,
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(customModesFile)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Post-process YAML to match standard Kilo format
	yamlString := string(yamlData)

	// Use folded scalar format for multi-line strings (>- instead of |-)
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |-", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |-", "customInstructions: >-")
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |", "customInstructions: >-")

	// Fix indentation to match standard format (2 spaces instead of 4)
	lines := strings.Split(yamlString, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "    ") {
			lines[i] = "  " + strings.TrimPrefix(line, "    ")
		}
	}
	yamlString = strings.Join(lines, "\n")

	// Save YAML file
	outputFile := filepath.Join(outputDir, filename)
	if err := os.WriteFile(outputFile, []byte(yamlString), 0644); err != nil {
		return "", fmt.Errorf("failed to write YAML file: %w", err)
	}

	return outputFile, nil
}

// saveSingleModeConfig saves a single Kilo Code mode configuration as YAML
func (c *Converter) saveSingleModeConfig(mode KiloMode, outputDir string) (string, error) {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create the custom modes file structure with single mode
	customModesFile := CustomModesFile{
		CustomModes: []KiloMode{mode},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(customModesFile)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Post-process YAML to match standard Kilo format
	yamlString := string(yamlData)

	// Use folded scalar format for multi-line strings (>- instead of |-)
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |-", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |-", "customInstructions: >-")
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |", "customInstructions: >-")

	// Fix indentation to match standard format (2 spaces instead of 4)
	lines := strings.Split(yamlString, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "    ") {
			lines[i] = "  " + strings.TrimPrefix(line, "    ")
		}
	}
	yamlString = strings.Join(lines, "\n")

	// Generate filename based on mode slug
	outputFile := filepath.Join(outputDir, mode.Slug+".yaml")
	if err := os.WriteFile(outputFile, []byte(yamlString), 0644); err != nil {
		return "", fmt.Errorf("failed to write YAML file: %w", err)
	}

	return outputFile, nil
}

// saveSingleModeConfigWithPath saves a single Kilo Code mode configuration as YAML with preserved folder structure
func (c *Converter) saveSingleModeConfigWithPath(mode KiloMode, inputPath, inputDir, outputDir string) (string, error) {
	// Calculate the relative path from the input directory
	relPath, err := filepath.Rel(inputDir, inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to calculate relative path: %w", err)
	}

	// Get the directory part of the relative path
	relDir := filepath.Dir(relPath)

	// Create the output directory structure
	var fullOutputDir string
	if relDir == "." {
		// File is in the root of input directory
		fullOutputDir = outputDir
	} else {
		fullOutputDir = filepath.Join(outputDir, relDir)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(fullOutputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create the custom modes file structure with single mode
	customModesFile := CustomModesFile{
		CustomModes: []KiloMode{mode},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(customModesFile)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Post-process YAML to match standard Kilo format
	yamlString := string(yamlData)

	// Use folded scalar format for multi-line strings (>- instead of |-)
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |-", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |-", "customInstructions: >-")
	yamlString = strings.ReplaceAll(yamlString, "roleDefinition: |", "roleDefinition:")
	yamlString = strings.ReplaceAll(yamlString, "customInstructions: |", "customInstructions: >-")

	// Fix indentation to match standard format (2 spaces instead of 4)
	lines := strings.Split(yamlString, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "    ") {
			lines[i] = "  " + strings.TrimPrefix(line, "    ")
		}
	}
	yamlString = strings.Join(lines, "\n")

	// Generate filename based on mode slug
	outputFile := filepath.Join(fullOutputDir, mode.Slug+".yaml")
	if err := os.WriteFile(outputFile, []byte(yamlString), 0644); err != nil {
		return "", fmt.Errorf("failed to write YAML file: %w", err)
	}

	return outputFile, nil
}

// convertFile converts a single file
func (c *Converter) convertFile(inputFile, outputDir string) (string, error) {
	mode, err := c.convertAgent(inputFile)
	if err != nil {
		return "", err
	}

	// For single file processing, we need to preserve the folder structure
	// Get the absolute path of the input file
	absInputFile, err := filepath.Abs(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path for input file: %w", err)
	}

	// Try to find a common base directory (claude-agents) to preserve relative structure
	var relativeOutputPath string

	// Check if the input file is within a claude-agents directory
	if strings.Contains(absInputFile, "claude-agents") {
		// Find the claude-agents part and everything after it
		parts := strings.Split(filepath.ToSlash(absInputFile), "/")
		claudeAgentsIndex := -1
		for i, part := range parts {
			if part == "claude-agents" {
				claudeAgentsIndex = i
				break
			}
		}

		if claudeAgentsIndex >= 0 && claudeAgentsIndex < len(parts)-1 {
			// Get the relative path from claude-agents onwards
			relativeParts := parts[claudeAgentsIndex+1:]
			relativeDir := strings.Join(relativeParts[:len(relativeParts)-1], "/")
			baseName := strings.TrimSuffix(parts[len(parts)-1], filepath.Ext(parts[len(parts)-1]))

			if relativeDir != "" {
				relativeOutputPath = filepath.Join(relativeDir, baseName+".yaml")
			} else {
				relativeOutputPath = baseName + ".yaml"
			}
		} else {
			// Fallback to just the base name
			baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
			relativeOutputPath = baseName + ".yaml"
		}
	} else {
		// If not in claude-agents directory, just use the base name
		baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
		relativeOutputPath = baseName + ".yaml"
	}

	// Create the full output directory path
	fullOutputDir := filepath.Join(outputDir, filepath.Dir(relativeOutputPath))
	outputFilename := filepath.Base(relativeOutputPath)

	modes := []KiloMode{*mode}
	return c.saveModeConfig(modes, fullOutputDir, outputFilename)
}

// convertDirectory converts all .md files in a directory
func (c *Converter) convertDirectory(inputDir, outputDir string, dryRun bool, singleFiles bool) error {
	var successful, total, sanitized int
	var allModes []KiloMode
	var issues []FileIssue

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}

		total++

		mode, wasSanitized, err := c.convertAgentWithStats(path)
		if err != nil {
			// Record the issue for diagnostic report
			issue := FileIssue{
				FilePath:    strings.TrimPrefix(path, inputDir+string(filepath.Separator)),
				IssueType:   "Conversion Error",
				Description: err.Error(),
				Suggestion:  "Check YAML frontmatter syntax and required fields",
			}
			issues = append(issues, issue)

			if dryRun {
				fmt.Printf("  ✗ %s → Error: %v\n", d.Name(), err)
			} else {
				fmt.Printf("✗ Failed to convert %s: %v\n", d.Name(), err)
			}
			return nil
		}

		if wasSanitized {
			sanitized++
		}

		if dryRun {
			if singleFiles {
				fmt.Printf("  ✓ %s → %s (in %s.yaml)\n", d.Name(), mode.Slug, mode.Slug)
			} else {
				fmt.Printf("  ✓ %s → %s (in custom_modes.yaml)\n", d.Name(), mode.Slug)
			}
		} else {
			if singleFiles {
				// Save individual file with preserved folder structure
				outputFile, err := c.saveSingleModeConfigWithPath(*mode, path, inputDir, outputDir)
				if err != nil {
					fmt.Printf("✗ Failed to save %s: %v\n", mode.Slug, err)
					return nil
				}
				fmt.Printf("✓ Converted %s → %s\n", d.Name(), outputFile)
			} else {
				allModes = append(allModes, *mode)
				fmt.Printf("✓ Converted %s → %s\n", d.Name(), mode.Slug)
			}
		}
		successful++

		return nil
	})

	if err != nil {
		return err
	}

	// Generate diagnostic report
	if err := GenerateDiagnosticReport(inputDir, issues, total, successful, sanitized); err != nil {
		fmt.Printf("Warning: Failed to generate diagnostic report: %v\n", err)
	}

	if dryRun {
		if singleFiles {
			fmt.Printf("Would convert %d files to individual YAML files\n", successful)
		} else {
			fmt.Printf("Would convert %d files to custom_modes.yaml\n", successful)
		}
		if sanitized > 0 {
			fmt.Printf("Note: %d files would require YAML sanitization\n", sanitized)
		}
	} else if !singleFiles && len(allModes) > 0 {
		// Only save combined file if not in single files mode
		outputFile, err := c.saveModeConfig(allModes, outputDir, "custom_modes.yaml")
		if err != nil {
			return fmt.Errorf("failed to save modes: %w", err)
		}
		fmt.Printf("\nConversion complete: %d/%d files converted successfully\n", successful, total)
		if sanitized > 0 {
			fmt.Printf("YAML sanitization applied to %d files\n", sanitized)
		}
		fmt.Printf("Output file: %s\n", outputFile)
	} else if singleFiles && successful > 0 {
		// Summary for single files mode
		fmt.Printf("\nConversion complete: %d/%d files converted successfully\n", successful, total)
		if sanitized > 0 {
			fmt.Printf("YAML sanitization applied to %d files\n", sanitized)
		}
		fmt.Printf("Output directory: %s\n", outputDir)
	}

	return nil
}
