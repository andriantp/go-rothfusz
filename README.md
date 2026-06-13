# Rothfusz

A lightweight Go library for calculating the **Heat Index** using the **Rothfusz Regression Equation**.

The library helps determine how hot the weather actually feels by combining air temperature and relative humidity.

## Features

* Pure Go implementation
* No external dependencies
* Calculates Heat Index using the Rothfusz Regression Equation
* Returns structured results ready for JSON serialization
* Provides Heat Index risk categories automatically
* Indicates whether conditions are considered comfortable
* Configurable validity thresholds


## Installation

```bash
go get github.com/yourusername/rothfusz
```

## Quick Example

```go
package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/andriantp/go-rothfusz/rothfusz"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("usage: go run . <temp> <rh>")
		return
	}

	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("failed to parse temp [%s]: %v", os.Args[1], err)
	}

	rh, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("failed to parse rh [%s]: %v", os.Args[2], err)
	}

	repo := rothfusz.NewRothfusz(
		26.7, // minimum valid temperature
		85,   // humidity threshold
	)

	result := repo.CalculateHeatIndex(temp, rh)

	js, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatalf("failed to construct result: %v", err)
	}

	log.Printf("result:\n%s", string(js))
}
```

Run:

```bash
go run . 32 70
```

Output:

```json
{
 "data": {
  "temp": 32,
  "rh": 70
 },
 "heatIndexResult": {
  "heatIndexC": 40.409273679555774,
  "category": "Extreme Caution",
  "comfortable": false
 }
}
```

## API

### Create Calculator

```go
calculator := rothfusz.NewRothfusz(
	26.7, // minimum valid temperature
	85,   // humidity threshold
)
```

The constructor allows you to configure the validity range used by the calculator.

### Calculate Heat Index

```go
result := calculator.CalculateHeatIndex(temp, rh)
```

Parameters:

* `temp` : Air temperature in Celsius.
* `rh` : Relative humidity in percent.

Returns a structured result containing the calculation outcome and additional metadata.

## Notes

The Rothfusz regression is generally considered valid when:

* Temperature ≥ 26.7°C (80°F)
* Relative humidity ≥ 40%

You may customize these thresholds depending on your application's requirements.

## Reference

Lans P. Rothfusz.

*The Heat Index "Equation" (or, More Than You Ever Wanted to Know About Heat Index)*.

National Weather Service Technical Attachment SR 90-23 (1990).

## License

MIT License.
