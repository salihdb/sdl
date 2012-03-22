package sdl_test

import "github.com/0xC3/sdl"

import "testing"

func TestInit(t *testing.T) {
   sdl.Quit() // shouldn't be needed.

   // init audio
   err := sdl.Init(sdl.InitAudio)
   if err != nil {
      t.Errorf("failed to initialize audio: %s.", err)
   }
   if sdl.WasInit(sdl.InitAudio) != sdl.InitAudio {
      t.Errorf("audio wasn't initialized.")
   }
   sdl.Quit()

   // init video
   err = sdl.Init(sdl.InitVideo)
   if err != nil {
      t.Errorf("failed to initialize video: %s.", err)
   }
   if sdl.WasInit(sdl.InitVideo) != sdl.InitVideo {
      t.Errorf("video wasn't initialized.")
   }
   sdl.Quit()

   // init without parachute
   err = sdl.Init(sdl.InitNoParachute)
   if err != nil {
      t.Errorf("failed to initialize without parachute: %s.", err)
   }
   sdl.Quit()

   // init everything
   err = sdl.Init(sdl.InitEverything)
   if err != nil {
      t.Errorf("failed to initialize everything: %s.", err)
   }
   if sdl.WasInit(sdl.InitEverything) != sdl.InitEverything {
      t.Errorf("everything wasn't initialized.")
   }
   sdl.Quit()
}

func TestInitSubSystem(t *testing.T) {
   sdl.Quit() // shouldn't be needed.

   inited := sdl.InitFlag(0)

   // init audio
   err := sdl.InitSubSystem(sdl.InitAudio)
   if err != nil {
      t.Errorf("failed to initialize audio: %s.", err)
   }
   inited |= sdl.InitAudio
   if sdl.WasInit(inited) != inited {
      t.Errorf("audio wasn't initialized.")
   }

   // init video
   err = sdl.InitSubSystem(sdl.InitVideo)
   if err != nil {
      t.Errorf("failed to initialize video: %s.", err)
   }
   inited |= sdl.InitVideo
   if sdl.WasInit(inited) != inited {
      t.Errorf("video wasn't initialized.")
   }

   sdl.Quit()
}

func TestQuit(t *testing.T) {
   sdl.Quit() // shouldn't be needed.

   // check after quit audio
   err := sdl.Init(sdl.InitAudio)
   if err != nil {
      t.Errorf("failed to initialize audio: %s.", err)
   }
   if sdl.WasInit(sdl.InitAudio) != sdl.InitAudio {
      t.Errorf("audio wasn't initialized.")
   }
   sdl.Quit()
   if sdl.WasInit(sdl.InitAudio) != 0 {
      t.Errorf("audio is still initialized after quit.")
   }

   // check after quit video
   err = sdl.Init(sdl.InitVideo)
   if err != nil {
      t.Errorf("failed to initialize video: %s.", err)
   }
   if sdl.WasInit(sdl.InitVideo) != sdl.InitVideo {
      t.Errorf("video wasn't initialized.")
   }
   sdl.Quit()
   if sdl.WasInit(sdl.InitVideo) != 0 {
      t.Errorf("video is still initialized after quit.")
   }

   // check after quit everything
   err = sdl.Init(sdl.InitEverything)
   if err != nil {
      t.Errorf("failed to initialize everything: %s.", err)
   }
   if sdl.WasInit(sdl.InitEverything) != sdl.InitEverything {
      t.Errorf("everything wasn't initialized.")
   }
   sdl.Quit()
   if sdl.WasInit(sdl.InitEverything) != 0 {
      t.Errorf("something is still initialized after quit.")
   }
}

func TestQuitSubSystem(t *testing.T) {
   sdl.Quit() // shouldn't be needed.

   // init everything
   err := sdl.Init(sdl.InitEverything)
   if err != nil {
      t.Errorf("failed to initialize everything: %s.", err)
   }
   inited := sdl.InitEverything
   if sdl.WasInit(sdl.InitEverything) != inited {
      t.Errorf("everything wasn't initialized.")
   }

   // check after quit audio
   sdl.QuitSubSystem(sdl.InitAudio)
   inited &^= sdl.InitAudio
   if sdl.WasInit(sdl.InitEverything) != inited {
      t.Errorf("audio is still initialized.")
   }

   // check after quit video
   sdl.QuitSubSystem(sdl.InitVideo)
   inited &^= sdl.InitVideo
   if sdl.WasInit(sdl.InitEverything) != inited {
      t.Errorf("video is still initialized.")
   }

   sdl.Quit()
}

func TestWasInit(t *testing.T) {
   sdl.Quit() // shouldn't be needed.

   // check after init audio
   err := sdl.Init(sdl.InitAudio)
   if err != nil {
      t.Errorf("failed to initialize audio: %s.", err)
   }
   if sdl.WasInit(sdl.InitAudio) != sdl.InitAudio {
      t.Errorf("WasInit believes audio wasn't initialized.")
   }
   sdl.Quit()

   // check after init video
   err = sdl.Init(sdl.InitVideo)
   if err != nil {
      t.Errorf("failed to initialize video: %s.", err)
   }
   if sdl.WasInit(sdl.InitVideo) != sdl.InitVideo {
      t.Errorf("WasInit believes video wasn't initialized.")
   }
   sdl.Quit()

   // check after init everything
   err = sdl.Init(sdl.InitEverything)
   if err != nil {
      t.Errorf("failed to initialize everything: %s.", err)
   }
   if sdl.WasInit(sdl.InitEverything) != sdl.InitEverything {
      t.Errorf("WasInit believes everything wasn't initialized.")
   }
   sdl.Quit()
}
