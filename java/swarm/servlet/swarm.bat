ECHO OFF
SET port=%1

IF "%1"=="" (
   SET port=8080
)

ECHO ON
call java -Xms512m -Xmx512m -Dswarm.http.port=%port% -jar target/wildfly-swarm-servlet-0.0.1-SNAPSHOT-swarm.jar

