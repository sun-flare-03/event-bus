# event-bus

[![Build Status](https://img.shields.io/github/actions/workflow/status/user/event-bus/ci.yml?branch=main)]()
[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**event-bus** in-process event bus with typed events and async handlers. Built with simplicity and performance in mind.

## Features

- Zero Dependencies: No external packages required for core functionality
- Graceful Shutdown: Clean shutdown handling with configurable drain timeout
- High Performance: Optimized for low-latency, high-throughput workloads

## Installation

```bash
go get github.com/user/event-bus@latest
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/user/event-bus"
)

func main() {
	client := eventbus.New(
		eventbus.WithTimeout(30 * time.Second),
	)

	result, err := client.Run(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
```

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
