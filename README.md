![OpenGl](opengl.png | height=200)
![Golang](golang.jpg | height=200)

# ppm-images-viewer

Simple portable pixmap images viever powered by Golang & OpenGl

# Getting started

To get the app running locally:

- Clone this repo
- Install required dependences:
    ```
    go get -u github.com/go-gl/gl/v4.1-core/gl
    go get -u github.com/go-gl/glfw/v3.2/glfw
    go get -u github.com/lmittmann/ppm
    ```
- From root directory run the app by `go run main.go`
- Application will scan gallery folder for *.ppm images and display the first of them. To navigate, just press **left** or **right arrow** on your **keyboard**.
