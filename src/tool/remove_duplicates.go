package tool

import (
    "io"
    "log"
    "path"
    "models"
    "encoding/xml"
    "io/ioutil"
)

// Struct interfaced by Tool
type RemoveDuplicates struct {}

// The name of the tool
func (tool * RemoveDuplicates) Name() string {
    return "remove-duplicates"
}

// How to use this tool
func (tool * RemoveDuplicates) Help() string {
    return "remove-duplicates [file-name]"
}

// Run this tool
//
// This tool will read a xml file and will try to remove duplicated entries.
// It will remove objects with the same name
func (tool * RemoveDuplicates) Run(args []string, out io.Writer) (err error) {

    if len(args) < 1 {
        log.Println("Not enough parameters, please follow this command. ")
        log.Println(tool.Help())
        return nil
    }

    // Read xml entries
    model := models.ObjectsModel{}
    err = model.Load(args[0])
    if err != nil {
        log.Println("Could not load xml file ",args)
        return err
    }

    log.Println(len(model.Data.Object.Items))

    // Remove duplicates
    hash := map[string]models.Item{}
    for _,item := range model.Data.Object.Items {
        hash[item.Name] = item
    }

    cleanObject := models.Object{}
    cleanObject.Items = make([]models.Item,len(hash))

    for _,item := range hash {
        cleanObject.Items = append(cleanObject.Items,item)
    }

    objDB := models.ObjectDB{}
    objDB.Object = cleanObject

    objModel := models.ObjectsModel{}
    objModel.Data = objDB

    log.Println(len(objModel.Data.Object.Items))

    // Write result
    newFileData,err := xml.Marshal(objModel.Data)
    if err != nil {
        log.Println("Could not marshall cleaned xml",args)
        return err
    }

    dir,file:= path.Split(args[0])
    err = ioutil.WriteFile(dir+"new_"+file,newFileData,0755)
    if err != nil {
        log.Println("Could not marshall cleaned xml",args)
        return err
    }
    return nil
}