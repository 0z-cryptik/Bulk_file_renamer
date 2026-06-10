# Bulk File Renamer

A simple CLI tool that renames multiple files in a folder by replacing a specified substring in their filenames.

## Features

* Preview all incoming filename changes before applying them
* Detects filename collisions within the batch rename operation
* Skips files that would overwrite existing files

## Notes

* Existing files will not be overwritten.
* If multiple files would be renamed to the same filename, the operation is aborted before any files are renamed.
* Only files in the specified folder are processed; subfolders are ignored.

## How To Use

### Run the program:

```bash
go run .
```

You will be prompted for:

1. The path to the target folder
2. The text to search for in filenames
3. The replacement text

The program will display a preview of all planned changes and ask for confirmation before renaming any files.

Enter Y to proceed or any other value to cancel.

Example

Given the following files:

```bash
photo_old.jpg
document_old.pdf
notes_old.txt
```

Input:

```bash
Text to replace: old
Replace with: new
```

Preview:

```bash
'photo_old.jpg' will be changed to 'photo_new.jpg'
'document_old.pdf' will be changed 'to document_new.pdf'
'notes_old.txt' will be changed to 'notes_new.txt'
```

After confirmation, the files will be renamed accordingly.

Build

To build an executable:

```bash
go build
```