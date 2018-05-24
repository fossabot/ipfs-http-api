package pubsub

import (
	"encoding/json"
)

// SubscriptionMessage contains the data and metadata from
// the IPFS message
type SubscriptionMessage struct {
	From     string          `json:"from"`
	Data     json.RawMessage `json:"data"`
	Seqno    string          `json:"seqno"`
	TopicIDs []string        `json:"topicIDs"`
}

// DataAsString will return the Data property as a string
func (m *SubscriptionMessage) DataAsString() (string, error) {
	b, err := m.DataAsBytes()
	return string(b), err
}

// DataAsBytes will return the Data property as a byte array
func (m *SubscriptionMessage) DataAsBytes() ([]byte, error) {
	var b []byte
	err := json.Unmarshal(m.Data, &b)

	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
