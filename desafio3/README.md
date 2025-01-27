# Executando o Comando `docker-compose up -d`

Para iniciar os serviços definidos no arquivo `docker-compose.yml` em segundo plano, siga os passos abaixo:

1. **Navegue até o diretório do projeto**:
    ```sh
    cd /home/wgreis/go/src/posgoexpert/desafio3
    ```

2. **Execute o comando `docker-compose up -d`**:
    ```sh
    docker-compose up -d
    ```

3. **Verifique se os contêineres estão em execução**:
    ```sh
    docker-compose ps
    ```

4. **Compile o serviço `service_a` (se necessario)**:
    ```sh
    cd desafio3/service_a
    go build -o ./bin/service_a ./cmd/api/main.go
    ```

5. **Compile o serviço `service_b` (se necessario)**:
    ```sh
    cd desafio3/service_b
    go build -o ./bin/service_b ./cmd/api/main.go
    ```

6. **Acesse o Zipkin para rastreamento distribuído**:
    Abra seu navegador e vá para [http://localhost:9411](http://localhost:9411) para acessar a interface do Zipkin.
