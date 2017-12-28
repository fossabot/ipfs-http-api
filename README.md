# ipfs

```go
import (
  ipfs "github.com/computes/ipfs-http-api"
)
```

## Usage

#### func DagGet

```go
func DagGet(ipfsURL url.URL, address string) (io.ReadCloser, error)
```

DagGet retrieves a dag object from IPFS

#### func DagGetBytes

```go
func DagGetBytes(ipfsURL url.URL, address string) ([]byte, error)
```

DagGetBytes retrieves a dag object from IPFS and reads the whole buffer into
memory

#### func DagPut

```go
func DagPut(ipfsURL url.URL, reader io.Reader) (string, error)
```

DagPut uploads a file to IPFS as a DAG Object and returns the address
