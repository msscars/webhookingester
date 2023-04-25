# moscars-webhookingester

## Local Development

- make changes
- ensure minikube is running
- `cd {root}`
- run `./push.sh` this will build the docker images and load these into minikube
- `cd .helm`
- `helm install -f values.yaml moscars-webhookingester .`