pipeline{
    agent any

    tools {
        go "Go 1.16"
    }

    stages{

        stage("build"){
            steps {
                echo "build application"
                sh "go run ."
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
            }
        }
    }

    post{
        always{
            emailext body: 'A Test EMail', recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']], subject: 'Test'
        }
    }
}