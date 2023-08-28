#!/usr/bin/env node

const fs = require('fs')
const path = require("path")

const { large_temp, small_temp } = require("./utils/directory")

//The Objective is to delete unused files in the temp and %temp% folders of the computer 
//If you want to use this CLI , read the doc 

const removeFiles = async (pathName) => {
    fs.readdir(pathName, { encoding: "utf-8" }, (err, files) => {
        if (err) return console.log(err)
        files.forEach(file => {
            const localPath = path.resolve(pathName, file)
            fs.rm(localPath, { recursive: true, force: false }, (err) => {
                //Those Files could not be deleted for reasons such as no authority to delete them or there is some ressources that needs them (It's fine if they are not deleted)
                console.log("Processing . . . ")
            })
        })
        console.log("Cache Cleared Successfully !")
    })
}
removeFiles(large_temp)
removeFiles(small_temp)



