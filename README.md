# ⚡ QR Code Generator

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v4.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Transform tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`transform` `data-processing` `cli` `golang`

---

## What is QR-Code-Generator?

**QR-Code-Generator** is a data transformation tool that converts, formats, and processes files between different formats.

## Features

- ✅ `buildGrid()` — Buildgrid
- ✅ `drawFinder()` — Drawfinder
- ✅ `renderSVG()` — Rendersvg
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/QR-Code-Generator.git
cd QR-Code-Generator

# Build
go build -o qr-code-generator .

# Run
./qr-code-generator <text> [out.svg]
```

### Or directly with `go run`:
```bash
go run main.go <text> [out.svg]
```

## Usage

```bash
# Basic usage
./qr-code-generator <text> [out.svg]

# With flags
./qr-code-generator <text> [out.svg] value <text> [out.svg]
```

### Example Output

```
$ ./qr-code-generator <text> [out.svg]
<text> [out.svg]
```

## Project Structure

```
QR-Code-Generator/
  main.go          # Entry point (81 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
