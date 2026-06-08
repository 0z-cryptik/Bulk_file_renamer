package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type renameParams struct {
  old_name string
  new_name string
  old_path string
  new_path string
}

func main(){
  var folder string
  var text_to_replace string
  var text_to_replace_with string

  fmt.Printf("Enter the path of the folder where you want to rename files: ")

  scanner := bufio.NewScanner(os.Stdin)

  if (scanner.Scan()){
	  folder = scanner.Text()
  }

  fmt.Printf("Enter the subtext you want replace: ")
  
  if (scanner.Scan()) {
	  text_to_replace = scanner.Text()
  }

  fmt.Printf("Enter the text you want to replace it with: ")

  if (scanner.Scan()){
    text_to_replace_with = scanner.Text()
  }

  if err := scanner.Err(); err != nil {
	  fmt.Fprintf(os.Stderr, "%v\n", err)
	  os.Exit(1)
  }

  if (text_to_replace == "" || text_to_replace_with == "" || folder == "") {
	  fmt.Fprintf(os.Stderr, "\nError: Folder path and/or any of the rename parameters cannot be empty.\n")
	  os.Exit(1)
  }

  files, err := os.ReadDir(folder)
  if (err != nil) {
	  fmt.Fprintf(os.Stderr, "%v\n", err)
	  os.Exit(1)
  }

  fmt.Printf("\n")

  var details_of_files_to_rename []renameParams

  for _, file := range(files) {
    if (file.IsDir()){
      continue
    }

    oldName := file.Name()
    newName := strings.ReplaceAll(oldName, text_to_replace, text_to_replace_with)

    if(oldName == newName) {
      continue
    }

    details := renameParams{
      old_name: oldName,
      new_name: newName,
      old_path: filepath.Join(folder, oldName),
      new_path: filepath.Join(folder, newName),
    }

    details_of_files_to_rename = append(details_of_files_to_rename, details)
    fmt.Printf("%s will be changed to %s\n", oldName, newName)
  }

  if (len(details_of_files_to_rename) < 1) {
    fmt.Printf("No files matched the subtext '%s'\n", text_to_replace)
    return
  }

  fmt.Printf("\nEnter 'Y' to proceed with the rename operation, enter anything else to terminate: ")

  var user_answer string

  _, err = fmt.Scan(&user_answer)
  if (err != nil){
    fmt.Fprintf(os.Stderr, "Error reading user input\nError: %v\n", err)
    os.Exit(1)
  }

  if (user_answer != "Y" && user_answer != "y"){
    return
  }

  files_successfully_renamed := 0
  fmt.Printf("\n")
  
  for _, file := range(details_of_files_to_rename) {

    // because if no error then the file already exist
    if _, err := os.Stat(file.new_path); err == nil {
      fmt.Printf("skipping %s because a file named %s already exists in this folder\n", file.old_name, file.new_name)
      continue
    }

    err := os.Rename(file.old_path, file.new_path)
    if (err != nil) {
      fmt.Fprintf(os.Stderr, 
        "Error encountered while performing rename operation. %d files successfully renamed\nERROR DETAILS: %v\n", 
        files_successfully_renamed, err)
      
      os.Exit(1)
    }

    fmt.Printf("%s successfully changed to %s\n", file.old_name, file.new_name)
    files_successfully_renamed++
  }

  fmt.Printf("\nOperation successful, %d files renamed\n\n", files_successfully_renamed)
}