package main

import (
	"errors"
	"flag"
	"fmt"
	CONFIG "go_image_processing_cli/config"
	"go_image_processing_cli/logs"
	"os"
)

var environmentType = CONFIG.GetEnvByKey("ENVIRONMENT")

type inputConfig struct {
	inputFilepath      string
	outputFilepath     string
	outputFiletype     string
	targetHeight       int
	targetWidth        int
	compressionQuality int
	croppingAllowed    bool
}

func readInput() (inputConfig, error) {
	// validate that we're getting the correct number of arguments
	if len(os.Args) < 2 {
		return inputConfig{}, errors.New("missing required arguments - please consult the docs")
	}

	// define option flags using the Flag package from the standard library
	// flag name | default value | short description
	outputFilepath := flag.String("output", "./output", "Choose output file path. \nDefaults to `./output`.")
	outputFiletype := flag.String("filetype", "jpg", "Choose output file type. \nDefaults to `jpg`.")
	targetHeight := flag.Int("height", 500, "Choose output image height (in pixels). \nDefaults to `500`.")
	targetWidth := flag.Int("width", 500, "Choose output image width (in pixels). \nDefaults to `500`.")
	compressionQuality := flag.Int("quality", 100, "Choose compression quality (0-100). \nDefaults to `100`.")
	croppingAllowed := flag.Bool("crop", false, "Allow cropping of the image when resizing?Defaults to `false`.")

	// parse all arguments from CLI
	flag.Parse()

	// get the first cli argument after flags have been passed ( this will be the input file )
	inputFilepath := flag.Arg(0)

	fmt.Println(inputFilepath)
	fmt.Println(*outputFilepath)
	fmt.Println(*outputFiletype)
	fmt.Println(*targetHeight)
	fmt.Println(*targetWidth)
	fmt.Println(*compressionQuality)
	fmt.Println(*croppingAllowed)

	// only allow jpgs, pngs, webp as the output file type
	if !(*outputFiletype == "jpg" || *outputFiletype == "jpeg" || *outputFiletype == "png" || *outputFiletype == "webp") {
		return inputConfig{}, errors.New("unfortunately this tool only supports jpg, jpeg, png, webp")
	}

	// only allow values between 1 and 100 for the quality option
	if *compressionQuality < 1 || *compressionQuality > 100 {
		return inputConfig{}, errors.New("please choose a value between 0 - 100 for the compression quality option")
	}

	// make sure the size options have been added
	if *targetHeight < 1 || *targetWidth < 1 {
		return inputConfig{}, errors.New("please specify height and width attributes for the output image (greater than 0)")
	}

	// populate struct with values and return
	return inputConfig{inputFilepath, *outputFilepath, *outputFiletype, *targetHeight, *targetWidth, *compressionQuality, *croppingAllowed}, nil
}

func main() {
	fmt.Printf("Hello World + %s \n", environmentType)

	parsedConfig, err := readInput()
	if err != nil {
		logs.ErrorLogger.Fatalln(err)
	}

	fmt.Println("CLI config retrieved")
	fmt.Println(parsedConfig)
}
