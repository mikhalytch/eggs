
# eggs 
extended [go][1] [generics][2], [scala][3]-like

---

[![codecov](https://codecov.io/gh/mikhalytch/eggs/branch/main/graph/badge.svg?token=U4I0VXG3KI)](https://codecov.io/gh/mikhalytch/eggs)

---

- computations are lazy where possible:
   1. immediate init
      ```go
      import "github.com/mikhalytch/eggs/try"
      num := try.From(strconv.Atoi("0")) // will execute Atoi immediately
      ```
   2. lazy init
      ```go 
      // execution is delayed to first non-functional call: Get, IsSuccess, etc...
      str := try.Lazy(func() (string, error) { return "0", nil })
      ```
   3. lazy calc
      ```go
      num = num.Map(func(i int) int { return i + 0 }) // delayed
      str = try.Map[int, string](num, func(i int) string { panic("delayed as well") })
      ```   

[1]:https://go.dev/
[2]:https://en.wikipedia.org/wiki/Generic_programming
[3]:https://scala-lang.org/

