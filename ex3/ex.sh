if curl -s -o /dev/null -w "%{http_code}" https://example.com | grep -q "200"; then
    echo "Site accessible"
else
    echo "Site indisponible"
fi