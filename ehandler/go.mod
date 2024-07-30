module github.com/AlanIsaacV/pkg/ehandler

go 1.22.5

//replace github.com/AlanIsaacV/pkg/gcp => ../gcp

require (
	github.com/AlanIsaacV/pkg/gcp v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.52.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/zerolog v1.33.0
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.55.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
