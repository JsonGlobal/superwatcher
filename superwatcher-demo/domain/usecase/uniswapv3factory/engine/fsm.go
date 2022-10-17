package engine

import (
	"sync"

	watcherengine "github.com/artnoi43/superwatcher/domain/usecase/engine"
	"github.com/artnoi43/superwatcher/superwatcher-demo/domain/entity"
)

type poolFactoryFSM struct {
	sync.RWMutex
	states map[entity.Uniswapv3FactoryWatcherKey]watcherengine.ServiceItemState
}

func NewPoolFactoryFSM() watcherengine.ServiceFSM[entity.Uniswapv3FactoryWatcherKey] {
	return &poolFactoryFSM{
		states: make(map[entity.Uniswapv3FactoryWatcherKey]watcherengine.ServiceItemState),
	}
}

func (fsm *poolFactoryFSM) SetServiceState(key entity.Uniswapv3FactoryWatcherKey, state watcherengine.ServiceItemState) {
	fsm.Lock()
	defer fsm.Unlock()

	fsm.states[key] = state
}

func (fsm *poolFactoryFSM) GetServiceState(key entity.Uniswapv3FactoryWatcherKey) watcherengine.ServiceItemState {
	fsm.RLock()
	defer fsm.RUnlock()

	return fsm.states[key]
}
