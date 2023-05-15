package main

import (
	"fmt"
	"log"
	"net/url"
)

type counter struct {
	value uint
}

func (c *counter) increment() {
	c.value++
}
func (c *counter) incrementDelta(delta uint) {
	c.value += delta
}

type usage struct {
	service string
	counter
}

func makeUsage(service string) usage {
	return usage{service, counter{}}

}

type pageviews struct {
	url *url.URL
	counter
}

func makePageviews(uri string) pageviews {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	return pageviews{u, counter{}}
}
func main() {
	usage := makeUsage("find")
	usage.increment()
	usage.increment()
	usage.increment()
	fmt.Printf("%s usage : %d\n", usage.service, usage.counter.value)

	pv := makePageviews("/doc/find")
	pv.incrementDelta(100)
	fmt.Printf("%s views : %d\n", pv.url, pv.counter.value)

}
