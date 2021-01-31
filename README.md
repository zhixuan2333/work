# zhixuan

## project
- go
    - [autobackup](./go/autobackup)
        - 一键backup
    - [vscode_ffmpeg_go](./go/vscode_electron)
        - 获取 ```vscode``` 对应的 ```electron``` 文件
    - [mc_heroku](./go/mc-heroku)
        - heroku minecraft

    <!-- - [line](./go/line)
        - line webhook server -->

## note

go mod init heroku


go mod tidy

heroku container:push web

heroku container:release web

heroku ps:scale web=1