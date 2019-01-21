#!/usr/bin/env bash
gcloud beta compute --project=uttara-1475266998216 instance-templates create go-template --machine-type=f1-micro --network=projects/uttara-1475266998216/global/networks/default --network-tier=PREMIUM --maintenance-policy=MIGRATE --service-account=213393127979-compute@developer.gserviceaccount.com --scopes=https://www.googleapis.com/auth/cloud-platform --tags=http-server,https-server --image=debian-9-stretch-v20181210 --image-project=debian-cloud --boot-disk-size=10GB --boot-disk-type=pd-standard --boot-disk-device-name=go-template