# 🌳 Music

This is a cli tool that uses playerctl to manage an instance of ncspot. I have put it into a command branch for inclusion into my c Bonzai stateful command tree.

If you just want to try it out, grab the release binary with curl and put into your PATH:

```bash
curl -L https://github.com/tr00datp00nar/music/releases/latest/download/music-linux-amd64 -o ~/.local/bin/music
curl -L https://github.com/tr00datp00nar/music/releases/latest/download/music-darwin-amd64 -o ~/.local/bin/music
curl -L https://github.com/tr00datp00nar/music/releases/latest/download/music-darwin-arm64 -o ~/.local/bin/music
curl -L https://github.com/tr00datp00nar/music/releases/latest/download/music-windows-amd64 -o ~/.local/bin/music
```
## A Note of Caution
The available functions in this package have some hard-coded values.

##### The `music open` command requires a `tmuxinator` session called 'music' to exist.

```bash
func playerOpen() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tmuxinator", "music"),
	)
	pipeline.Exec()
}
```

##### You must have an ncspot instance called 'ncspot' in order for the playerctl commands to work correctly.
```bash
func playerPlayPause() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("playerctl", "--player=ncspot", "play-pause"),
	)
	pipeline.Exec()
}

```
## Dependencies
This command uses https://github.com/hrkfdn/ncspot and https://github.com/altdesktop/playerctl and therefore carries a technical and legal dependency on them. See the respective repositories for licensing considerations and usage before deciding to use this command.

To learn more about Bonzai stateful command trees: https://github.com/rwxrob/bonzai

To see the original Bonzai stateful command tree z: https://github.com/rwxrob/z
