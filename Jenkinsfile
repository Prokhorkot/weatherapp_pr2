pipeline {
    agent any

    environment {
        WEATHER_API_KEY = credentials('WEATHER_API_KEY')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh '''
                go build -o weatherapp
                '''
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
