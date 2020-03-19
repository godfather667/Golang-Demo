# golang-demo

## Introduction
This is a program produced for an employment opportunity.

It will read a JSON file from www.data.gov and display it on a web page.

The assignment requirements required the inclusion of the following packages:
* https://github.com/labstack/echo
* https://github.com/go-resty/resty
* https://github.com/buger/jsonparser
* https://github.com/jinzhu/gorm

I provide two endpoints:
* localhost:8000/gov  -- Load Data and Build Database
* localhost:8000/rep  -- Read Database and Display Report Web Page

This program was compiled and tested on a windows system using **go version go1.13.5 windows/amd64**

To Run:

* Compile gov_data.go
* Run the executable with a web window at localhost:8000
* Enter the endpoint **/gov** - This acquire the data and build a Database
* Enter the endpoint **/rep** - This reads the database and displays a Web Page









