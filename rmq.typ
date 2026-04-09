#import "@preview/codly:1.3.0": *
#import "@preview/codly-languages:0.1.1": *
#show: codly-init.with()

#set page(
	columns: 2,
	flipped: true,
	margin: (top: 2pt, bottom: 2pt, left: 2pt, right: 2pt)
)
#set columns(gutter: 12pt)


#codly(languages: codly-languages)

== go.mod

#let text = read("go.mod")
#raw(text, block: true, lang: "go")

== person.proto

#let text = read("person.proto")
#raw(text, block: true, lang: "protobuf")

== common.go

#let text = read("common/common.go")
#raw(text, block: true, lang: "go")

#pagebreak()

== recv.go

#let text = read("recv/recv.go")
#raw(text, block: true, lang: "go")

#pagebreak()

== send.go

#let text = read("send/send.go")
#raw(text, block: true, lang: "go")

#pagebreak()

== Makefile

#let text = read("Makefile")
#raw(text, block: true, lang: "make")
