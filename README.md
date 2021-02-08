# Continuous Deployment Golang

System to use github actions to continuously deploy new versions of your Golang code.

Here is an example of the release type that will be done based on a commit messages:

**Deploying**

In our yml files samples, the idea is that if you commit new features or fixes to features branches, it will bump the patch version 1 value, unless you specify #minor or #major test on the commit message, then it will bump the minor or the major value.

If you commit to beta, it will bump the version as in features branches but also will compile and build the new executable into a new docker image. Also the executable compiled will also have build in the new version number that you can check on the right url. Also after it we will connect ssh to the server and re build and restart the docker compose running the images on the BETA server.

If you commit to master, will happen all the same as beta, but connection by ssh to the production server.

**Debugging**

Also by default contains the system to be able to debug inside of the docker container thanks to delve

**Bumping**
 1. Manual Bumping: Any commit message that includes #major, #minor, or
    #patch will trigger the respective version bump. If two or more are present, the highest-ranking one will take precedence.
 2. Automatic Bumping: If no #major, #minor or #patch tag is contained
    in the commit messages, it will bump whichever DEFAULT_BUMP is set
    to (which is minor by default). Disable this by setting DEFAULT_BUMP
    to none.
