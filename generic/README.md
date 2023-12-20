### 泛型

#### 自定义泛型数据类型
~ 代表底层是这个类型就行, 没有~ 就必须是同一类型
```go
type XXXX interface {
	~int | ~int8    
} 

```

#### 泛型结构体
```go
type XXX[T any] struct {
    Value T `json:"value"`
	f func(string) T
}

func (l XXX[T]) Get(key string) T {
	return l.f(key)
}
```

####
[slice处理的泛型包](golang.org/x/exp/slices)
