pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building'
                sh '''
                    $DIFF=$(git diff --name-only $GIT_PREVIOUS_COMMIT $GIT_COMMIT)
                    awk '{print $1}' $DIFF
                '''
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
