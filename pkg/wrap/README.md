# 概要
- golangの組み込み型（string, intなど）をラップしたパッケージです
- 処理で複雑なことはしていません

# usage
- 下記のように `import`することでパッケージ名を指定する必要がなくなります

```go
import . "github.com/oguri-souhei/goutil/pkg/wrap"

var wrapedString String = "foobar" // Stringはwrapパッケージの型
```