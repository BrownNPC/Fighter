package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)
import "github.com/nfnt/resize"

// frame index and it's path
type pair struct {
	index int
	path  string
}

var pairs []pair
var Flag_Mode *string
var Flag_OutputDir *string
var Flag_Downscale *bool

// prefix for the filename before the row and column info
var Flag_Prefix *string

type Mode struct {
	// name that would be used for the -mode flag
	name      string
	FileRegex string
	// what is the submatch for the incrementing number in this regex
	// eg 1.png would have the submatch 1
	NumberSubmatch int
	Prefix         string
}

var Modes = []Mode{
	{"bg", `^background_(\d+)\.png$`, 1, "stage"},
	{"sprite", `^(\d+)\.png$`, 1, "sprite"},
}

var SelectedMode Mode
var TotalModes string

func init() {
	// used for help
	for _, mode := range Modes {
		TotalModes += " " + mode.name + " "
	}
	TotalModes = fmt.Sprintf("{%s}", TotalModes)
	// parse -mode flag
	Flag_Mode = flag.String("mode", "bg", "-mode sprite")
	Flag_OutputDir = flag.String("o", "placeholder", "-o /path/to/output/directory")
	Flag_Prefix = flag.String("prefix", "placeholder", "-prefix walk")
	Flag_Downscale = flag.Bool("downscale", false, "will downscale the image to half res for anti aliasing")
	flag.Parse()
	// select the mode from the flags
	for _, mode := range Modes {
		if mode.name == *Flag_Mode {
			SelectedMode = mode
			break
		}
	}
	if *Flag_Prefix != "placeholder" {
		SelectedMode.Prefix = *Flag_Prefix
	}
	if SelectedMode.name == "" {
		fmt.Println("invalid mode, available modes are", TotalModes)
		os.Exit(0)
	}
}

func main() {
	root := flag.Arg(0)
	if root == "" {
		fmt.Println("provide frames directory as argument")
		return
	}
	fmt.Println("selected mode", SelectedMode.name, "out of", TotalModes)
	re := regexp.MustCompile(SelectedMode.FileRegex)

	dirs, err := os.ReadDir(root)
	if err != nil {
		Fatal(err)
	}
	for _, d := range dirs {
		p := filepath.Join(root, d.Name())
		m := re.FindStringSubmatch(d.Name())
		if m != nil {
			n, err := strconv.Atoi(m[SelectedMode.NumberSubmatch])
			if err != nil {
				Fatal("failed to convert to int", err)
			}
			pairs = append(pairs, pair{index: n, path: p})
		}
	}
	if len(pairs) == 0 {
		Fatal("no frames found in directory. Make sure there are files with the regex " + SelectedMode.FileRegex)
	}
	// sort by filename
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].index < pairs[j].index })
	if *Flag_OutputDir != "placeholder" {
		MakeCombinedTexture(*Flag_OutputDir)
	} else {
		MakeCombinedTexture(root)
	}
}

// combine all the frames into 1 big square sprite sheet.
func MakeCombinedTexture(saveDir string) {
	f, err := os.Open(pairs[0].path)
	if err != nil {
		panic(err)
	}
	frame, _, err := image.Decode(f)
	f.Close()
	if err != nil {
		Fatal(pairs[0].path, err)
	}

	size := frame.Bounds().Size()
	width, height := size.X, size.Y

	if *Flag_Downscale {
		// target size when downscaling each sprite by 2
		width = int(math.Ceil(float64(width) / 2.0))
		height = int(math.Ceil(float64(height) / 2.0))
	}

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
			Fatal("error opening frame:", p.path, err)
		}
		img, _, err := image.Decode(fr)
		fr.Close()
		if err != nil {
			Fatal("error decoding frame:", p.path, err)
		}

		// if downscale flag is set, resize each sprite individually using Lanczos3
		var sprite image.Image = img
		if *Flag_Downscale {
			origBounds := img.Bounds().Size()
			newW := uint(int(math.Ceil(float64(origBounds.X) / 2.0)))
			newH := uint(int(math.Ceil(float64(origBounds.Y) / 2.0)))
			sprite = resize.Resize(newW, newH, img, resize.Lanczos3)
		}

		col := i % cols
		row := i / cols
		x := col * width
		y := row * height

		rect := image.Rect(x, y, x+width, y+height)
		draw.Draw(sheet, rect, sprite, image.Point{}, draw.Over)
	}
	// stage_RowsxColumns_numFrames.png
	savePath := filepath.Join(saveDir,
		fmt.Sprintf(SelectedMode.Prefix+"_%dx%d_%dframes.png", rows, cols, numFrames))

	fout, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	if err := png.Encode(fout, sheet); err != nil {
		Fatal("error encoding png:", err)
	}
	fmt.Println("saved as", savePath)
}
func Fatal(err ...any) {
	fmt.Println(err...)
	os.Exit(0)
}
