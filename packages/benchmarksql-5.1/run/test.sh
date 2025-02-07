j=0
resultline=''
for i in `cat test.out`
do
$resultline=$resultline$i
done
echo $resultline
