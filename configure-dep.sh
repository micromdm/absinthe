#!/bin/bash

sudo defaults write com.apple.ManagedClient.cloudconfigurationd MCCloudConfigSessionURL http://127.0.0.1:8000/session
sudo defaults write com.apple.ManagedClient.cloudconfigurationd MCCloudConfigProfileURL http://127.0.0.1:8000/profile
sudo defaults write com.apple.ManagedClient.cloudconfigurationd MCCloudConfigCertificateURL http://127.0.0.1:8000/certificate.cer
