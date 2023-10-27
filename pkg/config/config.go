package config

import (
	"golang.org/x/exp/slices"
)

var keywords []string

func IsKeyword(word string) bool {
	return slices.Contains(keywords, word)
}

type Config struct {
	Frames    uint64 `short:"f" long:"frames"`  // frames per second
	Quality   uint8  `short:"q" long:"quality"` // quality, [0%, 100%]
	StartMs   uint64 `short:"s" long:"start"`   // start point (in milliseconds)
	EndMs     uint64 `short:"e" long:"end"`     // end point (in milliseconds)
	IFilename string `short:"i" long:"input"`   // input file/filepath (source)
	OFilename string `short:"o" long:"output"`  // output file/filepath (destination)
}

func (c *Config) SetShort(short string, value any) error {
	return nil
}

func (c *Config) SetLong(long string, value any) error {
	return nil
}

var DefaultConfig = Config{
	Frames:  10,
	Quality: 80,
}

func init() {
	// Generate keywords from "long" tags of Config struct using reflection
}
