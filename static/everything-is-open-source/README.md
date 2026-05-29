# tileseq

`tileseq` is a small single-file Go program (`tileseq.go`) that renders an
animated GIF on a grid of colored tiles. The canvas is subdivided into equally
sized square tiles; a piece of text is drawn onto that grid using a built-in
bitmap font, emerging out of a field of flickering random colors and then
freezing so it stays readable.

```
go run tileseq.go -text "GOLANG" -palette neon -output out.gif
```

## How it works

### Grid

The canvas (`-width` x `-height`, in pixels) is divided into square tiles of
`-tile` pixels. That gives `cols = width/tile` by `rows = height/tile` grid
cells; every cell is painted as a single flat color. All animation happens at
the tile level — each frame is a full `image.Paletted` repaint encoded into one
`image/gif` animation.

### Font and text layout

A built-in **5x7 bitmap font** covers A–Z, 0–9 and a little punctuation
(`. , ! ? - : '`). Lowercase input is upper-cased; unknown characters render as
`?`. Each font pixel is expanded into a `scale` x `scale` block of tiles
(`-scale`, or `0` to auto-fit the largest scale that fits the grid).

Text can be **multiple lines**:

- `-text` accepts literal `\n` for line breaks.
- `-file` reads text from a file, one animation line per file line (overrides
  `-text`). CRLF is normalized and a single trailing newline is ignored;
  interior blank lines are kept as vertical spacing.

Lines are each centered horizontally, the whole block is centered vertically,
`-gap` controls space between letters and `-linegap` the space between lines
(both in font-pixels). Layout produces a boolean **text mask** (which tiles are
"on") plus the per-character **cells** in reading order, used by the cursor.

### Phases

Each run is built from four sequential phases; frame counts are configurable:

1. **flux** (`-flux`) — every tile flickers through random palette colors.
2. **reveal** (`-reveal`) — the text is progressively drawn out of the noise.
   Each tile has a random reveal threshold; once the reveal progress passes it,
   the tile *locks* to its final color (text tiles become the text color,
   others settle onto a fixed background color). This makes letters emerge
   organically rather than all at once.
3. **scan** (`-scan`, frames *per character*; `0` disables) — after the text is
   fully settled, a **block cursor** sweeps over each character, line by line in
   reading order. The current cell is inverted (filled with the text color, the
   glyph punched out dark), giving a terminal-style scanning highlight.
4. **freeze** (`-freeze`) — the plain finished view is held still long enough to
   read.

### Color

- `-palette` selects a set of background flicker colors. Available palettes:
  `candy, cool, earth, forest, miami, mono, neon, ocean, pastel, rainbow,
  sunset, warm`.
- `-textcolor` sets the revealed-text color as a `#rrggbb` hex value or a name
  (`white, black, red, green, blue, yellow, cyan, magenta`); empty picks a
  sensible default per palette.
- `-bold` thickens the glyph strokes by dilating the text mask N times (8-
  connected). Because dilation also eats into letter counters, use `-scale 2`+
  and a slightly larger `-gap` when bolding so letters stay legible and
  separated.
- `-dim` (`0`..`1`, default `1` = off) darkens the background **behind and
  around** the text to make it pop; `-dimpad` sets how many tiles of halo around
  the letters are dimmed. Implemented by adding dimmed variants of each flicker
  color to the GIF palette and using them for settled tiles inside the halo.

### Animation timing

`-delay` is the per-frame delay in hundredths of a second (default `6` ≈ 60ms).
With the default `-scan 2`, the cursor spends ~120ms per character — a
comfortable scanning pace. `-seed` makes the random fluctuation reproducible.

## Flags

| flag | default | meaning |
|------|---------|---------|
| `-text` | `HELLO` | text to render; `\n` for line breaks |
| `-file` | | read text from a file (overrides `-text`) |
| `-width` / `-height` | `640` / `360` | canvas size in pixels |
| `-tile` | `16` | tile size in pixels |
| `-scale` | `0` | font scale in tiles per font-pixel (`0` = auto-fit) |
| `-gap` | `1` | gap between letters (font-pixels) |
| `-linegap` | `1` | vertical gap between lines (font-pixels) |
| `-bold` | `0` | stroke thickness (dilation passes) |
| `-palette` | `neon` | background palette (see list above) |
| `-textcolor` | | revealed-text color (`#rrggbb` or name) |
| `-dim` | `1.0` | background brightness behind text (`0`=black .. `1`=off) |
| `-dimpad` | `1` | halo padding around text, in tiles |
| `-flux` / `-reveal` / `-scan` / `-freeze` | `18` / `28` / `2` / `30` | phase frame counts (`-scan` is per character) |
| `-delay` | `6` | frame delay in 1/100s |
| `-seed` | `42` | random seed |
| `-output` | `tileseq.gif` | output GIF filename |

## Examples

```sh
# Three lines from a file, dimmed halo, white text
go run tileseq.go -file lines.txt -scale 2 -gap 2 -linegap 3 \
  -dim 0.3 -dimpad 2 -textcolor white -palette ocean -width 1200 -height 560

# Bold text with a slow scan cursor
go run tileseq.go -text "GO" -scale 3 -bold 1 -scan 4 -palette sunset
```

GIF size is mainly controlled via `-width`/`-height`/`-tile`; for further
optimization run the output through a tool like `gifsicle`.
