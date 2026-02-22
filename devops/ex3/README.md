# Containerized Deployment

This repository contains the source code of a **CRITICAL** website that we want to deploy
and make accesible at localhost:6666. We must also do this in a containerized manner so
as to assure consistency across environments, isolation of the application and a simplified
CI/CD pipeline.

What you must do is change `compose.yaml` and `Dockerfile` so that when you run:
```sh
docker compose up
```
You are able to access the web service through a call like:
```sh
curl localhost:6666/hi/lol/
```

## IMPORTANT
**ALL** your changes must be done to either `compose.yaml` or `Dockerfile` do **NOT** change
the source code of the website at all. Currently `compose.yaml` is set up but it doesn't
expose the app? Why? Also running build on `Dockerfile` throws some C errors... Why?

The app **MUST** be accesible at port 6666 (from outside the container).

## POST SCRIPTUM
There is one more file that you might wanna change. Which is it?

