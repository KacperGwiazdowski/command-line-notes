# Command Line Notes

A simple command-line tool for quickly appending timestamped notes to daily markdown files.

## Features

- Stores notes in `~/cln/` as daily `.md` files
- Each note is prepended with the current time
- Automatically creates the notes directory and daily file if they do not exist

## Usage

```sh
nn "Your note here"
```

This will append a line like:

```
15:04:05 - Your note here
```

to a file named `DD-MM-YYYY.md` in your `~/cln/` directory.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/command-line-notes.git
    cd command-line-notes
    ```

2. Build the binary:

    ```sh
    go build -o nn ./cmd/nn
    ```

3. (Optional) Install the binary:

    ```sh
    go install ./cmd/nn
    
    ```

## Requirements

- Go 1.22 or newer

## License

MIT License. See [LICENSE](LICENSE) for details.