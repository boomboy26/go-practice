### Go practice

<img  src="https://matobnews.b-cdn.net/wp-content/uploads/2020/11/go.png" width="200" />

#### Go run / Build

```
go run main.go
go build
```

#### print command

```
import "fmt"


fmt.Println(message)
```

#### Varible

```
var txt string
txt = "Some Text"

txt2 := "Some text"
```

#### Array / Slice

array is fixed size
slice not fix size
Array

```
arrString := [...]string{"John", "Ny", "Walker"}
```

Slice

```
slice := []int{2, 3, 4}
```

#### Condition

```
if value == "right" || value == "left" {
  fmt.Println("condition1")
} else if value == "direct" && isSomeFunc() {
  fmt.Println("condition2")
} else {
fmt.Println("condition3")
}
```

#### loop

```
for count := 0; count <= 10; count++ {
  fmt.Println("My counter is at", count)
}
// Range
 num := []string{0,1,2,3,4}
for i, val := range num {
  fmt.Println(i , val)
}
```

#### Structs

```
type Person struct {
  firstname string
  lastname string
}
```
