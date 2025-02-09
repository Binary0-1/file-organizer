# File Organizer

## Overview
File Organizer is a command-line tool that helps you automatically organize files into folders based on their extensions.

## Features
- Organizes files into folders according to their extensions
- Customizable configuration via a JSON file
- Supports dry-run mode to preview changes
- Color-coded terminal output for better visibility

## Installation
1. Clone the repository:
   ```sh
   git clone git@github.com:Binary0-1/file-organizer.git
   ```
2. Navigate to the project directory:
   ```sh
   cd file-organizer
   ```
3. Build the project:
   ```sh
   go build -o file-organizer
   ```

## Usage
Run the tool with:
```sh
./file-organizer -path "/your/directory" -config "config.json"
```

### Flags
- `-path`: Path to the directory to organize (default: current directory)
- `-config`: Path to the configuration file (default: `config.json`)
- `-dry-run`: Run without making actual changes

## Configuration
Modify `config.json` to customize how files are organized. Example:
```json
{
  ".txt": "TextFiles",
  ".jpg": "Images",
  ".mp4": "Videos"
}
```

## License
This project is licensed under the MIT License.

## Author
Binary0-1

