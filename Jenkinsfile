pipeline {
    agent any

    environment {
        // Здесь Jenkins будет использовать значение переменной окружения,
        // которое должно быть предварительно задано в Jenkins
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
                sh 'python3 -m venv venv'
                sh '. venv/bin/activate'
                sh 'pip install -r requirements.txt'
            }
        }

        stage('Get Weather') {
            steps {
                script {
                    // Запуск скрипта с использованием переменной окружения
                    sh 'python3 get_weather.py'
                }
            }
        }

        stage('Archive') {
            steps {
                // Архивация файла weather_report.txt
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
