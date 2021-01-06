## 跳转语句goto,break,continue的使用及区别

```go
goto语句可以无条件地转移到过程中指定的行。
通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱
goto对应(标签)既可以定义在for循环前面,也可以定义在for循环后面，当跳转到标签地方时，继续执行标签下面的代码。

func main() {
    //  放在for前面，此例会一直循环下去
    Loop:
    fmt.Println("test")
    for a:=0;a<5;a++{
        fmt.Println(a)
        if a>3{
            goto Loop
        }
    }
}

func main() {
    for a:=0;a<5;a++{
        fmt.Println(a)
        if a>3{
            goto Loop
        }
    }
    Loop:           //放在for后边
    fmt.Println("test")
}
```

```go
func main() {
    Loop:
    for j:=0;j<3;j++{
        fmt.Println(j)
        for a:=0;a<5;a++{
            fmt.Println(a)
            if a>3{
                break Loop
            }
        }
    }
}
//在没有使用loop标签的时候break只是跳出了第一层for循环
//使用标签后跳出到指定的标签,break只能跳出到之前，如果将Loop标签放在后边则会报错
//break标签只能用于for循环，跳出后不再执行标签对应的for循环    
```