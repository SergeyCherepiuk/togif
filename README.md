# In action

Converting WebM video file [[source](https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.webm)] to a GIF image and doubling the speed:

```bash
$ go run cmd/cli/main.go -o example/flower.gif -f 30 -s 2 -v example/flower.webm
Progress: [==================================================] 100.00% | ETA: 0s
```

<img src="example/flower.gif" style="width: 400px" />

# Options

You can get brief descriptions of all available options by running the program with `-h` or `--help` flag:

```bash
$ go run cmd/cli/main.go -h
```

Output on 02.11.2023:
```
A tool for converting videos into GIF images

Usage: togif [OPTIONS] [FILE]
* If input file is omitted stdin will be used

List of available options:
    -o, --output    Path to the output file (destination), if omitted stdout will be used
    -f, --frames    Sets the frames-rate of the resulting GIF image
    -s, --speed     Speeds up (s > 1.0) or slows down (s < 1.0) the output GIF
    -v, --verbose   TODO
    -h, --help      Provide information on existing options
```

# Installation and Usage

There are two main way of running the program described in the following subsections.

## Using Docker

The docker version of the tool **is not able** to work with files of the host machine. So instead of providing paths to input and output files you can pipe the contents of the file to a docker container and get the output from `stdout`.

Although the docker version of a program is limited in file operations, we encourage you to stick with it, since it doesn't require additional dependencies to be installed.

Build the docker image from `Dockerfile` by executing the command below (feel free to name it differently):

```bash
$ docker build -t togif:alpine
```

Run the container streaming in and out the contents of files (do not forget to change the name of image respectively it you did so running previous command):

```bash
$ cat ./origin.mp4 | docker run -i togif:alpine [OPTIONS] > ./output.gif
```

## On host

### Running

Running the program on the host machine requires some dependencies:

* [Golang compiler](https://go.dev/doc/install)
* [FFmpeg set of libraries](https://ffmpeg.org/download.html)

After all necessary software have been installed, you can run the tool by giving it filepaths as arguments:

```bash
$ go run cmd/cli/main.go [OPTIONS] -o ./origin.mp4 ./output.gif
```

or by piping files in and out:

```bash
$ cat ./origin.mp4 | go run cmd/cli/main.go [OPTIONS] > ./output.gif
```

### Building

Despite the fact that Golang compilation time is a <u>pure pleasure</u> building the program each time you want to run it is inconvenient.

Build the program once using `go build`:

```bash
$ go build -o cmd/cli/togif cmd/cli/main.go
```

Additionally you might want to copy it to `/usr/bin` directory, in order to be able to use it anywhere across you system:

```bash
$ sudo cp cmd/cli/togif /usr/bin
```
