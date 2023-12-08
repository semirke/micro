# Options

Micro stores all of the user configuration in its configuration directory.

Micro uses the `$XDG_CONFIG_HOME/micro` as the configuration directory. As per
the XDG spec, if `$XDG_CONFIG_HOME` is not set, `~/.config/micro` is used as
the config directory.

Here are the options that you can set:

* `autoindent`: when creating a new line use the same indentation as the
   previous line.

	default value: `on`

* `autosave`: micro will save the buffer every 8 seconds automatically. Micro
   also will automatically save and quit when you exit without asking. Be
   careful when using this feature, because you might accidentally save a file,
   overwriting what was there before.

	default value: `off`

* `colorcolumn`: if this is not set to 0, it will display a column at the
  specified column. This is useful if you want column 80 to be highlighted
  special for example.

	default value: `0`

* `colorscheme`: loads the colorscheme stored in
   $(configDir)/colorschemes/`option`.micro, This setting is `global only`.

	default value: `default`

	Note that the default colorschemes (default, solarized, and solarized-tc)
	are not located in configDir, because they are embedded in the micro binary.

	The colorscheme can be selected from all the files in the
	~/.config/micro/colorschemes/ directory. Micro comes by default with
	three colorschemes:

	You can read more about micro's colorschemes in the `colors` help topic
	(`help colors`).

* `cursorline`: highlight the line that the cursor is on in a different color
   (the color is defined by the colorscheme you are using).

	default value: `on`

* `cursornum`: highlight the row number at the cursor. Uses the color of cursor-line.

   default value: `on`

* `eofnewline`: micro will automatically add a newline to the file.

	default value: `true`

* `fakecursor`: forces micro to render the cursor using terminal colors rather
  than the actual terminal cursor. This is useful when the terminal's cursor is
  slow or otherwise unavailable/undesirable to use.

    default value: `false`

* `fastdirty`: this determines what kind of algorithm micro uses to determine if
   a buffer is modified or not. When `fastdirty` is on, micro just uses a
   boolean `modified` that is set to `true` as soon as the user makes an edit.
   This is fast, but can be inaccurate. If `fastdirty` is off, then micro will
   hash the current buffer against a hash of the original file (created when the
   buffer was loaded). This is more accurate but obviously more resource
   intensive. This option is only for people who really care about having
   accurate modified status.

	default value: `on`

* `fileformat`: this determines what kind of line endings micro will use for the
   file. UNIX line endings are just `\n` (lf) whereas dos line endings are
   `\r\n` (crlf). The two possible values for this option are `unix` and `dos`.
   The fileformat will be automatically detected and displayed on the statusline
   but this option is useful if you would like to change the line endings or if
   you are starting a new file.

	default value: `unix`

* `filetype`: sets the filetype for the current buffer. Set this option to
  `off` to completely disable filetype detection.

	default value: `unknown`. This will be automatically overridden depending
    on the file you open.

* `hlsearch`: highlight all instances of the searched text after a successful
   search. This highlighting can be temporarily turned off via the
   `UnhighlightSearch` action (triggered by the Esc key by default) or toggled
   on/off via the `ToggleHighlightSearch` action. Note that these actions don't
   change the `hlsearch` setting. As long as `hlsearch` is set to true, the next
   search will have the highlighting turned on again.

	default value: `on`

* `hltaberrors`: highlight tabs when spaces are expected, and spaces when tabs
   are expected. More precisely: if `tabstospaces` option is on, highlight
   all tab characters; if `tabstospaces` is off, highlight space characters
   in the initial indent part of the line.

	default value: `false`

* `hltrailingws`: highlight trailing whitespaces at ends of lines. Note that
   it doesn't highlight newly added trailing whitespaces that naturally occur
   while typing text. It highlights only nasty forgotten trailing whitespaces.

	default value: `false`

* `incsearch`: enable incremental search in "Find" prompt (matching as you type).

	default value: `true`

* `ignorecase`: perform case-insensitive searches.

	default value: `off`

* `indentchar`: sets the indentation character.

	default value: ` `

* `infobar`: enables the line at the bottom of the editor where messages are
   printed. This option is `global only`.

	default value: `on`

* `keepautoindent`: when using autoindent, whitespace is added for you. This
   option determines if when you move to the next line without any insertions
   the whitespace that was added should be deleted. By default the autoindent
   whitespace is deleted if the line was left empty.

	default value: `off`

* `keymenu`: display the nano-style key menu at the bottom of the screen. Note
   that ToggleKeyMenu is bound to `Alt-g` by default and this is displayed in
   the statusline. To disable the key binding, bind `Alt-g` to `None`.

	default value: `off`

* `mouse`: whether to enable mouse support. When mouse support is disabled,
   usually the terminal will be able to access mouse events which can be useful
   if you want to copy from the terminal instead of from micro (if over ssh for
   example, because the terminal has access to the local clipboard and micro
   does not).

	default value: `true`

* `multiopen`: specifies how to layout multiple files opened at startup.
   Most useful as a command-line option, like `-multiopen vsplit`. Possible
   values correspond to commands (see `> help commands`) that open files:
    * `tab`: open each file in a separate tab.
    * `vsplit`: open files side-by-side.
    * `hsplit`: open files stacked top to bottom.

	default value: `tab`

* `paste`: treat characters sent from the terminal in a single chunk as a paste
   event rather than a series of manual key presses. If you are pasting using
   the terminal keybinding (not Ctrl-v, which is micro's default paste
   keybinding) then it is a good idea to enable this option during the paste
   and disable once the paste is over. See `> help copypaste` for details about
   copying and pasting in a terminal environment.

    default value: `false`

* `parsecursor`: if enabled, this will cause micro to parse filenames such as
   file.txt:10:5 as requesting to open `file.txt` with the cursor at line 10
   and column 5. The column number can also be dropped to open the file at a
   given line and column 0. Note that with this option enabled it is not possible
   to open a file such as `file.txt:10:5`, where `:10:5` is part of the filename.
   It is also possible to open a file with a certain cursor location by using the
   `+LINE:COL` flag syntax. See `micro -help` for the command line options.

    default value: `false`

* `permbackup`: this option causes backups (see `backup` option) to be
   permanently saved. With permanent backups, micro will not remove backups when
   files are closed and will never apply them to existing files. Use this option
   if you are interested in manually managing your backup files.

    default value: `false`

* `pluginchannels`: contains all the channels micro's plugin manager will search
   for plugins in. A channel is simply a list of 'repository' json files which
   contain metadata about the given plugin. See the `Plugin Manager` section of
   the `plugins` help topic for more information.

	default value: `https://github.com/micro-editor/plugin-channel`

* `pluginrepos`: contains all the 'repositories' micro's plugin manager will
   search for plugins in. A repository consists of a `repo.json` file which
   contains metadata for a single plugin.

	default value: ` `

* `readonly`: when enabled, disallows edits to the buffer. It is recommended
   to only ever set this option locally using `setlocal`.

    default value: `false`

* `usectags`: when enabled and universal ctags is available, it'll be used
   to extract tags that'll be fed to the tags autocomplete.

    default value: `true`

* `reload`: controls the reload behavior of the current buffer in case the file
   has changed. The available options are `prompt`, `auto` & `disabled`.

   default value: `prompt`

* `rmtrailingws`: micro will automatically trim trailing whitespaces at ends of
   lines. Note: This setting overrides `keepautoindent`

	default value: `false`

* `ruler`: display line numbers.

	default value: `true`

* `relativeruler`: make line numbers display relatively. If set to true, all
   lines except for the line that the cursor is located will display the distance
   from the cursor's line.

	default value: `false`

* `savecursor`: remember where the cursor was last time the file was opened and
   put it there when you open the file again.

	default value: `off`

* `savehistory`: remember command history between closing and re-opening
   micro.

    default value: `on`

* `saveundo`: when this option is on, undo is saved even after you close a file
   so if you close and reopen a file, you can keep undoing.

	default value: `off`

* `scrollbar`: display a scroll bar

    default value: `off`

* `scrollbarchar`: specifies the character used for displaying the scrollbar

    default value: `|`

* `scrollmargin`: margin at which the view starts scrolling when the cursor
   approaches the edge of the view.

	default value: `3`

* `scrollspeed`: amount of lines to scroll for one scroll event.

	default value: `2`

* `softwrap`: should micro wrap lines that are too long to fit on the screen.

	default value: `off`

* `splitbottom`: when a horizontal split is created, should it be created below
   the current split?

	default value: `on`

* `splitright`: when a vertical split is created, should it be created to the
   right of the current split?

	default value: `true`

* `statusformatl`: format string definition for the left-justified part of the
   statusline. Special directives should be placed inside `$()`. Special
   directives include: `filename`, `modified`, `line`, `col`, `lines`,
   `percentage`, `opt`, `bind`.
   The `opt` and `bind` directives take either an option or an action afterward
   and fill in the value of the option or the key bound to the action.

    default value: `$(filename) $(modified)($(line),$(col)) $(status.paste)|
                    ft:$(opt:filetype) | $(opt:fileformat) | $(opt:encoding)`

* `statusline`: display the status line at the bottom of the screen.

	default value: `on`

* `syntax`: turns syntax on or off.

	default value: `on`

* `sucmd`: specifies the super user command. On most systems this is "sudo" but
   on BSD it can be "doas." This option can be customized and is only used when
   saving with su.

	default value: `sudo`

* `tabmovement`: navigate spaces at the beginning of lines as if they are tabs
   (e.g. move over 4 spaces at once). This option only does anything if
   `tabstospaces` is on.

	default value: `off`

* `tabhighlight`: inverts the tab characters' (filename, save indicator, etc)
  colors with respect to the tab bar.

	default value: false

* `tabreverse`: reverses the tab bar colors when active.

	default value: true

* `tabsize`: the size in spaces that a tab character should be displayed with.

	default value: `4`

* `numberedtabs`: Display the tab's number beside it in the tab bar.
	default value: `on`

* `showclock`: Display a clock in the status line.
	(May be off by up to 20 seconds to save on performance. )
	*Note:* Setting this causes micro to redraw the screen at least 4 times a
	minute, instead of just when it receives input. Hence why it is off by default.

	default value: `off`

* `12hourclock`: Display time in a 12-hour format, rather than the default 24-hour.

	default value: `off`

* `showseconds`: Display seconds in the status line clock.
	(This is a bit more resource intensive as it redraws the screen every second.)

	default value: `off`

* `termtitle`: defines whether or not your terminal's title will be set by micro when opened.
	default value: `off`

* `tabstospaces`: use spaces instead of tabs

	default value: `off`

* `termtitle`: defines whether or not your terminal's title will be set by micro
   when opened.

	default value: `off`

* `useprimary` (only useful on *nix): defines whether or not micro will use the
   primary clipboard to copy selections in the background. This does not affect
   the normal clipboard using Ctrl-C and Ctrl-V.

	default value: `on`

---

Default plugin options:

* `autoclose`: automatically close `{}` `()` `[]` `""` `''`. Provided by the
   `autoclose` plugin

	default value: `on`

* `ftoptions`: by default, micro will set some options based on the filetype. At
   the moment, micro will use tabs for makefiles and spaces for python and yaml
   files regardless of your settings. If you would like to disable this behavior
   turn this option off.

	default value: `on`

* `linter`: Automatically lint when the file is saved. Provided by the `linter`
   plugin.

	default value: `on`

Any option you set in the editor will be saved to the file
~/.config/micro/settings.json so, in effect, your configuration file will be
created for you. If you'd like to take your configuration with you to another
machine, simply copy the settings.json to the other machine.

## Settings.json file

The settings.json file should go in your configuration directory (by default
at `~/.config/micro`), and should contain only options which have been modified
from their default setting. Here is the full list of options in json format,
so that you can see what the formatting should look like.

```json
{
    "autoclose": true,
    "autoindent": true,
    "autosave": 0,
    "autosu": false,
    "backup": true,
    "backupdir": "",
    "basename": false,
    "clipboard": "external",
    "colorcolumn": 0,
    "colorscheme": "default",
    "comment": true,
    "cursorline": true,
    "diff": true,
    "diffgutter": false,
    "divchars": "|-",
    "divreverse": true,
    "encoding": "utf-8",
    "eofnewline": true,
    "fastdirty": false,
    "fileformat": "unix",
    "filetype": "unknown",
    "incsearch": true,
    "ftoptions": true,
    "ignorecase": true,
    "indentchar": " ",
    "infobar": true,
    "initlua": true,
    "keepautoindent": false,
    "keymenu": false,
    "linter": true,
    "literate": true,
    "matchbrace": true,
    "mkparents": false,
    "mouse": true,
    "parsecursor": false,
    "paste": false,
    "permbackup": false,
    "pluginchannels": [
        "https://raw.githubusercontent.com/micro-editor/plugin-channel/master/channel.json"
    ],
    "pluginrepos": [],
    "readonly": false,
    "relativeruler": false,
    "rmtrailingws": false,
    "ruler": true,
    "savecursor": false,
    "savehistory": true,
    "saveundo": false,
    "scrollbar": false,
    "scrollmargin": 3,
    "scrollspeed": 2,
    "smartpaste": true,
    "softwrap": false,
    "splitbottom": true,
    "splitright": true,
    "status": true,
    "statusformatl": "$(filename) $(modified)($(line),$(col)) $(status.paste)| ft:$(opt:filetype) | $(opt:fileformat) | $(opt:encoding)",
    "statusformatr": "$(bind:ToggleKeyMenu): bindings, $(bind:ToggleHelp): help",
    "statusline": true,
    "sucmd": "sudo",
    "syntax": true,
    "tabmovement": false,
    "tabhighlight": true,
    "tabreverse": false,
    "tabsize": 4,
    "tabstospaces": false,
    "useprimary": true,
    "xterm": false
}
```

## Global and local settings

You can set these settings either globally or locally. Locally means that the
setting won't be saved to `~/.config/micro/settings.json` and that it will only
be set in the current buffer. Setting an option globally is the default, and
will set the option in all buffers.

The `colorscheme` option is global only, and the `filetype` option is local
only. To set an option locally, use `setlocal` instead of `set`.

In the `settings.json` file you can also put set options locally by specifying a
glob. Here is an example which has `tabstospaces` on for all files except Go
files, and `tabsize` 4 for all files except Ruby files:

```json
{
	"*.go": {
		"tabstospaces": false
	},
	"*.rb": {
		"tabsize": 2
	},
	"tabstospaces": true,
	"tabsize": 4
}
```

As you can see it is quite easy to set options locally using the `settings.json`
file.
