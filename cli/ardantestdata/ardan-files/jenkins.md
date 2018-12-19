# Jenkins

[Jenkins](https://jenkins.io/) is used to build and deploy projects. The exact branches and procedures
may differ a little by project.

Ardan Labs build server is at https://jenkins.ardanlabs.com.

## Credentials

Credentials for github repository and other systems must be added to Jenkins
before projects can be setup.

  - Select *Credentials* from left, then the *(global)* scope. Then *Add Credentials* from left.
  - Put in the username and password. Make the id something that is short and unique.
    Make the description the same.
  - Docker registry credentials must use email as username.
  - AWS credentials use the AccessKey as the username and SecretKey as the password.

## Projects

After credentials have been added then build projects can be created and setup.

  - If there are going to be multiple builds for a project you might want to create a folder.
    Select *New Item -> Folder* from top left..
  - For the actual build project select *New Item -> Multibranch Pipeline* from top righ*New Item -> Multibranch Pipeline* from top right.
  - Add a short name and display name. Make them the same.
  - Add a good description about the project.
  - Add `source` in the next section. Select Bitbucket or Github and then the appropriate
    credential that has been previously setup.
  - Add the behaviors.
    - Discover branches - Exclude branches that are also filed as PRs.
    - Discover pull request from forks
    - Discover pull request from origin
    - Filter by name (with wildcards). Include: \* Exclude: feature/\*
  - The rest of the options can be left alone.

## Github and Bitbucket Webhooks

Webhooks need to be setup for each repository so they are built on merging to master or
pull requests.

### BitBucket

  - In Jenkins
    - *Manage Jenkins -> Configure System*, scroll down to *Bitbucket Endpoints*. Select
      *Manage hooks*, the credentials and save.
  - In BitBucket.
    - Under the repoistory settings add a *Webhook*.
    - Title `Jenkins` and URL `https://jenkins.ardanlabs.com/bitbucket-scmsource-hook/notify`.
    - View all the triggers and select `Merged` along with `Push`.
    - https://support.cloudbees.com/hc/en-us/articles/115000053051-How-to-Trigger-Multibranch-Jobs-from-BitBucket-Server-

### GitHub

  - In GitHub
    - In the repository *Settings -> Add webhook*.
    - For the Jenkins hook url use `https://jenkins.ardanlabs.com/github-webhook/`.
    - The secret is in LastPass under `ardan-jenkins`.

## Nodes

Then Jenkins master node is an AWS EC2 instance that is managed through [Cloud66](https://cloud66.com/)
and runs in Docker.

There is another MacOS node that runs on an old laptop in the office. It issued to build and test
ReactNative projects. The emulators required for integration tests need to run on bare metal.

### Build on MacOS

To have the project built on MacOS simply specify it when defining the `node` in the `Jenkinsfile`
for the project.

```
// Run on the slave with name or tag MacOS.
node('MacOS') {
  // Your stages go here.
  ...
}
```

### Setting Up Slave

Setting up a slave Jenkins node for builds basically entails creating a user on the slave and ssh keys.
The Jenkins master installs the agent automatically via ssh.

  - Copy private key onto Jenkins master at `/var/lib/jenkins/.ssh/id_rsa`. Then add a new
    credential of kind `SSH Username with private key` and select to get the private key from a
    file on master. Use the same `/var/jenkins_home/.ssh/id_rsa` as the path.
  - Do not verify with known_hosts.
  - Add node with name MacOS with short description.
  - Specify the remote root as `/Users/jenkins`
  - Finally press the *Launch Agent*  button.

### Access Server

The Jenkins host is an AWS EC2 instance managed through Cloud66. To ssh into the host the Cloud66
toolbelt is required.

```
cx ssh -s 'Jenkins' -e production dockerhost
```

### Backups

Backups are done weekly and uploaded to S3 https://s3.console.aws.amazon.com/s3/buckets/com.ardanlabs.backups/jenkins/?region=us-east-1&tab=overview
by this container running on the host https://github.com/ardanlabs/cron.
