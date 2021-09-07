# Go CLI Image Processing Tool

> Simple Go tool to quickly resize/optimise one or many images via CLI
>
> Distributes a cross-OS executable to be installed locally.
>
> Add command to $PATH to enable cli use from anywhere as well as directly.

**In Progress**

## Commands

- Process an image (tool in $PATH)

```bash
abc [--flags] path/file.jpg
```

- Process an image (local to executable)

```bash
abc.exe [--flags] path/file.jpg
```

- Get a list of available options

```bash
abc --help
```

## Flags (options)

> Flags should be passed in the format --flagname=flagvalue => `--filetype=png`

- output `string`

  - Choose output file path
  - Defaults to `./output`

- filetype `string`

  - Choose output file type
  - Defaults to `jpg`

- height `int`

  - Choose output image height (in pixels)
  - Defaults to `500`

- width `int`

  - Choose output image width (in pixels)
  - Defaults to `500`

- quality `int`

  - Choose compression quality (0-100)
  - Defaults to `100`

- crop `bool`
  - Allow cropping of the image when resizing?
  - Defaults to `false`
