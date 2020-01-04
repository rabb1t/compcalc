# Cartridge resonance frequency calculator

For all the vinyl enthusiasts out there here is the simple and easy solution to deal with the tonearm, cartridge and headshell compatibility problem.

With the provided program you can easily find the best matched cartridge/tonearm/headshell combinations.

### Prerequisites

To successfully compile the program you need to install the following dependency package:

```
go get gopkg.in/yaml.v2
```

### Installing

Build the same way as you build every single golang program:

```
go build compcalc.go
```

### Usage

Fill the provided [example YAML file](compcalc.yaml) with desired cartridges, tonearms and headshells.
Run program:

```
./compcalc
```

For the additional usage options see the builtin help information:

```
./compcalc -h
```

## Built With

* [go-yaml](https://github.com/go-yaml/yaml) - YAML support for the Go language.

## Authors

* **Pavel A. Raykov** - *Initial commit* - [rabb1t](https://github.com/rabb1t)

## License

This project is licensed under the Description WTFPL License - see the [LICENSE](LICENSE) file for details
