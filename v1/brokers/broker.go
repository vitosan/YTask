package brokers

import (
	"github.com/gojuukaze/YTask/v1/ymsg"
)

type BrokerInterface interface {
	Get(queryName string) (ymsg.Message, error)
	Send(queryName string,msg ymsg.Message) error

}
