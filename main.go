package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
  var old_path string
  var new_name string

  fmt.Printf("Enter the path of the file you want to rename: ")

  scanner := bufio.NewScanner(os.Stdin)

  if (scanner.Scan()){
	old_path = scanner.Text()
  }

  fmt.Printf("Enter the new name you want for the file: ")
  
  if (scanner.Scan()) {
	new_name = scanner.Text()
  }

  if err := scanner.Err(); err != nil {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
  }

  if (old_path == "" || new_name == "") {
	fmt.Fprintf(os.Stderr, "Error: File path and new name cannot be empty.\n")
	os.Exit(1)
  }

  file, err := os.Stat(old_path)
  if (err != nil) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
  }

  if (file.IsDir()){
	fmt.Fprintf(os.Stderr, "This is a folder, not a file, please enter a file path\n")
	os.Exit(1)
  }


  folder := filepath.Dir(old_path)

  new_path := filepath.Join(folder, new_name)
  err = os.Rename(old_path, new_path)

  if (err != nil) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
  }

  fmt.Printf("File name successfully changed to %s\n", new_name)
}