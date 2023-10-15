# birdfeeder

## Local Development

Run the terraform
Run `gcloud init` and switch to the birdfeeder1 project the terraform created

Run `go build main.go`

## Troubleshooting

if auth is not working
delete the cred file: `rm ~/.config/gcloud/application_default_credentials.json`
then regenerate it: `gcloud auth application-default login`
