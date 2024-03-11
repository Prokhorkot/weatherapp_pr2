pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }

    environment {
        WEATHER_API_KEY = credentials('WEATHER_API_KEY')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Test') {
            steps {
                sh 'go version'
                sh 'go test ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o weatherapp'
            }
        }

        stage('Archive Binary') {
            steps {
                archiveArtifacts artifacts: 'weatherapp', onlyIfSuccessful: true
            }
        }
    }

    post {
        success {
            echo 'Бинарный файл успешно создан и архивирован.'
        }
        failure {
            echo 'Не удалось создать или архивировать бинарный файл.'
        }
    }
}
