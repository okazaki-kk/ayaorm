make testgen
count=`git diff --name-only | wc -l`
if [ $count -eq 0 ]; then
  echo "No changes"
else
  echo "Changes"
  exit 1
fi
