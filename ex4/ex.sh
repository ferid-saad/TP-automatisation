USAGE=$(df / | awk 'NR==2 {print $5}' | sed 's/%//')
if [ $USAGE -gt 90 ]; then
    echo "Espace disque faible (<10%)"
else
    echo "Espace disque OK"
fi