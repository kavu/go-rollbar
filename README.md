# Ratchet.io Go Notifier

This is _very_ experimental [Go](http://golang.org) package for semi-automatic sending of the application errors to the [Ratchet.io](https://ratchet.io) error tracking service.

## Installation

Plain and simple:

```
go get github.com/kavu/go-ratchetio
```

## Example

```go
package main

import (
  ratchetio "github.com/kavu/go-ratchetio"
  "log"
  "time"
)

func main() {
  defer ratchetio.CapturePanics()

  ratchetio.Config.APIKey = "<your API key>"

  log.Println("Working...")

  go func() {
    defer ratchetio.CapturePanics()
    a := make([]int, 10)
    log.Println(a[110])
  }()

  log.Println("Working normally...")

  time.Sleep(10 * time.Second)

  log.Println("Done!")
}

```

## Contributing

Just open pull request or ping me directly on e-mail, if you want to discuss some ideas.

## One more thing

This package was done in one evening like an experiment and Proof of Concept, so, please, don't judge me so harsh.

## License

Copyright (c) 2012 Max Riveiro <kavu13@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
