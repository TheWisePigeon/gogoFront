cd frontend
npm run build
cd ..
git add .
git commit -m "deploy: $1"
git push
echo "App deployed :)"