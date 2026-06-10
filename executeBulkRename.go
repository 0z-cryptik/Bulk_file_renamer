package main

import (
	"fmt"
	"os"
)

func perform_bulk_rename_op(file_details []renameParams) (int, error) {
  files_successfully_renamed := 0
  fmt.Printf("\n")

  for _, file := range(file_details) {
    _, err := os.Stat(file.new_path)
    // because if no error then the file already exist
    if (err == nil) {
      fmt.Printf("skipping %s because a file named %s already exists in this folder\n", file.old_name, file.new_name)
      continue
    }

    //some other error
    if (!os.IsNotExist(err)){
      return files_successfully_renamed, fmt.Errorf("Error checking file %w", err)
    }

    err = os.Rename(file.old_path, file.new_path)
    if (err != nil) {
      return files_successfully_renamed, fmt.Errorf("Error encountered while performing rename operation. %d files successfully renamed\nERROR DETAILS: %w", 
        files_successfully_renamed, err)
    }

    fmt.Printf("'%s' successfully changed to '%s'\n", file.old_name, file.new_name)
    files_successfully_renamed++
  }

  return files_successfully_renamed, nil
}