## FizzBuzz Kata with Go

**Goal**: Convert integers into a string with specific rules

**Rules**: 
- Values divisible by 3 are converted to "Fizz"
- Values divisible by 5 are converted to "Buzz"
- Values divisible by 3 and 5 are converted to "FizzBuzz"
- All other values are converted to their string representations

**Test cases**:
```
1 -> "1"
2 -> "2"
3 -> "Fizz"
4 -> "4"
5 -> "Buzz"
6 -> "Fizz"
7 -> "7"
8 -> "8"
9 -> "Fizz"
10 -> "Buzz"
11 -> "11"
12 -> "Fizz"
13 -> "13"
14 -> "14"
15 -> "FizzBuzz"
```

### Extension 1 - FizzBuzzBang

**Rules**: 
- The same rules as for FizzBuzz, but with the following additions: 
- Values divisible by 7 are converted to "Bang"
- Values divisible by 3 and 7 are converted to "FizzBang"
- Values divisible by 5 and 7 are converted to "BuzzBang"
- Values divisible by 3, 5 and 7 are converted to "FizzBuzzBang"
- All other values are converted to their string representations

**Test cases**:
```
1 -> "1"
2 -> "2"
3 -> "Fizz"
4 -> "4"
5 -> "Buzz"
6 -> "Fizz"
7 -> "Bang"
8 -> "8"
9 -> "Fizz"
10 -> "Buzz"
11 -> "11"
12 -> "Fizz"
13 -> "13"
14 -> "Bang"
15 -> "FizzBuzz"
21 -> "FizzBang"
35 -> "BuzzBang"
105 -> "FizzBuzzBang"
```

### Extension 2 - Changing rules

**Rules**: 
- The same rules as for FizzBuzz but with the following additions: 
- Values which contain the digit 3 are also converted to "Fizz"
- Values which contain the digit 5 are also converted to "Buzz"
- Values which contain the digits 3 and 5 are also converted to "FizzBuzz"
- All other values are converted to their string representations

**Test cases**:
```
1 -> "1"
2 -> "2"
3 -> "Fizz"
4 -> "4"
5 -> "Buzz"
6 -> "Fizz"
7 -> "7"
8 -> "8"
9 -> "Fizz"
10 -> "Buzz"
11 -> "11"
12 -> "Fizz"
13 -> "Fizz"
14 -> "7"
15 -> "FizzBuzz"
31 -> "Fizz"
52 -> "Buzz"
53 -> "FizzBuzz"
```