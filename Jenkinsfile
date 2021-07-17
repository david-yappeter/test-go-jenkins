pipeline{
    agent any

    stages{
        stage("build"){
            steps{
                echo "build application"
            }
        }

        stage("test"){
            step{
                echo "testing application"
            }
        }

        stage("deploy") {
            step{
                echo "deploy application"
            }
        }
    }
}