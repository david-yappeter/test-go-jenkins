def gv

pipeline{
    agent any

    tools {
        go "Go 1.16"
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
                sh "export CGO_ENABLED=0"
                sh "echo $CGO_ENABLED"
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
    }
}