# drone-teams

Drone plugin for sending teams webhook notifications

```bash
export DRONE_BUILD_STATUS=Failed
export DRONE_BUILD_ACTION=test
export DRONE_REPO_NAME=test
export DRONE_COMMIT_AUTHOR=test
export DRONE_COMMIT_MESSAGE=test
export DRONE_COMMIT_LINK=test
export PLUGIN_WEBHOOK=<WEBHOOK ENDPOINT>

cd cmd/drone-teams; go build; ./drone-teams 
```
