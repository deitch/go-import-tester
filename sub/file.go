package sub

import (
	"log"

	"github.com/containerd/containerd"

	"github.com/deitch/go-import-tester/initter"

	_ "github.com/deitch/go-import-tester/noinit"
)

// Const is a constant
const Const = "Hello, Sub World!"

func Func() {
	_, err := containerd.New("")
	if err != nil {
		log.Fatal(err)
	}

}

func Init() {
	_ = initter.Const
}
