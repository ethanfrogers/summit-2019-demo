# fetch-artifact

`fetch-artifact` is a utility tool for for fetching artifacts from Spinnaker's Clouddriver artifact APIs. This tools is meant to be used as a
pre-job `initContainer` for fetching and artifact and populating a workspace for downstream container runs i.e. more initContainers or the
job's final container.