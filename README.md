# turtle
A simple app that returns delayed results. The intent of this application is that sometimes when you're writing UI tests, you might have a test that timesout and doesn't wait long enough. To specifically test these kinds of tests, I wrote an application that is specifically slow so that you can test timeout issues.

# To Run
This is a golang application. You can run the app with:
```
go install
go run main.go
```

