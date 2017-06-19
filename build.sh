#! /bin/bash -eux

docker build -t gcr.io/pubg-assistant/api-ai-discord-bot .
gcloud docker -- push gcr.io/pubg-assistant/api-ai-discord-bot
# docker run -e DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN -e APIAI_DEVELOPER_ACCESS_TOKEN=$APIAI_DEVELOPER_ACCESS_TOKEN gcr.io/pubg-assistant/api-ai-discord-bot
