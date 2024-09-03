# lscontent

`lscontent` is a command-line tool that prints the relative path and content of each file in a specified directory. It supports various options for customizing the output and file selection.

I use this for AI assistant projects where I need to quickly check the content of a directory.

## Features

- List files in a directory
- Optionally include subdirectories
- Respect .gitignore rules (optional)
- Customize output format

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/lscontent.git
   ```
3. Navigate to the project directory:
   ```
   cd lscontent
   ```
4. Install dependencies:
   ```
   go get github.com/sabhiram/go-gitignore
   ```
5. Build the project:
   ```
   go build
   ```

## Usage

```
lscontent [-r] [-i] [-f format] <DIR>
```

Options:
- `-r`: Also print subdirectories
- `-c`: Copy output to clipboard
- `-i`: Enable .gitignore rules
- `-f`: Use custom format (default: "*{path}*:\n```{suffix}\n{content}\n```")

## Custom Format

You can use the following placeholders in your custom format:
- `{path}`: Relative path of the file
- `{suffix}`: File extension (without the dot)
- `{content}`: File content
- `{newline}`: A newline character (`\n`)

Example:
```
lscontent -f "File: {path}{newline}Type: {suffix}{newline}---{newline}{content}{newline}---{newline}" /path/to/directory
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
