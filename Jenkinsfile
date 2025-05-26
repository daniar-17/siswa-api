pipeline {
    agent any

    environment {
        IMAGE_NAME = "daniar17/siswa-api"
        IMAGE_TAG = "latest"
    }

    stages {
        stage('Cleanup') {
            steps {
                cleanWs()
            }
        }

        stage('Checkout') {
            steps {
                // Pakai git langsung
                git url: 'https://github.com/daniar-17/siswa-api.git', branch: 'main'
            }
        }
        
        stage('Copy .ENV') {
            steps {
                withCredentials([file(credentialsId: 'siswa-api-env', variable: 'ENV_FILE')]) {
                    sh 'cp $ENV_FILE .env'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME:$IMAGE_TAG .'
            }
        }

        stage('Push Docker Image') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'docker-hub-credentials',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh '''
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $IMAGE_NAME:$IMAGE_TAG
                    '''
                }
            }
        }
        
        stage('Deploy') {
            steps {
                sh '''
                    docker container stop golang-app || true
                    docker container rm golang-app || true
                    docker image rm $IMAGE_NAME:$IMAGE_TAG || true
                    docker pull $IMAGE_NAME:$IMAGE_TAG
                    docker-compose up -d app
                '''
            }
        }

    }

    // Hapus post/slackSend jika plugin Slack belum terinstall
    post {
        success {
            echo 'Build sukses!'
        }
        failure {
            echo 'Build gagal!'
        }
    }
}