pipeline{
    agent any

    stages{

        stage("build"){
            steps {
                echo env
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