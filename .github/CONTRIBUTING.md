# Contributing to claude2kilo

Thank you for your interest in contributing to the Claude Agent to Kilo Code Mode Converter! This guide will help you contribute effectively while maintaining a positive community environment.

## Before You Contribute

1. **Read our [Code of Conduct](.github/CODE_OF_CONDUCT.md)** - All interactions must follow our community standards
2. **Search existing issues** - Check if your suggestion or bug report already exists
3. **Use appropriate templates** - Follow the provided issue and PR templates

## Types of Contributions

### Core Converter Improvements
- Bug fixes in conversion logic
- Performance optimizations
- Enhanced intelligent analysis features
- YAML sanitization improvements
- Diagnostic reporting enhancements

### New Features
- Additional output formats
- New intelligent analysis capabilities
- Enhanced icon selection logic
- Improved content analysis patterns
- Extended CLI functionality

### Infrastructure & Tooling
- Build system improvements
- GitHub Actions enhancements
- Documentation updates
- Testing infrastructure
- Cross-platform compatibility

## Contribution Process

### 1. Issues First
- **Always create an issue before starting work** on significant changes
- Use the appropriate issue template
- Provide clear, detailed descriptions
- Include relevant examples or sample files

### 2. Pull Requests
- Fork the repository and create a feature branch
- Follow existing Go code style and conventions
- Include tests where appropriate
- Reference the related issue in your PR description
- Use clear, descriptive commit messages

### 3. Review Process
- All PRs require review from maintainers
- Address feedback promptly and professionally
- Be patient - reviews may take time

## Development Guidelines

### Code Quality Standards
- **Go formatting**: Use `gofmt` and `go vet`
- **Error handling**: Proper error propagation and user-friendly messages
- **Documentation**: Clear comments for complex logic
- **Testing**: Include unit tests for new functionality
- **Performance**: Consider memory usage and processing efficiency

### Architecture Principles
- **Modularity**: Keep components separate and focused
- **Extensibility**: Design for easy addition of new features
- **Reliability**: Handle edge cases and malformed input gracefully
- **User Experience**: Provide clear feedback and helpful error messages

## Content Guidelines

### What We Accept
- ✅ Constructive feedback and suggestions
- ✅ Well-researched feature requests
- ✅ Clear bug reports with reproduction steps
- ✅ Professional, respectful communication
- ✅ Documentation improvements
- ✅ Performance optimizations
- ✅ Cross-platform compatibility fixes

### What We Don't Accept
- ❌ Hate speech, discrimination, or harassment
- ❌ Spam, promotional content, or off-topic posts
- ❌ Personal attacks or inflammatory language
- ❌ Duplicate or low-effort submissions
- ❌ Breaking changes without discussion
- ❌ Copyright infringement

## Quality Standards

### For Code Contributions
- Clear, maintainable Go code
- Proper error handling and validation
- Comprehensive testing coverage
- Performance considerations
- Cross-platform compatibility

### For Documentation
- Clear, concise writing
- Accurate technical information
- Consistent formatting and style
- Practical examples and use cases

## Development Setup

### Prerequisites
- Go 1.21 or later
- Make (for build automation)
- Git for version control

### Getting Started
```bash
# Clone your fork
git clone https://github.com/Romboter/claude2kilo.git
cd claude2kilo

# Build the project
make clean
make

# Run tests
go test ./...

# Test with sample files
./bin/claude2kilo-linux -input sample.md -output ./test-output/ -dry-run
```

### Testing Your Changes
- Test with various input formats
- Verify cross-platform compatibility
- Check error handling with malformed files
- Validate output format correctness
- Test CLI options and edge cases

## Community Guidelines

### Communication
- **Be respectful** - Treat all community members with dignity
- **Be constructive** - Focus on improving the project
- **Be patient** - Allow time for responses and reviews
- **Be helpful** - Share knowledge and assist others

### Collaboration
- **Give credit** - Acknowledge others' contributions
- **Share knowledge** - Help others learn and grow
- **Stay focused** - Keep discussions on topic
- **Follow up** - Respond to feedback and requests

## Getting Help

- 📖 **Documentation**: Check the README and code comments
- 💬 **Discussions**: Use GitHub Discussions for questions and brainstorming
- 🐛 **Issues**: Report bugs or request features through issue templates
- 📧 **Direct Contact**: Reach out to maintainers for sensitive matters

## Recognition

Contributors who consistently provide high-quality submissions and maintain professional conduct will be:
- Acknowledged in release notes
- Given priority review for future contributions
- Potentially invited to become maintainers

## Enforcement

Violations of these guidelines may result in:
1. **Warning** - First offense or minor issues
2. **Temporary restrictions** - Suspension of contribution privileges
3. **Permanent ban** - Severe or repeated violations

Reports of violations should be made through:
- GitHub's built-in reporting tools
- Issues tagged with `moderation`
- Direct contact with maintainers

---

Thank you for helping make this project a welcoming, productive environment for everyone!