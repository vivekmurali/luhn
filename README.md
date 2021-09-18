# luhn

The Luhn algorithm or Luhn formula, also known as the "modulus 10" or "mod 10" algorithm, is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers and so on. 

Implementation of the Luhn algorithm. This package generates the checksum as well as validates if a given number is a Luhn number.


## USAGE

Generate the whole number

```go
num, _ := Generate("123")
fmt.Println(num)
```

Generate only the checkssum

```go
num, _ := GenerateNumber("123")
fmt.Println(num)
```

Validate a given number

```go
fmt.Println(Validate("49927398716"))
```
