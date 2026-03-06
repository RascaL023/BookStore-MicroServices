consul agent -dev &

sleep 3

consul services register ./writer-service.json
