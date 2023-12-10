#!/bin/sh

container_name="mpc-db"

if [ "$( docker container inspect -f '{{.State.Running}}' $container_name )" = "true" ]; then
    echo "Dumping database..."
    docker compose exec database sh -c 'mariadb-dump -uroot -ple_mdp_de_ouf --databases mpc-db > dump.sql'
    echo "Dump done"
else
    echo "Container mpc-db is not running, please run it first (docker compose up -d database)" >&2
fi
