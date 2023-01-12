package main

import (
	"fmt"

	"github.com/pipeline-prototype/entity"
	"github.com/pipeline-prototype/node"
	channellib "gopkg.in/eapache/channels.v1"
)

func main() {

	dt := entity.Contents[map[string]int]{Data: map[string]int{"red": 5}}

	producer := node.NewProducer(dt, entity.Options{
		SrcId:  "test",
		DestId: "test",
		Owner:  "test",
	})
	err := producer.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	// v := producer.Out()
	// for _, e := range v {
	// 	fmt.Println(e)
	// }

	processor := node.Newprocessor("owner")
	// get the sending channel
	ch, err := producer.GetBufferedOutputChannel()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err := processor.InjectChannel(ch, true); err != nil {
		fmt.Println(err)
	}

	if err := processor.InjectChannel(channellib.NewInfiniteChannel(), false); err != nil {
		fmt.Println(err)
	}

	if err := processor.Run(); err != nil {
		fmt.Println(err)
	}

	v := processor.Out()
	for _, e := range v {
		fmt.Println(e)
	}

}
