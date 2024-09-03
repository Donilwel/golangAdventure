## once
[sync.Once](https://golang.org/pkg/sync/#Once) решает задачу ленивой инициализации.
Реализовал once без использования пакета `sync`
```go
type Once struct {}

func (o *Once) Do(f func()) {}
```
`Do` вызывает функцию `f` тогда и только тогда, когда `Do` вызывается впервые для данного экземпляра `Once`.
