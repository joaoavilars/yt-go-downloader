# 🎥 ytbaixar

**ytbaixar** é uma ferramenta de linha de comando desenvolvida em Go para Windows, que permite baixar vídeos e áudios do YouTube de forma simples e interativa, utilizando o `yt-dlp` e `ffmpeg` como base.

---

## 🚀 Funcionalidades

- Baixa vídeos em MP4 com opções de qualidade (HD, Full HD, 4K).
- Extrai áudios em MP3 com qualidade de 128kbps, 196kbps ou 320kbps.
- Interface no terminal interativa: você escolhe o formato ao rodar o comando.
- Nome do arquivo de saída com o título do vídeo original.

---

## 💻 Como usar

### 1. Baixe o executável compilado

Você precisa ter os seguintes arquivos na mesma pasta:

- `ytbaixar.exe` → o programa em Go
- `yt-dlp.exe` → ferramenta para baixar do YouTube
- `ffmpeg.exe` → para extrair e converter áudio em MP3

### 2. Execute no terminal:

```bash
ytbaixar.exe https://www.youtube.com/watch?v=EXEMPLO123
``` 
O programa irá exibir os formatos disponíveis:

1. mp4 - HD (22)
2. mp4 - FHD (137+140)
3. mp3 - 128kbps (140)
4. mp3 - 320kbps (251)


📦 Requisitos
Sistema: Windows

yt-dlp - https://github.com/yt-dlp/yt-dlp/releases

ffmpeg - https://ffmpeg.org/download.html

## 🛠️ Compilando o projeto
Se quiser compilar o projeto manualmente com Go:

1. Clone o repositório

2. Baixe os binários e coloque na pasta:
yt-dlp.exe
ffmpeg.exe

3. Compile com Go

```bash
go build -o ytbaixar.exe main.go
```

4. Estrutura do projeto:
```bash
ytbaixar/
├── main.go
├── ytbaixar.exe
├── yt-dlp.exe
└── ffmpeg.exe
```