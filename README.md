# AeroAnalyzer

This was an application built live on [my Twitch Channel](https://www.twitch.tv/jeremymorgan) using only ChatGPT to build the code. [View the replay of the stream here](https://www.youtube.com/watch?v=Rt7CuhFv8xM)

**Note**: This not my code!! 99% of it was written by ChatGPT using a series of prompts. 

It's a Web application to display drones for comparison when considering a purchase. 

## How it works: 

**AeroAnalyzer Web** is a Web API written in Go. It uses the [net/http package](https://pkg.go.dev/net/http) is very primitive. It selects the records from a [SQLite](https://www.sqlite.org/index.html) database on the file system. 

To run it, type

```
go main go
``` 

or 

```
go build
```

It will run as a server that listens on `port 8080` though you can change it to whatever you like. 

a request to `http://localhost:8080/drone` will return a JSON list of objects to parse. 

It can be run as a service on a server, or container. It could probably also run as a Lambda or some other kind of serverless service. 

**AeroAnalyzer Client** 

This is the web interface portion of the application, and it consists of a single HTML file, which contains vanilla JavaScript to call the API and display the JSON on the screen. It also contains some CSS. 

It can be served up on any static site generator service, or traditional web hosting. 

### Possible other names for this product:

I asked ChatGPT to come up with some names, and I chose the corniest one. 

Here were the runner ups:

```
Here are a few suggestions for naming an application that displays and sorts data about drones:

DroneDash
SkySort
FlightFinder
AeroAnalyzer
WingWatch
DroneDeck
AirScope
HoverHub
SkySpy
FlyFocus
Of course, the best name for the application would depend on the specific goals and features of the application. It's important to choose a name that is descriptive, easy to remember, and evokes a sense of what the application does.
```

