package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		input      = flag.String("input", "", "Input file (.md) or directory containing Claude Code sub-agent files")
		output     = flag.String("output", "./kilo-modes", "Output directory for Kilo Code mode files")
		dryRun     = flag.Bool("dry-run", false, "Show what would be converted without creating files")
		singleFile = flag.Bool("single-files", false, "Output each mode to individual YAML files instead of custom_modes.yaml (directory mode only)")
		help       = flag.Bool("help", false, "Show help message")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Claude Code Sub-agent to Kilo Code Mode Converter\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Convert a single file\n")
		fmt.Fprintf(os.Stderr, "  %s -input ai-engineer.md -output ./kilo-modes/\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n  # Convert all files in current directory to custom_modes.yaml\n")
		fmt.Fprintf(os.Stderr, "  %s -input . -output ./kilo-modes/\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n  # Convert all files to individual YAML files\n")
		fmt.Fprintf(os.Stderr, "  %s -input ./claude-agents/ -output ./kilo-modes/ -single-files\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n  # Dry run to see what would be converted\n")
		fmt.Fprintf(os.Stderr, "  %s -input ./claude-agents/ -output ./converted-modes/ -dry-run\n", os.Args[0])
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *input == "" {
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\nError: input is required\n")
		os.Exit(1)
	}

	converter := NewConverter()

	// Check if input exists
	inputInfo, err := os.Stat(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Input path '%s' does not exist\n", *input)
		os.Exit(1)
	}

	if inputInfo.IsDir() {
		// Convert directory
		if *dryRun {
			fmt.Printf("Dry run mode - showing what would be converted:\n")
		}

		if err := converter.convertDirectory(*input, *output, *dryRun, *singleFile); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Convert single file
		if !strings.HasSuffix(strings.ToLower(*input), ".md") {
			fmt.Fprintf(os.Stderr, "Error: Input file must have .md extension\n")
			os.Exit(1)
		}

		if *dryRun {
			mode, err := converter.convertAgent(*input)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			baseName := strings.TrimSuffix(filepath.Base(*input), filepath.Ext(*input))
			fmt.Printf("Would convert %s to:\n", filepath.Base(*input))
			fmt.Printf("  - %s (in %s.yaml)\n", mode.Slug, baseName)
		} else {
			outputFile, err := converter.convertFile(*input, *output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ Converted %s\n", filepath.Base(*input))
			fmt.Printf("  → %s\n", outputFile)
		}
	}
}
