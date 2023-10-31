package gif

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/SergeyCherepiuk/togif/pkg"
	"github.com/SergeyCherepiuk/togif/pkg/config"
)

func Convert(config config.Config) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i", "-",
		"-f", "gif",
		"-vf", fmt.Sprintf("fps=%d", config.Frames),
		config.OutputPath,
	) // TODO: Make use of config.Quality option

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("%s: %s: %v", pkg.CLI_NAME, pkg.CONVERSION_STAGE, err)
	}

	go func() {
		buf := bufio.NewWriterSize(stdin, 65536)
		buf.ReadFrom(config.Input) // TODO: Track progress using return value
		stdin.Close()
	}()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %s: %v", pkg.CLI_NAME, pkg.CONVERSION_STAGE, err)
	}

	return nil
}
