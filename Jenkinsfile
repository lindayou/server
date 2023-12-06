pipeline {
    agent any
    stages {
        stage('拉取git代码') {
            steps {
                echo 'mvn clean install'
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