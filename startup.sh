#! /bin/bash -eux

# Do you want to check the startup script log?
# sudo journalctl -u google-startup-scripts.service

export DISCORD_BOT_TOKEN=$(curl http://metadata/computeMetadata/v1/project/attributes/discord-bot-token -H "Metadata-Flavor: Google")
export APIAI_DEVELOPER_ACCESS_TOKEN=$(curl http://metadata/computeMetadata/v1/project/attributes/apiai-developer-access-token -H "Metadata-Flavor: Google")

# METADATA=http://metadata.google.internal/computeMetadata/v1
# SVC_ACCT=$METADATA/instance/service-accounts/default
# ACCESS_TOKEN=$(curl -H 'Metadata-Flavor: Google' $SVC_ACCT/token | cut -d'"' -f 4)
# docker login -u _token -p $ACCESS_TOKEN https://gcr.io

# avoid unbound variable error
export HOME=/home/vvakame
mkdir -p $HOME
/usr/share/google/dockercfg_update.sh

docker run --rm -e DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN -e APIAI_DEVELOPER_ACCESS_TOKEN=$APIAI_DEVELOPER_ACCESS_TOKEN gcr.io/pubg-assistant/api-ai-discord-bot:latest
