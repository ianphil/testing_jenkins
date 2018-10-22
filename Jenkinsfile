#!groovy

pipeline {
    agent any
    environment {
        GIT_REPO = 'git@github.com:iphilpot/testing_jenkins.git'
        CREDENTIALSID = 'JenkinsSSH'
    }
    stages {
        stage('Build') {
            steps {
                echo "Building ${GIT_COMMIT}"
                script {
                    output = sh(returnStdout: true, script: 'git diff --name-only $GIT_PREVIOUS_COMMIT $GIT_COMMIT | cut -d "/" -f 1 | sort | uniq').trim()
                    outputClass = output.getClass()
                }
                echo "output=$outputClass";
            }
        }
        stage('Test') {
            steps {
                echo 'Testing'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying'
            }
        }
    }
}

def sparseCheckout() {
  def scmVars = checkout scm
  def commitHash = scmVars.GIT_COMMIT
  
  checkout scm: [
    $class: 'GitSCM', 
    branches: [[name: commitHash]], 
    doGenerateSubmoduleConfigurations: false, 
    extensions: [[
      $class: 'SparseCheckoutPaths', 
      sparseCheckoutPaths: [[
        path: 'bar'
      ]]
    ]], 
      submoduleCfg: [], 
      userRemoteConfigs: [[
        credentialsId: "${CREDENTIALSID}", 
        url: "${GIT_REPO}"
      ]]
  ]
}
