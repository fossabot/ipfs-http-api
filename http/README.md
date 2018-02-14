# http
--
    import "github.com/computes/ipfs-http-api/http"


## Usage

#### func  Get

```go
func Get(getURL string) (io.ReadCloser, error)
```
Get will make an http Get request and handle common errors
