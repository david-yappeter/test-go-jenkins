def gv

pipeline{
    agent any

    tools {
        go "Go 1.16"
    }

    environment {
        CGO_ENABLED = 0
        DOCKER_FILE_PATH = "Dockerfile"
        DOCKER_PASS = credentials("DOCKER_PASS")
    }
    parameters {
        string(name: 'GO_IMAGE_NAME', defaultValue: '', description: 'go image name')
        string(name: 'GO_IMAGE_TAG', defaultValue: 'latest', description: 'go image tag')
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

        stage("build"){
            steps {
                sh "echo build application"
                sh "go build ."
            }
        }

        stage("containerize") {
            steps {
                sh "docker build -f $DOCKER_FILE_PATH -t ${params.GO_IMAGE_NAME}:${params.GO_IMAGE_TAG} ."

                script {
                    echo gv.buildApp()
                }
            }
        }

        stage("docker push") {
            steps {
                sh "/tmp/script/push.sh"
            }
        }

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