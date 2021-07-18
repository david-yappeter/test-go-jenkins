def gv

pipeline{
    agent any

    parameters{
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

        stage("build"){
            steps {
                echo "build application"
                sh "go run ."

                script {
                    echo gv.buildApp()
                }
            }
        }

        stage("test"){
            steps {
                echo "testing application"
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