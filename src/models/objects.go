package models

import (
    "log"
    "encoding/xml"
    "io/ioutil"
)


// Struct interfaced by Tool
type ObjectsModel struct {
    Data        ObjectDB `xml:"ObjectDB"`
}

type ObjectDB struct {
    Object      Object  `xml:"Objects"`
}

// Object List
type Object struct {
    Items       []Item  `xml:"item"`
}

// Item list
type Item struct {
    Name        string  `xml:"name"`
    Type        string  `xml:"type"`
    GameType    string  `xml:"game_type"`
    Path        string  `xml:"path"`
}

// Load objects from xml file.
//
// @param string xml file name
// @return ObjectsModel
func (objects *ObjectsModel) Load(file string) (err error){
    log.Println("Reading file ",file)
    xmlData, err := ioutil.ReadFile(file)

    if err != nil {
        log.Println("Could not read file: ",file)
        return err
    }

    log.Println(len(xmlData))

    objects.Data = ObjectDB{}
    xml.Unmarshal(xmlData, &objects.Data)
    return nil
}

