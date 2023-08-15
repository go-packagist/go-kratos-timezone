# go-kratos-timezone

[![Go Version](https://badgen.net/github/release/go-packagist/go-kratos-timezone/stable)](https://github.com/go-packagist/go-kratos-timezone/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/go-kratos-timezone)](https://pkg.go.dev/github.com/go-packagist/go-kratos-timezone)
[![codecov](https://codecov.io/gh/go-packagist/go-kratos-timezone/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/go-kratos-timezone)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/go-kratos-timezone)](https://goreportcard.com/report/github.com/go-packagist/go-kratos-timezone)
[![tests](https://github.com/go-packagist/go-kratos-timezone/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/go-kratos-timezone/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/go-kratos-timezone
```

## Usage

```go
package main

import (
	kratos "github.com/go-kratos/kratos/v2"
	timezone "github.com/go-packagist/go-kratos-timezone"
)

func main() {
	app := kratos.New(
		kratos.BeforeStart(
			timezone.Timezone(timezone.WithTimezone("UTC")),
		),
	)

	app.Run()
}

```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.