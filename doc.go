/*
Package sdl contains bindings to the SDL library.

The application must quit the SDL library when finished with it:
   err := sdl.Init(sdl.InitVideo)
   if err != nil {
      return err
   }
   defer sdl.Quit()
*/
package sdl
