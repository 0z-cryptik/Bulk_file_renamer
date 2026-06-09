package main

import (
	"fmt"
	"os"
)

func main(){
  user_input, err := get_user_input()
  if(err != nil){
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  folder := user_input.folder
  subtext_to_match := user_input.text_to_replace
  replacement_text := user_input.text_to_replace_with

  details_of_files_to_rename, err := prep_files_for_rename(folder, subtext_to_match, replacement_text)
  if(err != nil){
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  if (len(details_of_files_to_rename) < 1) {
    fmt.Printf("No files matched the subtext '%s'\n", subtext_to_match)
    return
  }

  fmt.Print("\n> Enter 'Y' to proceed with the rename operation, enter anything else to terminate: ")

  var user_answer string

  _, err = fmt.Scan(&user_answer)
  if (err != nil){
    fmt.Fprintf(os.Stderr, "Error reading user input\nError: %v\n", err)
    os.Exit(1)
  }

  if (user_answer != "Y" && user_answer != "y"){
    return
  }

  files_successfully_renamed, err := perform_bulk_rename_op(details_of_files_to_rename)
  if(err != nil){
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  fmt.Printf("\nOperation successful, %d files renamed successfully\n", files_successfully_renamed)
}