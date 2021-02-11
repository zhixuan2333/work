# zhixuan

## project
- go
    - [autobackup](./go/autobackup)
        - 一键backup
    - [vscode_ffmpeg_go](./go/vscode_electron)
        - 获取 ```vscode``` 对应的 ```electron``` 文件
    <!-- - [line](./go/line)
        - line webhook server -->

## note 

go mod init heroku

go mod tidy

heroku container:push web

heroku container:release web

heroku ps:scale web=1

function ClickConnect() {
console.log("Working");
document
  .querySelector('#top-toolbar > colab-connect-button')
  .shadowRoot.querySelector('#connect')
  .click()
}
setInterval(ClickConnect, 60000)
