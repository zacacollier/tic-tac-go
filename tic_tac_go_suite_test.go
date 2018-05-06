package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTicTacGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TicTacGo Suite")
}
