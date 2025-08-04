package main

import (
	"testing"
)

func TestNewConverter(t *testing.T) {
	c := NewConverter()
	if c == nil {
		t.Fatal("Expected NewConverter to return non-nil")
	}
}

func TestGenerateSlug(t *testing.T) {
	c := NewConverter()
	got := c.generateSlug("Test Name 123!")
	want := "test-name-123"
	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestParseFrontmatter_Valid(t *testing.T) {
	c := NewConverter()
	input := `---
name: test-agent
description: test desc
model: opus
---
Some markdown content.`
	agent, md, err := c.parseFrontmatter(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if agent.Name != "test-agent" || agent.Description != "test desc" {
		t.Errorf("Unexpected agent fields: %+v", agent)
	}
	if md != "Some markdown content." {
		t.Errorf("Unexpected markdown: %q", md)
	}
}

func TestParseFrontmatter_Invalid(t *testing.T) {
	c := NewConverter()
	input := `---
name: test-agent
---
No desc.`
	_, _, err := c.parseFrontmatter(input)
	if err == nil {
		t.Error("Expected error for missing description, got nil")
	}
}

func TestDetermineGroups(t *testing.T) {
	c := NewConverter()
	groups := c.determineGroups("AI Engineer", "", "")
	if len(groups) == 0 {
		t.Error("Expected non-empty groups")
	}
}

func TestDetermineFileRestrictions(t *testing.T) {
	c := NewConverter()
	re, desc := c.determineFileRestrictions("Architect Reviewer", "", "")
	if re != "\\.md$" || desc != "Markdown files only" {
		t.Errorf("Expected markdown restriction, got %q, %q", re, desc)
	}
}

func TestConvertAgent_FileNotFound(t *testing.T) {
	c := NewConverter()
	_, err := c.convertAgent("nonexistentfile.md")
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}
}
