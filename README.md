# streaming-video-with-watchlist

In this tutorial, we will build a live streaming application that displays the online presence of the users you are interested in and who are currently streaming a video.

[View tutorial](https://pusher.com/tutorials/build-a-streaming-platform-and-watch-videos-with-your-friends/)

## Features used

* User authentication
* Watchlist Online Status
* Subscription count event

## Technologies used

* Pusher Channels
* Go
* JavaScript (Vue)

## Prerequisites

* An IDE of your choice, e.g. Visual Studio Code.
* Go (version >= 0.10.x) installed on your computer.
* JavaScript (Vue) installed on your computer.
* Pusher app.

## Getting Started

To get started with the project, make sure you have all the prequiisites above.

1. Clone the project to your machine.
2. Update the Pusher keys in the `.env` and `./static/dashboard.html` files.
3. Run the command: `$ go mod init`.
4. Run the command: `$ go mod tidy`.
5. Run the command: `$ go get github.com/joho/godotenv`.
6. Run the command: `$ go get github.com/pusher/pusher-http-go/v5`.
7. Run the command: `$ go run server.go`.
8. Visit `http://localhost:8090` to see application in action.
