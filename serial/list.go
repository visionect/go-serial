package serial

type SerialPort struct {
    Name        string
    FullName    string
}

func List() []SerialPort {
    return listInternal()
}
