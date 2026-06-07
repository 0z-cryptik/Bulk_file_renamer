package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main(){
  var folder string
  var text_to_change string
  var text_to_replace string

  fmt.Printf("Enter the path of the folder where you want to rename files: ")

  scanner := bufio.NewScanner(os.Stdin)

  if (scanner.Scan()){
	  folder = scanner.Text()
  }

  fmt.Printf("Enter the subtext you want replace: ")
  
  if (scanner.Scan()) {
	  text_to_change = scanner.Text()
  }

  fmt.Printf("Enter the text you want to replace it with: ")

  if (scanner.Scan()){
    text_to_replace = scanner.Text()
  }

  if err := scanner.Err(); err != nil {
	  fmt.Fprintf(os.Stderr, "%v\n", err)
	  os.Exit(1)
  }

  if (text_to_change == "" || text_to_replace == "" || folder == "") {
	  fmt.Fprintf(os.Stderr, "\nError: Folder path and/or any of the rename parameters cannot be empty.\n")
	  os.Exit(1)
  }

  files, err := os.ReadDir(folder)
  if (err != nil) {
	  fmt.Fprintf(os.Stderr, "%v\n", err)
	  os.Exit(1)
  }

  fmt.Printf("\n")

  for _, file := range(files) {
    if (file.IsDir()){
      continue
    }

    old_name := file.Name()
    new_name := strings.ReplaceAll(old_name, text_to_change, text_to_replace)

    if(old_name == new_name) {
      continue
    }

    fmt.Printf("%s will be changed to %s\n", old_name, new_name)
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
  
  for _, file := range(files) {
    if (file.IsDir()){
      continue
    }

    old_name := file.Name()
    new_name := strings.ReplaceAll(old_name, text_to_change, text_to_replace)

    if (old_name == new_name){
      continue
    }

    old_path := filepath.Join(folder, old_name)
    new_path := filepath.Join(folder, new_name)

    // because if no error then the file already exist
    if _, err = os.Stat(new_path); err == nil {
      fmt.Printf("skipping %s because a file named %s already exists in this folder\n", old_name, new_name)
      continue
    }

    err := os.Rename(old_path, new_path)
    if (err != nil) {
      fmt.Fprintf(os.Stderr, 
        "Error encountered while performing rename operation. %d files successfully renamed\nERROR DETAILS: %v\n", 
        files_successfully_renamed, err)
      
      os.Exit(1)
    }

    fmt.Printf("%s successfully changed to %s\n", old_name, new_name)
    files_successfully_renamed++
  }

  fmt.Printf("\nOperation successful, %d files renamed\n\n", files_successfully_renamed)
}