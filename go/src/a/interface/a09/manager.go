package main

import "errors"

// StateManager struct.
type StateManager struct {
	// 已经添加的状态
	stateByName map[string]State

	// 状态改变时的回调
	OnChange func(from, to State)

	// 当前状态
	curr State
}

// Get func.
func (sm *StateManager) Get(name string) State {
	if v, ok := sm.stateByName[name]; ok {
		return v
	}
	return nil
}

// Add func.
func (sm *StateManager) Add(s State) {
	name := StateName(s)
	s.(interface {
		setName(name string)
	}).setName(name)
	if sm.Get(name) != nil {
		panic("Duplicate state: " + name)
	}
	sm.stateByName[name] = s
}

// NewStateManager func.
func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

// ErrStateNotFound var.
var ErrStateNotFound = errors.New("State not found")

// ErrForbidSameStateTransit var.
var ErrForbidSameStateTransit = errors.New("Forbid same state transit")

// ErrCannotTransitToState var.
var ErrCannotTransitToState = errors.New("Cannot transit to state")

// CurrState func.
func (sm *StateManager) CurrState() State {
	return sm.curr
}

// CanCurrTransitTo func.
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil {
		return true
	}
	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}
	return sm.curr.CanTransitTo(name)
}

// Transit func.
func (sm *StateManager) Transit(name string) error {
	next := sm.Get(name)
	if next == nil {
		return ErrStateNotFound
	}

	prev := sm.curr
	if sm.curr != nil {
		if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}
		sm.curr.OnEnd()
	}
	sm.curr = next
	sm.curr.OnBegin()
	if sm.OnChange != nil {
		sm.OnChange(prev, sm.curr)
	}
	return nil
}
