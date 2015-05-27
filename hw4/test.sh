i=0
while [ $i != 100 ]
do
    ./hw4 >> concurrency.data
    i=$(($i+1))
done
echo "$i"
