#!/bin/bash

# Copia .env
[ -f .env ] || cp .env.example .env

# Instalar dependências usando o Bun
bun install

# Iniciar o servidor de desenvolvimento
bun run dev --host 0.0.0.0
