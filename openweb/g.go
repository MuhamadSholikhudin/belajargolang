
// Go program to illustrate how to rename
// and move a file in default directory
package main
    
import (
    "log"
    "os"
)
    
func main() {
   
    // Rename and Remove a file
    // Using Rename() function
    Original_Path := "gfg.txt"
    New_Path := "gfg.bat"
    e := os.Rename(Original_Path, New_Path)
    if e != nil {
        log.Fatal(e)
    }
      
}