package serial

type SerialPort struct {
    Name    string
    Path    string
}

func List() []SerialPort {
    return listInternal()
}
