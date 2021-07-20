def gv

pipeline{
    agent any

    parameters{
        string(name: "PARAM_A", defaultValue: "0", description: "")
        string(name: "PARAM_B", defaultValue: "0", description: "")
        choice(name: "VERSION", choices:["1.1.0", "1.1.1", "1.1.2"], description: "")
    }

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
                echo "build application"
                sh "go run . ${params.PARAM_A} ${params.PARAM_B}"

                script {
                    echo gv.buildApp()
                }
            }
        }

        stage("deploy") {
            steps {
                echo "deploy application"
                echo "deploying with ${params.VERSION}"
            }
        }
    }

    post{
        always{
            emailext body: 'A Test EMail', recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']], subject: 'Test'
        }
    }
}