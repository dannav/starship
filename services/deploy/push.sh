# push searchd
scp bin/searchd dev@167.99.7.84:/home/dev
cat ~/serverpass | ssh -tt dev@167.99.7.84 "sudo mv /home/dev/searchd /opt/starship/searchd"

# push tikad
scp bin/tikad dev@167.99.7.84:/home/dev
cat ~/serverpass | ssh -tt dev@167.99.7.84 "sudo mv /home/dev/tikad /opt/starship/tikad"

# restart each service on remote server
cat ~/serverpass | ssh -tt dev@167.99.7.84 "sudo systemctl restart searchd.service"
cat ~/serverpass | ssh -tt dev@167.99.7.84 "sudo systemctl restart tikad.service"