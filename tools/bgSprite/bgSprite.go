package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/fs"
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
		log.Fatal("provide stage frames root directory as arg")
	}
	re := regexp.MustCompile(`^background_(\d+)\.png$`)

	as := os.DirFS(".")
	err := fs.WalkDir(as, root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		name := d.Name()
		m := re.FindStringSubmatch(name)
		if m != nil {
			n, err := strconv.Atoi(m[1])
			if err != nil {
				return err
			}
			pairs = append(pairs, pair{index: n, path: p})
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
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
	totalWidth := len(pairs) * width

	numFrames := len(pairs)
	// try to make it square.
	cols := int(math.Ceil(math.Sqrt(float64(numFrames))))
	rows := int(math.Ceil(float64(numFrames) / float64(cols)))

	// texture containing all the frames
	sheet := image.NewRGBA(image.Rect(0, 0, totalWidth, height))
	// paste each frame side-by-side
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
		// destination rectangle for this frame
		offset := i * width
		rect := image.Rect(offset, 0, offset+width, height)

		draw.Draw(sheet, rect, img, image.Point{}, draw.Over)
	}
	// N.png where N is the number of frames
	savePath := filepath.Join(saveDir, strconv.Itoa(len(pairs))+".png")
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
