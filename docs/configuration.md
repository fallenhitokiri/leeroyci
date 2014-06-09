# Configuration
The following configuration shows all options Leeroy supports. We will go through them one by one. I leave as many values from my test configuration in there as possible. For a minimal configuration you can read the "Quickstart" section in the Readme. If you do not use a feature you do not have to add the configuration options for it.

     {
       "URL": "https://localhost:8082/",
       "Secret": "superdupersecret",
       "Cert": "/Users/tizi/tmp/leeroy.crt",
       "Key": "/Users/tizi/tmp/leeroy.key",
       "BuildLogPath": "/Users/tizi/tmp/build.json",
       "EmailFrom": "Leeroy CI",
       "EmailHost": "smtp.gmail.com",
       "EmailPort": 587,
       "EmailUser": "onlyaspamaccount@gmail.com",
       "EmailPassword": "foo",
       "GitHubKey": "bar",
       "SlackChannel": "#devel",
       "SlackEndpoint": "baz",
       "Repositories": [
         {
           "URL": "https://github.com/fallenhitokiri/pushtest",
           "CommentPR": true,
           "ClosePR": true,
           "Commands": [
             {
               "Name": "pass",
               "Execute": "/Users/tizi/tmp/leeroy/test.sh"
             },
             {
               "Name": "fail",
               "Execute": "/Users/tizi/tmp/leeroy/test2.sh"
             }
           ],
           "Notify": [
             {
               "Service": "slack",
               "Arguments": []
             }
           ]
         }
       ]
     }

- `URL` specify the URL you want Leeroy to use. If you use https as scheme you have to configure a cert and key.
- `Secret` the secret that has to be added to the callback URL. If the secret is wrong no actions will be triggered
- `Cert` and `Key` full path to SSL certificate and key. Required if your URL scheme is `https`
- `BuildLogPath` full path where Leeroy can write the build log to. Every build will be written.
- `Email*` host, part and credentials for your mail server so Leeroy can send you notifications if builds were successful or failed
- `GitHubKey` your personal access token
- `SlackChannel` to which channel should build statuses be posted
- `SlackEndpoint` the endpoint you configured in Slack
- `Repositories` list of all repositories Leeroy will run builds for

#### Repositories
- `URL` URL on which your repository is hosted. It is required that it matches your repository URL or Leeroy will not run any builds
- `CommentPR` if you open a pull request Leeroy will post a comment with the build status for HEAD
- `ClosePR` if the build for HEAD failed Leeroy will close a pull request
- `Commands` list of commands to run when a build is triggered (push / PR)
- `Notify` list of notifications which will be triggered after a build finished

#### Commands
- `Name` a name for the command. Just so you can easily identify what failed if you run multiple commands
- `Execute` script to execute

#### Notify
- `Service` name of the service
- `Arguments` arguments to pass

Currently supported notifications are `slack` and `email`. Slack takes no arguments, email expects a list of email addresses to send the build status to. Remember that the person who pushed or opened the pull request will always get an email, so only configure additional people here.

##### Email
If you want to notify `ops@example.tld` and `devops@example.tld` when someone triggered a build so they immediately know not to deploy a branch you can use the following configuration

     {
       "Service": "email",
       "Arguments": ["ops@example.tld", "devops@example.tld"]
     }

## SSL
If you want to use a self-signed certificate make sure to disable GitHubs SSL verification. You can generate a certificate and key with the following command

      openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout leeroy.key -out leeroy.crt