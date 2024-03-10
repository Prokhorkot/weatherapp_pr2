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

        stage('Setup') {
            steps {
                sh '''
                python3 -m venv venv
                . venv/bin/activate
                pip install -r requirements.txt
                '''
            }
        }

        stage('Get Weather') {
            steps {
                script {
                    sh 'python3 weatherapp.py'
                }
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'weather_report.txt', onlyIfSuccessful: true
            }
        }
    }

    post {
        success {
            echo 'Погодный отчет успешно создан.'
        }
        failure {
            echo 'Не удалось создать погодный отчет.'
        }
    }
}
