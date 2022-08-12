# act-artifact-server
An example Artifact server used with Act

# Description

This is a simple Githu Artifact Server that can be used in conjunction with [act](https://github.com/nektos/act).

# How to use ?

1. Launch the artifact server

```bash
docker run --name artifact-server -d -p 1234:1234 elron/act-artifact-server

```

2. In the project root directory create (or add) to your `.env` file

```bash
ACTIONS_CACHE_URL=http://localhost:1234/
ACTIONS_RUNTIME_URL=http://localhost:1234/
ACT_ARTIFACT_SERVER_TOKEN="-" #! This is necessary as act will always to expose this env var to the workflow
```
