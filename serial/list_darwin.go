package serial

import (
    "path/filepath"
    )

func listInternal() []SerialPort {
    files, _ := filepath.Glob("/dev/tty.*")

    list := make([]SerialPort, len(files))

    for i, file := range files {
        list[i] = SerialPort{filepath.Base(file), file}
    }

    return list
}
