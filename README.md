# GoLang BaseLinker API

go-baselinker is a Go client library for accessing the [BaseLinker](https://baselinker.com/) service.

Currently, **library is BETA version** and testing only on Go version 1.12.

Library not implement each of method from BaseLinker API [Documentation](https://api.baselinker.com/) yet.

## Usage

```go
import "github.com/oxess/go-baselinker"
```

Construct a new BaseLinker client. Base API url is publish in [Documentation](https://api.baselinker.com/).

For generate api key must go to "My account" -> "API" -> type name of application and click "Generate token".

```go
baseLinkerUrl := "https://api.baselinker.com/connector.php"
baseLinkerToken := "xxx"

baseLinker := baselinker.NewBaseLinker(baseLinkerUrl, baseLinkerToken)
```

### Example

Get logs kind of "create new order", "change order status" from the journal.

```go
journalListParameters := baselinker.GetJournalListParameters{
    Types: []int{
        baselinker.LogTypeCreateNewOrder, 
        baselinker.LogTypeChangeOrderStatus
    },
}
logs, err := baseLinker.GetJournal(journalListParameters)
```

Get order by the id.

```go
order, err := baseLinker.GetOrder(orderId)
```

## Errors

BaseLinker return field "status" with values: "SUCCESS" or "ERROR".
When status is "ERROR" then it add fields: "error_message" and "error_code".

Library supports errors by return error object but with custom field "code".
Each error allow check code and compare it with baselinker error code list.
Example:

```go

order, err := baseLinker.GetOrder(orderId)
if nil != err {
    if err.CodeError() == ErrorCodeAccountBlocked {
        panic("Your %s account is blocked!", baseLinkerToken)
    } else {
        panic(err.Error()) // <-- here is baselinker error message
    } 
}

```

## Developed methods 

Name       | Description
---------- | -----------
GetJournal | Get list of logs
GetOrders  | Get list of orders
GetOrder   | Get single order by the id

## Versioning 

Versioning is base on [semver](https://semver.org/). 
The new version is release by a new tag.

## Licence

This library is distributed under the BSD 3 Licence, see LICENSE for more information.

## Contributing

Author of this library is [Mikołaj Jeziorny](https://mikolaj-jeziorny.pl).

If you want to help with development it - **fork me!**