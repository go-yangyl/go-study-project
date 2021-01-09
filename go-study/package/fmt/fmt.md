## fmt



fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分

实现接口fmt.Fprintln()传入的interface

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

func main() {
	w := new(W)
	fmt.Fprintln(w, 1)

}

type W struct {
}

func (w *W) Write(p []byte) (n int, err error) {
	return 1, nil
}

```



#### 格式化占位符

```go
func Printf(format string, a ...interface{}) (n int, err error) 
```



| 占位符 | 说明             |
| ------ | ---------------- |
| %v     | 值的默认格式表示 |
| %T     | 类型             |
| %%     | 百分号           |
| %t     | True false       |
| %s     | String []byte    |
|        |                  |
|        |                  |
|        |                  |
|        |                  |

