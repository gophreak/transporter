# Transporter

The transporter library is a middle layer for allowing transport of payloads via a configured method. Currently supporting SMS but with the intention of being able to push messages to multiple SMS providers, to interface with Queue systems such as RabbitMQ and Kafka.

## Getting Started

Clone the repository and use it in your own projects by accessing functions inside of the `transporter.go`. The Notification interface should be used for interacting with the underlying adaptors.

To begin using this library, please import the transport mechanisms required into your main function

```
_ "github.com/gophreak/transporter/sms"
```

This will run the `init()` and register all packaged vendors for SMS.

To get the adaptor from the list of available adaptors, you should use the following code line:
```
adaptor, e := transporter.GetAdaptor("sms", "nexmo")
```

This will get the `sms` transport using the `nexmo` library. Once you have acquired the appropriate adaptor, call the Setup function and pass through the specific configuration for the adaptor:

```
adaptor.Setup([]byte(`{
    "api_key": "********",
    "api_secret": "****************"
}`))

```

Then you are free to use the adaptor for pushing messages:
```
err := adaptor.Push("developers", "+447902040506", []byte("Hello, World!"), transporter.Table{})
```

### Prerequisites

This library depends on Go's dependency manager `dep` to install vendor tools.

```
dep ensure
```

## Running the tests

To test this system you can natively use Go's built in testing tools

```
go test -v ./...
```

## Contributing

When contributing to this library, please ensure that you follow best pratices for go linting. The following linting tools have been used in this project:
 * [go vet](https://golang.org/cmd/vet/) - Finds potential errors that would otherwise compile.
 * [deadcode](https://github.com/tsenart/deadcode) - Finds unused code.
 * [gocyclo](https://github.com/alecthomas/gocyclo) - Finds the cyclomatic complexity of functions, and reports if over 10.
 * [golint](https://github.com/golang/lint) - Finds styling issues.
 * [ineffassign](https://github.com/gordonklaus/ineffassign/blob/master/list) - Find assignments to existing variables that are not used.
 * [goconst](https://github.com/jgautheron/goconst) - Finds repeated strings that could be constants.
 * [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - Finds any code which could be simplified.
 * [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - find both obvious and subtle bugs statically.
 * [gofmt](https://golang.org/cmd/gofmt/) - Similar to linting, finds any formatting issue with code.
 * [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Finds missing or unreferenced packages.
 * [interfacer](https://github.com/mvdan/interfacer) - Finds and suggest narrower interfaces.
 * [gas](https://github.com/GoASTScanner/gas) - Finds security problems by scanning the Go AST.

## Versioning

We use [SemVer](http://semver.org/) for versioning.

Current version is `1.0.0`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
