
```go
package main

    import "flag"

    func main(){
        var name string
	    flag.StringVar(&name, name, "默认值为", "注释")
	    flag.String(name, "默认值为", "注释")
    }
    
```