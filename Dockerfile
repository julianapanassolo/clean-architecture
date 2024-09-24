FROM mysql:latest

# Define o usuário e a senha do banco de dados
ENV MYSQL_ROOT_PASSWORD=social
ENV MYSQL_DATABASE=social

# Crie o diretório de dados
RUN mkdir -p /var/lib/mysql/

# Executa o script de inicialização do banco de dados
CMD ["mysqld"]