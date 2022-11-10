# Email Sender

Compose emails from a a template file and a customers csv file.
Output to a directory and invalid customers to a error csv file

```shell
./send_email

 -customers string
        (required) path to customers csv file
  -error string
        (required) path to errors csv file
  -out string
        (required) path to output emails directory
  -template string
        (required) path to email template json file

```

## How to run

### From source

requires go v1.19+

1. Build source into file `./send_email`

```shell
go build -o ./send_email cmd/send_email/main.go
```

2. Run it, can use the provided test data in `etc/test/`

```shell
./send_email -customers etc/test/customers.csv -out etc/test/out -template etc/test/email_template.json -error etc/test/errors.csv
```

3. Check the result in `etc/test/out` and `etc/test/errors.csv`

Remember to point to a different error file each running time

### Using docker

1. Build image and name it `email_sender`

```shell
docker build -t email_sender .
```

2. Prepare a folder that contains your `customers.csv`, `email_template.json`, `out` directory, error file `errors.csv`.
   We can use the `etc/test` folder as an example

3. Run docker with mount directory `etc/test` . Remember to remove the `etc/test/errors.csv` file if any.

```shell
docker run -v /$(pwd)/etc/test/:/app/test/  email_sender_v1  ./send_email -customers test/customers.csv \
-out test/out/ \
-template test/email_template.json \
-error test/errors.csv
```

4. Check result in `out` and `errors.csv`

## Test

```shell
make test

==> Running tests...
ok      github.com/ziggy192/email_sender        0.513s  coverage: 71.0% of statements
```

## Format code

```shell
make fmt
```

## Check lint
```shell
make lint
```
