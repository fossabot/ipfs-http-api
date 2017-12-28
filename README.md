# ipfs

```go
import (
  DAG "github.com/computes/ipfs-http-api/dag"
)
```

## Usage

#### func DAG.Get

```go
func DAG.Get(ipfsURL url.URL, address string) (io.ReadCloser, error)
```

DAG.Get retrieves a dag object from IPFS

#### func DAG.GetBytes

```go
func DAG.GetBytes(ipfsURL url.URL, address string) ([]byte, error)
```

DAG.GetBytes retrieves a dag object from IPFS and reads the whole buffer into
memory

#### func DAG.Put

```go
func DAG.Put(ipfsURL url.URL, reader io.Reader) (string, error)
```

DAG.Put uploads a file to IPFS as a DAG Object and returns the address
