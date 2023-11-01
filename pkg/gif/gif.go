package gif

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/SergeyCherepiuk/togif/pkg"
	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/internal"
	"github.com/SergeyCherepiuk/togif/pkg/progress"
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

	size, _ := internal.FileSize(&config.Input)
	stdin, ch := progress.NewWriteCloser(stdin)
	buf := bufio.NewWriterSize(stdin, 65536)

	go func() {
		buf.ReadFrom(config.Input)
		stdin.Close()
	}()
	go progress.Display(ch, size)

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s: %s: %v", pkg.CLI_NAME, pkg.CONVERSION_STAGE, err)
	}

	fmt.Fprint(os.Stdout, string(output))
	return nil
}
