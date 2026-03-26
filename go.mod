module example/hello

go 1.26.1

require (
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.3.0
)

require golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect

replace example.com/greetings => ../greetings
