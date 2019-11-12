# demo-jobs

The YAML files in this folder represent different [Custom Job]() configurations. Each of these can be dropped into Orca's
configuration (`orca-local.yml`) and used within pipelines. 

## Requirements
* Spinnaker 1.17+
* Kubernetes V2 provider account configured and enabled
* [Git Repo artifact]() configured and enabled


## Job Description

### Kaniko (kaniko-build-image)

The Kaniko stage uses [Kaniko]() to build and push Docker images. Using this stage enables you to push image building
into the pipeline that deploys it. It uses an `initContainer` to fetch the Git repo at `REPO` and write it to a shared
volume. The single `container` runs Kaniko against the source code files in this shared volume, building a new Docker
image and pushing it to your desired registry.

*Note: You will need to mount a secret called `docker-creds` who's 
contents are your credentials for your Docker registry.*

### Run Script (run-script)
 
The Run Script stage enables users a script that is stored in a Git repo by providing their Repo URL, a Docker image,
and a command to run. It works in the same way the above Kaniko stage works. First, an `initContainer` fetches your source
code to a shared volume. Then, the user defined image and command are executed. This stage is useful if your users aren't
yet familiar with Kubernetes and have some one-off scripts they need to run with their deployments. 