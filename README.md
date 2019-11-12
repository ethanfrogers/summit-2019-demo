# summit-2019-demo

## About this project

This project contains the source code for the demo I gave at the 2019 Spinnaker Summit in San Diego, CA. The topic of
this talk was about using the Run Job and Webhook stages to extend Spinnaker in a low-friction way.  In this repo you'll
find the source code for my demo application, a Kubernetes deployment manifest, the JSON for the demo pipeline, and 
YAML files containing the Custom Job and Webhook configurations. 

## Sample App

The sample application is written using Go and provides a set of unit and integration tests.

Unit tests can be run via `go test ./... --short`. 

Integration tests can be run via `go test ./integration --endpoint={some url}`. The goal of these tests is to validate
a deployment and are run as after this application has been deployed to staging.
