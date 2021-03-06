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

        // stage("test echo file") {
        //     steps {
        //         sh "/tmp/script/tester.sh"
        //     }
        // }

        stage("deploy") {
            steps {
                // echo "/tmp/script/deploy.sh"
                script {
                    def deployOutput = sh(script: "/tmp/script/deploy.sh", returnStdout: true).trim()
                    println("disk_size = ${deployOutput}")
                }

                // sh "docker-compose up -d"
            }
        }

        // stage("archive") {
        //     steps {
        //         echo "Archive artifact"
        //         sh '''
        //             pwd
        //             ls -l
        //         '''
        //         archiveArtifacts artifacts: 'myapp'
        //     }
        // }
    }

    post {
        success {
            archiveArtifacts artifacts: 'myapp'
        }
    }
}