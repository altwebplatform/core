# Contributing to AltWebPlatform

## Development Environment

[Install minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) and run it: 

`minikube start`

Save your access credentials locally (in the altwebplatform/core directory):

`kubectl config view > .kubeconfig`

[Install CockroachDB](https://www.cockroachlabs.com/docs/install-cockroachdb.html) and run it: 

`cockroach start --http-port 9080 --insecure --background`
`cockroach sql --insecure -d "root@localhost:26257" -e "CREATE DATABASE altwebplatform"`

Install glide:

`brew install glide`

Start AWP: 

`go run main.go`

### Other Considerations

- To add or update a Go dependency:
  - `glide get github.com/user/package` will fetch it and add to glide.yaml.  Sometimes running `glide install` will help `glide get` function.
  - `git add` any new files
  - Create a PR with all the changes.

## Style Guide

[Style Guide](STYLE.md)

## Code Review Workflow

+ All contributors need to sign the [Contributor License Agreement](https://cla-assistant.io/altwebplatform/core).

+ Create a local feature branch to do work on, ideally on one thing at a time.
  If you are working on your own fork, see [this tip](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html)
  on forking in Go, which ensures that Go import paths will be correct.

  `git checkout -b update-readme`

+ Hack away and commit your changes locally using `git add` and `git commit`.
  Remember to write tests! The following are helpful for running specific
  subsets of tests:

  ```shell
  sh scripts/test.sh
  ```

  When you're ready to commit, be sure to write a Good Commit Message™. Consult
  https://github.com/erlang/otp/wiki/Writing-good-commit-messages if you're
  not sure what constitutes a Good Commit Message™.
  In addition to the general rules referenced above, please also prefix your
  commit subject line with the affected package, if one can easily be chosen.
  For example, the subject line of a commit mostly affecting the
  `web/` package might read: "web/: made great again".
  Commits which affect many packages as a result of a shared dependency change
  should probably begin their subjects with the name of the shared dependency.
  Finally, some commits may need to affect many packages in a way which does
  not point to a specific package; those commits may begin with "*:" or "all:"
  to indicate their reach.

+ Run the linters, code generators, and unit test suites locally:

  ```
  sh scripts/test.sh
  ````

  This will take several minutes.

+ When you’re ready for review, groom your work: each commit should pass tests
  and contain a substantial (but not overwhelming) unit of work. You may also
  want to `git fetch origin` and run
  `git rebase -i --exec "make lint test" origin/master` to make sure you're
  submitting your changes on top of the newest version of our code. Next, push
  to your fork:

  `git push -u <yourfork> update-readme`

+ Then [create a pull request using GitHub’s UI](https://help.github.com/articles/creating-a-pull-request). If you know of
  another GitHub user particularly suited to reviewing your pull request, be
  sure to mention them in the pull request body. If you possess the necessary
  GitHub privileges, please also [assign them to the pull request using
  GitHub's UI](https://help.github.com/articles/assigning-issues-and-pull-requests-to-other-github-users/).
  This will help focus and expedite the code review process.

+ Address feedback by amending your commits. If your change contains multiple
  commits, address each piece of feedback by amending that commit to which the
  particular feedback is aimed. Wait (or ask) for new feedback on those
  commits if they are not straightforward. An `LGTM` ("looks good to me") by
  someone qualified is usually posted when you're free to go ahead and merge.
  Most new contributors aren't allowed to merge themselves; in that case, we'll
  do it for you.
