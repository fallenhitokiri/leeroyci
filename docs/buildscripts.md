# Buildscripts
Leeroy requires you to bring your own build scripts. This is actually pretty easy and allows you to run whatever you want. The output of your scripts will be saved to the build log and displayed on the webinterface.

If your scripts works when you run it manually it should also "just work" with Leeroy. For GitHub you have to make sure to add your SSH key to the deployment keys of your repository if it is a private one.

Your scripts get two arguments. The first one is the repository URL, the second one the branch name to which was pushed.

## Django
This is a script I use to run tests for a Django project.

     #! /bin/bash		
     cd /home/ec2-user/test
     git fetch
     git checkout $2
     git pull
     source /home/ec2-user/test/.env
     python /home/ec2-user/test/manage.py test