package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor =json.NewProcessor()

type Test struct {

}

func init() {
	Processor.Register(&Test{})
}
