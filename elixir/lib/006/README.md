### Sum square difference

#### Problem

https://projecteuler.net/problem=6

The sum of the squares of the first ten natural numbers is,

`1^2 + 2^2 + ... + 10^2 = 385`

The square of the sum of the first ten natural numbers is,

`(1 + 2 + ... + 10)^2 = 552 = 3025`

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.


#### Solution

```
x = Enum.sum(1..100) |> :math.pow(2)
y = Enum.reduce(1..100, fn(x, acc) -> acc + :math.pow(x, 2) end)

IO.puts x - y
```