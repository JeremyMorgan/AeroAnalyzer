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


- what would you name an application that displays and sorts data about drones?
- how do you start a new go application?
- how do I create a web api with go?
- how do you run a go program?
- how to do you try to run a statement in go, and show and error if it fails?
- how do I find an error with http.ListenAndServe?
- how do I find errors with http.NewServeMux?
- Can you give me a list of the ten most popular camera drones, show their flight time, sensor size, weight, top speed, and cost
- how do you access data in sqlite with go
- can you write a sql query for sqlite?
- create a struct in go based on this sqlite table  
```
CREATE TABLE IF NOT EXISTS "Drone" (
        "Id"    INTEGER,
        "Name"  TEXT,
        "FlightTime"    INTEGER,
        "SensorSize"    TEXT,
        "WeightMetric"  NUMERIC,
        "WeightImperial"        NUMERIC,
        "TopSpeedMetric"        NUMERIC,
        "TopSpeedImperial"      NUMERIC,
        "Cost"  NUMERIC,
        PRIMARY KEY("Id" AUTOINCREMENT)
);
```
- how do I insert data from rows retrieved from a sqlite query into a struct in go?
- can you generate a boilerplate html file for starting a web page
- can you generate a boilerplate html file for starting a web page using bootstrap and make it look pretty?
- do you have some newer links to the Bootstrap CDN? those ones do not work
- How do I make a call to a web api with vanilla javascript and display the data from JSON on the webpage?
- How do I make a call to a web api with vanilla javascript and display the data from JSON on the webpage?  

Here's the JSON: 
```
[
    {
        "id": 1,
        "name": "DJI Mavic 2 Pro",
        "flight_time": 31,
        "sensor_size": "1 inch (20 MP)",
        "weight_metric": 907,
        "weight_imperial": 2,
        "top_speed_metric": 72,
        "top_speed_imperial": 45,
        "cost": 1499
    },
    {
        "id": 2,
        "name": "Phantom 4 Pro",
        "flight_time": 30,
        "sensor_size": "1 inch (20 MP)",
        "weight_metric": 1388,
        "weight_imperial": 3.1,
        "top_speed_metric": 72,
        "top_speed_imperial": 45,
        "cost": 1499
    },
    {
        "id": 3,
        "name": "DJI Mavic Air 2",
        "flight_time": 34,
        "sensor_size": "1/2 inch (48 MP)",
        "weight_metric": 570,
        "weight_imperial": 1.3,
        "top_speed_metric": 68,
        "top_speed_imperial": 42,
        "cost": 799
    },
    {
        "id": 4,
        "name": "DJI Mini 2",
        "flight_time": 31,
        "sensor_size": "1/2.3 inch (12 MP)",
        "weight_metric": 249,
        "weight_imperial": 0.5,
        "top_speed_metric": 57.6,
        "top_speed_imperial": 36,
        "cost": 449
    },
    {
        "id": 5,
        "name": "Autel Robotics Evo II",
        "flight_time": 40,
        "sensor_size": "1 inch (20 MP)",
        "weight_metric": 1127,
        "weight_imperial": 2.5,
        "top_speed_metric": 72,
        "top_speed_imperial": 45,
        "cost": 1495
    },
    {
        "id": 6,
        "name": "Parrot Anafi",
        "flight_time": 25,
        "sensor_size": "1/2.4 inch (21 MP)",
        "weight_metric": 320,
        "weight_imperial": 0.7,
        "top_speed_metric": 55,
        "top_speed_imperial": 34,
        "cost": 699
    },
    {
        "id": 7,
        "name": "Yuneec Typhoon H Pro",
        "flight_time": 25,
        "sensor_size": "1/2.3 inch (12 MP)",
        "weight_metric": 1950,
        "weight_imperial": 4.3,
        "top_speed_metric": 70,
        "top_speed_imperial": 43,
        "cost": 1699
    },
    {
        "id": 8,
        "name": "Holy Stone HS100D",
        "flight_time": 15,
        "sensor_size": "1/2.5 inch (1080P HD)",
        "weight_metric": 700,
        "weight_imperial": 1.5,
        "top_speed_metric": 45,
        "top_speed_imperial": 28,
        "cost": 279
    },
    {
        "id": 9,
        "name": "DJI Mavic Pro",
        "flight_time": 27,
        "sensor_size": "1/2.3 inch (12 MP)",
        "weight_metric": 734,
        "weight_imperial": 1.6,
        "top_speed_metric": 64.8,
        "top_speed_imperial": 40,
        "cost": 999
    },
    {
        "id": 10,
        "name": "DJI Inspire 2",
        "flight_time": 27,
        "sensor_size": "Micro Four Thirds (various options)",
        "weight_metric": 4280,
        "weight_imperial": 9.4,
        "top_speed_metric": 94,
        "top_speed_imperial": 58,
        "cost": 3299
    }
]
```
- how do I fix an error that localhost has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present, in my go web api application?
- how do I style html tables with bootstrap so they look better?
- how do I put table-striped into rows, and tablebordered into rows of my table?
- how do I get table-striped to work with the fetch javascript code you generated earlier?
- the table rows are not striped? Why?
- how do I use different stripe colors?
- how do I make the table only take up 80% of the page?
- how do I center a table with css?
- how do I center text in css?
- can you generate a pretty background for my webpage using CSS? 
- how do you make the backround scale to 100% of the page and not tile
- how do I change the color of my table header in bootstrap?
- can you generate me a 16x16 pixel favicon in svg that looks like a drone?
- can you generate me a 16 pixel by 16 pixel favicon in svg that looks like a drone flying?


