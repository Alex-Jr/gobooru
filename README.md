# logbooru

Projeto de uma galeria de arquivos como imagens, vídeos e documentos.

## Motivação
Eu era um usuário do projeto [szurubooru](https://github.com/rr-/szurubooru). Ele é bem potente, mas está sem manutenção e não tem algumas features que eu preciso. 

Então, me surgiu a ideia de fazer uma solução prática e performática para entusiastas de self-hosted (pessoas que tem servidores em casa) conseguirem organizar seus arquivos e acessá-los facilmente por navegadores ou APIs.



![Homepage](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/home.png?raw=true)


## Funcionalidade

- Cadastro de um "Post" e seus metadados.
- Agrupamento de vários "Post" em uma "Pool".
- Adicionar "Tags" para categorizar "Post" facilitando sua busca.
- Geração de thumbnail automática a partir de vídeos e imagens.
- Deduplicação de arquivos iguais e similares.
- CRUD de todas as entidades (Post, Pool, Tags...).
- Guaxinins!

## Funcionalidades Planejadas

- Demo online
- Sistema de Permissão
- Melhorar a detecção de duplicadas.
- Comentários
- Favoritos
- Score
- Anotações
- Backups automáticos

## Tecnologias
- GoLang + Echo (Back)
- Typescript + React (Front)
- PostgreSQL
- Docker 
- CloudFlare Tunnel

## Como executar?
- Copie o .env.example para .env
- Altere os valores necessários
- Execute o comando
```
docker compose up -d
```

## Telas

### Listagem de Posts
![PostList](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/post-list.png?raw=true)

### Posts
![Post](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/post.png?raw=true)


### Cadastro de Post
![Post](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/post-new.png?raw=true)

### Listagem de Pools
![PoolList](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/pool-list.png?raw=true)

### Pool
![Pool](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/pool.png?raw=true)


### Listagem de tags
![TagList](https://github.com/Alex-Jr/gobooru/blob/df69f959502be5e488d59844ef0ff34f5381cb80/screenshot/tag-list.png?raw=true)
