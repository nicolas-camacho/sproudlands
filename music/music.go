package music

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	music       rl.Music
	musicVolume float32
	musicPaused bool
)

func SetInitialValues() {
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("resources/music/AveryFarm.mp3")
	musicPaused = false
	musicVolume = 0.1
	rl.SetMusicVolume(music, musicVolume)
	rl.PlayMusicStream(music)
}

func Unload() {
	rl.CloseAudioDevice()
	rl.UnloadMusicStream(music)
}

func InputHandler() {
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}
}

func Update() {
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
}
