## Configuring Notifications
Setting up notifications is relatively easy, you just have to make sure to add all arguments a service expects and stick with the format `key:::value:::::key::value`.

##### Email
No need for any arguments. It will use the mailserver that is already configured.

##### Slack
Slack expects two arguments:

- `channel` channel to post to
- `endpoint` your Slack notification endpoint

##### HipChat
HipChat expects two arguments:

- `channel` channel to post to
- `key` access key for HipChat

##### Campfire
Campfire expects three arguments:

- `id` your Campfire ID
- `room` room to post to
- `key` access key for Campfire

## SSL
If you want to use a self-signed certificate make sure to disable GitHubs SSL verification. You can generate a certificate and key with the following command

      openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout leeroy.key -out leeroy.crt

You can add it in `Admin - Leeroy Config` if you did not setup SSL during the initial setup. You have to restart Leeroy to activate SSL.