pipeline {
    agent any

    environment {
        APP_NAME = "student-app"
    }

    stages {
        stage('Checkout Code') {
            steps {
                checkout scm
            }
        }

        stage('Prepare Files') {
            steps {
                withCredentials([
                    file(credentialsId: 'env-file', variable: 'ENV_FILE'),
                    file(credentialsId: 'pem-file', variable: 'PEM_FILE')
                ]) {
                    sh '''
                    sudo cp $ENV_FILE .env
                    sudo cp $PEM_FILE global-bundle.pem
                    '''
                }
            }
        }

        stage('Build Docker') {
            steps {
                sh 'docker build -t student-app .'
            }
        }

        stage('Run App') {
            steps {
                sh '''
                docker stop student-app || true
                docker rm student-app || true
                docker run -d -p 8082:8082 \
                --env-file .env \
                --name student-app student-app
                '''
            }
        }

        stage('Verify Health') {
            steps {
                sh 'docker ps | grep student-app'
                echo "✅ Pipeline successfully completed."
            }
        }
    }

    post {
        success {
            echo "🎉 Deployment Successful!"
        }
        failure {
            echo "❌ Deployment Failed!"
        }
    }
}