package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type renameParams struct {
  old_name, new_name, old_path, new_path string
}

type userInput struct {
  folder, text_to_replace, text_to_replace_with string
}

func get_user_input() (userInput, error){
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
	  return userInput{}, err
  }

  var user_input userInput
  user_input.folder = folder
  user_input.text_to_replace = text_to_replace
  user_input.text_to_replace_with = text_to_replace_with

  return user_input, nil
}

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

    fmt.Printf("%s successfully changed to %s\n", file.old_name, file.new_name)
    files_successfully_renamed++
  }

  return files_successfully_renamed, nil
}

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

  files_successfully_renamed, err := perform_bulk_rename_op(details_of_files_to_rename)
  if(err != nil){
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
  }

  fmt.Printf("Operation successful, %d files renamed successfully", files_successfully_renamed)
}