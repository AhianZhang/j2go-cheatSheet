# j2go-cheetSheet
Senior Java Programmer to Junior Golang Programmer cheat sheet
java 转 go 快速对照表

# 数组
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```go
// 如果 Handler 没有实现 http.Handler，会在运行时报错
type Handler struct {
  // ...
}
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
```

</td><td>

```go
type Handler struct {
  // ...
}
// 用于触发编译期的接口的合理性检查机制
// 如果 Handler 没有实现 http.Handler，会在编译期报错
var _ http.Handler = (*Handler)(nil)
func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```

</td></tr>
</tbody></table>
