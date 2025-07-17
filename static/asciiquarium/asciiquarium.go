package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
)

const Version = "1.1-go"

type Depth struct {
	GUIText    int
	GUI        int
	Shark      int
	FishStart  int
	FishEnd    int
	Seaweed    int
	Castle     int
	WaterLine3 int
	WaterGap3  int
	WaterLine2 int
	WaterGap2  int
	WaterLine1 int
	WaterGap1  int
	WaterLine0 int
	WaterGap0  int
}

var depths = Depth{
	GUIText:    0,
	GUI:        1,
	Shark:      2,
	FishStart:  3,
	FishEnd:    20,
	Seaweed:    21,
	Castle:     22,
	WaterLine3: 2,
	WaterGap3:  3,
	WaterLine2: 4,
	WaterGap2:  5,
	WaterLine1: 6,
	WaterGap1:  7,
	WaterLine0: 8,
	WaterGap0:  9,
}

type Position struct {
	X, Y, Z int
}

type Velocity struct {
	DX, DY, DZ float64
	Speed      float64
}

type Entity struct {
	ID           string
	Type         string
	Shape        []string
	ColorMask    []string
	Pos          Position
	Vel          Velocity
	Width        int
	Height       int
	Frame        int
	MaxFrames    int
	DieTime      time.Time
	DieOffscreen bool
	Physical     bool
	DefaultColor tcell.Color
	Transparent  rune
	DieFrame     int
	CurrentFrame int
}

type Aquarium struct {
	screen      tcell.Screen
	width       int
	height      int
	entities    []*Entity
	paused      bool
	startTime   time.Time
	entityID    int
	newFish     bool
	newMonster  bool
	classicMode bool
}

func NewAquarium(classicMode bool) (*Aquarium, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := screen.Init(); err != nil {
		return nil, err
	}

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	screen.Clear()

	width, height := screen.Size()

	return &Aquarium{
		screen:      screen,
		width:       width,
		height:      height,
		entities:    make([]*Entity, 0),
		startTime:   time.Now(),
		newFish:     !classicMode,
		newMonster:  !classicMode,
		classicMode: classicMode,
	}, nil
}

func (a *Aquarium) nextEntityID() string {
	a.entityID++
	return fmt.Sprintf("entity_%d", a.entityID)
}

func (a *Aquarium) AddEntity(e *Entity) {
	if e.ID == "" {
		e.ID = a.nextEntityID()
	}
	a.entities = append(a.entities, e)
}

func (a *Aquarium) RemoveEntity(id string) {
	for i, e := range a.entities {
		if e.ID == id {
			a.entities = append(a.entities[:i], a.entities[i+1:]...)
			break
		}
	}
}

func (a *Aquarium) Run() {
	defer a.screen.Fini()

	// Initialize environment
	a.addEnvironment()
	a.addCastle()
	a.addAllSeaweed()
	a.addAllFish()
	a.addRandomObject()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !a.paused {
				a.animate()
			}
			a.draw()

		default:
			// Handle input
			if a.screen.HasPendingEvent() {
				ev := a.screen.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					switch ev.Key() {
					case tcell.KeyEscape, tcell.KeyCtrlC:
						return
					case tcell.KeyRune:
						switch ev.Rune() {
						case 'q', 'Q':
							return
						case 'r', 'R':
							a.restart()
						case 'p', 'P':
							a.paused = !a.paused
						}
					}
				case *tcell.EventResize:
					a.width, a.height = a.screen.Size()
					a.restart()
				}
			}
		}
	}
}

func (a *Aquarium) restart() {
	a.entities = a.entities[:0]
	a.screen.Clear()
	a.addEnvironment()
	a.addCastle()
	a.addAllSeaweed()
	a.addAllFish()
	a.addRandomObject()
}

func (a *Aquarium) animate() {
	toRemove := make([]string, 0)

	for _, e := range a.entities {
		// Update position
		e.Pos.X += int(e.Vel.DX)
		e.Pos.Y += int(e.Vel.DY)

		// Update animation frame
		if e.MaxFrames > 1 {
			e.Frame = (e.Frame + 1) % e.MaxFrames
		}

		// Check if entity should die
		if e.DieOffscreen && (e.Pos.X < -e.Width || e.Pos.X > a.width+e.Width) {
			toRemove = append(toRemove, e.ID)
			continue
		}

		if !e.DieTime.IsZero() && time.Now().After(e.DieTime) {
			toRemove = append(toRemove, e.ID)
			continue
		}

		if e.DieFrame > 0 {
			e.CurrentFrame++
			if e.CurrentFrame >= e.DieFrame {
				toRemove = append(toRemove, e.ID)
				continue
			}
		}

		// Special behaviors
		switch e.Type {
		case "fish":
			if rand.Intn(100) > 97 {
				a.addBubble(e)
			}
		case "seaweed":
			// Sway animation
			phase := float64(time.Now().UnixNano()/1000000) * e.Vel.Speed
			e.Vel.DX = math.Sin(phase) * 0.5
		}
	}

	// Remove dead entities
	for _, id := range toRemove {
		a.RemoveEntity(id)
	}

	// Randomly add new objects
	if rand.Intn(1000) < 2 {
		a.addRandomObject()
	}
}

func (a *Aquarium) draw() {
	a.screen.Clear()

	// Sort entities by depth
	for z := 30; z >= 0; z-- {
		for _, e := range a.entities {
			if e.Pos.Z == z {
				a.drawEntity(e)
			}
		}
	}

	a.screen.Show()
}

func (a *Aquarium) drawEntity(e *Entity) {
	frameIndex := 0
	if e.MaxFrames > 1 {
		frameIndex = e.Frame
	}

	shape := e.Shape[frameIndex]
	colorMask := ""
	if len(e.ColorMask) > frameIndex {
		colorMask = e.ColorMask[frameIndex]
	}

	lines := strings.Split(shape, "\n")
	colorLines := strings.Split(colorMask, "\n")

	for y, line := range lines {
		if e.Pos.Y+y < 0 || e.Pos.Y+y >= a.height {
			continue
		}

		colorLine := ""
		if y < len(colorLines) {
			colorLine = colorLines[y]
		}

		for x, r := range line {
			if e.Pos.X+x < 0 || e.Pos.X+x >= a.width {
				continue
			}

			if r == e.Transparent {
				continue
			}

			color := e.DefaultColor
			if x < len(colorLine) {
				color = a.getColor(rune(colorLine[x]))
			}

			style := tcell.StyleDefault.Foreground(color)
			a.screen.SetContent(e.Pos.X+x, e.Pos.Y+y, r, nil, style)
		}
	}
}

func (a *Aquarium) getColor(colorCode rune) tcell.Color {
	switch colorCode {
	case '1':
		return tcell.ColorBlue
	case '2':
		return tcell.ColorGreen
	case '3':
		return tcell.ColorCyan
	case '4':
		return tcell.ColorRed
	case '5':
		return tcell.ColorMagenta
	case '6':
		return tcell.ColorYellow
	case '7':
		return tcell.ColorWhite
	case 'c', 'C':
		return tcell.ColorTeal
	case 'r', 'R':
		return tcell.ColorMaroon
	case 'y', 'Y':
		return tcell.ColorOlive
	case 'b', 'B':
		return tcell.ColorNavy
	case 'g', 'G':
		return tcell.ColorGreen
	case 'm', 'M':
		return tcell.ColorPurple
	case 'W':
		return tcell.ColorWhite
	default:
		return tcell.ColorWhite
	}
}

func (a *Aquarium) addEnvironment() {
	waterLineSegments := []string{
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~",
		"^^^^ ^^^  ^^^   ^^^    ^^^^      ",
		"^^^^      ^^^^     ^^^    ^^     ",
		"^^      ^^^^      ^^^    ^^^^^^  ",
	}

	segmentSize := len(waterLineSegments[0])
	segmentRepeat := a.width/segmentSize + 1

	for i, segment := range waterLineSegments {
		repeatedSegment := strings.Repeat(segment, segmentRepeat)
		entity := &Entity{
			Type:         "waterline",
			Shape:        []string{repeatedSegment},
			Pos:          Position{X: 0, Y: i + 5, Z: depths.WaterLine0 - i},
			DefaultColor: tcell.ColorCyan,
			Physical:     true,
		}
		a.AddEntity(entity)
	}
}

func (a *Aquarium) addCastle() {
	castleImage := `               T~~
               |
              /^\
             /   \
 _   _   _  /     \  _   _   _
[ ]_[ ]_[ ]/ _   _ \[ ]_[ ]_[ ]
|_=__-_ =_|_[ ]_[ ]_|_=-___-__|
 | _- =  | =_ = _    |= _=   |
 |= -[]  |- = _ =    |_-=_[] |
 | =_    |= - ___    | =_ =  |
 |=  []- |-  /| |\   |=_ =[] |
 |- =_   | =| | | |  |- = -  |
 |_______|__|_|_|_|__|_______|`

	lines := strings.Split(castleImage, "\n")
	entity := &Entity{
		Type:         "castle",
		Shape:        []string{castleImage},
		Pos:          Position{X: a.width - 32, Y: a.height - len(lines), Z: depths.Castle},
		DefaultColor: tcell.ColorYellow,
		Width:        32,
		Height:       len(lines),
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addAllSeaweed() {
	seaweedCount := a.width / 15
	for i := 0; i < seaweedCount; i++ {
		a.addSeaweed()
	}
}

func (a *Aquarium) addSeaweed() {
	height := rand.Intn(4) + 3
	leftSide := make([]string, height)
	rightSide := make([]string, height)

	for i := 0; i < height; i++ {
		if i%2 == 0 {
			leftSide[i] = "("
			rightSide[i] = " )"
		} else {
			leftSide[i] = " )"
			rightSide[i] = "("
		}
	}

	shape := ""
	for i := 0; i < height; i++ {
		if i%2 == 0 {
			shape += "(\n"
		} else {
			shape += " )\n"
		}
	}

	x := rand.Intn(a.width-2) + 1
	y := a.height - height
	speed := rand.Float64()*0.05 + 0.25

	entity := &Entity{
		Type:         "seaweed",
		Shape:        []string{shape},
		Pos:          Position{X: x, Y: y, Z: depths.Seaweed},
		Vel:          Velocity{Speed: speed},
		DefaultColor: tcell.ColorGreen,
		DieTime:      time.Now().Add(time.Duration(rand.Intn(4*60)+8*60) * time.Second),
		Width:        2,
		Height:       height,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addBubble(fish *Entity) {
	bubblePos := fish.Pos
	if fish.Vel.DX > 0 {
		bubblePos.X += fish.Width
	}
	bubblePos.Y += fish.Height / 2
	bubblePos.Z--

	entity := &Entity{
		Type:         "bubble",
		Shape:        []string{".", "o", "O", "O", "O"},
		Pos:          bubblePos,
		Vel:          Velocity{DX: 0, DY: -1, Speed: 0.1},
		MaxFrames:    5,
		DieOffscreen: true,
		DefaultColor: tcell.ColorCyan,
		Width:        1,
		Height:       1,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addAllFish() {
	screenSize := (a.height - 9) * a.width
	fishCount := screenSize / 350
	for i := 0; i < fishCount; i++ {
		a.addFish()
	}
}

func (a *Aquarium) addFish() {
	fishShapes := []string{
		`   \
  / \
>=_('>
  \_/
   /`,
		`  /
 / \
<')_=<
 \_/
  \`,
		`     ,
     }\
\  .'  '\
}}<   ( 6>
/  ',  .'
     }/
     '`,
		`    ,
   /\{
 /'  '.  /
<6 )   >\{\{
 '.  ,'  \
   \{
    '`,
	}

	fishColors := []string{
		`   2
  1 1
663745
  111
   3`,
		`  2
 111
547366
 111
  3`,
		`     2
     22
6  11  11
661   7 45
6  11  11
     33
     3`,
		`    2
   22
 11  11  6
54 7   166
 11  11  6
   33
    3`,
	}

	fishIndex := rand.Intn(len(fishShapes))
	speed := rand.Float64()*2 + 0.25
	depth := rand.Intn(depths.FishEnd-depths.FishStart) + depths.FishStart

	if fishIndex%2 == 1 {
		speed *= -1
	}

	maxHeight := 9
	minHeight := a.height - 5
	y := rand.Intn(minHeight-maxHeight) + maxHeight

	x := 1
	if speed < 0 {
		x = a.width - 2
	}

	entity := &Entity{
		Type:         "fish",
		Shape:        []string{fishShapes[fishIndex]},
		ColorMask:    []string{fishColors[fishIndex]},
		Pos:          Position{X: x, Y: y, Z: depth},
		Vel:          Velocity{DX: speed, DY: 0},
		DefaultColor: tcell.ColorWhite,
		DieOffscreen: true,
		Physical:     true,
		Width:        7,
		Height:       5,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addRandomObject() {
	objects := []func(){
		a.addShark,
		a.addShip,
		a.addWhale,
	}

	if a.newMonster {
		objects = append(objects, a.addMonster)
	}

	objectFunc := objects[rand.Intn(len(objects))]
	objectFunc()
}

func (a *Aquarium) addShark() {
	sharkShape := `                              __
                             ( '\
  ,??????????????????????????)   '\
;' '.????????????????????????(     '\__
 ;   '.?????????????__..---''          '~~~~-._
  '.   '.____...--''                       (b  '--._
    >                     _.-'      .((      ._     )
  .'.-'--...__         .-'     -.___.....-(|/|/|/|/'
 ;.'?????????'. ...----'.___.',,,_______......---'
 '???????????'-'`

	dir := rand.Intn(2)
	x := -53
	y := rand.Intn(a.height-(10+9)) + 9
	speed := 2.0

	if dir == 1 {
		speed *= -1
		x = a.width - 2
	}

	entity := &Entity{
		Type:         "shark",
		Shape:        []string{sharkShape},
		Pos:          Position{X: x, Y: y, Z: depths.Shark},
		Vel:          Velocity{DX: speed, DY: 0},
		DefaultColor: tcell.ColorWhite,
		DieOffscreen: true,
		Width:        53,
		Height:       10,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addShip() {
	shipShape := `     |    |    |
    )_)  )_)  )_)
   )___))___))___)\
  )____)____)_____)\\\
_____|____|____|____\\\\\__
\                   /`

	dir := rand.Intn(2)
	x := -24
	speed := 1.0

	if dir == 1 {
		speed *= -1
		x = a.width - 2
	}

	entity := &Entity{
		Type:         "ship",
		Shape:        []string{shipShape},
		Pos:          Position{X: x, Y: 0, Z: depths.WaterGap1},
		Vel:          Velocity{DX: speed, DY: 0},
		DefaultColor: tcell.ColorYellow,
		DieOffscreen: true,
		Width:        24,
		Height:       6,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addWhale() {
	whaleShape := `        .-----:
      .'       '.
,????/       (o) \
\'.]/          ,__)`

	dir := rand.Intn(2)
	x := -18
	speed := 1.0

	if dir == 1 {
		speed *= -1
		x = a.width - 2
	}

	entity := &Entity{
		Type:         "whale",
		Shape:        []string{whaleShape},
		Pos:          Position{X: x, Y: 0, Z: depths.WaterGap2},
		Vel:          Velocity{DX: speed, DY: 0},
		DefaultColor: tcell.ColorBlue,
		DieOffscreen: true,
		Width:        18,
		Height:       4,
	}
	a.AddEntity(entity)
}

func (a *Aquarium) addMonster() {
	monsterShape := `                                                          ____
            __??????????????????????????????????????????/   o  \
          /    \????????_?????????????????????_???????/     ____ >
  _??????|  __  |?????/   \????????_????????/   \????|     |
 | \?????|  ||  |????|     |?????/   \?????|     |???|     |`

	dir := rand.Intn(2)
	x := -64
	speed := 2.0

	if dir == 1 {
		speed *= -1
		x = a.width - 2
	}

	entity := &Entity{
		Type:         "monster",
		Shape:        []string{monsterShape},
		Pos:          Position{X: x, Y: 2, Z: depths.WaterGap2},
		Vel:          Velocity{DX: speed, DY: 0},
		DefaultColor: tcell.ColorGreen,
		DieOffscreen: true,
		Width:        64,
		Height:       5,
	}
	a.AddEntity(entity)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	classicMode := false
	if len(os.Args) > 1 && os.Args[1] == "-c" {
		classicMode = true
	}

	aquarium, err := NewAquarium(classicMode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing aquarium: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Asciiquarium %s - Press 'q' to quit, 'p' to pause, 'r' to restart\n", Version)
	time.Sleep(2 * time.Second)

	aquarium.Run()
}
