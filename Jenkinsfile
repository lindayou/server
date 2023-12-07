pipeline {
    agent any
    environment{
    version = 'v4'
    }
    stages {
        stage('拉取git代码') {
            steps {
                git branch: 'main', credentialsId: '941d2cbe-0abe-40ad-87a3-43e4e68825c4', url: 'git@github.com:lindayou/server.git'
            }
        }
        stage('通过go build 构建项目') {
            steps {
                sh '''cd /var/lib/jenkins/workspace/server
                '''
            }
        }
        stage('通过docker制作镜像') {
            steps {
                sh '''cd /var/lib/jenkins/workspace/server
                docker build -t registry.cn-hangzhou.aliyuncs.com/yeppy/zero:${version} .

                '''
            }
        }

         stage('推送到阿里云') {
             steps {
                 sh '''cd /var/lib/jenkins/workspace/server
                    docker push registry.cn-hangzhou.aliyuncs.com/yeppy/zero:${version}
                         '''
                    }
                }
         stage('通知K8s服务器更新镜像') {
                      steps {
                 sshPublisher(publishers: [sshPublisherDesc(configName: 'k8s', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'pubsh-zero.sh $version', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: '', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
                             }
                         }
    }
}
