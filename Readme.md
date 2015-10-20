# Leeroy CI
Leeroy is a self hosted, continuous integration, build and deployment service. It is designed to be easy to setup and will not require an additional ops person to keep running. It runs on your own server, so you can create the test environment you want, exactly mirroring production, without having to trust anyone else to keep your source code or eventually database images with sensitive information safe.

![Leeroy](https://raw.github.com/fallenhitokiri/leeroyci/master/assets/leeroy.png)

## Integrations
Currently Leeroy plays nicely with GitHub. Integrations for GitLab and Bitbucket are planned.

## Features
- bring your own build / test scripts
- comment on GitHub pull requests
- close GitHub pull requests if the build for HEAD fails
- send notifications about the build via email
- post results to a Slack, Campfire or HipChat channel
- see all builds on an acceptable designed - read bootstrap - webinterface
- continuous deployment using your own deployment scripts - deploy to whichever environment you want.
- search for branches, commits and repositories

## Quickstart
For now please check out the master branch of this repository and run it via `go run leeroy.go`. Binaries will be available with the first stable release.
Master is always considered to be stable and ready for production.

### Build Script
Before you start make sure you have a script that is able to run tests for your repository. Two arguments are passed to your build script, the repository URL (first argument) and the branch name (second argument) to which was pushed. Let us use a really simple one for now

     #! /bin/bash
     ls

We assume this script is saved in `/home/ec2-user/test.sh`. See `docs/buildscripts.md` for more information and sample scripts.

### Configuration
To set the path for the SQLite database you can use the environment variable `DATABASE_URL`. The format is `sqlite3 /path/to/leeory.sqlite3`.

Once Leeroy is running go to port `8082` in your web browser and click through the setup assistant. The user you create will automatically be an administrator. If you specify an SSL certificate you have to restart Leeroy after completing the setup.

To configure a repository click on `Admin -> Repository Management -> Add Repository`. After adding the repository you can add commands and notifications on the repository detail page you are redirected to. The access key needs permissions to update the status of your commits, comment on PRs and close them if you want to use that feature.

For a command you can select a kind, name, branch and script to run. If a branch is specified the command will only run when you push to the specific branch. For the script please specify the full path.

The order to run commands is

1. tests
2. builds
3. deploy

The script runner exists on the first failed step, so if tests fail builds and deploys will never run.

On GitHub you have to setup a webhook pointing to `http://yourhost:8082/callback/github/` - or `https://…` if you configured SSL and make sure it sends "push" and "pull" events.

## Planned Features
While Leeroy is working and doing its job it is far from being feature complete. Before version 1.0 will be released the following features will be finished

- GitLab integration
- Bitbucket integration
- support for custom templates and notifications
- website with a browsable documentation and more default snippets

## Contributing
Feel free to open issues about bugs or features you want to see or open pull requests. Beside using `go fmt` and `go vet` on your code please try to keep the code length around 80 characters. This is no hard limit. If a line is 86 characters long but easy to read and understand there is no need to break it into multiple lines.

## License
Leeroy is released under the MIT license.
