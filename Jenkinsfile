pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                /* CORRECT */
                echo 'hello world'
                sh('docker build -f Dockerfile -t my_sample_app:latest .')
                sh('docker rm my_app')
                sh('docker run --name=my_app -d -p 9992:9992 my_sample_app:latest')
            }
        }
    }
}