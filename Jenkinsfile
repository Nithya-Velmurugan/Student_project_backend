pipeline {
    agent any

    environment {
        // App name matching your docker-compose container_name
        APP_NAME = "student-app"
    }

    stages {
        stage('Checkout Code') {
            steps {
                // Pulls the latest code from your Git repository branch
                checkout scm
            }
        }

        stage('Verify Environment') {
            steps {
                // Since .env and global-bundle.pem should NEVER be committed to Git,
                // Jenkins must ensure they exist before trying to use docker-compose.
                // (You can inject them using Jenkins "Credentials Binding Plugin" or manually copy them)
                script {
                    if (!fileExists('.env')) {
                        error "🚨 FATAL ERROR: .env file is missing from the Jenkins workspace!"
                    }
                    if (!fileExists('global-bundle.pem')) {
                        error "🚨 FATAL ERROR: global-bundle.pem is missing from the Jenkins workspace!"
                    }
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
                // Ensure the container actually started up.
                sh 'docker ps | grep student-app'
                echo "hellow your pipeline successfully completed."
            }
        }
    }

    post {
        success {
            echo "🎉 Deployment Successful! Application is listening on AWS port 8080."
        }
        failure {
            echo "❌ Deployment Failed! Check the Jenkins logs in the AWS dashboard."
        }
    }
}
