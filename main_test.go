package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "tic_tac_go"
)

var _ = Describe("Main", func() {

	Describe("UnpackSlice", func() {
		It("should unpack a slice into strings", func() {
			s := []string{"a", "b", "c"}
			var a, b, c string
			UnpackSlice(s, &a, &b, &c)
			Expect(a).To(Equal(s[0]))
			Expect(b).To(Equal(s[1]))
			Expect(c).To(Equal(s[2]))
		})
	})

	Describe("CheckForWin", func() {
		It("should find a horizontal win", func() {
			By("checking for a horizontal win")
			g := MakeBoard()
			for i := 0; i <= 2; i++ {
				g.MakeMove(1, i)
			}
			win, index, result := g.CheckForWin()
			Expect(win).To(Equal(Row))
			Expect(index).To(Equal(1))
			Expect(result).To(Equal(true))
		})
		It("should find a vertical win", func() {
			By("checking for a vertical win")
			g := MakeBoard()
			for i := 0; i <= 2; i++ {
				g.MakeMove(i, 1)
			}
			win, index, result := g.CheckForWin()
			Expect(win).To(Equal(Column))
			Expect(index).To(Equal(1))
			Expect(result).To(Equal(true))
		})
		It("should find a diagonal win", func() {
			By("checking for a diagonal win")
			g := MakeBoard()
			g.MakeMove(0, 2)
			g.MakeMove(1, 1)
			g.MakeMove(2, 0)
			win, index, result := g.CheckForWin()
			Expect(win).To(Equal(Diagonal))
			Expect(index).To(Equal(1))
			Expect(result).To(Equal(true))
		})
	})

})
