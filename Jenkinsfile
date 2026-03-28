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
                    cp $ENV_FILE .env
                    cp $PEM_FILE global-bundle.pem
                    '''
                }
            }
        }

        stage('Build Docker') {
            steps {
                sh 'sudo docker build -t student-app .'
            }
        }

        stage('Run App') {
            steps {
                sh '''
                sudo docker stop student-app || true
                sudo docker rm student-app || true
                sudo docker run -d -p 8082:8082 \
                --env-file .env \
                --name student-app student-app
                '''
            }
        }

        stage('Verify Health') {
            steps {
                sh 'sudo docker ps | grep student-app'
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