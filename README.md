# Api client for sipuni telephony

## Install

    go get -u github.com/bzdvdn/sipuni-go/sipuni

## Usage

=======

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/bzdvdn/sipuni-go/sipuni"
)

func main() {
	client := sipuni.NewClient("<id>", "<secreteKey>")
}
```

#### Make calls

- <a href="https://help.sipuni.com/articles/134-182-113--sozdanie-zvonka-na-nomer-s-pomoshyu-api/">Documentation</a>

```go
// make basic call
client.Call.MakeCall("79379992", "201", 0, 0)

// make tree call
client.Call.MakeTreeCall("79379992", "201", "000658610", 0, 0)

// make external call
client.Call.MakeExternalCall("79379992", "79379993", "201", "202")

```

#### Make voice call

- <a href="https://help.sipuni.com/articles/134-182-107--generaciya-golosovogo-zvonka-s-pomoshyu-api/I">Documentation</a>

```go

client.Call.MakeVoiceCall("79379992", "text 123", "Anna_n", "201")

```

#### Hangup call

- <a href="https://help.sipuni.com/articles/134-182-108--zapros-na-zavershenie-zvonka/">Documentation</a>

```go
client.Call.HangUpCall("0286966c7e35e1ec2565e9d37e77edb1")

```

#### Statistic

- <a href="https://help.sipuni.com/articles/134-182-112--poluchenie-statistiki-po-zvonkam-zapisej-razgovorov-i-statusov-sotrudnikov/">Documentation</a>

```go

// export requests
exReq := sipuni.ExportRequest{
    From: time.Now().AddDate(0, 0, -1),
    To:   time.Now(),
}
csvData, err := client.Statistic.Export(exReq)

// export all
csvData, err := client.Statistic.ExportAll(10, "asc", 1)

// get call record
record, err := client.Statistic.getRecord("123")

```

### TODO

- examples
- tests

### License

MIT
