package music

import (
	scriptish "github.com/ganbarodigital/go_scriptish"
)

func playerOpen() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tmuxinator", "music"),
	)
	pipeline.Exec()
}

func playerPlayPause() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "play-pause"),
	)
	pipeline.Exec()
}

func playerStop() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "stop"),
	)
	pipeline.Exec()
}

func playerNext() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "next"),
	)
	pipeline.Exec()
}

func playerPrevious() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "previous"),
	)
	pipeline.Exec()
}

func playerShow() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "metadata", "--format", "Now Playing:\n{{title}} - {{artist}} - {{album}}"),
		scriptish.ToStdout(),
	)
	pipeline.Exec()
}

func playerShuffle() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "shuffle", "'Toggle'"),
	)
	pipeline.Exec()
}

func playerLoopTrack() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "loop", "'Track'"),
	)
	pipeline.Exec()
}

func playerLoopPlaylist() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "loop", "'Playlist'"),
	)
	pipeline.Exec()
}

func playerLoopNone() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "loop", "'None'"),
	)
	pipeline.Exec()
}
