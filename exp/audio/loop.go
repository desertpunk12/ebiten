// Copyright 2016 Hajime Hoshi
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

package audio

import (
	"io"
)

type Loop struct {
	stream ReadSeekCloser
	size   int64
}

func NewLoop(stream ReadSeekCloser, size int64) ReadSeekCloser {
	return &Loop{
		stream: stream,
		size:   size,
	}
}

func (l *Loop) Read(b []byte) (int, error) {
	n, err := l.stream.Read(b)
	if err == io.EOF {
		if _, err := l.Seek(0, 0); err != nil {
			return 0, err
		}
		err = nil
	}
	return n, err
}

func (l *Loop) Seek(offset int64, whence int) (int64, error) {
	next := int64(0)
	switch whence {
	case 0:
		next = offset
	case 1:
		current, err := l.stream.Seek(0, 1)
		if err != nil {
			return 0, err
		}
		next = current + offset
	case 2:
		panic("audio: whence must be 0 or 1 for a loop stream")
	}
	next %= l.size
	pos, err := l.stream.Seek(next, 0)
	if err != nil {
		return 0, err
	}
	return pos, nil
}

func (l *Loop) Close() error {
	return l.stream.Close()
}