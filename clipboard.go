package sdl

// #cgo pkg-config: sdl2
// #include <SDL2/SDL.h>
import "C"

import "unsafe"

// GetClipboardText returns text from the clipboard.
//
// Note: A Window must be created before calling this function.
func GetClipboardText() (text string, err error) {
   s := C.SDL_GetClipboardText()
   if s == nil {
      return "", getError()
   }
   defer C.SDL_free(unsafe.Pointer(s))
   text = C.GoString(s)
   return text, nil
}

// HasClipboardText returns true if the clipboard exists and contains a
// non-empty text string.
//
// Note: A Window must be created before calling this function.
func HasClipboardText() bool {
   if C.SDL_HasClipboardText() == C.SDL_TRUE {
      return true
   }
   return false
}

// SetClipboardText puts text into the clipboard.
//
// Note: A Window must be created before calling this function.
func SetClipboardText(text string) (err error) {
   if C.SDL_SetClipboardText(C.CString(text)) != 0 {
      return getError()
   }
   return nil
}
