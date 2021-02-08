echo 'Restarting docker compose...'
ssh ec2-user@apidev.circles.berlin "
ls -la
whoami
"
echo 'Restarting docker compose completed successfully'
echo "docker-compose -f docker-compose-download.yml down -v && docker-compose rm -f && docker-compose pull && docker-compose up -d && curl localhost"
echo "docker-compose -f docker-compose-download.yml down -v && docker-compose -f docker-compose-download.yml rm -f && docker-compose -f docker-compose-download.yml pull && docker-compose -f docker-compose-download.yml up -d && curl localhost"
