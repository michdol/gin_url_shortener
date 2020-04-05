# gin_url_shortener

### Installation
Requires docker installed

```
cp ~/.ssh/id_rsa.pub WORKDIR/certificates/dev-certificate.pub
docker-compose up -d
```

### Install Go
```
ssh root@localhost -p 2228

cd /tmp
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
sudo tar -xvf go1.11.linux-amd64.tar.gz
sudo mv go /usr/local
```

### Setup env variables
Add those lines to the end of ~/.profile and run `source ~/.profile`
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

### Install dependencies
```
go get github.com/go-sql-driver/mysql
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
```

### Usage

```
go run main.go
```

###### Create entry
```
curl -X POST localhost:8080/v1/urls -d '{"url": "https://www.nintendo.com/switch/"}'
{"id":16,"url":"https://www.nintendo.com/switch/"}
```

###### Use short url
```
http://localhost:8080/r/16
301
```

###### Delete entry
```
curl -X DELETE localhost:8080/v1/urls/16
204
```
