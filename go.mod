module github.com/etclabscore/go-openrpc-reflect

go 1.13

require github.com/open-rpc/meta-schema v0.0.42

require (
	github.com/alecthomas/jsonschema v0.0.2
	github.com/etclabscore/go-jsonschema-walk v0.0.5
	github.com/ethereum/go-ethereum v1.9.12
	github.com/go-openapi/spec v0.19.7
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/gjson v1.6.0
)

replace github.com/open-rpc/meta-schema => github.com/meowsbits/meta-schema v0.0.42

replace github.com/alecthomas/jsonschema => github.com/etclabscore/go-jsonschema-reflect v0.0.2
