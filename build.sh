GOOS=linux GOARCH=amd64 go build -o app

scp -i elgaml.pem .env admin@13.60.191.29:/home/admin/far/
scp -i elgaml.pem cert.pem admin@13.60.191.29:/home/admin/far/
scp -i elgaml.pem key.pem admin@13.60.191.29:/home/admin/far/
scp -i elgaml.pem ./app admin@13.60.191.29:/home/admin/far/

