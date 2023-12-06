pipeline {
    agent any
    stages {
        stage('拉取git代码') {
            steps {
                sh 'mvn clean install'
            }
        }
        stage('通过go build 构建项目') {
            steps {
                sh 'mvn test'
            }
        }
        stage('通过docker制作镜像') {
            steps {
                sh 'mvn deploy'
            }
        }
         stage('推送到阿里云') {
                    steps {
                        sh 'mvn deploy'
                    }
                }
    }
}