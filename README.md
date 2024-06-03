### Objetivo

Converter vídeos no formato MP4 em um formato mais adequado para o playback de vídeos na Internet: MPEG-DASH.

### Fluxo

- Recebe uma mensagem via RabbitMQ informando qual vídeo deve ser convertido.
- Faz download do vídeo no Google Cloud Storage
- Fragmenta o Vídeo
- Converte o vídeo para MPEG-DASH
- Faz o upload do vídeo no Google Cloud Storage
- Envia uma notificação via fila com as informações do vídeo convertido ou informando erro na conversão.
- Em caso de erro, a mensagem orginal enviada vi RabbitMQ será rejeitada e encaminhada diretamente a uma Dead Letter Exchange.

### Pontos Importantes

- Sistema processa diversas mensagens de forma paralela / concorrente.
- Um simples MP4 quando convertido para MPEG-DASH é segmentado em múltiplos arquivos de áudio e vídeo, logo o processo de upload não é apenas um único arquivo.
- O processo de upload também ocorre de forma concorrente.