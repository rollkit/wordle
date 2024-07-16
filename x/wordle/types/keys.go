package types

const (
	// ModuleName defines the module name
	ModuleName = "wordle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wordle"
)

var (
	ParamsKey = []byte("p_wordle")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
