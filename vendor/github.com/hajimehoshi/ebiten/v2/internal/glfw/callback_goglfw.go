// Copyright 2021 The Ebiten Authors
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

//go:build windows

package glfw

import (
	"github.com/hajimehoshi/ebiten/v2/internal/goglfw"
)

func ToCharModsCallback(cb func(window *Window, char rune, mods ModifierKey)) CharModsCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window, char rune, mods goglfw.ModifierKey) {
		cb((*Window)(window), char, ModifierKey(mods))
	}
}

func ToCloseCallback(cb func(window *Window)) CloseCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window) {
		cb((*Window)(window))
	}
}

func ToDropCallback(cb func(window *Window, names []string)) DropCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window, names []string) {
		cb((*Window)(window), names)
	}
}

func ToFramebufferSizeCallback(cb func(window *Window, width int, height int)) FramebufferSizeCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window, width int, height int) {
		cb((*Window)(window), width, height)
	}
}

func ToMonitorCallback(cb func(monitor *Monitor, event PeripheralEvent)) MonitorCallback {
	if cb == nil {
		return nil
	}
	return func(monitor *goglfw.Monitor, event goglfw.PeripheralEvent) {
		cb((*Monitor)(monitor), PeripheralEvent(event))
	}
}

func ToScrollCallback(cb func(window *Window, xoff float64, yoff float64)) ScrollCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window, xoff float64, yoff float64) {
		cb((*Window)(window), xoff, yoff)
	}
}

func ToSizeCallback(cb func(window *Window, width int, height int)) SizeCallback {
	if cb == nil {
		return nil
	}
	return func(window *goglfw.Window, width int, height int) {
		cb((*Window)(window), width, height)
	}
}
