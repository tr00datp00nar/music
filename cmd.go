package music

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//	     ────────────────────────────────────────────────────────────
//									Main Cmd
//	     ────────────────────────────────────────────────────────────
var Cmd = &Z.Cmd{
	Name:        `music`,
	Usage:       `COMMAND`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_music),
	Description: help.D(_music),

	Commands: []*Z.Cmd{
		loopCmd,
		nextCmd,
		openCmd,
		playCmd,
		previousCmd,
		showCmd,
		shuffleCmd,
		stopCmd,
		help.Cmd,
	},
	Other: []Z.Section{
		{`Dependencies`, `
				This {{cmd .Name}} command uses github.com/hrkfdn/ncspot and github.com/altdesktop/playerctl and
				therefore carries a technical and legal dependency on them. See the respective repositories for
				licensing considerations and usage before deciding to use this command.
			`,
		},
	},
}

//	     ────────────────────────────────────────────────────────────
//									Core Cmds
//	     ────────────────────────────────────────────────────────────
var openCmd = &Z.Cmd{
	Name:        `open`,
	Usage:       `[help]`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_open),
	Description: help.D(_open),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerOpen()
		return nil
	},
}

var playCmd = &Z.Cmd{
	Name:        `play`,
	Aliases:     []string{`p`, `pause`},
	Usage:       `[help]`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_playpause),
	Description: help.D(_playpause),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerPlayPause()
		return nil
	},
}

var stopCmd = &Z.Cmd{
	Name:        `stop`,
	Usage:       `[help|stop]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_stop),
	Description: help.D(_stop),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerStop()
		return nil
	},
}

var previousCmd = &Z.Cmd{
	Name:        `previous`,
	Aliases:     []string{`prev`},
	Usage:       `[help|previous]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_previous),
	Description: help.D(_previous),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerPrevious()
		return nil
	},
}

var nextCmd = &Z.Cmd{
	Name:        `next`,
	Usage:       `[help|next]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_next),
	Description: help.D(_next),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerNext()
		return nil
	},
}

var showCmd = &Z.Cmd{
	Name:        `show`,
	Usage:       `[help|show]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_show),
	Description: help.D(_show),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerShow()
		return nil
	},
}

var shuffleCmd = &Z.Cmd{
	Name:        `shuffle`,
	Usage:       `[help|shuffle]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_shuffle),
	Description: help.D(_shuffle),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerShuffle()
		return nil
	},
}

//	     ────────────────────────────────────────────────────────────
//									Loop Cmds
//	     ────────────────────────────────────────────────────────────
var loopCmd = &Z.Cmd{
	Name:        `loop`,
	Usage:       `COMMAND`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_loop),
	Description: help.D(_loop),

	Commands: []*Z.Cmd{
		help.Cmd,
		loopTrackCmd,
		loopPlaylistCmd,
		loopNoneCmd,
	},
}

var loopTrackCmd = &Z.Cmd{
	Name:        `track`,
	Usage:       `[help|track]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_looptrack),
	Description: help.D(_looptrack),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerLoopTrack()
		return nil
	},
}

var loopPlaylistCmd = &Z.Cmd{
	Name:        `playlist`,
	Usage:       `[help|playlist]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_loopplaylist),
	Description: help.D(_loopplaylist),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerLoopPlaylist()
		return nil
	},
}

var loopNoneCmd = &Z.Cmd{
	Name:        `none`,
	Usage:       `[help|none]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_loopnone),
	Description: help.D(_loopnone),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		playerLoopNone()
		return nil
	},
}
