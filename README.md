# protoc-gen-pojo

A protoc plugin for generating Java POJO classes from Protobuf files.

## Getting Started

To install the package:

```bash
go get github.com:sagroskin/protoc-gen-pojo.git
```

Generate Java classes with default options:

```bash
protoc --pojo_out=:. <your_file>.proto
```

Generate Java classes with string enums:

```bash
protoc --pojo_out=string_enums=true:. <your_file>.proto
```

## Run without installing module:

```bash
git clone https://github.com/sagroskin/protoc-gen-pojo.git
cd protoc-gen-pojo
go build
protoc --plugin=./protoc-gen-pojo.exe --proto_path=./ignore/proto --pojo_out=./ignore/java example.proto
```

## Developer Environment Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to run the project:

- Go

### Installing

```bash
go install
```

## Running Tests

```bash
go test
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://stash.sv2.trulia.com/users/sergea/repos/graphql-poc/branches).

## Authors

- **Serge Agroskin** - _Initial work_

Based on [protoc-gen-tstypes](<[LICENSE](https://github.com/tmc/grpcutil/tree/master/protoc-gen-tstypes)>)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
