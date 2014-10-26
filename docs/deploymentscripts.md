# Deployment scripts
Leeroy requires you to bring your own deployment scripts. This is actually pretty easy and allows you to run whatever you want. The output of your scripts will be saved to the build log and displayed on the webinterface.

If your scripts work when you run them manually it should also "just work" with Leeroy.

Your scripts get two arguments. The first one is the repository URL, the second one the branch name to which was pushed.

## Deploying to AWS Elastic Beanstalk
The following scripts demonstrates one way to deploy to AWS EB

     #! /bin/bash
     cd /home/ec2-user/test-deploy
     git fetch
     get reset --hard origin/master
     git aws.push
