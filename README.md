# ipfs
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcomputes%2Fipfs-http-api.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcomputes%2Fipfs-http-api?ref=badge_shield)



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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fcomputes%2Fipfs-http-api.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fcomputes%2Fipfs-http-api?ref=badge_large)