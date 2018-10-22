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
                    output.eachLine { line ->
                        println line
                        switch (line) {
                            case "foo":
                                env.foo = "true"
                            case "bar":
                                env.bar = "true"
                            default:
                                env.foobar = "true"
                        }
                    }
                }
                // echo "output=$outputClass";
            }
        }
        stage('Deploy Foo') {
            when { environment name: 'foo', value: 'true' }
            steps {
                echo 'Deploy Foo'
            }
        }
        stage('Deploy Bar') {
            when { environment name: 'bar', value: 'true' }
            steps {
                echo 'Deploy Bar'
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
