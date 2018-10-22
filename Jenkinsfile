#!groovy

pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "Building ${GIT_COMMIT}"
                script {
                    output = sh(returnStdout: true, script: 'git diff --name-only $GIT_PREVIOUS_COMMIT $GIT_COMMIT | cut -d "/" -f 1 | sort | uniq').trim()
                    output.eachLine { line ->
                        println line
                        if (line == "foo"){
                            env.foo = "true"
                        }
                        if (line == "bar"){
                            env.bar = "true"
                        }
                    }
                }
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
