
# eggs 
extended [go][1] [generics][2], [scala][3]-like

project moved to [codeberg.org][4]

---

[![codecov](https://codecov.io/gh/mikhalytch/eggs/branch/main/graph/badge.svg?token=U4I0VXG3KI)](https://codecov.io/gh/mikhalytch/eggs)

---

Things to note: 
1. computations are not lazy:
   ```go
   try.Trie(os.ReadFile("file")). // will execute without any .unsafeRunSync()
                   Map(bytes.ToLower) 
   ```
   it's a long road, but we'll get there

[1]:https://go.dev/
[2]:https://en.wikipedia.org/wiki/Generic_programming
[3]:https://scala-lang.org/
[4]:https://codeberg.org/mikhalytch/eggs

