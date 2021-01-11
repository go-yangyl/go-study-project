## unsafe



#### 什么是unsafe

前面所说的指针是类型安全的，但它有很多限制。Go 还有非类型安全的指针，这就是 unsafe 包提供的 unsafe.Pointer。在某些情况下，它会使代码更高效，当然，也更危险。

unsafe 包用于 Go 编译器，在编译阶段使用。从名字就可以看出来，它是不安全的，官方并不建议使用。我在用 unsafe 包的时候会有一种不舒服的感觉，可能这也是语言设计者的意图吧。

但是高阶的 Gopher，怎么能不会使用 unsafe 包呢？它可以绕过 Go 语言的类型系统，直接操作内存。例如，一般我们不能操作一个结构体的未导出成员，但是通过 unsafe 包就能做到。unsafe 包让我可以直接读写内存，还管你什么导出还是未导出。



#### unsafe的实现原理

```go
type ArbitraryType int

type Pointer *ArbitraryType
```

从命名来看，`Arbitrary` 是任意的意思，也就是说 Pointer 可以指向任意类型，实际上它类似于 C 语言里的 `void*`。



unsafe 包还有其他三个函数：

```go
func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr
```

1. 任何类型的指针和 unsafe.Pointer 可以相互转换。
2. uintptr 类型和 unsafe.Pointer 可以相互转换。
3. 任何指针都可以转换为`unsafe.Pointer`
4. `unsafe.Pointer`可以转换为任何指针
5. `uintptr`可以转换为`unsafe.Pointer`
6. `unsafe.Pointer`可以转换为`uintptr`



### string转[]byte

```go

func StringToByte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
```



### []byte转string

```go
func ByteToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&stringHeader))
}
```