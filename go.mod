module github.com/myriadeinc/whetstone

go 1.15

replace github.com/myriadeinc/whetstone/server => ./server

require (
	github.com/myriadeinc/whetstone/server v0.0.0-20201218192446-6380adb0e4d2
	github.com/rs/zerolog v1.20.0
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
)
