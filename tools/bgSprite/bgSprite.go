package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

// frame index and it's path
type pair struct {
	index int
	path  string
}

var pairs []pair

func init() {
	flag.Parse()
}
func main() {
	root := flag.Arg(0)
	if root == "" {
		fmt.Println("provide stage frames directory")
		return
	}
	re := regexp.MustCompile(`^background_(\d+)\.png$`)

	dirs, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dirs {
		p := filepath.Join(root, d.Name())
		m := re.FindStringSubmatch(d.Name())
		if m != nil {
			n, err := strconv.Atoi(m[1])
			if err != nil {
				log.Fatalln("failed to convert to int", err)
			}
			pairs = append(pairs, pair{index: n, path: p})
		}
	}
	if len(pairs) == 0 {
		log.Fatal("no frames found in directory. Make sure there are files with the convention background_1.png and so on.")
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].index < pairs[j].index })
	MakeCombinedTexture(root)
}

// combine all the frames into 1 long sprite sheet.
func MakeCombinedTexture(saveDir string) {
	f, err := os.Open(pairs[0].path)
	if err != nil {
		panic(err)
	}
	frame, _, err := image.Decode(f)
	if err != nil {
		log.Panicln(pairs[0].path, err)
	}

	size := frame.Bounds().Size()
	width, height := size.X, size.Y

	numFrames := len(pairs)
	// try to make it square.
	cols := int(math.Ceil(math.Sqrt(float64(numFrames))))
	rows := int(math.Ceil(float64(numFrames) / float64(cols)))
	totalWidth := cols * width
	totalHeight := rows * height
	// texture containing all the frames
	sheet := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))
	// paste each frame into a set
	for i, p := range pairs {
		fr, err := os.Open(p.path)
		if err != nil {
			log.Panicln("error opening frame:", p.path, err)
		}
		img, _, err := image.Decode(fr)
		fr.Close()
		if err != nil {
			log.Panicln("error decoding frame:", p.path, err)
		}

		col := i % cols
		row := i / cols
		x := col * width
		y := row * height

		rect := image.Rect(x, y, x+width, y+height)
		draw.Draw(sheet, rect, img, image.Point{}, draw.Over)
	}
	// stage_RowsxColumns_numFrames.png
	savePath := filepath.Join(saveDir,
		fmt.Sprintf("stage_%dx%d_%dframes.png", rows, cols, numFrames))

	f, err = os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, sheet); err != nil {
		log.Panicln("error encoding png:", err)
	}
	fmt.Println("saved as", savePath)
}
