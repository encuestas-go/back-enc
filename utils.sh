#!/bin/bash

# SeLinux
sudo chcon -Rt svirt_sandbox_file_t ./mysql-data

#zip
# zip -r 2024-06-11_database.zip mysql-data/
# mv 2024-06-11_database.zip backups
