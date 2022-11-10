# Email Sender

Compose emails from a template and save to file

## How to run

### From source

requires go v1.19+

```shell
go build -o ./send_email cmd/send_email/main.go
```

```shell
-customers customers.csv -out ./out/ -template email_template.json -error err.csv
```

### Using docker

1. Build image and name it `email_sender`
```shell
docker build -t email_sender .
```

2. Prepare a folder that contains your `customers.csv`, `email_template.json`, `out` directory, error file `errors.csv`.
We can use the `etc/test` folder as an example

3. Run docker with mount directory `etc/test` 
```shell
docker run -v /$(pwd)/etc/test/:/app/test/  email_sender_v1  ./send_email -customers test/customers.csv \
-out test/out/ \
-template test/email_template.json \
-error test/errors.csv
```

4. Check result in `out` and `errors.csv`
