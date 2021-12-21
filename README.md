## msignal 📶 - Basic Utilities for OS Signal Management

### Usage

```go
func main() {
	syscall_ctx := msignal.DefaultCtx()

Loop:
	for {
		select {
		case <-syscall_ctx.Done():
			break Loop
		default:
			fmt.Println("Hello World!")
			time.Sleep(time.Millisecond * 200)
		}
	}
}
```
