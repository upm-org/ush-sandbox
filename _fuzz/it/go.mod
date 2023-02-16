module local.tld/fuzz

go 1.13

replace github.com/upm-org/ush => ./../..

require (
	cloud.google.com/go/firestore v1.0.0 // indirect
	github.com/containerd/containerd v1.5.18 // indirect
	github.com/dvyukov/go-fuzz v0.0.0-20191008232133-fdaa9b19a67d
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/fuzzitdev/fuzzit/v2 v2.4.73
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/pierrec/lz4 v2.3.0+incompatible // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/stephens2424/writerset v1.0.2 // indirect
	github.com/upm-org/ush v2.6.4+incompatible
)
