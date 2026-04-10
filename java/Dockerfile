FROM maven:3.9.9-eclipse-temurin-17

WORKDIR /app

COPY src/ /app/

EXPOSE 8080

CMD ["sh", "-c", "if [ -f mvnw ]; then chmod +x mvnw && ./mvnw spring-boot:run; elif [ -f pom.xml ]; then mvn spring-boot:run; else echo 'Add your Spring Boot project under java/src and start again.'; tail -f /dev/null; fi"]
