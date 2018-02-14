# pubsub
--
    import "github.com/computes/ipfs-http-api/pubsub"


## Usage

#### func  IsDisconnectError

```go
func IsDisconnectError(err error) bool
```
IsDisconnectError tests to see if an error is a disconnect error. This can be
used to implement reconnection logic. This error will also sent if the
subscription.Close method was called.

#### func  Publish

```go
func Publish(ipfsURL url.URL, topic, payload string) error
```
Publish will publish the content to a given URL

#### type DisconnectError

```go
type DisconnectError struct{}
```

DisconnectError is returned when a pubsub sub connection is severed on the
server side

#### func (*DisconnectError) Error

```go
func (e *DisconnectError) Error() string
```

#### type Subscription

```go
type Subscription struct {
	Errors   chan error
	Messages chan []byte
}
```

Subscription is a stateful connection to IPFS

#### func  Subscribe

```go
func Subscribe(ipfsURL url.URL, topic string) (*Subscription, error)
```
Subscribe will subscribe to a given topic and returns a channel for messages and
a channel for errors

#### func (*Subscription) Close

```go
func (s *Subscription) Close() error
```
Close closes an open connection. This will return an error if the connection has
already been closed.

#### func (*Subscription) Connect

```go
func (s *Subscription) Connect() error
```
Connect establishes an IPFS connection. This method will panic if it is called
after Close because it'll try to write to closed channels
