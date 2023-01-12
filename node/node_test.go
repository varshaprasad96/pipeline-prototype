package node

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pipeline-prototype/entity"
)

type testdata struct {
	d string
}

// TODO: find a way to declare and test generics as interface{} is not accepted!
// entity.Contents[testdata]{Data: testdata{d: "hi"}},
var _ = Describe("TestGenericsOnProducer", func() {
	var options entity.Options
	It("Testing string input", func() {
		data := entity.Contents[string]{Data: "hi"}
		options = entity.Options{
			SrcId:  "test1",
			DestId: "test1",
			Owner:  "test",
		}

		producer := NewProducer(data, options)
		err := producer.Run()
		Expect(err).NotTo(HaveOccurred())

		res := make([]interface{}, 0)
		v := producer.Out()
		res = append(res, v...)
		Expect(len(res)).To(BeEquivalentTo(4))
		for _, r := range res {
			Expect(r).To(BeEquivalentTo(data))
		}
	})
	It("Testing custom data type input", func() {
		data := entity.Contents[testdata]{Data: testdata{d: "hi"}}
		options = entity.Options{
			SrcId:  "test2",
			DestId: "test2",
			Owner:  "test",
		}

		producer := NewProducer(data, options)
		err := producer.Run()
		Expect(err).NotTo(HaveOccurred())

		res := make([]interface{}, 0)
		v := producer.Out()
		res = append(res, v...)
		Expect(len(res)).To(BeEquivalentTo(4))
		for _, r := range res {
			Expect(r).To(BeEquivalentTo(data))
		}
	})
})

func TestEventhandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Node Suite")
}
