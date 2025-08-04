package main

import (
	"reflect"
	"testing"
)

func TestClaudeAgentFields(t *testing.T) {
	agent := ClaudeAgent{Name: "n", Description: "d", Model: "m", Tools: []string{"t1", "t2"}}
	if agent.Name != "n" || agent.Description != "d" || agent.Model != "m" {
		t.Error("ClaudeAgent fields not set correctly")
	}
	if !reflect.DeepEqual(agent.Tools, []string{"t1", "t2"}) {
		t.Error("ClaudeAgent tools not set correctly")
	}
}

func TestKiloModeFields(t *testing.T) {
	mode := KiloMode{Slug: "s", Name: "n", IconName: "i", RoleDefinition: "r", WhenToUse: "w", Description: "d", Groups: []string{"g"}, CustomInstructions: "c", Source: "src", FileRegex: "f", OriginalModel: "o"}
	if mode.Slug != "s" || mode.Name != "n" || mode.IconName != "i" || mode.RoleDefinition != "r" || mode.WhenToUse != "w" || mode.Description != "d" || mode.Source != "src" || mode.FileRegex != "f" || mode.OriginalModel != "o" {
		t.Error("KiloMode fields not set correctly")
	}
	if !reflect.DeepEqual(mode.Groups, []string{"g"}) {
		t.Error("KiloMode groups not set correctly")
	}
}

func TestCustomModesFile(t *testing.T) {
	m := KiloMode{Name: "n"}
	cmf := CustomModesFile{CustomModes: []KiloMode{m}}
	if len(cmf.CustomModes) != 1 || cmf.CustomModes[0].Name != "n" {
		t.Error("CustomModesFile not set correctly")
	}
}

func TestIconSelectorStruct(t *testing.T) {
	is := &IconSelector{}
	if is == nil {
		t.Error("IconSelector struct not instantiable")
	}
}

func TestContentAnalyzerStruct(t *testing.T) {
	ca := &ContentAnalyzer{}
	if ca == nil {
		t.Error("ContentAnalyzer struct not instantiable")
	}
}

func TestConverterStruct(t *testing.T) {
	c := &Converter{}
	if c == nil {
		t.Error("Converter struct not instantiable")
	}
}
