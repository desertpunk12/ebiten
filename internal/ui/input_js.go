// Copyright 2015 Hajime Hoshi
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

// +build js

package ui

var keyCodeToKey = map[int]Key{
	32: KeySpace,
	37: KeyLeft,
	39: KeyRight,
	38: KeyUp,
	40: KeyDown,
}

var currentInput input

func IsKeyPressed(key Key) bool {
	return currentInput.isKeyPressed(key)
}

func IsMouseButtonPressed(button MouseButton) bool {
	return currentInput.isMouseButtonPressed(button)
}

func CursorPosition() (x, y int) {
	return currentInput.cursorPosition()
}

func (i *input) keyDown(key int) {
	k := keyCodeToKey[key]
	i.keyPressed[k] = true
}

func (i *input) keyUp(key int) {
	k := keyCodeToKey[key]
	i.keyPressed[k] = false
}