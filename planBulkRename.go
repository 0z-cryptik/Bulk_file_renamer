package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func prep_files_for_rename(folder string, sub_str string, replacement_str string) ([]renameParams, error) {
  if (sub_str == "" || replacement_str == "" || folder == "") {
	  return nil, fmt.Errorf("Error: Folder path and/or any of the rename parameters cannot be empty.")
  }

  files, err := os.ReadDir(folder)
  if (err != nil) {
	  return nil, fmt.Errorf("Error reading folder: %w", err)
  }

  fmt.Printf("\n")

  var details_of_files_to_rename []renameParams
  seen_path := make(map[string]bool)

  for _, file := range(files) {
    if (file.IsDir()){
      continue
    }

    oldName := file.Name()
    newName := strings.ReplaceAll(oldName, sub_str, replacement_str)

    if(oldName == newName) {
      continue
    }

    details := renameParams{
      old_name: oldName,
      new_name: newName,
      old_path: filepath.Join(folder, oldName),
      new_path: filepath.Join(folder, newName),
    }
    
    if(seen_path[details.new_path]){
      return nil, fmt.Errorf("Batch Collision Error, multiple files attempting to rename to %s", details.new_name)
    }

    seen_path[details.new_path] = true

    details_of_files_to_rename = append(details_of_files_to_rename, details)
    fmt.Printf("%s will be changed to %s\n", oldName, newName)
  }

  return details_of_files_to_rename, nil
}