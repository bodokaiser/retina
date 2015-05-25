# Retina

Automated retinopathy detection.

## Example

    go run ./cmd/httpd/main.go

Start the server, optionally customize address to listen via `--listen`.

    curl -H "Content-Type: application/json" -X POST -d '{"input":"<encoded as base64>"}' http://localhost:3000/

As result you should get a json encoded response which contains the `output`
attribute which is the image encoded as base64 with marked symptoms.

## License

Copyright © 2015 Bodo Kaiser <i@bodokaiser.io>