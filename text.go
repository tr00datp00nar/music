package music

import _ "embed"

//go:embed text/en/music.md
var _music string

//go:embed text/en/loop.md
var _loop string

//go:embed text/en/looptrack.md
var _looptrack string

//go:embed text/en/loopplaylist.md
var _loopplaylist string

//go:embed text/en/loopnone.md
var _loopnone string

//go:embed text/en/next.md
var _next string

//go:embed text/en/open.md
var _open string

//go:embed text/en/playpause.md
var _playpause string

//go:embed text/en/previous.md
var _previous string

//go:embed text/en/show.md
var _show string

//go:embed text/en/shuffle.md
var _shuffle string

//go:embed text/en/stop.md
var _stop string
