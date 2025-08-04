package main

import (
	"testing"
)

func TestNewIconSelector(t *testing.T) {
	is := NewIconSelector()
	if is == nil {
		t.Fatal("Expected NewIconSelector to return non-nil")
	}
}

func TestSelectIcon_ExactRole(t *testing.T) {
	is := NewIconSelector()
	icon := is.SelectIcon("architect", "", "")
	if icon != "codicon-type-hierarchy-sub" {
		t.Errorf("Expected codicon-type-hierarchy-sub, got %s", icon)
	}
}

func TestSelectIcon_DomainKeyword(t *testing.T) {
	is := NewIconSelector()
	icon := is.SelectIcon("", "This is an AI agent", "")
	if icon != "codicon-robot" {
		t.Errorf("Expected codicon-robot, got %s", icon)
	}
}

func TestSelectIcon_CharacteristicKeyword(t *testing.T) {
	is := NewIconSelector()
	icon := is.SelectIcon("", "", "This engineer automates tasks")
	if icon != "codicon-run-all" && icon != "codicon-gear" {
		t.Errorf("Expected codicon-run-all or codicon-gear, got %s", icon)
	}
}

func TestSelectIcon_Fallback(t *testing.T) {
	is := NewIconSelector()
	icon := is.SelectIcon("support", "", "")
	if icon != "codicon-person" {
		t.Errorf("Expected codicon-person, got %s", icon)
	}
}

func TestSelectIcon_DefaultFallback(t *testing.T) {
	is := NewIconSelector()
	icon := is.SelectIcon("unknownrole", "", "")
	if icon != "codicon-gear" {
		t.Errorf("Expected codicon-gear, got %s", icon)
	}
}
