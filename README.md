# QR Code Generator

![CI](https://github.com/Qyroxen/QR-Code-Generator/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/QR-Code-Generator/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/QR-Code-Generator?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/QR-Code-Generator)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/QR-Code-Generator)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/QR-Code-Generator?style=social)](https://github.com/Qyroxen/QR-Code-Generator/stargazers)

## What is it?

QR Code Generator is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/QR-Code-Generator.git
cd QR-Code-Generator
go build -o qrcodegenerator .

# Run
./qrcodegenerator --help
```

## CLI Usage

```bash
# Basic usage
./qrcodegenerator

# With flags
./qrcodegenerator --verbose --output json

# Get help
./qrcodegenerator --help
```

## Examples

```bash
# Example 1
./qrcodegenerator example1

# Example 2
./qrcodegenerator example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o qrcodegenerator .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/QR-Code-Generator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/QR-Code-Generator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/QR-Code-Generator/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/QR-Code-Generator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/QR-Code-Generator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/QR-Code-Generator" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/QR-Code-Generator/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/QR-Code-Generator" alt="Pull Requests">
  </a>
</p>
