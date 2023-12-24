package types

const (
	// ModuleName defines the module name
	ModuleName = "min"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_min"
)

var (
	ParamsKey = []byte("p_min")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
