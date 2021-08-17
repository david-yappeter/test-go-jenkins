def gv

pipeline{
    agent any

    tools {
        go "Go 1.16"
    }

    environment {
        CGO_ENABLED = 0
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

                script {
                    echo gv.buildApp()
                }
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
                archiveArtifacts artifacts: './myapp'
            }
        }
    }
}