pipeline {
    agent any
    stages {
        stage('拉取git代码') {
            steps {
                git branch: 'main', credentialsId: '941d2cbe-0abe-40ad-87a3-43e4e68825c4', url: 'git@github.com:lindayou/server.git'
            }
        }
        stage('通过go build 构建项目') {
            steps {
                echo 'mvn test'
            }
        }
        stage('通过docker制作镜像') {
            steps {
                echo 'mvn deploy'
            }
        }
         stage('推送到阿里云') {
                    steps {
                        echo 'mvn deploy'
                    }
                }
    }
}