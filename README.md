# API.AI - Discord Bridge BOT

```
$ gcloud compute --project "pubg-assistant" instance-templates create "api-ai-discord-bot-template" \
    --machine-type "f1-micro" \
    --metadata-from-file startup-script=startup.sh \
    --maintenance-policy "TERMINATE" --preemptible \
    --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring.write","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
    --image "cos-beta-60-9592-23-0" --image-project "cos-cloud" --boot-disk-size "10" --boot-disk-type "pd-standard" --boot-disk-device-name "api-ai-discord-bot-template"

$ gcloud compute --project "pubg-assistant" instance-groups managed create "api-ai-discord-bot" \
    --zone "asia-east1-a" --base-instance-name "api-ai-discord-bot" --template "api-ai-discord-bot-template" --size "1"
```

```
gcloud compute --project "pubg-assistant" instance-groups managed delete "api-ai-discord-bot" --zone "asia-east1-a" -q && \
gcloud compute --project "pubg-assistant" instance-templates delete "api-ai-discord-bot-template" -q
```
