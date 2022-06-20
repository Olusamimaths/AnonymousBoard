# AnonymousBoard
Anonymous Board API written in GoLang, project description is here https://github.com/freeCodeCamp/freeCodeCamp/blob/main/curriculum/challenges/english/09-information-security/information-security-projects/anonymous-message-board.md

Link to the Youtube tutorial [Youtube Playlist](https://www.youtube.com/watch?v=V3T5d1TZ3wA&list=PLBeQxJQNprbjDQ4fJi58jOcZYA8VYauVh) by [Kelvin Mai](https://github.com/kelvin-mai), credits to him for providing the free resource.


## Starting the Project
Run `docker-compose up` to start the database then run `go run main.go` to start the go app

# Architetural Diagrams
To further help in understanding the lessons from the video tutorial, I have independently created the following Pseudo-Architetural Diagrams

## Modules
![image](https://user-images.githubusercontent.com/31919592/174629352-498667ea-8ea4-44a6-9d11-8950f7c58fcd.png)

## main.go 
Relationship with other modules
![image](https://user-images.githubusercontent.com/31919592/174629680-fc9778c6-9a49-42d7-b9e4-2bc6044e2c19.png)

## Flow in main.go<br/>
<img src="https://user-images.githubusercontent.com/31919592/174629842-31dcf7ae-2394-4753-ae46-6142d5915cc4.png" width="400" center>


# Class Diagrams
Although Go doesn't have classes, I have created a class diagram that models the structs and interfaces in the project
![image](https://user-images.githubusercontent.com/31919592/174631005-953ab2ca-d664-4299-a72d-4644ef210949.png)
