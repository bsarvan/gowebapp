pipeline {
  agent any
    
  environment {
      IMAGE_NAME = 'webapp'
      IMAGE_VERSION = '1.0'
  }


  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }

  stages {
    stage('Clone') {
        steps {
            echo "Cloning the WebApp Repository"
            checkout scm
        }
    }
    
    stage('Build') {
        steps {
            sh 'sudo docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .'
        }
    }
    stage('Scan') {
      steps {
        sh 'trivy image ${IMAGE_NAME}:${IMAGE_VERSION}'
      }
    }
  }
}
