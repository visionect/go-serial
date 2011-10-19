// Copyright 2011 Aaron Jacobs. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file contains OS-specific constants and types.

package serial

import "io"
import "os"

// termios types
type cc_t byte
type speed_t uint64
type tcflag_t uint64

// sys/termios.h
const (
	B9600 = 9600
	B14400 = 14400
	B19200 = 19200

	CS5 = 0x00000000
	CS6 = 0x00000100
	CS7 = 0x00000200
	CS8 = 0x00000300
	CLOCAL = 0x00008000
	CREAD = 0x00000800
	IGNPAR = 0x00000004

	NCCS = 20;

	VMIN = tcflag_t(16);
	VTIME = tcflag_t(17);
)

type termios struct {
	c_iflag tcflag_t
	c_oflag tcflag_t
	c_cflag tcflag_t
	c_lflag tcflag_t
	c_cc [NCCS]cc_t
	c_ispeed speed_t
	c_ospeed speed_t
}

func openInternal(options OpenOptions) (io.ReadWriteCloser, os.Error) {
	return nil, os.NewError("Not implemented.")
}
