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
                    if (output.contains("foo")) {
                        env.foo = "true"
                    }
                    if (output.contains("bar")) {
                        env.bar = "true"
                    }
                }
            }
        }
        stage('Deploy Foo') {
            when { environment name: 'foo', value: 'true' }
            steps {
                sparseCheckout('foo')
                echo 'Deploy Foo'
                sh 'ls -all'                
            }
        }
        stage('Deploy Bar') {
            when { environment name: 'bar', value: 'true' }
            steps {
                sparseCheckout('bar')
                echo 'Deploy Bar'
                sh 'ls -all'
            }
        }
    }
}

def sparseCheckout(sparsePath) {
  def scmVars = checkout scm
  def commitHash = scmVars.GIT_COMMIT
  
  checkout scm: [
    $class: 'GitSCM', 
    branches: [[name: commitHash]], 
    doGenerateSubmoduleConfigurations: false, 
    extensions: [[
      $class: 'SparseCheckoutPaths', 
      sparseCheckoutPaths: [[
        path: sparsePath
      ]]
    ]], 
      submoduleCfg: [], 
      userRemoteConfigs: [[
        credentialsId: "${CREDENTIALSID}", 
        url: "${GIT_REPO}"
      ]]
  ]
}
