package main

import "fmt"

// IdleState struct.
type IdleState struct {
	StateInfo
}

// OnBegin func.
func (is *IdleState) OnBegin() {
	fmt.Println("Idle state begin")
}

// OnEnd func.
func (is *IdleState) OnEnd() {
	fmt.Println("Idle state end")
}

// MoveState struct.
type MoveState struct {
	StateInfo
}

// OnBegin func.
func (ms *MoveState) OnBegin() {
	fmt.Println("Move state begin")
}

// EnableSameTransit func.
func (ms *MoveState) EnableSameTransit() bool {
	return true
}

// JumpState struct.
type JumpState struct {
	StateInfo
}

// OnBegin func.
func (js *JumpState) OnBegin() {
	fmt.Println("Jump state begin")
}

// CanTransitTo func.
func (js *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

// 有限状态机实现
// go run src/a/interface/a09/{info.go,manager.go,state.go} src/a/interface/a09/main.go
func main() {
	sm := NewStateManager()
	sm.OnChange = func(from, to State) {
		fmt.Printf("%s -> %s\n\n", StateName(from), StateName(to))
	}
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	toReport(sm, "IdleState")
	toReport(sm, "MoveState")
	toReport(sm, "MoveState")
	toReport(sm, "JumpState")
	toReport(sm, "JumpState")
	toReport(sm, "IdleState")
}

func toReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("Failed: %s -> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}
