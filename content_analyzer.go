package main

import (
	"fmt"
	"regexp"
	"strings"
)

// NewContentAnalyzer creates a new content analyzer with predefined patterns
func NewContentAnalyzer() *ContentAnalyzer {
	return &ContentAnalyzer{
		rolePatterns: map[string]string{
			// Frontend and UI development (high priority - check first)
			"frontend|react|ui|ux|component|responsive|web": "when you need frontend development, UI/UX work, or web application building. Perfect for React components, responsive design, user interfaces, or client-side development",

			// Backend and API development (high priority)
			"backend|api|server|database|microservice|rest|graphql": "when you need backend development, API design, or server-side programming. Ideal for building APIs, managing databases, or creating server applications",

			// Debugging and troubleshooting
			"debug|troubleshoot|error|bug|issue|diagnostic|investigate": "when you're troubleshooting issues, investigating errors, or diagnosing problems. Specialized in systematic debugging, adding logging, analyzing stack traces, and identifying root causes before applying fixes",

			// Testing and quality assurance
			"test|testing|qa|quality|unit|integration|e2e|automation|coverage": "when you need comprehensive testing, quality assurance, or test automation. Perfect for creating test suites, setting up CI pipelines, or ensuring code quality",

			// Security and compliance
			"security|audit|vulnerability|compliance|penetration|owasp|auth|encryption": "when you need security reviews, vulnerability assessments, or compliance checks. Ideal for implementing secure authentication, conducting security audits, or ensuring regulatory compliance",

			// Legal and regulatory
			"legal|privacy|gdpr|ccpa|terms|policy|compliance|regulatory|law": "when you need legal documentation, privacy policies, or regulatory compliance. Perfect for drafting terms of service, privacy policies, or ensuring legal compliance",

			// Marketing and content
			"market|content|blog|social|seo|email|campaign|copy|brand": "when you need marketing content, social media posts, or content strategy. Ideal for creating blog posts, email campaigns, or SEO-optimized content",

			// Finance and trading
			"finance|trading|quant|risk|portfolio|investment|financial": "when you need financial analysis, trading strategies, or risk management. Perfect for quantitative finance, portfolio optimization, or market analysis",

			// Performance and optimization (lower priority to avoid conflicts)
			"performance-engineer|optimize-specialist|speed-expert|cache-expert|benchmark|profile": "when you need performance optimization, scalability improvements, or system tuning. Specialized in profiling applications, implementing caching strategies, and optimizing bottlenecks",

			// Architecture and design
			"architect|design|pattern|structure|system|planning|specification": "when you need system design, architecture planning, or technical documentation. Perfect for creating technical specifications, designing system architecture, or planning complex projects",

			// Code review and analysis
			"review|reviewer|analyze|analysis|inspect|examine|evaluate": "when you need code review, quality assurance, or technical analysis. Ideal for reviewing code changes, analyzing system performance, or conducting technical evaluations",
		},
		domainPatterns: map[string]string{
			// AI/ML and data science
			"ai|llm|ml|machine|learning|data|analytics|neural|vector|embedding|rag|prompt": "Specialized in AI/ML development, LLM integration, data analysis, or machine learning workflows",

			// Frontend and UI/UX
			"frontend|react|ui|ux|web|html|css|javascript|component|responsive|mobile": "Perfect for frontend development, UI/UX design, React components, or web application building",

			// Backend and API development
			"backend|api|server|database|sql|microservice|rest|graphql|endpoint": "Ideal for backend development, API design, database management, or server-side programming",

			// Cloud and infrastructure
			"cloud|aws|azure|gcp|terraform|kubernetes|docker|devops|infrastructure|deploy": "Expert in cloud infrastructure, DevOps automation, containerization, or deployment strategies",

			// Mobile development
			"mobile|ios|android|swift|kotlin|react-native|flutter|app": "Specialized in mobile application development, cross-platform solutions, or native app creation",

			// Game development
			"game|unity|unreal|3d|graphics|rendering|physics|gameplay": "Perfect for game development, 3D graphics, game engine programming, or interactive applications",
		},
		actionPatterns: map[string]string{
			"build|create|implement|develop|construct":  "building and implementing solutions",
			"optimize|improve|enhance|tune|refactor":    "optimization and performance improvements",
			"analyze|review|audit|inspect|evaluate":     "analysis and review tasks",
			"design|architect|plan|structure|model":     "design and architecture planning",
			"automate|streamline|integrate|orchestrate": "automation and integration tasks",
			"monitor|track|observe|measure|report":      "monitoring and reporting activities",
			"migrate|modernize|upgrade|transform":       "migration and modernization projects",
			"document|write|draft|create|generate":      "documentation and content creation",
		},
		fallbackPattern: "general development tasks and code implementation",
	}
}

// findBestMatch finds the best matching pattern and returns its description
func (ca *ContentAnalyzer) findBestMatch(text string, patterns map[string]string) (string, int) {
	bestMatch := ""
	bestScore := 0

	for pattern, description := range patterns {
		keywords := strings.Split(pattern, "|")
		score := 0

		for _, keyword := range keywords {
			keyword = strings.TrimSpace(keyword)
			if strings.Contains(text, keyword) {
				// Weight matches by frequency and keyword length
				frequency := strings.Count(text, keyword)
				keywordWeight := len(keyword) // Longer keywords get higher weight
				score += frequency * keywordWeight
			}
		}

		if score > bestScore {
			bestScore = score
			bestMatch = description
		}
	}

	return bestMatch, bestScore
}

// extractFromProactiveStatement attempts to extract existing "Use PROACTIVELY for..." statements
func (ca *ContentAnalyzer) extractFromProactiveStatement(description string) string {
	// Look for "Use PROACTIVELY for..." patterns
	proactivePattern := regexp.MustCompile(`(?i)use\s+proactively\s+for\s+([^.]+)`)
	if match := proactivePattern.FindStringSubmatch(description); len(match) > 1 {
		useCase := strings.TrimSpace(match[1])
		return fmt.Sprintf("Use this mode when you need %s", useCase)
	}
	return ""
}

// generateWhenToUseStatement creates a comprehensive "when to use" statement
func (ca *ContentAnalyzer) generateWhenToUseStatement(name, description, content string) string {
	// First, try to extract from existing "Use PROACTIVELY for..." statements
	if proactiveStatement := ca.extractFromProactiveStatement(description); proactiveStatement != "" {
		return proactiveStatement + "."
	}

	// Combine all text for analysis
	allText := strings.ToLower(fmt.Sprintf("%s %s %s", name, description, content))

	// Find primary use case from role patterns
	primaryUse, primaryScore := ca.findBestMatch(allText, ca.rolePatterns)

	// Find domain specialization
	domainSpec, domainScore := ca.findBestMatch(allText, ca.domainPatterns)

	// Find action patterns
	actionPattern, actionScore := ca.findBestMatch(allText, ca.actionPatterns)

	// Build the statement
	var statement strings.Builder

	if primaryUse != "" && primaryScore > 0 {
		statement.WriteString(fmt.Sprintf("Use this mode %s", primaryUse))
	} else if actionPattern != "" && actionScore > 0 {
		statement.WriteString(fmt.Sprintf("Use this mode when you need %s", actionPattern))
	} else {
		statement.WriteString(fmt.Sprintf("Use this mode for %s", ca.fallbackPattern))
	}

	// Add domain specialization if found and significant, but avoid duplication
	if domainSpec != "" && domainScore > 5 && !strings.Contains(primaryUse, "Specialized in") && !strings.Contains(primaryUse, strings.Split(domainSpec, " ")[2]) {
		statement.WriteString(fmt.Sprintf(". %s", domainSpec))
	}

	return statement.String() + "."
}

// generateDescription creates a short description for the agent using ContentAnalyzer
func generateDescription(name, description, content string) string {
	analyzer := NewContentAnalyzer()
	text := strings.ToLower(fmt.Sprintf("%s %s %s", name, description, content))

	// Create mapping from patterns to concise descriptions (3-5 words)
	shortDescriptions := map[string]string{
		// Role-based patterns
		"debug|troubleshoot|error|bug|issue|diagnostic|investigate":                            "Debug and troubleshoot",
		"test|testing|qa|quality|unit|integration|e2e|automation|coverage":                     "Testing and QA",
		"security|audit|vulnerability|compliance|penetration|owasp|auth|encryption":            "Security and auditing",
		"legal|privacy|gdpr|ccpa|terms|policy|compliance|regulatory|law":                       "Legal and compliance",
		"market|content|blog|social|seo|email|campaign|copy|brand":                             "Marketing and content",
		"finance|trading|quant|risk|portfolio|investment|financial":                            "Financial analysis",
		"performance-engineer|optimize-specialist|speed-expert|cache-expert|benchmark|profile": "Performance optimization",
		"architect|design|pattern|structure|system|planning|specification":                     "Architecture and design",
		"review|reviewer|analyze|analysis|inspect|examine|evaluate":                            "Code review",
		"frontend|react|ui|ux|component|responsive|web":                                        "Frontend development",
		"backend|api|server|database|microservice|rest|graphql":                                "Backend development",
	}

	// Domain-based patterns
	domainDescriptions := map[string]string{
		"ai|llm|ml|machine|learning|data|analytics|neural|vector|embedding|rag|prompt": "AI and ML",
		"cloud|aws|azure|gcp|terraform|kubernetes|docker|devops|infrastructure|deploy": "DevOps and cloud",
		"mobile|ios|android|swift|kotlin|react-native|flutter|app":                     "Mobile development",
		"game|unity|unreal|3d|graphics|rendering|physics|gameplay":                     "Game development",
		"database|sql|nosql|mongodb|postgres|mysql":                                    "Database management",
	}

	// Find best match from role patterns first (higher priority)
	bestDescription := ""
	bestScore := 0

	for pattern, desc := range shortDescriptions {
		_, score := analyzer.findBestMatch(text, map[string]string{pattern: desc})
		if score > bestScore {
			bestScore = score
			bestDescription = desc
		}
	}

	// If no strong role match, try domain patterns
	if bestScore == 0 {
		for pattern, desc := range domainDescriptions {
			_, score := analyzer.findBestMatch(text, map[string]string{pattern: desc})
			if score > bestScore {
				bestScore = score
				bestDescription = desc
			}
		}
	}

	// Fallback to default if no matches
	if bestDescription == "" {
		bestDescription = "Development specialist"
	}

	return bestDescription
}
