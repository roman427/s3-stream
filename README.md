# S3-Stream

S3-Stream is a program that streams data to Amazon S3.

## Building requirements

* [Go](https://golang.org/) -  main dependency

## Compiling and running

1. Check if you have go installed on your machine by typing:

    ```
    $ go version
    ```

2. Run a *build.sh* script:
    ```
    $ ./build.sh
    ```

3. If you don't prefer running scripts or you are using Windows, then run an app directly:
    ```
    $ go run ./cmd/web/main.go
    ```

4. After all, executable will be created in bin/ folder, if you include _--help_ flag to the installer, you can view available flags.
    ```
    $ ./s3-stream --help
    ```

## Description

This project is an entry for [freelancer.com](https://www.freelancer.com/contest/Golang-Stream-Data-to-S-1757553) contest. Description of the project is in docs/desc.pdf.

## Design Notes

A project structure is inspired by [manfred](https://manfred.life/golang-project-layout).

Project is built using famous package [cobra](https://github.com/spf13/cobra). 

Instead of standard net/http, [fasthttp](https://github.com/valyala/fasthttp) is used, because of the project requirements.

I've used only / route for the server, because no other routes are needed.