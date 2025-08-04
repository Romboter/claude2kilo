package main

import "strings"

// NewIconSelector creates a new icon selector with predefined mappings
func NewIconSelector() *IconSelector {
	return &IconSelector{
		exactRoleMap: map[string]string{
			"architect":            "codicon-type-hierarchy-sub",
			"debugger":             "codicon-bug",
			"code-reviewer":        "codicon-code-review",
			"security-auditor":     "codicon-shield",
			"database-admin":       "codicon-database",
			"devops":               "codicon-gear",
			"prompt-engineer":      "codicon-wand",
			"ai-engineer":          "codicon-robot",
			"ml-engineer":          "codicon-beaker",
			"data-scientist":       "codicon-graph-line",
			"data-engineer":        "codicon-database",
			"frontend-developer":   "codicon-browser",
			"backend-developer":    "codicon-server",
			"mobile-developer":     "codicon-device-mobile",
			"ui-ux-designer":       "codicon-paintcan",
			"test-automator":       "codicon-beaker-stop",
			"performance-engineer": "codicon-pulse",
			"deployment-engineer":  "codicon-rocket",
			"network-engineer":     "codicon-broadcast",
			"cloud-architect":      "codicon-cloud",
			"incident-responder":   "codicon-warning",
			"legal-advisor":        "codicon-law",
			"business-analyst":     "codicon-briefcase",
			"content-marketer":     "codicon-megaphone",
			"customer-support":     "codicon-person",
		},
		domainKeywords: map[string]string{
			// AI/ML Domain
			"ai":        "codicon-robot",
			"llm":       "codicon-copilot",
			"ml":        "codicon-beaker",
			"machine":   "codicon-beaker",
			"learning":  "codicon-lightbulb",
			"neural":    "codicon-circuit-board",
			"embedding": "codicon-symbol-array",
			"vector":    "codicon-symbol-array",
			"rag":       "codicon-search",
			"prompt":    "codicon-wand",

			// Web/Frontend
			"frontend":   "codicon-browser",
			"react":      "codicon-symbol-interface",
			"ui":         "codicon-paintcan",
			"ux":         "codicon-paintcan",
			"css":        "codicon-color-mode",
			"html":       "codicon-code",
			"web":        "codicon-globe",
			"mobile":     "codicon-device-mobile",
			"responsive": "codicon-device-mobile",

			// Backend/Infrastructure
			"backend":    "codicon-server",
			"api":        "codicon-plug",
			"server":     "codicon-server-process",
			"database":   "codicon-database",
			"sql":        "codicon-table",
			"cloud":      "codicon-cloud",
			"devops":     "codicon-gear",
			"docker":     "codicon-package",
			"kubernetes": "codicon-organization",
			"terraform":  "codicon-tools",
			"aws":        "codicon-cloud",
			"azure":      "codicon-azure",

			// Security
			"security":      "codicon-shield",
			"audit":         "codicon-verified",
			"compliance":    "codicon-law",
			"penetration":   "codicon-bug",
			"vulnerability": "codicon-warning",

			// Testing/Quality
			"test":        "codicon-beaker",
			"testing":     "codicon-beaker-stop",
			"qa":          "codicon-pass",
			"quality":     "codicon-verified",
			"debug":       "codicon-bug",
			"performance": "codicon-pulse",

			// Data/Analytics
			"data":          "codicon-graph-line",
			"analytics":     "codicon-pie-chart",
			"etl":           "codicon-arrow-swap",
			"pipeline":      "codicon-arrow-right",
			"warehouse":     "codicon-database",
			"visualization": "codicon-graph-scatter",

			// Languages
			"python":     "codicon-python",
			"javascript": "codicon-symbol-method",
			"typescript": "codicon-symbol-interface",
			"java":       "codicon-symbol-class",
			"golang":     "codicon-symbol-method",
			"rust":       "codicon-gear",
			"cpp":        "codicon-symbol-structure",
			"csharp":     "codicon-symbol-class",
			"php":        "codicon-code",
		},
		characteristicKeywords: map[string]string{
			// Functional roles
			"architect":  "codicon-type-hierarchy-sub",
			"engineer":   "codicon-gear",
			"developer":  "codicon-code",
			"specialist": "codicon-star-full",
			"expert":     "codicon-verified",
			"pro":        "codicon-star-full",
			"admin":      "codicon-person",
			"manager":    "codicon-briefcase",
			"lead":       "codicon-organization",
			"senior":     "codicon-mortar-board",

			// Action-oriented
			"build":        "codicon-tools",
			"deploy":       "codicon-rocket",
			"monitor":      "codicon-eye",
			"optimize":     "codicon-pulse",
			"automate":     "codicon-run-all",
			"integrate":    "codicon-plug",
			"migrate":      "codicon-arrow-swap",
			"modernize":    "codicon-lightbulb-autofix",
			"troubleshoot": "codicon-search",
			"review":       "codicon-eye",
			"document":     "codicon-book",
			"analyze":      "codicon-inspect",
		},
		fallbackMap: map[string]string{
			"development": "codicon-code",
			"engineering": "codicon-gear",
			"design":      "codicon-paintcan",
			"analysis":    "codicon-inspect",
			"management":  "codicon-briefcase",
			"support":     "codicon-person",
			"research":    "codicon-telescope",
			"consulting":  "codicon-comment-discussion",
			"default":     "codicon-gear",
		},
		validIcons: createValidIconsSet(),
	}
}

// createValidIconsSet creates a set of all valid codicon names
func createValidIconsSet() map[string]bool {
	return map[string]bool{
		"codicon-account": true, "codicon-activate-breakpoints": true, "codicon-add": true,
		"codicon-archive": true, "codicon-arrow-both": true, "codicon-arrow-circle-down": true,
		"codicon-arrow-circle-left": true, "codicon-arrow-circle-right": true, "codicon-arrow-circle-up": true,
		"codicon-arrow-down": true, "codicon-arrow-left": true, "codicon-arrow-right": true,
		"codicon-arrow-small-down": true, "codicon-arrow-small-left": true, "codicon-arrow-small-right": true,
		"codicon-arrow-small-up": true, "codicon-arrow-swap": true, "codicon-arrow-up": true,
		"codicon-attach": true, "codicon-azure-devops": true, "codicon-azure": true,
		"codicon-beaker-stop": true, "codicon-beaker": true, "codicon-bell-dot": true,
		"codicon-bell-slash-dot": true, "codicon-bell-slash": true, "codicon-bell": true,
		"codicon-blank": true, "codicon-bold": true, "codicon-book": true, "codicon-bookmark": true,
		"codicon-bracket-dot": true, "codicon-bracket-error": true, "codicon-briefcase": true,
		"codicon-broadcast": true, "codicon-browser": true, "codicon-bug": true, "codicon-calendar": true,
		"codicon-call-incoming": true, "codicon-call-outgoing": true, "codicon-case-sensitive": true,
		"codicon-chat-sparkle": true, "codicon-check-all": true, "codicon-check": true,
		"codicon-checklist": true, "codicon-chevron-down": true, "codicon-chevron-left": true,
		"codicon-chevron-right": true, "codicon-chevron-up": true, "codicon-chip": true,
		"codicon-chrome-close": true, "codicon-chrome-maximize": true, "codicon-chrome-minimize": true,
		"codicon-chrome-restore": true, "codicon-circle-filled": true, "codicon-circle-large-filled": true,
		"codicon-circle-large": true, "codicon-circle-slash": true, "codicon-circle-small-filled": true,
		"codicon-circle-small": true, "codicon-circle": true, "codicon-circuit-board": true,
		"codicon-clear-all": true, "codicon-clippy": true, "codicon-close-all": true, "codicon-close": true,
		"codicon-cloud-download": true, "codicon-cloud-upload": true, "codicon-cloud": true,
		"codicon-code-oss": true, "codicon-code-review": true, "codicon-code": true, "codicon-coffee": true,
		"codicon-collapse-all": true, "codicon-color-mode": true, "codicon-combine": true,
		"codicon-comment-discussion": true, "codicon-comment-draft": true, "codicon-comment-unresolved": true,
		"codicon-comment": true, "codicon-compass-active": true, "codicon-compass-dot": true,
		"codicon-compass": true, "codicon-copilot-blocked": true, "codicon-copilot-error": true,
		"codicon-copilot-in-progress": true, "codicon-copilot-large": true, "codicon-copilot-not-connected": true,
		"codicon-copilot-snooze": true, "codicon-copilot-success": true, "codicon-copilot-unavailable": true,
		"codicon-copilot-warning-large": true, "codicon-copilot-warning": true, "codicon-copilot": true,
		"codicon-copy": true, "codicon-coverage": true, "codicon-credit-card": true, "codicon-dash": true,
		"codicon-dashboard": true, "codicon-database": true, "codicon-debug-all": true,
		"codicon-debug-alt-small": true, "codicon-debug-alt": true, "codicon-debug-breakpoint-conditional-unverified": true,
		"codicon-debug-breakpoint-conditional": true, "codicon-debug-breakpoint-data-unverified": true,
		"codicon-debug-breakpoint-data": true, "codicon-debug-breakpoint-function-unverified": true,
		"codicon-debug-breakpoint-function": true, "codicon-debug-breakpoint-log-unverified": true,
		"codicon-debug-breakpoint-log": true, "codicon-debug-breakpoint-unsupported": true,
		"codicon-debug-console": true, "codicon-debug-continue-small": true, "codicon-debug-continue": true,
		"codicon-debug-coverage": true, "codicon-debug-disconnect": true, "codicon-debug-line-by-line": true,
		"codicon-debug-pause": true, "codicon-debug-rerun": true, "codicon-debug-restart-frame": true,
		"codicon-debug-restart": true, "codicon-debug-reverse-continue": true, "codicon-debug-stackframe-active": true,
		"codicon-debug-stackframe": true, "codicon-debug-start": true, "codicon-debug-step-back": true,
		"codicon-debug-step-into": true, "codicon-debug-step-out": true, "codicon-debug-step-over": true,
		"codicon-debug-stop": true, "codicon-debug": true, "codicon-desktop-download": true,
		"codicon-device-camera-video": true, "codicon-device-camera": true, "codicon-device-mobile": true,
		"codicon-diff-added": true, "codicon-diff-ignored": true, "codicon-diff-modified": true,
		"codicon-diff-multiple": true, "codicon-diff-removed": true, "codicon-diff-renamed": true,
		"codicon-diff-single": true, "codicon-diff": true, "codicon-discard": true,
		"codicon-edit-session": true, "codicon-edit-sparkle": true, "codicon-edit": true,
		"codicon-editor-layout": true, "codicon-ellipsis": true, "codicon-empty-window": true,
		"codicon-error-small": true, "codicon-error": true, "codicon-exclude": true,
		"codicon-expand-all": true, "codicon-export": true, "codicon-extensions-large": true,
		"codicon-extensions": true, "codicon-eye-closed": true, "codicon-eye": true,
		"codicon-feedback": true, "codicon-file-binary": true, "codicon-file-code": true,
		"codicon-file-media": true, "codicon-file-pdf": true, "codicon-file-submodule": true,
		"codicon-file-symlink-directory": true, "codicon-file-symlink-file": true, "codicon-file-zip": true,
		"codicon-file": true, "codicon-files": true, "codicon-filter-filled": true, "codicon-filter": true,
		"codicon-flag": true, "codicon-flame": true, "codicon-fold-down": true, "codicon-fold-up": true,
		"codicon-fold": true, "codicon-folder-active": true, "codicon-folder-library": true,
		"codicon-folder-opened": true, "codicon-folder": true, "codicon-game": true, "codicon-gear": true,
		"codicon-gift": true, "codicon-gist-secret": true, "codicon-gist": true, "codicon-git-commit": true,
		"codicon-git-compare": true, "codicon-git-fetch": true, "codicon-git-merge": true,
		"codicon-git-pull-request-closed": true, "codicon-git-pull-request-create": true,
		"codicon-git-pull-request-done": true, "codicon-git-pull-request-draft": true,
		"codicon-git-pull-request-go-to-changes": true, "codicon-git-pull-request-new-changes": true,
		"codicon-git-pull-request": true, "codicon-git-stash-apply": true, "codicon-git-stash-pop": true,
		"codicon-git-stash": true, "codicon-github-action": true, "codicon-github-alt": true,
		"codicon-github-inverted": true, "codicon-github-project": true, "codicon-github": true,
		"codicon-globe": true, "codicon-go-to-editing-session": true, "codicon-go-to-file": true,
		"codicon-go-to-search": true, "codicon-grabber": true, "codicon-graph-left": true,
		"codicon-graph-line": true, "codicon-graph-scatter": true, "codicon-graph": true,
		"codicon-gripper": true, "codicon-group-by-ref-type": true, "codicon-heart-filled": true,
		"codicon-heart": true, "codicon-history": true, "codicon-home": true, "codicon-horizontal-rule": true,
		"codicon-hubot": true, "codicon-inbox": true, "codicon-indent": true, "codicon-info": true,
		"codicon-insert": true, "codicon-inspect": true, "codicon-issue-draft": true,
		"codicon-issue-reopened": true, "codicon-issues": true, "codicon-italic": true,
		"codicon-jersey": true, "codicon-json": true, "codicon-kebab-vertical": true, "codicon-key": true,
		"codicon-keyboard-tab-above": true, "codicon-keyboard-tab-below": true, "codicon-keyboard-tab": true,
		"codicon-law": true, "codicon-layers-active": true, "codicon-layers-dot": true, "codicon-layers": true,
		"codicon-layout-activitybar-left": true, "codicon-layout-activitybar-right": true,
		"codicon-layout-centered": true, "codicon-layout-menubar": true, "codicon-layout-panel-center": true,
		"codicon-layout-panel-dock": true, "codicon-layout-panel-justify": true, "codicon-layout-panel-left": true,
		"codicon-layout-panel-off": true, "codicon-layout-panel-right": true, "codicon-layout-panel": true,
		"codicon-layout-sidebar-left-dock": true, "codicon-layout-sidebar-left-off": true,
		"codicon-layout-sidebar-left": true, "codicon-layout-sidebar-right-dock": true,
		"codicon-layout-sidebar-right-off": true, "codicon-layout-sidebar-right": true,
		"codicon-layout-statusbar": true, "codicon-layout": true, "codicon-library": true,
		"codicon-lightbulb-autofix": true, "codicon-lightbulb-empty": true, "codicon-lightbulb-sparkle": true,
		"codicon-lightbulb": true, "codicon-link-external": true, "codicon-link": true,
		"codicon-list-filter": true, "codicon-list-flat": true, "codicon-list-ordered": true,
		"codicon-list-selection": true, "codicon-list-tree": true, "codicon-list-unordered": true,
		"codicon-live-share": true, "codicon-loading": true, "codicon-location": true,
		"codicon-lock-small": true, "codicon-lock": true, "codicon-magnet": true, "codicon-mail-read": true,
		"codicon-mail": true, "codicon-map-filled": true, "codicon-map-vertical-filled": true,
		"codicon-map-vertical": true, "codicon-map": true, "codicon-markdown": true, "codicon-mcp": true,
		"codicon-megaphone": true, "codicon-mention": true, "codicon-menu": true, "codicon-merge": true,
		"codicon-mic-filled": true, "codicon-mic": true, "codicon-milestone": true, "codicon-mirror": true,
		"codicon-mortar-board": true, "codicon-move": true, "codicon-multiple-windows": true,
		"codicon-music": true, "codicon-mute": true, "codicon-new-file": true, "codicon-new-folder": true,
		"codicon-newline": true, "codicon-no-newline": true, "codicon-note": true,
		"codicon-notebook-template": true, "codicon-notebook": true, "codicon-octoface": true,
		"codicon-open-preview": true, "codicon-organization": true, "codicon-output": true,
		"codicon-package": true, "codicon-paintcan": true, "codicon-pass-filled": true, "codicon-pass": true,
		"codicon-percentage": true, "codicon-person-add": true, "codicon-person": true, "codicon-piano": true,
		"codicon-pie-chart": true, "codicon-pin": true, "codicon-pinned-dirty": true, "codicon-pinned": true,
		"codicon-play-circle": true, "codicon-play": true, "codicon-plug": true, "codicon-preserve-case": true,
		"codicon-preview": true, "codicon-primitive-square": true, "codicon-project": true,
		"codicon-pulse": true, "codicon-python": true, "codicon-question": true, "codicon-quote": true,
		"codicon-radio-tower": true, "codicon-reactions": true, "codicon-record-keys": true,
		"codicon-record-small": true, "codicon-record": true, "codicon-redo": true, "codicon-references": true,
		"codicon-refresh": true, "codicon-regex": true, "codicon-remote-explorer": true, "codicon-remote": true,
		"codicon-remove": true, "codicon-replace-all": true, "codicon-replace": true, "codicon-reply": true,
		"codicon-repo-clone": true, "codicon-repo-fetch": true, "codicon-repo-force-push": true,
		"codicon-repo-forked": true, "codicon-repo-pinned": true, "codicon-repo-pull": true,
		"codicon-repo-push": true, "codicon-repo": true, "codicon-report": true, "codicon-request-changes": true,
		"codicon-robot": true, "codicon-rocket": true, "codicon-root-folder-opened": true,
		"codicon-root-folder": true, "codicon-rss": true, "codicon-ruby": true, "codicon-run-above": true,
		"codicon-run-all-coverage": true, "codicon-run-all": true, "codicon-run-below": true,
		"codicon-run-coverage": true, "codicon-run-errors": true, "codicon-save-all": true,
		"codicon-save-as": true, "codicon-save": true, "codicon-screen-full": true, "codicon-screen-normal": true,
		"codicon-search-fuzzy": true, "codicon-search-sparkle": true, "codicon-search-stop": true,
		"codicon-search": true, "codicon-send-to-remote-agent": true, "codicon-send": true,
		"codicon-server-environment": true, "codicon-server-process": true, "codicon-server": true,
		"codicon-settings-gear": true, "codicon-settings": true, "codicon-share": true, "codicon-shield": true,
		"codicon-sign-in": true, "codicon-sign-out": true, "codicon-smiley": true, "codicon-snake": true,
		"codicon-sort-precedence": true, "codicon-source-control": true, "codicon-sparkle-filled": true,
		"codicon-sparkle": true, "codicon-split-horizontal": true, "codicon-split-vertical": true,
		"codicon-squirrel": true, "codicon-star-empty": true, "codicon-star-full": true, "codicon-star-half": true,
		"codicon-stop-circle": true, "codicon-surround-with": true, "codicon-symbol-array": true,
		"codicon-symbol-boolean": true, "codicon-symbol-class": true, "codicon-symbol-color": true,
		"codicon-symbol-constant": true, "codicon-symbol-enum-member": true, "codicon-symbol-enum": true,
		"codicon-symbol-event": true, "codicon-symbol-field": true, "codicon-symbol-file": true,
		"codicon-symbol-interface": true, "codicon-symbol-key": true, "codicon-symbol-keyword": true,
		"codicon-symbol-method-arrow": true, "codicon-symbol-method": true, "codicon-symbol-misc": true,
		"codicon-symbol-namespace": true, "codicon-symbol-numeric": true, "codicon-symbol-operator": true,
		"codicon-symbol-parameter": true, "codicon-symbol-property": true, "codicon-symbol-ruler": true,
		"codicon-symbol-snippet": true, "codicon-symbol-string": true, "codicon-symbol-structure": true,
		"codicon-symbol-variable": true, "codicon-sync-ignored": true, "codicon-sync": true,
		"codicon-table": true, "codicon-tag": true, "codicon-target": true, "codicon-tasklist": true,
		"codicon-telescope": true, "codicon-terminal-bash": true, "codicon-terminal-cmd": true,
		"codicon-terminal-debian": true, "codicon-terminal-linux": true, "codicon-terminal-powershell": true,
		"codicon-terminal-tmux": true, "codicon-terminal-ubuntu": true, "codicon-terminal": true,
		"codicon-text-size": true, "codicon-three-bars": true, "codicon-thumbsdown-filled": true,
		"codicon-thumbsdown": true, "codicon-thumbsup-filled": true, "codicon-thumbsup": true,
		"codicon-tools": true, "codicon-trash": true, "codicon-triangle-down": true, "codicon-triangle-left": true,
		"codicon-triangle-right": true, "codicon-triangle-up": true, "codicon-twitter": true,
		"codicon-type-hierarchy-sub": true, "codicon-type-hierarchy-super": true, "codicon-type-hierarchy": true,
		"codicon-unfold": true, "codicon-ungroup-by-ref-type": true, "codicon-unlock": true,
		"codicon-unmute": true, "codicon-unverified": true, "codicon-variable-group": true,
		"codicon-verified-filled": true, "codicon-verified": true, "codicon-versions": true,
		"codicon-vm-active": true, "codicon-vm-connect": true, "codicon-vm-outline": true,
		"codicon-vm-running": true, "codicon-vm": true, "codicon-vr": true, "codicon-vscode-insiders": true,
		"codicon-vscode": true, "codicon-wand": true, "codicon-warning": true, "codicon-watch": true,
		"codicon-whitespace": true, "codicon-whole-word": true, "codicon-window": true,
		"codicon-word-wrap": true, "codicon-workspace-trusted": true, "codicon-workspace-unknown": true,
		"codicon-workspace-untrusted": true, "codicon-zoom-in": true, "codicon-zoom-out": true,
	}
}

// SelectIcon chooses the best icon for an agent based on name and description
func (is *IconSelector) SelectIcon(name, description, content string) string {
	// Normalize inputs
	normalizedName := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, "-", " "), "_", " "))
	normalizedDesc := strings.ToLower(description)
	normalizedContent := strings.ToLower(content)

	// Check exact role match first
	for role, icon := range is.exactRoleMap {
		if strings.Contains(normalizedName, role) {
			if is.validIcons[icon] {
				return icon
			}
		}
	}

	// Score-based selection
	iconScores := make(map[string]int)

	// Score domain keywords
	for keyword, icon := range is.domainKeywords {
		score := 0
		if strings.Contains(normalizedName, keyword) {
			score += 10
		}
		if strings.Contains(normalizedDesc, keyword) {
			score += 5
		}
		if strings.Contains(normalizedContent, keyword) {
			score += 2
		}
		if score > 0 {
			iconScores[icon] += score
		}
	}

	// Score characteristic keywords
	for keyword, icon := range is.characteristicKeywords {
		score := 0
		if strings.Contains(normalizedName, keyword) {
			score += 8
		}
		if strings.Contains(normalizedDesc, keyword) {
			score += 3
		}
		if strings.Contains(normalizedContent, keyword) {
			score += 1
		}
		if score > 0 {
			iconScores[icon] += score
		}
	}

	// Find highest scoring icon
	var bestIcon string
	var bestScore int
	for icon, score := range iconScores {
		if score > bestScore && is.validIcons[icon] {
			bestScore = score
			bestIcon = icon
		}
	}

	if bestIcon != "" {
		return bestIcon
	}

	// Fallback logic
	for category, icon := range is.fallbackMap {
		if strings.Contains(normalizedName+" "+normalizedDesc, category) {
			if is.validIcons[icon] {
				return icon
			}
		}
	}

	return "codicon-gear" // Ultimate fallback
}
