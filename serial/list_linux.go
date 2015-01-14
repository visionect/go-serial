package serial

import (
    "io/ioutil"
    "os"
    "path"
    "path/filepath"
)

func listInternal() []SerialPort {
    base := "/dev/serial/by-id"
    files, _ := ioutil.ReadDir(base)

    list := make([]SerialPort, len(files))

    for i, file := range files {
        link, _ := os.Readlink(path.Join(base , file.Name()))
        list[i] = SerialPort{filepath.Base(link), path.Join(base, link)}
    }

    return list
}
