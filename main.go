package contractschema

import (
	_ "embed"
)

//go:embed schema/hpse-contract-schema.json
var ContractSchema string
