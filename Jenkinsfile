pipeline {
  agent any
    parameters {
        environment {
            IMAGE_NAME = 'webapp'
            IMAGE_VERSION = '1.0'
        }
    }

  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }

  stages {
    stage('Clone') {
        echo "Cloning the WebApp Repository"
        checkout scm
    }
    
    stage ('Build') {
        steps {
            sh 'docker build -t ${IMAGE_NAME}:${IMAGE_VERSION}'
        }
    }
    stage('Scan') {
      steps {
        sh 'trivy image jenkins/jenkins'
      }
    }
  }
}