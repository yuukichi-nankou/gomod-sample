# gomod-sample
sample code for go modules


## Case 1 Use VScode remote containers 

``` 
# git clone https://github.com/microsoft/vscode-remote-try-go

<< Open remote container >>

# git clone https://github.com/yuukichi-nankou/gomod-sample
# go run gomod-sample/standalone/hello.go 
``` 


## Case 2 Use own Dockerfile   
``` 
# git clone https://github.com/yuukichi-nankou/godev

<< Open remote container >>

# git clone https://github.com/yuukichi-nankou/gomod-sample
# cd gomod-sample
# go mod init github.com/yuukichi-nankou/gomod-sample
# go build gomod/hello.go 
# .hello

<< forward port >>
1. Press F1 and run the Remote-Containers: Forward Port from Container... command.
2. Select port 5000.
3. Click "Open Browser" in the notification that appears to access the web app on this new port.
``` 

## Case 3 Use other container image
```
# git clone https://github.com/yuukichi-nankou/blog-gae-go

<< Open remote container >>

<< hugo >> 

<< exit and deploy>>
gcloud app deploy
```

## Case 4 Deploy container
```
# docker exec -it 923ad9b69d6b ash
# PORT=8080 && docker run -p 8080:${PORT} -e PORT=${PORT} monpote

```