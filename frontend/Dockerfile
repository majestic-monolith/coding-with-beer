# Usar a imagem oficial do Node.js
FROM node:22.14

WORKDIR /app

# Instalar Bun
RUN apt-get update && apt-get install -y curl unzip
RUN curl -fsSL https://bun.sh/install | bash
ENV BUN_INSTALL="/root/.bun"
ENV PATH="$BUN_INSTALL/bin:$PATH"

# Expor a porta 5173 (porta padrão do Vite)
EXPOSE 5173

# Comando para executar o script start.sh
CMD ["sh", "/app/start.sh"] 