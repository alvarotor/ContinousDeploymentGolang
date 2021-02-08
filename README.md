# Continous Deployment Golang

System to use github actions to continuously deploy new versions of your Golang code.

Here is an example of the release type that will be done based on a commit messages:

Deploying
In our yml files smaples, the idea is that if you commit new features or fixes to features branch, it will bump the patch version 1 value, unless you specify #minor or #major values, if you commit to beta, it will bump the version as in features

Bumping
Manual Bumping: Any commit message that includes #major, #minor, or #patch will trigger the respective version bump. If two or more are present, the highest-ranking one will take precedence.

Automatic Bumping: If no #major, #minor or #patch tag is contained in the commit messages, it will bump whichever DEFAULT_BUMP is set to (which is minor by default). Disable this by setting DEFAULT_BUMP to none.
