package main

import (
	"bufio"
	"fmt"
  "strings"
)

func clean_path(path string) string{
  path = strings.TrimSpace(path)
  path = strings.Trim(path, `"`)
  path = strings.Trim(path, `'`)
  return path
}

func get_user_input(scanner *bufio.Scanner) (userInput, error){
  var folder string
  var text_to_replace string
  var text_to_replace_with string

  fmt.Print("> Enter the path of the folder where you want to rename files: ")
  if (scanner.Scan()){
	  folder = scanner.Text()
  }

  fmt.Printf("> Enter the subtext you want replace: ")
  if (scanner.Scan()) {
	  text_to_replace = scanner.Text()
  }

  fmt.Printf("> Enter the text you want to replace it with: ")
  if (scanner.Scan()){
    text_to_replace_with = scanner.Text()
  }

  if err := scanner.Err(); err != nil {
	  return userInput{}, err
  }

  var user_input userInput
  user_input.folder = clean_path(folder)
  user_input.text_to_replace = text_to_replace
  user_input.text_to_replace_with = text_to_replace_with

  return user_input, nil
}