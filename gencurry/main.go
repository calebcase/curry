package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/dave/jennifer/jen"
)

func render(w io.Writer) error {
	file := jen.NewFile("curry")

	file.HeaderComment("This file is generated - do not edit.")
	file.Line()

	// These are somewhat arbitrary, but I haven't seen too many functions
	// in the wild that have more than 5 return parameters or more than 10
	// arguments.
	for rets := 1; rets <= 5; rets++ {
		for args := 2; args <= 10; args++ {
			fchain := func(i int) func(g *jen.Group) {
				return func(g *jen.Group) {
					tail := g.Func().Params(
						jen.Id(fmt.Sprintf("a%d", i)).Id(fmt.Sprintf("A%d", i)),
					)

					i++

					for ; i < args; i++ {
						tail = tail.Func().Params(
							jen.Id(fmt.Sprintf("a%d", i)).Id(fmt.Sprintf("A%d", i)),
						)
					}

					tail.ParamsFunc(func(g *jen.Group) {
						for j := 0; j < rets; j++ {
							g.Id(fmt.Sprintf("R%d", j))
						}
					})
				}
			}

			file.Func().Id(fmt.Sprintf("A%dR%d", args, rets)).TypesFunc(func(g *jen.Group) {
				var t *jen.Statement

				for i := 0; i < args; i++ {
					t = g.Id(fmt.Sprintf("A%d", i))
				}

				for j := 0; j < rets; j++ {
					t = g.Id(fmt.Sprintf("R%d", j))
				}

				t.Any()
			}).ParamsFunc(func(g *jen.Group) {
				g.Id("f").Func().ParamsFunc(func(g *jen.Group) {
					for i := 0; i < args; i++ {
						g.Id(fmt.Sprintf("a%d", i)).Id(fmt.Sprintf("A%d", i))
					}
				}).ParamsFunc(func(g *jen.Group) {
					for j := 0; j < rets; j++ {
						g.Id(fmt.Sprintf("R%d", j))
					}
				})
			}).ParamsFunc(fchain(0)).BlockFunc(func(g *jen.Group) {
				k := args - 1

				end := jen.ReturnFunc(fchain(k)).BlockFunc(func(g *jen.Group) {
					g.Return(jen.Id("f").CallFunc(func(g *jen.Group) {
						for i := 0; i < args; i++ {
							g.Id(fmt.Sprintf("a%d", i))
						}
					}))
				})

				k--

				for ; k >= 0; k-- {
					end = jen.ReturnFunc(fchain(k)).BlockFunc(func(g *jen.Group) {
						g.Add(end)
					})
				}

				g.Add(end)
			}).Line()
		}
	}

	return file.Render(w)
}

func main() {
	buf := &bytes.Buffer{}
	if err := render(buf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./curry.go", buf.Bytes(), 0644); err != nil {
		panic(err)
	}
}
