# Minikube Basics

We use Kubernetes for orchestrating division online's services, and minikube
is the easiest way to run a local, single node Kubernetes cluster on your own
machine to practice with. Before touching any of the exercises below, you need
a running cluster.

## Setting up the cluster

1. Install a container runtime if you don't already have one (Docker is
   fine, since you should already have it from the devops exercises).
2. Install kubectl, the command line tool used to talk to a Kubernetes
   cluster:
   ```sh
   curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
   sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
   ```
3. Install minikube:
   ```sh
   curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
   sudo install minikube-linux-amd64 /usr/local/bin/minikube
   ```
4. Start the cluster:
   ```sh
   minikube start
   ```
   This downloads a small VM/container image and boots a one node cluster
   inside it. It can take a few minutes the first time.
5. Confirm the cluster is up:
   ```sh
   kubectl get nodes
   ```
   You should see a single node listed with status `Ready`.
6. Point kubectl at minikube (usually done automatically by `minikube start`,
   but worth double checking):
   ```sh
   kubectl config current-context
   ```
   This should print `minikube`.

Keep the cluster running for both exercises below. When you are done, you can
tear it down with `minikube stop` or `minikube delete`.

## Exercise 1: deploy a basic image from Docker Hub

Deploy the `nginxdemos/hello` image (a tiny web server that just replies with
some basic info about the request) onto your minikube cluster and make it
reachable from your own machine.

What you need to do:

1. Create a Deployment that runs the `nginxdemos/hello` image.
2. Create a Service that exposes that Deployment.
3. Use `minikube service` (or `kubectl port-forward`) to reach the service
   from your browser or with `curl`, and confirm you get a response back.

You are free to do this with plain `kubectl create`/`kubectl expose` commands,
or by writing your own YAML manifests and applying them with `kubectl apply
-f`. Either is fine, but you should understand what a Pod, a Deployment and a
Service are and how they relate to each other by the end of this exercise.

## Exercise 2: add a secret to the cluster

Real applications need to consume credentials (API keys, database passwords,
tokens, etc.) without hard coding them into images or manifests. Kubernetes
Secrets exist for exactly this.

What you need to do:

1. Create a Secret in your cluster holding at least one key/value pair, for
   example a fake `API_KEY`. You can do this either from the command line
   with `kubectl create secret generic ...` or from a YAML manifest.
2. Update the Deployment from exercise 1 (or create a new one) so that the
   Secret is injected into the container as an environment variable.
3. Confirm the value actually reaches the container, for example with:
   ```sh
   kubectl exec <pod-name> -- printenv API_KEY
   ```

Do not commit the actual secret value anywhere, only the commands or
manifests you used (with a placeholder value) if you want to keep a record of
what you did.
