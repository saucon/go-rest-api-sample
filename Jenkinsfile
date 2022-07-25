pipeline {
    agent any
    environment {
            ENV_GORESTSAMPLE = credentials('go-rest-api-sample')
        }
    stages {
        stage('Example') {
            steps {
                echo 'hello world'
                sh('ls -al')
                sh('cat .env')
                sh('cp \$ENV_GORESTSAMPLE .')
                sh('cat .env')
                sh('docker build -f Dockerfile -t my_sample_app:latest .')
                sh('docker rm -f my_app')
                sh('docker run --name=my_app -d -p 9992:9992 my_sample_app:latest')
            }
        }
    }
}