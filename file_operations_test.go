package main

import (
	"os"
	"testing"
)

type dummyConverter struct{ Converter }

type dummyMode struct{}

func TestSaveModeConfig_CreatesFile(t *testing.T) {
	c := NewConverter()
	modes := []KiloMode{{Slug: "test", Name: "Test", IconName: "codicon-gear", RoleDefinition: "desc", WhenToUse: "", Description: "desc", Groups: []string{"read"}, CustomInstructions: "", Source: "project", OriginalModel: "opus"}}
	dir := t.TempDir()
	file, err := c.saveModeConfig(modes, dir, "test.yaml")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if _, err := os.Stat(file); err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}
}

func TestSaveSingleModeConfig_CreatesFile(t *testing.T) {
	c := NewConverter()
	mode := KiloMode{Slug: "test", Name: "Test", IconName: "codicon-gear", RoleDefinition: "desc", WhenToUse: "", Description: "desc", Groups: []string{"read"}, CustomInstructions: "", Source: "project", OriginalModel: "opus"}
	dir := t.TempDir()
	file, err := c.saveSingleModeConfig(mode, dir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if _, err := os.Stat(file); err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}
}

func TestSaveSingleModeConfigWithPath_CreatesFile(t *testing.T) {
	c := NewConverter()
	mode := KiloMode{Slug: "test", Name: "Test", IconName: "codicon-gear", RoleDefinition: "desc", WhenToUse: "", Description: "desc", Groups: []string{"read"}, CustomInstructions: "", Source: "project", OriginalModel: "opus"}
	dir := t.TempDir()
	file, err := c.saveSingleModeConfigWithPath(mode, "input/agent.md", "input", dir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if _, err := os.Stat(file); err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}
}

func TestSaveModeConfig_ErrorOnBadDir(t *testing.T) {
	c := NewConverter()
	modes := []KiloMode{{Slug: "test", Name: "Test", IconName: "codicon-gear", RoleDefinition: "desc", WhenToUse: "", Description: "desc", Groups: []string{"read"}, CustomInstructions: "", Source: "project", OriginalModel: "opus"}}
	_, err := c.saveModeConfig(modes, string([]byte{0}), "test.yaml")
	if err == nil {
		t.Error("Expected error for bad directory, got nil")
	}
}
