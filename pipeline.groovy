pipeline {
    
    agent any 
    
    tools { 
        go "Default"
    }
    
    stages {
        stage('Checkout') { 
            steps {
                git branch: 'main', credentialsId: 'github', url: 'https://github.com/my-team-it/stunning-adventure'
            }
        }
        stage('Build') { 
            steps {
                //Build docker image based on Dockerfile
                //Push image to container registry
                sh 'docker build -t [docker-hub-username]/grinfield:latest .'
                sh 'docker push [docker-hub-username]/grinfield'
            }
        }
        stage('Test') { 
            steps {
                sh 'make test'
            }
        }
        stage('Security') { 
            steps {
                //Scan security with docker scan command
                echo 'Security Scan'
                sh 'docker scout cves [docker-hub-username]/grinfield'
            }
        }
        stage('Deploy') { 
            steps {
                echo 'Deploy'
            }
        }
        stage('Monitor') { 
            steps {
                echo 'Monitor in grafana server'
            }
        }
    }
}