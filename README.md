# flexhash
Easy Struct data manipulation in Golang

#### Extract Method
```go
type MyStruct struct {
    Field1 string
    Field2 int
}

a := MyStruct { Field1: "a", Field2: "b" }


result = Extract(a, "Field1")
// 
// flexhash.H{
//     "Field1":"a"
// } 
//
```