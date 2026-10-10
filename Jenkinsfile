pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Pipeline is running'
            }
        }
        stage('List files') {
            steps {
                sh 'ls -la'
            }
        }
    }
}
