package messaging

import (
	"sync"
)

var MinionStateCollector *Minions

type MinionInfo struct {
	Name                  string   `json:"name"`
	MessageFromClient     string   `json:"message_from_client"`
	MessageToClient       string   `json:"message_to_client"`
	CommunicationsChannel chan any `json:"-"`
	TriggerServer         bool     `json:"trigger_server"`
	TriggerClient         bool     `json:"trigger_client"`
	IsConnected           bool     `json:"is_connected"`
}

type MinionState struct {
	lock       sync.RWMutex
	minionInfo []*MinionInfo
}

func (m *MinionState) AddMinion(mi *MinionInfo) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.minionInfo = append(m.minionInfo, mi)
}

func NewMinionState() *MinionState {
	return &MinionState{
		minionInfo: []*MinionInfo{},
	}
}

type Minions struct {
	minionState *MinionState
}

func NewMinions() *Minions {
	return &Minions{
		minionState: NewMinionState(),
	}
}

func (m *Minions) AddMinion(mi *MinionInfo) {
	m.minionState.AddMinion(mi)
}

func (m *Minions) GetMinions() []string {
	// log.Printf("- m.minions: %+v", m.minionState.minionInfo)
	var names []string
	for _, v := range m.minionState.minionInfo {
		if v.IsConnected {
			names = append(names, string(v.Name))
		}
	}
	// log.Printf("- m.GetMinions(): %s", names)
	return names

}
