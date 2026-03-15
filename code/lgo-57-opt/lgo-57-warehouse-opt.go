// Optimal Warehouse Placement using Nelder-Mead (Downhill Simplex)
//
// Given N customers at random locations with varying demand (weight),
// find the warehouse location that minimizes total weighted distance.
//
// Outputs: static images (SVG or PNG) and an animated GIF of the optimization.
//
// Usage: go run lgo-57-warehouse-opt.go [-n 8] [-seed 42] [-format svg] [-iter 200] [-prefix warehouse] [-width 800] [-height 600]
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	padX = 60
	padY = 50
)

var (
	flagN           = flag.Int("n", 8, "number of customers")
	flagSeed        = flag.Int64("seed", 42, "random seed")
	flagFormat      = flag.String("format", "svg", "static image format: svg or png")
	flagIter        = flag.Int("iter", 200, "maximum Nelder-Mead iterations")
	flagPrefix      = flag.String("prefix", "warehouse", "output filename prefix")
	flagWidth       = flag.Int("W", 900, "canvas width in pixels")
	flagHeight      = flag.Int("H", 900, "canvas height in pixels")
	flagDelay       = flag.Int("delay", 8, "delay between GIF frames in centiseconds (100 = 1s)")
	flagPlacement   = flag.String("placement", "uniform", "customer placement: uniform, clustered, ring, random, diagonal")
	flagNumClusters = flag.Int("numclusters", 2, "number of cluster")
)

// Point is a location in the unit square [0,1]x[0,1].
type Point struct{ X, Y float64 }

// px maps a logical point to pixel coordinates.
func (p Point) px() (float64, float64) {
	return float64(padX) + p.X*float64(*flagWidth-2*padX),
		float64(padY) + p.Y*float64(*flagHeight-2*padY)
}

func (p Point) pxi() (int, int) {
	x, y := p.px()
	return int(x), int(y)
}

// Customer has a location and a demand weight.
type Customer struct {
	Loc    Point
	Weight float64
	Label  string
}

// --- Objective function ---

// objective computes total weighted Euclidean distance from p to all customers.
func objective(p Point, cs []Customer) float64 {
	var s float64
	for _, c := range cs {
		dx, dy := p.X-c.Loc.X, p.Y-c.Loc.Y
		s += c.Weight * math.Sqrt(dx*dx+dy*dy)
	}
	return s
}

// --- Customer placement strategies ---

func clamp01(v float64) float64 { return max(0.02, min(0.98, v)) }

// placeUniform distributes customers with a minimum distance between them.
func placeUniform(n int, rng *rand.Rand) []Point {
	const minDist = 0.12
	var ps []Point
	for attempts := 0; len(ps) < n && attempts < n*1000; attempts++ {
		p := Point{0.05 + 0.9*rng.Float64(), 0.05 + 0.9*rng.Float64()}
		tooClose := false
		for _, q := range ps {
			dx, dy := p.X-q.X, p.Y-q.Y
			if math.Sqrt(dx*dx+dy*dy) < minDist {
				tooClose = true
				break
			}
		}
		if !tooClose {
			ps = append(ps, p)
		}
	}
	return ps
}

// placeClustered groups customers around well-separated cluster centers.
// Customers are assigned round-robin so each cluster is populated evenly.
func placeClustered(n int, rng *rand.Rand) []Point {
	nc := *flagNumClusters
	// Place centers with generous separation so clusters are visually distinct.
	const minCenterDist = 0.4
	var centers []Point
	for attempts := 0; len(centers) < nc && attempts < nc*1000; attempts++ {
		p := Point{0.15 + 0.7*rng.Float64(), 0.15 + 0.7*rng.Float64()}
		tooClose := false
		for _, c := range centers {
			dx, dy := p.X-c.X, p.Y-c.Y
			if math.Sqrt(dx*dx+dy*dy) < minCenterDist {
				tooClose = true
				break
			}
		}
		if !tooClose {
			centers = append(centers, p)
		}
	}
	// Round-robin assignment ensures each cluster gets customers.
	ps := make([]Point, n)
	for i := range ps {
		c := centers[i%len(centers)]
		ps[i] = Point{
			X: clamp01(c.X + rng.NormFloat64()*0.07),
			Y: clamp01(c.Y + rng.NormFloat64()*0.07),
		}
	}
	return ps
}

// placeRing arranges customers roughly along a circle.
func placeRing(n int, rng *rand.Rand) []Point {
	ps := make([]Point, n)
	for i := range ps {
		angle := 2*math.Pi*float64(i)/float64(n) + rng.NormFloat64()*0.15
		r := 0.35 + rng.NormFloat64()*0.04
		ps[i] = Point{
			X: clamp01(0.5 + r*math.Cos(angle)),
			Y: clamp01(0.5 + r*math.Sin(angle)),
		}
	}
	return ps
}

// placeRandom uses pure uniform random placement, no spacing constraints.
func placeRandom(n int, rng *rand.Rand) []Point {
	ps := make([]Point, n)
	for i := range ps {
		ps[i] = Point{0.05 + 0.9*rng.Float64(), 0.05 + 0.9*rng.Float64()}
	}
	return ps
}

// placeDiagonal scatters customers along the main diagonal with some spread.
func placeDiagonal(n int, rng *rand.Rand) []Point {
	ps := make([]Point, n)
	for i := range ps {
		t := 0.1 + 0.8*rng.Float64()
		offset := rng.NormFloat64() * 0.07
		ps[i] = Point{
			X: clamp01(t + offset),
			Y: clamp01(t - offset),
		}
	}
	return ps
}

// --- Customer generation ---

func generateCustomers(n int, rng *rand.Rand) []Customer {
	var points []Point
	switch *flagPlacement {
	case "clustered":
		points = placeClustered(n, rng)
	case "ring":
		points = placeRing(n, rng)
	case "random":
		points = placeRandom(n, rng)
	case "diagonal":
		points = placeDiagonal(n, rng)
	default:
		points = placeUniform(n, rng)
	}
	labels := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cs := make([]Customer, len(points))
	for i, p := range points {
		cs[i] = Customer{
			Loc:    p,
			Weight: 1 + rng.Float64()*9,
			Label:  string(labels[i%26]),
		}
	}
	return cs
}

// --- Nelder-Mead optimization ---

func nelderMead(cs []Customer, start [3]Point, maxIter int) (Point, [][3]Point) {
	f := func(p Point) float64 { return objective(p, cs) }
	s := start
	v := [3]float64{f(s[0]), f(s[1]), f(s[2])}

	sortSimplex := func() {
		for i := 0; i < 2; i++ {
			for j := i + 1; j < 3; j++ {
				if v[j] < v[i] {
					s[i], s[j] = s[j], s[i]
					v[i], v[j] = v[j], v[i]
				}
			}
		}
	}

	var hist [][3]Point
	for iter := 0; iter < maxIter; iter++ {
		sortSimplex()
		hist = append(hist, s)

		// Convergence: max edge length of simplex
		maxD := 0.0
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 3; j++ {
				dx, dy := s[i].X-s[j].X, s[i].Y-s[j].Y
				if d := math.Sqrt(dx*dx + dy*dy); d > maxD {
					maxD = d
				}
			}
		}
		if maxD < 1e-8 {
			break
		}

		// Centroid of best two vertices
		cx := (s[0].X + s[1].X) / 2
		cy := (s[0].Y + s[1].Y) / 2

		// Reflection (alpha=1)
		r := Point{cx + (cx - s[2].X), cy + (cy - s[2].Y)}
		rv := f(r)

		switch {
		case rv >= v[0] && rv < v[1]:
			s[2], v[2] = r, rv

		case rv < v[0]:
			// Expansion (gamma=2)
			e := Point{cx + 2*(r.X-cx), cy + 2*(r.Y-cy)}
			if ev := f(e); ev < rv {
				s[2], v[2] = e, ev
			} else {
				s[2], v[2] = r, rv
			}

		default:
			// Contraction (rho=0.5)
			ct := Point{cx + 0.5*(s[2].X-cx), cy + 0.5*(s[2].Y-cy)}
			if cv := f(ct); cv < v[2] {
				s[2], v[2] = ct, cv
			} else {
				// Shrink (sigma=0.5)
				for i := 1; i < 3; i++ {
					s[i] = Point{
						s[0].X + 0.5*(s[i].X-s[0].X),
						s[0].Y + 0.5*(s[i].Y-s[0].Y),
					}
					v[i] = f(s[i])
				}
			}
		}
	}

	sortSimplex()
	hist = append(hist, s)
	return s[0], hist
}

// --- Drawing primitives ---

func iabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fillCircle(img draw.Image, cx, cy, r int, col color.Color) {
	b := img.Bounds()
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				px, py := cx+dx, cy+dy
				if px >= b.Min.X && px < b.Max.X && py >= b.Min.Y && py < b.Max.Y {
					img.Set(px, py, col)
				}
			}
		}
	}
}

func drawLine(img draw.Image, x0, y0, x1, y1 int, col color.Color) {
	dx := iabs(x1 - x0)
	dy := iabs(y1 - y0)
	sx, sy := 1, 1
	if x0 > x1 {
		sx = -1
	}
	if y0 > y1 {
		sy = -1
	}
	err := dx - dy
	b := img.Bounds()
	for {
		if x0 >= b.Min.X && x0 < b.Max.X && y0 >= b.Min.Y && y0 < b.Max.Y {
			img.Set(x0, y0, col)
		}
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func thickLine(img draw.Image, x0, y0, x1, y1, thickness int, col color.Color) {
	t := thickness / 2
	for dy := -t; dy <= t; dy++ {
		for dx := -t; dx <= t; dx++ {
			drawLine(img, x0+dx, y0+dy, x1+dx, y1+dy, col)
		}
	}
}

// --- Bitmap font for GIF text overlay ---

const (
	glyphW = 5
	glyphH = 7
)

// 5x7 bitmap glyphs; each row is a uint8, bits 4..0 map to pixels left-to-right.
var glyphs = map[byte][glyphH]uint8{
	'0': {0x0E, 0x11, 0x13, 0x15, 0x19, 0x11, 0x0E},
	'1': {0x04, 0x0C, 0x04, 0x04, 0x04, 0x04, 0x0E},
	'2': {0x0E, 0x11, 0x01, 0x02, 0x04, 0x08, 0x1F},
	'3': {0x0E, 0x11, 0x01, 0x06, 0x01, 0x11, 0x0E},
	'4': {0x02, 0x06, 0x0A, 0x12, 0x1F, 0x02, 0x02},
	'5': {0x1F, 0x10, 0x1E, 0x01, 0x01, 0x11, 0x0E},
	'6': {0x06, 0x08, 0x10, 0x1E, 0x11, 0x11, 0x0E},
	'7': {0x1F, 0x01, 0x02, 0x04, 0x08, 0x08, 0x08},
	'8': {0x0E, 0x11, 0x11, 0x0E, 0x11, 0x11, 0x0E},
	'9': {0x0E, 0x11, 0x11, 0x0F, 0x01, 0x02, 0x0C},
	'.': {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04},
	' ': {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
}

func drawText(img draw.Image, x, y, scale int, s string, c color.Color) {
	b := img.Bounds()
	for _, ch := range []byte(s) {
		glyph, ok := glyphs[ch]
		if !ok {
			x += (glyphW + 1) * scale
			continue
		}
		for row := 0; row < glyphH; row++ {
			for bit := 0; bit < glyphW; bit++ {
				if glyph[row]&(1<<(glyphW-1-bit)) != 0 {
					for sy := 0; sy < scale; sy++ {
						for sx := 0; sx < scale; sx++ {
							px, py := x+bit*scale+sx, y+row*scale+sy
							if px >= b.Min.X && px < b.Max.X && py >= b.Min.Y && py < b.Max.Y {
								img.Set(px, py, c)
							}
						}
					}
				}
			}
		}
		x += (glyphW + 1) * scale
	}
}

// --- SVG output ---

func writeSVG(path string, cs []Customer, opt *Point, simplex *[3]Point, title string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	pr := func(format string, a ...any) { fmt.Fprintf(f, format, a...) }

	pr(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d">
<style>text{font-family:system-ui,sans-serif}</style>
<rect width="%d" height="%d" fill="#fafafa" rx="4"/>
`, *flagWidth, *flagHeight, *flagWidth, *flagHeight)

	// Title
	pr(`<text x="%d" y="30" text-anchor="middle" font-size="15" font-weight="bold" fill="#333">%s</text>
`, *flagWidth/2, title)

	// Grid
	for i := 0; i <= 10; i++ {
		x := float64(padX) + float64(i)/10*float64(*flagWidth-2*padX)
		y := float64(padY) + float64(i)/10*float64(*flagHeight-2*padY)
		pr(`<line x1="%.0f" y1="%d" x2="%.0f" y2="%d" stroke="#eee"/>
`, x, padY, x, *flagHeight-padY)
		pr(`<line x1="%d" y1="%.0f" x2="%d" y2="%.0f" stroke="#eee"/>
`, padX, y, *flagWidth-padX, y)
	}

	// Connection lines from optimal to customers
	if opt != nil {
		ox, oy := opt.px()
		for _, c := range cs {
			cx, cy := c.Loc.px()
			opacity := 0.15 + 0.25*(c.Weight/10)
			sw := 1 + c.Weight/5
			pr(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#e74c3c" stroke-width="%.1f" opacity="%.2f"/>
`, ox, oy, cx, cy, sw, opacity)
		}
	}

	// Simplex triangle
	if simplex != nil {
		ax, ay := simplex[0].px()
		bx, by := simplex[1].px()
		cx, cy := simplex[2].px()
		pr(`<polygon points="%.0f,%.0f %.0f,%.0f %.0f,%.0f" fill="rgba(46,204,113,0.08)" stroke="#2ecc71" stroke-width="2"/>
`, ax, ay, bx, by, cx, cy)
	}

	// Customer circles with labels
	for _, c := range cs {
		cx, cy := c.Loc.px()
		r := 4 + c.Weight*1.6
		pr(`<circle cx="%.1f" cy="%.1f" r="%.1f" fill="#4a90d9" stroke="#2c5aa0" stroke-width="1.5"/>
`, cx, cy, r)
		pr(`<text x="%.1f" y="%.1f" text-anchor="middle" fill="#333" font-weight="bold" font-size="11">%s</text>
`, cx, cy-r-5, c.Label)
		pr(`<text x="%.1f" y="%.1f" text-anchor="middle" fill="#999" font-size="9">w=%.1f</text>
`, cx, cy+r+12, c.Weight)
	}

	// Optimal location marker
	if opt != nil {
		ox, oy := opt.px()
		pr(`<circle cx="%.1f" cy="%.1f" r="14" fill="none" stroke="#e74c3c" stroke-width="1.5" opacity="0.4"/>
`, ox, oy)
		pr(`<circle cx="%.1f" cy="%.1f" r="7" fill="#e74c3c" stroke="#c0392b" stroke-width="2"/>
`, ox, oy)
		pr(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#fff" stroke-width="2"/>
`, ox-3.5, oy, ox+3.5, oy)
		pr(`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="#fff" stroke-width="2"/>
`, ox, oy-3.5, ox, oy+3.5)
	}

	pr("</svg>\n")
	return nil
}

// --- PNG output ---

func writePNG(path string, cs []Customer, opt *Point) error {
	img := image.NewRGBA(image.Rect(0, 0, *flagWidth, *flagHeight))
	draw.Draw(img, img.Bounds(), image.NewUniform(color.RGBA{250, 250, 250, 255}), image.Point{}, draw.Src)

	var (
		blue    = color.RGBA{74, 144, 217, 255}
		red     = color.RGBA{231, 76, 60, 255}
		connCol = color.RGBA{231, 76, 60, 60}
	)

	if opt != nil {
		ox, oy := opt.pxi()
		for _, c := range cs {
			cx, cy := c.Loc.pxi()
			drawLine(img, ox, oy, cx, cy, connCol)
		}
	}

	for _, c := range cs {
		cx, cy := c.Loc.pxi()
		r := int(4 + c.Weight*1.6)
		fillCircle(img, cx, cy, r, blue)
	}

	if opt != nil {
		ox, oy := opt.pxi()
		fillCircle(img, ox, oy, 7, red)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

// --- GIF output ---

var gifPalette = color.Palette{
	color.RGBA{250, 250, 250, 255}, // 0: background
	color.RGBA{238, 238, 238, 255}, // 1: grid
	color.RGBA{74, 144, 217, 255},  // 2: customer
	color.RGBA{44, 90, 160, 255},   // 3: customer border
	color.RGBA{231, 76, 60, 255},   // 4: optimal
	color.RGBA{46, 204, 113, 255},  // 5: simplex
	color.RGBA{255, 180, 50, 255},  // 6: trail
	color.RGBA{210, 210, 210, 255}, // 7: connection
	color.Black,                    // 8
	color.White,                    // 9
	color.RGBA{180, 180, 180, 255}, // 10: dim text
}

func newGIFFrame() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, *flagWidth, *flagHeight), gifPalette)
	for i := range img.Pix {
		img.Pix[i] = 0 // background color index
	}
	return img
}

func drawGIFScene(cs []Customer, tri *[3]Point, best *Point, trail []Point, val float64) *image.Paletted {
	var (
		img = newGIFFrame()

		green  = gifPalette[5]
		blue   = gifPalette[2]
		red    = gifPalette[4]
		orange = gifPalette[6]
		conn   = gifPalette[7]
		dim    = gifPalette[10]
	)

	// Connection lines from best to customers
	if best != nil {
		bx, by := best.pxi()
		for _, c := range cs {
			cx, cy := c.Loc.pxi()
			drawLine(img, bx, by, cx, cy, conn)
		}
	}

	// Simplex triangle
	if tri != nil {
		for i := 0; i < 3; i++ {
			j := (i + 1) % 3
			ax, ay := tri[i].pxi()
			bx, by := tri[j].pxi()
			thickLine(img, ax, ay, bx, by, 2, green)
		}
		for i := 0; i < 3; i++ {
			px, py := tri[i].pxi()
			fillCircle(img, px, py, 3, green)
		}
	}

	// Trail of previous best points
	for _, p := range trail {
		px, py := p.pxi()
		fillCircle(img, px, py, 2, orange)
	}

	// Customers
	for _, c := range cs {
		cx, cy := c.Loc.pxi()
		r := int(4 + c.Weight*1.5)
		fillCircle(img, cx, cy, r, blue)
	}

	// Best point
	if best != nil {
		bx, by := best.pxi()
		fillCircle(img, bx, by, 6, red)
	}

	// Objective value in bottom margin
	drawText(img, padX, *flagHeight-padY+8, 2, fmt.Sprintf("%.8f", val), dim)

	return img
}

func writeGIF(path string, cs []Customer, hist [][3]Point, opt Point) error {
	var (
		targetFrames = 60
		step         = max(1, len(hist)/targetFrames)

		frames []*image.Paletted
		delays []int
		trail  []Point
	)

	for i := 0; i < len(hist); i += step {
		tri := hist[i]
		best := tri[0]
		val := objective(best, cs)
		trail = append(trail, best)
		frame := drawGIFScene(cs, &tri, &best, trail, val)
		frames = append(frames, frame)
		delays = append(delays, *flagDelay)
	}

	// Final frame: hold on the result
	lastTri := hist[len(hist)-1]
	trail = append(trail, opt)
	final := drawGIFScene(cs, &lastTri, &opt, trail, objective(opt, cs))
	frames = append(frames, final)
	delays = append(delays, 300) // 3s hold

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return gif.EncodeAll(f, &gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: 0, // loop forever
	})
}

// --- TSV output ---

func writeTSV(path string, cs []Customer, opt Point) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintln(f, "label\tx\ty\tweight")
	for _, c := range cs {
		fmt.Fprintf(f, "%s\t%.6f\t%.6f\t%.2f\n", c.Label, c.Loc.X, c.Loc.Y, c.Weight)
	}
	fmt.Fprintf(f, "OPT\t%.6f\t%.6f\t\n", opt.X, opt.Y)
	return nil
}

// --- Main ---

func log(format string, a ...any) { fmt.Fprintf(os.Stderr, format, a...) }

func main() {
	flag.Parse()

	rng := rand.New(rand.NewSource(*flagSeed))
	cs := generateCustomers(*flagN, rng)

	log("generated %d customers:\n", len(cs))
	for _, c := range cs {
		log("  %s: (%.3f, %.3f) weight=%.1f\n", c.Label, c.Loc.X, c.Loc.Y, c.Weight)
	}

	// Initial simplex: large triangle covering the space
	start := [3]Point{{0.1, 0.1}, {0.9, 0.15}, {0.5, 0.9}}
	opt, hist := nelderMead(cs, start, *flagIter)

	log("\noptimal warehouse location: (%.4f, %.4f)\n", opt.X, opt.Y)
	log("total weighted distance:    %.4f\n", objective(opt, cs))
	log("nelder-mead iterations:     %d\n\n", len(hist)-1)

	ext := *flagFormat
	name := func(suffix string) string { return *flagPrefix + "-" + suffix }

	// 1. Customers only
	fn := name("customers." + ext)
	var err error
	if ext == "svg" {
		err = writeSVG(fn, cs, nil, nil, "Warehouse Placement — Customers")
	} else {
		err = writePNG(fn, cs, nil)
	}
	if err != nil {
		log("error writing %s: %v\n", fn, err)
		os.Exit(1)
	}
	log("wrote %s\n", fn)

	// 2. Result with optimal location
	fn = name("result." + ext)
	if ext == "svg" {
		err = writeSVG(fn, cs, &opt, nil, "Warehouse Placement — Optimal Location")
	} else {
		err = writePNG(fn, cs, &opt)
	}
	if err != nil {
		log("error writing %s: %v\n", fn, err)
		os.Exit(1)
	}
	log("wrote %s\n", fn)

	// 3. Animated optimization
	fn = name("optimization.gif")
	if err = writeGIF(fn, cs, hist, opt); err != nil {
		log("error writing %s: %v\n", fn, err)
		os.Exit(1)
	}
	log("wrote %s\n", fn)

	// 4. Customer data as TSV
	fn = name("data.tsv")
	if err = writeTSV(fn, cs, opt); err != nil {
		log("error writing %s: %v\n", fn, err)
		os.Exit(1)
	}
	log("wrote %s\n", fn)
}
