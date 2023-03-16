# IntruSearch

Intrusearch is a library written in Go that provides a simple, fast and reliable way to search for data in a set of files.

## Prerequisites

- Install [GoLang](https://go.dev/doc/install) (1.20+)

## Installation

To install Intrusearch, simply run the following command:

```shell
go get github.com/IntruderLabs/intrusearch
```

Alternatively, you can clone the repository and install it manually:

```shell
git clone https://github.com/IntruderLabs/intrusearch.git
cd intrusearch
go install
```

## Usage

Once installed, you can start using Intrusearch in your Go code:

```go
import "github.com/IntruderLabs/intrusearch"

func main() {
    // Create a new searcher
    searcher, err := intrusearch.New("/path/to/directory")
    if err != nil {
        // Handle error
    }

    // Search for a string in the files
    results, err := searcher.Search("hello world")
    if err != nil {
        // Handle error
    }

    // Print the results
    for _, result := range results {
        fmt.Println(result.Filename, result.LineNumber, result.Line)
    }
}
```

## Contributing

If you would like to contribute to Intrusearch, please fork the repository and submit a pull request. Before submitting a pull request, please make sure to run the tests:

```go
go test -v
```

## Contact

Felipe Rios - [@rios0rios0](https://rios0rios0.github.io/tabs/about) - `frios at intruderlabs dot com dot br`

Project Link: [https://github.com/IntruderLabs/intrusearch](https://github.com/IntruderLabs/intrusearch)
