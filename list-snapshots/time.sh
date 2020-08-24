for ((i = 0; i < 500; i++))
do
	date
	./list-snapshots | grep readyToUse
	sleep 1
done
