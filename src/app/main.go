package main

/**
 * Interface for blitzkrieg project
 *
 * TODO: Create tool to remove duplicates from xml
 * TODO: Run tools without window
 * TODO: Open window in windows
 */
import (
    "os"
    "fmt"
    "log"
    "flag"
    "tool"
)

// Name of program
const PROGRAM_NAME = "Blitzkirg Support project"


// Populate tools.
//
// Return a map with tool name and the instantiation of it
func PopulateTools() (tls map[string] tool.Tool){

    tls = make(map[string]tool.Tool,1)

    t := &tool.RemoveDuplicates{}
    tls[t.Name()] = t

    // Add another
    //t = tools.<Name of tool>{}
    //tls.put[t.Name()] = t
    return tls
}


// Print information about how run all tools
func Help(tls map[string]tool.Tool) {

    fmt.Println("Possible tools:")

    for _, tool := range tls {
        fmt.Println(tool.Help())
    }
}

/**
 * Start program open tools.
 */
func main(){
    log.Println("Init ",PROGRAM_NAME)

    flag.Parse()
    tls := PopulateTools()

    toolName := flag.Arg(0)

    if t,ok := tls[toolName]; ok {
        // Tool exists
        log.Println("Running tool ",toolName)
        args := flag.Args()[1:]
        err := t.Run(args,os.Stdout)
        if err != nil {
            log.Println("ERROR: ",err.Error())
        } else {
            log.Println("Well done!")
        }
    } else {
        Help(tls)
    }
}