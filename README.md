# Zero-Dependency Todo Application

A completely custom-built Todo application in Go. This project is designed to explore software engineering fundamentals and demonstrate a deep understanding of computer science and software development by avoiding Go's standard library imports entirely. Instead of relying on built-in packages like fmt or os, this application utilizes custom-engineered internal packages to handle all system interactions, formatting, and I/O operations.

## Architecture & Project Structure

To simulate a professional engineering environment, the project separates execution logic from internal libraries.

```text
todo-application/
├── cmd/
│   └── todo/
│       └── main.go       # The main entry point of the application
├── takumifmt/            # Custom string formatting and printing package
├── takumios/             # Custom OS interaction, file parsing, and I/O package
├── docs/                 # Project documentation and RFCs
├── go.mod                # Module definition
└── README.md
```

## Core Features

* **Zero Standard-Library Imports:** Built strictly without importing Go's standard library to demonstrate a deep understanding of memory management, file descriptors, and system calls.
* **Custom I/O (takumios):** All file reads, writes, and command-line argument parsing are handled through a proprietary package.
* **Custom Formatting (takumifmt):** String manipulation and standard output rendering are engineered from the ground up.
* **Persistent Storage:** State is saved and retrieved using a custom serialization format.

## Getting Started

### Prerequisites

* [Go](https://golang.org/doc/install) (1.x or later)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/todo-application.git
   ```

2. Navigate to the project directory:

   ```bash
   cd todo-application
   ```

### Usage

Execute the application from the entry point:

```bash
# Add a task
go run cmd/todo/main.go add "Review system architecture"

# List tasks
go run cmd/todo/main.go list

# Complete a task
go run cmd/todo/main.go complete 1
```

## Development Roadmap

* [ ] Phase 1: Implement memory structs and command parser.
* [ ] Phase 2: Connect `takumifmt` for stdout rendering.
* [ ] Phase 3: Wire up `takumios` for state persistence.
* [ ] Phase 4: Build a custom testing harness.

## License

MIT License

