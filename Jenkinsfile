pipeline{
    agent any

    environment {
        SERVER_CREDENTIALS = credentials('global')
    }

    stages{

        stage("build"){
            steps {
                echo "${SERVER_CREDENTIALS}"
                echo "build application"
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