### 字符串合并

```go
// 第一种 + 合并
// 第二种 fmt.Sprintf
// 第三种 strings.builder
  var buf strings.Builder
	buf.WriteString("h")
	buf.WriteString("e")
	fmt.Println(buf.String())
// 第四种bytes.Buffer
  var buf1 bytes.Buffer
	buf1.WriteString("h")
	buf1.WriteString("e")
	fmt.Println(buf1.String())
```



### 字符串转int

```go
  var str = "10"
	res, _ := strconv.Atoi(str)
	fmt.Println(res)
```



### 字符串转int64

```go
  var str = "10"
	res, _ := strconv.ParseInt(str, 10, 64) // 十进制 int64
	fmt.Println(res)
```



### int转string

```go
  var i = 1
	res := strconv.Itoa(i)
	fmt.Println(res)
```



### int64转string

```go
  var i int64 = 1
	res := strconv.FormatInt(i, 10)
	fmt.Println(res)
```



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

