# rcrd - Shell Session Recorder

rcrd is a tool that records your shell sessions, capturing both input and output. It provides a seamless way to log your terminal activities with minimal disruption to your workflow.

## Features

- Records shell input and output to timestamped log files
- Custom naming for log files
- Visual indicator (red dot and prompt) when recording is active
- Custom ":exit" command to end recording session

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/rcrd.git
   cd rcrd
   ```
3. Build the project using the Makefile:
   ```
   make build
   ```

## Usage

To start a recording session:

```
./rcrd
```

To start a recording session with a custom name:

```
./rcrd -n mysession
```
or
```
./rcrd --name mysession
```

This will create a log file named "mysession.txt" in the ~/.rcrd directory.

While in a recording session:
- Your prompt will turn red to indicate active recording.
- A red dot will appear at the end of your prompt.
- Use ":exit" to end the recording session (instead of the standard "exit" command).

## Makefile Commands

The project includes a Makefile for common tasks:

- `make build`: Compiles the rcrd program.
- `make install`: Builds rcrd and installs it to /usr/local/bin (may require sudo).
- `make clean`: Removes the compiled binary and cleans the Go build cache.
- `make run`: Builds and then runs the rcrd program.

Example:
```
make build
make install
```

## Configuration

rcrd uses your default shell (as specified by the $SHELL environment variable). Ensure your .zshrc (or equivalent) is properly configured to work with rcrd.

## Logs

Logs are stored in the ~/.rcrd directory. Each log file is named with a timestamp or the custom name you provided.

## Troubleshooting

If you encounter any issues:
1. Ensure your Go installation is up to date.
2. Check that your .zshrc file is correctly configured.
3. Make sure you have write permissions for the ~/.rcrd directory.

For more help, please open an issue on the GitHub repository.
