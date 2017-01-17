package fileManager

import (
	"os"
	"log"
    "bufio"
    "errors"
	"project_elevator/elevator"
)

const filename = "requests.txt"

func WriteOrdersToFile(e elevator.ElevatorData) {
	var slice [][]byte

    for i := range(e.Orders) {
        var data []byte
        for c := range(e.Orders[i]) {
            data = append(data, byte(e.Orders[i][c]))
        }
        slice = append(slice, data)
    }

    WriteSliceToFile(slice)
}

func ReadOrdersFromFile() ([elevator.N_FLOORS][elevator.N_BUTTONS]int, error) {

    f, err := os.Open(filename)
    if err != nil {return [elevator.N_FLOORS][elevator.N_BUTTONS]int{}, err}

    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)
    
    var requests [elevator.N_FLOORS][elevator.N_BUTTONS]int

    row := 0
    for scanner.Scan() {
        line := scanner.Bytes()
        for col := range(line) {
            if row >= elevator.N_FLOORS || col >= elevator.N_BUTTONS {
                return [elevator.N_FLOORS][elevator.N_BUTTONS]int{}, errors.New("Invalid data in file.")
            }   else {
                    requests[row][col] = int(line[col])
                }
        } 
        row++
    }

    return requests, nil
}

func WriteSliceToFile(requests [][]byte) {
    f, err := os.Create(filename)
    check(err)

    defer f.Close()

    for r := range(requests) {
        var data []byte
        for c := range(requests[r]) {
            data = append(data, requests[r][c])
        }

        _, err := f.Write(data)
        check(err)
        _, err = f.WriteString("\n")
        check(err)
    }

    f.Sync()
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}