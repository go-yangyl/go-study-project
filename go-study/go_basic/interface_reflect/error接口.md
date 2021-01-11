## error接口

#### error的基本用法

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

创建一个 error 最简单的方法就是调用 errors.New 函数，它会根据传入的错误信息返回一个新的 error，示例代码如下：

```go
package main
import (
    "errors"
    "fmt"
    "math"
)
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, errors.New("math: square root of negative number")
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
```

#### 自定义错误类型

除了上面的 errors.New 用法之外，我们还可以使用 error 接口自定义一个 Error() 方法，来返回自定义的错误信息。

```go
package main
import (
    "fmt"
    "math"
)
type dualError struct {
    Num     float64
    problem string
}
func (e dualError) Error() string {
    return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
}
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, dualError{Num: f}
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
```