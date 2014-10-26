# Leeroy CI
Leeroy is a self hosted, continuous integration, build and deployment service. It is designed to be easy to setup and will not require an additional ops person to keep running. It runs on your own server, so you can create the test environment you want, exactly mirroring production, without having to trust anyone else to keep your source code or eventually database images with sensitive information safe.

![Leeroy](https://raw.github.com/fallenhitokiri/leeroyci/master/assets/leeroy.jpg)

## Integrations
Currently Leeroy plays nicely with GitHub. Integrations for GitLab and Bitbucket are planned.

## Features
- bring your own build / test scripts
- comment on GitHub pull requests
- close GitHub pull requests if the build for HEAD fails
- send notifications about the build via email
- post results to a Slack, Campfire or HipChat channel
- see all builds on an acceptable designed, read bootstrap, webinterface.
- get all builds, branchs or single commits as JSON
- continuous deployment

## Quickstart
For now please check out the master branch of this repository and run it via `go run leeroy.go`. Binaries will be available in some days.
Master is always considered to be stable and ready for production.

### Build Script
Before you start make sure you have a script that is able to run tests for your repository. Two arguments are passed to your build script, the repository URL (first argument) and the branch name (second argument) to which was pushed. Let us use a really simple one for now

     #! /bin/bash
     ls

We assume this script is saved in `/home/ec2-user/test.sh`. See `docs/buildscripts.md` for more information and sample scripts.

### Configuration
If you put your configuration in the same directory as the executable you can name it `leeroy.json` and it will be found automatically. If you put it somewhere else you can specify a configuration file with the `-config` flag.

#### GitHub
You need a GitHub personal access token and a test repository.

In your repository setup a webhook pointing to `http://yourhost:8082/callback/github/superdupersecret`. Content type should be `application/json` and as events select `Push` and `Pull Request`.

You can obtain a personal access token under the `Applications` menu in your settings. Permissions to access your repositories is enough.

We use the most basic configuration for now

     {
       "URL": "http://0.0.0.0:8082/",
       "Secret": "superdupersecret",
       "BuildLogPath": "/tmp/build.json",
       "GitHubKey": "foobar",
       "Repositories": [
          {
            "URL": "https://github.com/fallenhitokiri/pushtest",
            "Commands": [
                {
                    "Name": "test.sh",
                    "Execute": "/home/ec2-user/test.sh"
                }
            ]
          }
       ]
     }

Make sure to replace the GitHubKey with your personal access token. For more details and all configuration option please read `docs/configuration.md`.

### Webinterface
When you run Leeroy and push to the repository you should new see a build when visiting the webinterface. There are several views which show you the status of your various builds.

![success](https://raw.github.com/fallenhitokiri/leeroyci/master/docs/success.png)

- `/status/` lists all builds in a chronological order, newest first.
- `/status/repo/<hex>` lists all builds for a repository - `<hex>` is the repository
URL in hexadecimal.
- `/status/branch/<hex>/<branch>` lists all builds for a specific branch of a
repository.
- `/status/commit/<hex>/<sha1>` shows the build for a specifc commit of a repository.
- `/status/badge/<hex>/<branch>` returns a SVG with the status of the last build for a specific branch

By appending `?format=json` to one of the URLs a JSON object will be returned making it easy to integrate Leeroy with other tools.

## Planned Features
While Leeroy is working and doing its job it is far from being feature complete. Before version 1.0 will be released the following features will be finished

- GitLab integration
- Bitbucket integration
- support for custom templates and notifications
- user authentication (status pages)
- configuration of repositories through the webinterface
- website with a browsable documentation and more default snippets

## Contributing
Feel free to open issues about bugs or features you want to see or open pull requests. Beside using `go fmt` and `go vet` on your code please try to keep the code length around 80 characters. This is no hard limit. If a line is 86 characters long but easy to read and understand there is no need to break it into multiple lines.

## License
Leeroy is released under the MIT license.
