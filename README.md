# True Least Recently Used (LRU) Cache -
Least Recently Used Cache impletmented with Golang

- 1. If an item already exists, we need to remove it and add it to the beginning
- 2. An order of items is maintained
- 3. Deletion happens at the tail and addition happens at the head
- 4. If queue size limit is reached, remove the last item before adding new element
