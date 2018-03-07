# ipfs


## Installation

```shell
dep ensure --add github.com/computes/ipfs-http-api
```

## Packages

* [dag](dag)
* [pin](pin)
* [pubsub](pubsub)


## Usage

```go
import "github.com/computes/ipfs-http-api"
```

#### func  Cat

```go
func Cat(ipfsURL url.URL, address string) (io.ReadCloser, error)
```
Cat returns a reader for the data in IPFS located at address

#### func  ID

```go
func ID(ipfsURL url.URL) (io.ReadCloser, error)
```

ID returns a reader of the IPFS node info
