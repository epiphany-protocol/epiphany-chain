go build -o polygon-edge main.go
mv polygon-edge /usr/local/bin
scp polygon-edge root@8.219.62.228:/usr/local/bin
scp polygon-edge root@8.219.129.249:/usr/local/bin
scp polygon-edge root@8.219.143.221:/usr/local/bin