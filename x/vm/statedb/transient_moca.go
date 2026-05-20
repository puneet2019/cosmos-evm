package statedb

import "github.com/ethereum/go-ethereum/common"

// GetTransientState implements EIP-1153 transient storage reads.
// moca compat: required by moca's go-ethereum fork vm.StateDB interface.
func (s *StateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	if m, ok := s.transientStorage[addr]; ok {
		return m[key]
	}
	return common.Hash{}
}

// SetTransientState implements EIP-1153 transient storage writes.
// moca compat: no journaling — transient state is not reverted on snapshot rollback.
func (s *StateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	if s.transientStorage == nil {
		s.transientStorage = make(map[common.Address]map[common.Hash]common.Hash)
	}
	if s.transientStorage[addr] == nil {
		s.transientStorage[addr] = make(map[common.Hash]common.Hash)
	}
	s.transientStorage[addr][key] = value
}
