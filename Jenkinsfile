def gv

pipeline{
    agent any

    tools {
        go "Go 1.16"
    }

    environment {
        CGO_ENABLED = 0
        DOCKER_FILE_PATH = "Dockerfile"
        TESTER = "abcde"
    }
    parameters {
        string(name: 'GO_IMAGE_NAME', defaultValue: '', description: 'go image name')
    }

    stages{

        stage("init") {
            steps {
                script {
                    gv = load "script.groovy"
                }
            }
        }

        stage("test") {
            steps {
                sh "go test -v ./..."
            }
        }

        stage("test env") {
            steps {
                sh "/tmp/script/tester.sh"
            }
        }

        stage("build"){
            steps {
                sh "echo build application"
                sh "go build ."
            }
        }

        // stage("containerize") {
        //     steps {
        //         sh "docker build -f $DOCKER_FILE_PATH -t ${params.GO_IMAGE_NAME} ."

        //         script {
        //             echo gv.buildApp()
        //         }
        //     }
        // }

        stage("deploy") {
            steps {
                echo "deploy application"
            }
        }

        stage("archive") {
            steps {
                echo "Archive artifact"
                sh '''
                    pwd
                    ls -l
                '''
                archiveArtifacts artifacts: 'myapp'
            }
        }
    }
}