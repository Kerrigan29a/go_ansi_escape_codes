#!/bin/bash
# Nothing to see here, just a utility I use to create new releases

CURRENT_VERSION=$(cat VERSION)
TO_UPDATE=(
    VERSION
)

echo -n "Current version is $CURRENT_VERSION, select new version: "
read NEW_VERSION
echo "********************************************************************************"
echo "Creating version $NEW_VERSION"
echo "********************************************************************************"

echo "********************************************************************************"
echo "Starting release $NEW_VERSION"
echo "********************************************************************************"
git flow release start $NEW_VERSION

for file in "${TO_UPDATE[@]}"
do
    echo "********************************************************************************"
    echo "Patching $file"
    echo "********************************************************************************"
    if [[ $(uname) == 'Darwin' ]]; then
        sed -i "" -e "s/$CURRENT_VERSION/$NEW_VERSION/g" $file
    else
        sed -i -e "s/$CURRENT_VERSION/$NEW_VERSION/g" $file
    fi
    git add $file
done

git commit -m "Releasing $NEW_VERSION"
echo "********************************************************************************"
echo "Finishing release $NEW_VERSION"
echo "********************************************************************************"
git flow release finish $NEW_VERSION

echo "********************************************************************************"
echo "Pushing release $NEW_VERSION"
echo "********************************************************************************"
git push origin $NEW_VERSION

echo "********************************************************************************"
echo "DONE"
echo "********************************************************************************"
