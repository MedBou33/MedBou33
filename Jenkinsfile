pipeline {
    agent any

    stages {
        stage('Format'){
            steps {
                sh 'test -z "$(gofmt -l .)"'
            }
        }
        stage('Vet') {
            steps {
                sh 'go vet ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        stage('Build'){
            steps {
                sh 'go build -o bin/app .'
            }
        }
    }
    post {
        always {
            echo 'Pipeline finished'
        }
        success {
            echo 'All stages passed'
        }
        failure {
            echo 'Something failed, check the console output'
        }
    }
}
