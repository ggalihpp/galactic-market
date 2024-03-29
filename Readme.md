# GALAXY MERCHANT TRADING GUIDE

## Problem Description 
You decided to give up on earth after the latest financial collapse left 99.99% of the earth's
population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your
account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell
common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy
requires you to convert numbers and units, and you decided to write a program to help you.The
numbers used for intergalactic transactions follows similar convention to the roman numerals and
you have painstakingly collected the appropriate translation between them. Roman numerals are
based on seven symbols: 

|Symbol|Value|
|---|---|
|I| 1|
|V| 5|
|X| 10|
|L| 50|
|C| 100|
|D| 500|
|M| 1,000| 

Numbers are formed by combining symbols together and adding the values. For example, MMVI is
1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the
largest values. When smaller values precede larger values, the smaller values are subtracted from
the larger values, and the result is added to the total. For example MCMXLIV = 1000 + (1000 −
100) + (50 − 10) + (5 − 1) = 1944. 

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They
may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.)
"D", "L", and "V" can never be repeated.
"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can
be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.
Only one small-value symbol may be subtracted from any large-value symbol.
A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of
1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately.
In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

## Test case
#### Input:
```
glob is I
prok is V
pish is X
tegj is L 
glob glob Silver is 34 Credits
glob prok Gold is 57800 Credits
pish pish Iron is 3910 Credits
how much is pish tegj glob glob ?
how many Credits is glob prok Silver ?
how many Credits is glob prok Gold ?
how many Credits is glob prok Iron ?
how much wood could a woodchuck chuck if a woodchuck could chuck wood ? 
```

### Expecting Output:
```
pish tegj glob glob is 42
glob prok Silver is 68 Credits
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
I have no idea what you are talking about 
```
---
# Project Guidelines

## Prerequisites
Golang version 1.11.x or later that support *GO MODULES*
check the golang version with `go version` command

## Installation
note: This installation guide assumed your os is unix based.

clone the project
```
git clone https://github.com/ggalihpp/galactic-market.git
```
go to project folder
```
cd galactic-market
```
build the project with GO MODULE enabled
```
GO111MODULE=on go build .
```
*galactic-market* binary will created run with
```
./galactic-market
```

## Unit Test
Run this command to run unit test, if this is your first time running the project enable the GO MODULE
```
GO111MODULE=on go test
```
The test will running test case [above](#test-case)

to change the test case, change the value ov variable `inventoryInput` and `transactionInput` inside `main_test.go` file

## Development
To run the project directly use `run` command (If run for the first time, activate go module)
```
go run main.go
```

You can use DEBUG mode to print the error of the input instead of printing default message
```
DEBUG=true go run main.go
```