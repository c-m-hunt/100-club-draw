# 100 Club Draw

A very basic CLI application which will give you details of entries and also make random draws.

## Installation
```
go build -o hundclub && mv hundclub ~/go/bin
```

## Usage
```
hundclub -h
```

You must point the application to a CSV file containing the entries. The CSV file must have the following columns:
* Number
* Name

```
number,name
1,Alice
2,Bob
3,Carol
4,David
```

## Example
```
hundclub entries entries.csv
```

```
huindclub draw entries.csv
```