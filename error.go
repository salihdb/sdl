package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

import "errors"

// getError returns a message about the last error that occurred.
func getError() (err error) {
   return  errors.New(C.GoString(C.SDL_GetError()))
}
