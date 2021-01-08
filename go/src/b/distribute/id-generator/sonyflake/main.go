package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sony/sonyflake"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2021-01-07")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}
	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(id)
}

func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}
	return machineID, nil
}

func checkMachineID(machineID uint16) bool {
	res, err := saddMachineIDToRedisSet()
	if err != nil || res == 0 {
		return true
	}
	if err := saveMachineIDToLocalFile(machineID); err != nil {
		return true
	}
	return false
}

func generateMachineID() (merchineID uint16, err error) {
	var i uint16
	return i, nil
}

func readMachineIDFromFile() uint16 {
	var i uint16
	return i
}

func saddMachineIDToRedisSet() (uint16, error) {
	return 0, nil
}

func saveMachineIDToLocalFile(merchineID uint16) error {
	return nil
}
