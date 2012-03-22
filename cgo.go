package sdl

// #include <SDL2/SDL.h>
import "C"

import "encoding/binary"
import "errors"
import "image"
import "log"
import "unsafe"

// goRect converts a C SDL_Rect to a go image.Rectangle.
func goRect(cRect C.SDL_Rect) (rect image.Rectangle) {
   return image.Rect(int(cRect.x), int(cRect.y), int(cRect.x + cRect.w), int(cRect.y + cRect.h))
}

// goRect converts a go image.Rectangle to a C SDL_Rect.
func cRect(rect image.Rectangle) (cRect *C.SDL_Rect) {
   if rect == image.ZR {
      return nil
   }
   cRect = new(C.SDL_Rect)
   cRect.x = C.int(rect.Min.X)
   cRect.y = C.int(rect.Min.Y)
   cRect.w = C.int(rect.Max.X - rect.Min.X)
   cRect.h = C.int(rect.Max.Y - rect.Min.Y)
   return cRect
}

// nativeEndian is set to the native endian byte order of the system.
var nativeEndian binary.ByteOrder

// initNativeEndianness sets the native endian byte order of the system.
func initNativeEndianness() (err error) {
   var i int32 = 0x01020304
   p := unsafe.Pointer(&i)
   pb := (*byte)(p)
   b := *pb
   if b == 0x01 {
      nativeEndian = binary.BigEndian
      return nil
   } else if b == 0x04 {
      nativeEndian = binary.LittleEndian
      return nil
   }
   return errors.New("unable to calculate native endianness.")
}

func init() {
   err := initNativeEndianness()
   if err != nil {
      log.Fatalln(err)
   }
}
